# 🐛 Docker로 Blue/Green 배포시 `no space left on device` 오류

## ⚠️ 문제 상황
- 발생 날짜: 2025-04-08  
- 발생 환경 linux, docker(blue/green), next.js
- 재현 방법 
1. 도커가 사용중인 디스크 용량이 꽉 찬 상태에서 배포시 용량이 부족해 빌드가 실패함.

---

## 🔍 원인 분석
- Blue/Green 배포시 이미지나 정적 데이터가 누적되어 디스크 용량이 꽉 차게되는 현상 발생

---

## 🛠 해결 방법
- `deploy.sh` 에 기존 Blue/Green 이미지 자동 정리 추가
- 빌드시 기존 이미지를 캐시로 이용하는 현상을 해결하기 위해 `docker-compose.yml` 와 `Dockerfile` 에 `BUILD_TAG` 를 활용해 캐시 무효화
- 캐시 무효화시 프로젝트 빌드가 끝난 후에 하는게 좋음(빌드 속도 이슈)

#### `deploy.sh`
```sh
# 현재 시각을 기반으로 고유한 빌드 태그 생성 (캐시 무효화 목적)
BUILD_TAG=$(date +%Y%m%d%H%M%S)

# BUILD_TAG 는 deploy.sh → docker-compose.yml → Dockerfile ARG로 전달되어 캐시를 무력화하고 이미지 구분에 활용됨
BUILD_TAG=$BUILD_TAG docker-compose build $TARGET_CONTAINER
BUILD_TAG=$BUILD_TAG docker-compose up -d $TARGET_CONTAINER

# 종료될 컨테이너가 사용 중인 이미지 ID 조회 (docker-compose rm 컨테이너 삭제보다 먼저 선언되어야함)
IMAGE_ID=$(docker inspect --format='{{.Image}}' $ACTIVE_CONTAINER)

echo "Deleting image used by $ACTIVE_CONTAINER..."
if docker ps -q --filter ancestor="$IMAGE_ID" | grep -q .; then
  echo "Image $IMAGE_ID is still in use by a running container. Skipping delete."
else
  echo "Deleting image $IMAGE_ID used by $ACTIVE_CONTAINER"
  docker rmi "$IMAGE_ID"
fi

echo "Cleaning up dangling images..."
# 태그가 없고 사용되지 않는 이미지(<none>) 정리
docker image prune -f
```

#### docker-compose.yml
```
build:
  args:
    # 빌드 시점의 고유 태그 전달 (Dockerfile에 ARG로 전달됨)
    BUILD_TAG: ${BUILD_TAG}
```

#### Dockerfile
```
# 외부로부터 전달받을 빌드 인자 선언
ARG BUILD_TAG
# 환경변수 설정 (필수는 아님)
ENV BUILD_TAG=$BUILD_TAG
# 빌드 태그를 파일로 저장(캐시 무효화)
RUN echo "Build ID: $BUILD_TAG" > /build_info.txt
```

---

## 🚀 결과
```
Deleting image used by container-name...
14:09:35   Deleting image sha256:42c25880fc7a6fd9d5dfb01287e47b707cd6e8993d... used by container_name
14:09:37   Untagged: image-name:latest
14:09:37   Deleted: sha256:42c25880fc7a6fd9d5dfb...
14:09:37   Cleaning up dangling images...
14:09:38   Deleted Images:
14:09:38   deleted: sha256:7bf83cdffb7844b3d5b33...
14:09:38   Total reclaimed space: 1.128GB
Removing ... done
```

- Blue/Green 배포시 기존 컨테이너가 사용했던 이미지도 삭제됨 (도커 디스크 용량 증가 ❌)
