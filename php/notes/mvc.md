# PHP 간단한 MVC 패턴

## 1. MVC 패턴 개요
MVC(Model-View-Controller)는 애플리케이션의 구조를 체계적으로 분리하는 디자인 패턴입니다.
- **Model**: 데이터 및 비즈니스 로직을 관리
- **View**: 사용자 인터페이스(UI) 및 출력 처리
- **Controller**: 사용자의 요청을 처리하고 Model과 View를 연결

## 2. 프로젝트 구조
```
/project_root
│── index.php          # 프론트 컨트롤러
│── controllers
│   └── HomeController.php
│── models
│   └── User.php
│── views
│   └── home.php
│── core
│   ├── Controller.php  # 기본 컨트롤러
│   ├── Model.php       # 기본 모델
│   └── Router.php      # 라우팅 처리
```

## 3. `index.php` (메인 진입점)
```php
require_once 'core/Router.php';
$router = new Router();
$router->run();
```

## 4. `core/Router.php` (라우팅 처리)
```php
class Router {
    public function run() {
        $controllerName = isset($_GET['controller']) ? $_GET['controller'] : 'Home';
        $action = isset($_GET['action']) ? $_GET['action'] : 'index';

        $controllerFile = "controllers/" . $controllerName . "Controller.php";
        if (file_exists($controllerFile)) {
            require_once $controllerFile;
            $controllerClass = $controllerName . "Controller";
            $controller = new $controllerClass();
            if (method_exists($controller, $action)) {
                $controller->$action();
            } else {
                echo "메서드 없음: $action";
            }
        } else {
            echo "컨트롤러 없음: $controllerName";
        }
    }
}
```

## 5. `controllers/HomeController.php` (컨트롤러)
```php
require_once "models/User.php";
class HomeController {
    public function index() {
        $userModel = new User();
        $users = $userModel->getUsers();
        require_once "views/home.php";
    }
}
```

## 6. `models/User.php` (모델)
```php
class User {
    public function getUsers() {
        return [
            ["name" => "Alice", "email" => "alice@example.com"],
            ["name" => "Bob", "email" => "bob@example.com"]
        ];
    }
}
```

## 7. `views/home.php` (뷰)
```php
<!DOCTYPE html>
<html>
<head>
    <title>사용자 목록</title>
</head>
<body>
    <h1>사용자 목록</h1>
    <ul>
        <?php foreach ($users as $user): ?>
            <li><?= $user['name'] ?> (<?= $user['email'] ?>)</li>
        <?php endforeach; ?>
    </ul>
</body>
</html>
```

## 8. 실행 방법
웹 브라우저에서 다음과 같이 실행합니다:
```
http://localhost/project_root/index.php?controller=Home&action=index
```

이렇게 하면 `HomeController`의 `index` 메서드가 실행되고 `User` 모델에서 데이터를 가져와 `home.php` 뷰에 전달합니다.

---

이 간단한 MVC 구조를 기반으로 기능을 확장하면 PHP 웹 애플리케이션을 체계적으로 개발할 수 있습니다. 😊
```
