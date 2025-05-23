# 🛠 Kubernetes 클러스터 구축 가이드 (Rocky Linux 8 + containerd)

## 📌 시스템 사양

| 역할     | 사양                         |
|----------|------------------------------|
| Master   | RAM 8GB / CPU 8 Core         |
| Worker   | RAM 16GB / CPU 8 Core        |
| OS       | Rocky Linux 8                |
| 런타임   | containerd                   |

---

## 📁 공통 작업 (모든 노드)

## 1. 계정 생성 (선택)
#### 1-1. 사용자 계정 생성 (예: kubeadmin/kube1234)
```bash
sudo useradd -m kubeadmin
sudo passwd kubeadmin
```

#### 1-2. sudo 권한 부여
```bash
sudo usermod -aG wheel kubeadmin
```

#### 1-3. SSH 접속 허용 (선택)
```bash
sudo su - kubeadmin
```

## 2. 초기 서버 설정
```bash
sudo dnf update -y
sudo hostnamectl set-hostname k8s-master # 마스터 노드에서 실행
sudo hostnamectl set-hostname k8s-worker-1 # 워커 노드에서 실행 (1~n 형태)
```

### 3. SELinux 비활성화
```bash
sudo setenforce 0
sudo sed -i 's/^SELINUX=.*/SELINUX=permissive/' /etc/selinux/config
```

### 4. Swap 비활성화
```bash
sudo swapoff -a
sudo sed -i '/swap/d' /etc/fstab
```

### 5. 방화벽 설정 (선택)
```bash
sudo firewall-cmd --permanent --add-port=6443/tcp     # kube-apiserver
sudo firewall-cmd --permanent --add-port=10250/tcp    # kubelet
sudo firewall-cmd --permanent --add-port=30000-32767/tcp  # NodePort
sudo firewall-cmd --reload
```

### 6. containerd 설치

#### 6-1. 필수 패키지 설치
```bash
sudo dnf install -y yum-utils device-mapper-persistent-data lvm2
```

#### 6-2. Docker 저장소 추가
```bash
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
```

#### 6-3. containerd 설치
```bash
sudo dnf install -y containerd.io
```

#### 6-4. 기본 설정 파일 생성 및 수정
```bash
sudo mkdir -p /etc/containerd
containerd config default | sudo tee /etc/containerd/config.toml
sudo sed -i 's/SystemdCgroup = false/SystemdCgroup = true/' /etc/containerd/config.toml
```

#### 6-5. containerd 서비스 시작
```bash
sudo systemctl daemon-reexec
sudo systemctl enable --now containerd
```

### 7. Kubernetes 설치

#### 7-1. Kubernetes 저장소 추가
```bash
cat <<EOF | sudo tee /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://pkgs.k8s.io/core:/stable:/v1.29/rpm/
enabled=1
gpgcheck=1
gpgkey=https://pkgs.k8s.io/core:/stable:/v1.29/rpm/repodata/repomd.xml.key
EOF
```

#### 7-2. kubeadm, kubelet, kubectl 설치
```bash
sudo dnf install -y kubelet kubeadm kubectl
sudo systemctl enable --now kubelet
```

---

## 🧭 Master Node 설정

커널 모듈 및 sysctl 설정
```
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF

sudo modprobe br_netfilter

cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-iptables = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward = 1
EOF

sudo sysctl --system
```
### 1. 클러스터 초기화
```bash
sudo kubeadm init --pod-network-cidr=10.244.0.0/16 \
  --cri-socket=unix:///run/containerd/containerd.sock
```

### 2. kubeconfig 설정
```bash
mkdir -p $HOME/.kube
sudo cp /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

### 3. CNI 플러그인 설치 (Flannel)
```bash
kubectl apply -f https://raw.githubusercontent.com/flannel-io/flannel/master/Documentation/kube-flannel.yml
```

### 4. Worker 노드 조인 명령어 확인
```bash
kubeadm token create --print-join-command
```

---

## 🚀 Worker Node 설정

### 1. Master에서 받은 조인 명령어 실행
```bash
sudo kubeadm join <MASTER_IP>:6443 \
  --token <TOKEN> \
  --discovery-token-ca-cert-hash sha256:<HASH> \
  --cri-socket=unix:///run/containerd/containerd.sock
```

---

## ✅ 클러스터 확인 (Master 노드에서)
```bash
kubectl get nodes
```

---

## 🧩 후속 작업 제안

- Go API 배포: Deployment + Service 구성
- Ingress 설정: 도메인 연결 또는 리버스 프록시 구성
- 모니터링: Prometheus + Grafana (Helm 설치 권장)
- 인증서 관리: cert-manager + Let's Encrypt
- GitOps/CD: TeamCity 또는 ArgoCD 기반 자동화 구성
