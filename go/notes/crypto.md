# 🔐 Go 언어 암호화와 해싱

Go는 표준 라이브러리를 통해 다양한 암호화 및 해시 기능을 제공합니다.  
`crypto`, `encoding`, `golang.org/x/crypto` 등을 사용하면 안전한 데이터 처리를 구현할 수 있습니다.

---

## 1️⃣ 해싱 (Hashing)

### 1) `SHA-256` 해시

```go
import (
    "crypto/sha256"
    "fmt"
)

func main() {
    data := []byte("hello")
    hash := sha256.Sum256(data)
    fmt.Printf("SHA256: %x\n", hash)
}
```

✔ `sha256.Sum256()`은 32바이트 고정 길이 해시를 반환  
✔ `%x`는 16진수 출력  

---

### 2) `SHA-1` / `SHA-512` / `MD5`

```go
crypto/sha1    // sha1.New()
crypto/sha512  // sha512.Sum512()
crypto/md5     // md5.Sum()
```

✔ `MD5`, `SHA-1`은 보안상 권장되지 않음 (빠른 식별 용도에만 사용)

---

## 2️⃣ HMAC (해시 기반 메시지 인증 코드)

```go
import (
    "crypto/hmac"
    "crypto/sha256"
    "fmt"
)

func main() {
    key := []byte("secret-key")
    message := []byte("hello")

    mac := hmac.New(sha256.New, key)
    mac.Write(message)
    result := mac.Sum(nil)

    fmt.Printf("HMAC: %x\n", result)
}
```

✔ 메시지 무결성/인증 검증 용도로 사용  
✔ 수신 측에서 동일한 키로 검증 가능  

---

## 3️⃣ 대칭키 암호화 (AES)

```go
import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
)

func encryptAES(plaintext, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
    return ciphertext, nil
}
```

✔ `key`는 16, 24, 32바이트 (AES-128, AES-192, AES-256)  
✔ `iv`(초기화 벡터)는 매번 랜덤으로 생성  
✔ `CFB`, `CBC`, `GCM` 등 다양한 모드 사용 가능  

---

## 4️⃣ 비대칭키 암호화 (RSA)

```go
import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
)

func main() {
    privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
    publicKey := &privateKey.PublicKey

    message := []byte("secret")
    label := []byte("")
    hash := sha256.New()

    // 암호화
    ciphertext, _ := rsa.EncryptOAEP(hash, rand.Reader, publicKey, message, label)

    // 복호화
    decrypted, _ := rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, label)

    fmt.Println("복호화된 메시지:", string(decrypted))
}
```

✔ RSA는 공개키(Public Key)로 암호화, 개인키(Private Key)로 복호화  
✔ OAEP는 안전한 패딩 방식  

---

## 5️⃣ 비밀번호 해싱 (bcrypt)

```bash
go get golang.org/x/crypto/bcrypt
```

```go
import (
    "golang.org/x/crypto/bcrypt"
)

func main() {
    password := []byte("secret")

    // 해싱
    hash, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

    // 비교
    err := bcrypt.CompareHashAndPassword(hash, []byte("secret"))
    fmt.Println("비밀번호 일치 여부:", err == nil)
}
```

✔ bcrypt는 느리게 설계되어 무차별 대입 공격에 강함  
✔ `DefaultCost`는 10 (보안성과 속도 균형)  
✔ 비교 시 반드시 `CompareHashAndPassword` 사용  

---

## 🎯 정리

| 기능 | 패키지 | 설명 |
|------|--------|------|
| SHA-256 해싱 | `crypto/sha256` | 일반 데이터 해싱 |
| HMAC | `crypto/hmac` | 키 기반 메시지 무결성 |
| AES 암호화 | `crypto/aes` | 대칭키 방식 |
| RSA 암/복호화 | `crypto/rsa` | 공개키 기반 |
| bcrypt | `golang.org/x/crypto/bcrypt` | 비밀번호 해싱 |
