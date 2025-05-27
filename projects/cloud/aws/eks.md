
## 1. AWS EKS 클러스터 생성
### 1-1. EKS 클러스터 생성
AWS Console → EKS → “클러스터 생성”  
이름: go-api-cluster  
버전: 가장 최신 (ex. 1.32)  
VPC: 새로 만들기 (NAT gateways: 1 per AZ)  
서브넷: 퍼블린 서브넷 2개 이상 선택  

### 1-2. VPC/서브넷 태그 설정 (가장 그대로 ELB 업데이트 위해 필요)
서브넷 (Public Subnet 에 설정)  
태그 1: kubernetes.io/cluster/go-api-cluster = owned  
태그 2: kubernetes.io/role/elb = 1  
서브넷 설정  
"Enable auto-assign public IPv4 address" 보안 IP 자동할당 확인 및 활성화  

### 1-3. 클러스터 연결 (kubectl)
```
aws eks update-kubeconfig --region ap-northeast-2 --name go-api-cluster
kubectl get nodes  # → Ready 2개 확인
```

## 2. Go API Docker 이미지 만들고 ECR에 업로드

### 2-1. Go API 프로젝트 생성
```
// main.go
package main
import (
    "fmt"
    "net/http"
)
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Go API is running")
    })
    http.ListenAndServe(":8080", nil)
}

```

```
go mod init go-api
```

### 2-2. Dockerfile 작성
```
FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go mod tidy && go build -o main

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]

```

### 2-3. 이미지 빌드 및 ECR 푸시

```
# 빌드
docker build -t go-api .

# ECR 로그인
aws ecr create-repository --repository-name go-api  # 이미 있으면 에러 무시해도 됨
aws ecr get-login-password --region ap-northeast-2 | docker login --username AWS --password-stdin 134606631094.dkr.ecr.ap-northeast-2.amazonaws.com

# 푸시
docker tag go-api:latest 134606631094.dkr.ecr.ap-northeast-2.amazonaws.com/go-api:latest
docker push 134606631094.dkr.ecr.ap-northeast-2.amazonaws.com/go-api:latest

```

## 3. EKS에 Go API 배포

### 3-1. Deployment
```
# go-api-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-api
spec:
  replicas: 2
  selector:
    matchLabels:
      app: go-api
  template:
    metadata:
      labels:
        app: go-api
    spec:
      containers:
      - name: go-api
        image: 134606631094.dkr.ecr.ap-northeast-2.amazonaws.com/go-api:latest
        ports:
        - containerPort: 8080

```

### 3-2. Service
```
# go-api-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: go-api
spec:
  selector:
    app: go-api
  ports:
  - port: 80
    targetPort: 8080
  type: ClusterIP

```

```
kubectl apply -f go-api-deployment.yaml
kubectl apply -f go-api-service.yaml
```


## 4. Ingress Controller 설치
```
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update

helm install ingress-nginx ingress-nginx/ingress-nginx \
  --namespace ingress-nginx \
  --create-namespace \
  --set controller.publishService.enabled=true \
  --set controller.service.annotations."service\.beta\.kubernetes\.io/aws-load-balancer-scheme"="internet-facing"

```

```
kubectl get svc -n ingress-nginx
# → EXTERNAL-IP로 AWS LoadBalancer 주소 확인
```

## 5. Ingress 리소스 작성 (도메인 연동)
```
# go-api-ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: go-api-ingress
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  ingressClassName: nginx
  rules:
  - host: api.yourdomain.com # 테스트 필요시 주석 후 http: 앞에 - 추가 하면 EXTERNAL -IP 로 확인 가능
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: go-api
            port:
              number: 80

```

```
kubectl apply -f go-api-ingress.yaml

```

## 6. 도메인 (가비아) 연결
AWS ELB 주소: kubectl get svc -n ingress-nginx

가비아 DNS 설정:

호스트: api

타입: CNAME

값: k8s-ingressn-xxxxx.elb.ap-northeast-2.amazonaws.com.

이미 가비아에서 SSL이 적용되어 있다면, HTTPS도 자동 작동

## 7. 테스트
```
curl http://api.yourdomain.com/
# 또는
curl https://api.yourdomain.com/
```
