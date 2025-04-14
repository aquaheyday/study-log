# 🐘 PHP

이 폴더는 **PHP 학습 과정**에서 정리한 자료, 예제 코드, 프로젝트를 보관하는 공간입니다.  

---

## 📂 디렉토리 구성

| 폴더명 | 설명 |
|---|---|
| [notes](./notes) | 개념 정리 및 학습 노트 |
| [examples](./examples) | 주요 기능별 예제 코드 |

---

## 📋 PHP 개념 정리 목록

### 📌 기본 개념
| 주제 | 파일명 | 설명 |
|---|---|---|
| PHP 개요 | [php-intro.md](./notes/php-intro.md) | PHP란? 특징과 장점 |
| 개발 환경 설정 | [setup-php.md](./notes/setup-php.md) | PHP 설치, XAMPP, Composer 활용 |
| 변수와 데이터 타입 | [variables.md](./notes/variables.md) | 숫자, 문자열, 배열, 객체 등 |
| 연산자와 표현식 | [operators.md](./notes/operators.md) | 산술, 비교, 논리, 비트 연산자 |
| 제어문 | [control-flow.md](./notes/control-flow.md) | if, switch, for, while 사용법 |

### 🔲 함수와 객체지향 프로그래밍
| 주제 | 파일명 | 설명 |
|---|---|---|
| 함수 사용법 | [functions.md](./notes/functions.md) | 함수 정의, 매개변수, 반환값, 익명 함수 |
| 클래스와 객체 | [oop.md](./notes/oop.md) | 클래스, 객체, 상속, 다형성, 인터페이스 |
| 네임스페이스와 오토로딩 | [namespaces.md](./notes/namespaces.md) | 네임스페이스 활용, PSR-4 오토로딩 |
| 예외 처리 | [exceptions.md](./notes/exceptions.md) | try-catch, throw, 사용자 정의 예외 |

### 🔄 데이터 처리
| 주제 | 파일명 | 설명 |
|---|---|---|
| 파일 입출력 | [file-io.md](./notes/file-io.md) | 파일 읽기, 쓰기, 삭제 |
| 데이터베이스 연동 | [database.md](./notes/database.md) | MySQL, PDO, Prepared Statement |
| 세션과 쿠키 | [session-cookie.md](./notes/session-cookie.md) | 사용자 상태 관리, 보안 고려 사항 |

### 🌍 웹 개발
| 주제 | 파일명 | 설명 |
|---|---|---|
| 폼 처리 및 입력 검증 | [form-handling.md](./notes/form-handling.md) | GET, POST 요청 처리, 필터링 |
| HTTP 요청 및 API 연동 | [http-api.md](./notes/http-api.md) | cURL, file_get_contents 활용 |
| PHP와 JavaScript 연동 | [php-js.md](./notes/php-js.md) | AJAX, JSON, WebSocket |

### 🚀 고급 개념
| 주제 | 파일명 | 설명 |
|---|---|---|
| 정규 표현식 | [regex.md](./notes/regex.md) | preg_match, preg_replace 활용 |
| 멀티스레딩 및 병렬 처리 | [multithreading.md](./notes/multithreading.md) | pthreads, async 처리 |
| 고급 함수형 프로그래밍 | [functional-programming.md](./notes/functional-programming.md) | 클로저, 콜백, 고차 함수 |
| Composer 패키지 관리 | [composer.md](./notes/composer.md) | 의존성 관리, autoload 설정 |

### 📊 PHP와 데이터 과학
| 주제 | 파일명 | 설명 |
|---|---|---|
| GD와 이미지 처리 | [gd-library.md](./notes/gd-library.md) | 이미지 생성, 조작, 썸네일 생성 |
| PHP 데이터 분석 | [data-analysis.md](./notes/data-analysis.md) | PHP로 데이터 처리 및 시각화 |
| PDF 및 엑셀 처리 | [pdf-excel.md](./notes/pdf-excel.md) | PHPSpreadsheet, TCPDF 활용 |

### 🔒 보안 및 성능 최적화
| 주제 | 파일명 | 설명 |
|---|---|---|
| PHP 보안 기본 | [security-basics.md](./notes/security-basics.md) | XSS, SQL Injection 방어 |
| 해싱과 암호화 | [encryption.md](./notes/encryption.md) | password_hash, OpenSSL |
| 캐싱과 성능 최적화 | [caching.md](./notes/caching.md) | OPcache, Redis, Memcached |

### 🛠️ 테스트 및 배포
| 주제 | 파일명 | 설명 |
|---|---|---|
| 단위 테스트 | [testing.md](./notes/testing.md) | PHPUnit 활용법 |
| 패키징과 배포 | [packaging.md](./notes/packaging.md) | phar, Docker, CI/CD 배포 |

---

## 📋 예제 목록

| 주제 | 파일명 | 설명 |
|---|---|---|
| 변수와 자료형 | [variables.php](./examples/variables.php) | 변수 선언, 자료형 변환 등 기초 예제 |
| 연산자 | [operators.php](./examples/operators.php) | 산술, 비교, 논리 연산자 활용 예제 |
| 제어문 | [control-flow.php](./examples/control-flow.php) | if문, switch문, 반복문 예제 |
| 함수 | [functions.php](./examples/functions.php) | 함수 정의, 매개변수, 반환값 예제 |
| 배열과 연관 배열 | [arrays.php](./examples/arrays.php) | 배열 생성, 다차원 배열 등 |
| 파일 입출력 | [file-handling.php](./examples/file-handling.php) | 파일 열기, 쓰기, 읽기 |
| 폼 데이터 처리 | [form-handling.php](./examples/form-handling.php) | GET/POST 처리 예제 |
| 세션과 쿠키 | [session-cookie.php](./examples/session-cookie.php) | 세션 관리 및 쿠키 설정 예제 |
| 예외 처리 | [exceptions.php](./examples/exceptions.php) | try-catch, 사용자 정의 예외 |
| 데이터베이스 연동 | [database.php](./examples/database.php) | PDO로 DB 연결 및 쿼리 실행 예제 |

---

## 📚 참고 자료
- [PHP 공식 문서](https://www.php.net/manual/en/) - PHP의 공식 매뉴얼 및 가이드  
- [PHP The Right Way](https://phptherightway.com/) - PHP 개발을 위한 모범 사례 및 권장 방법론  
- [Composer 공식 문서](https://getcomposer.org/doc/) - PHP 패키지 관리 도구 Composer 사용법  
- [PHP-FIG PSR 표준](https://www.php-fig.org/psr/) - PHP 코드 스타일 및 표준 가이드

