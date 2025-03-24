# 📄 PHP PDF 및 엑셀 처리

PHP는 웹에서 **리포트, 영수증, 통계자료** 등을 PDF 또는 엑셀로 출력할 수 있습니다.  
이를 위해 보통 다음과 같은 오픈소스 라이브러리를 사용합니다.

| 작업 유형 | 대표 라이브러리               |
|-----------|-------------------------------|
| PDF 생성  | `FPDF`, `TCPDF`, `DOMPDF`     |
| Excel 처리| `PhpSpreadsheet`              |

---

## 1️⃣ PDF 생성 - FPDF 사용

#### 1. 설치 (Composer 또는 수동)

```bash
composer require setasign/fpdf
```

또는 직접 다운로드: [http://www.fpdf.org/](http://www.fpdf.org/)

#### 2. 사용 예시

```php
<?php
require 'vendor/autoload.php'; // 또는 fpdf.php

$pdf = new FPDF();
$pdf->AddPage();
$pdf->SetFont('Arial', 'B', 16);
$pdf->Cell(40, 10, 'Hello PDF!');
$pdf->Output();
?>
```

✔ `AddPage()` → 새 페이지 추가  
✔ `SetFont()` → 폰트 설정  
✔ `Cell()` → 텍스트 출력  
✔ `Output()` → 브라우저로 출력 또는 저장  

---

## 2️⃣ PDF 생성 - DOMPDF (HTML 기반 PDF)

#### 1. 설치

```bash
composer require dompdf/dompdf
```

#### 2. HTML → PDF 변환 예시

```php
<?php
use Dompdf\Dompdf;

require 'vendor/autoload.php';

$dompdf = new Dompdf();
$dompdf->loadHtml("<h1>안녕하세요, PDF입니다</h1>");
$dompdf->setPaper("A4", "portrait");
$dompdf->render();
$dompdf->stream("example.pdf");
?>
```

✔ HTML/CSS 기반으로 PDF 생성 → 웹 디자인 감각 그대로 PDF 출력 가능  

---

## 3️⃣ Excel 처리 - PhpSpreadsheet

#### 1. 설치

```bash
composer require phpoffice/phpspreadsheet
```

#### 2. 엑셀 파일 읽기

```php
<?php
use PhpOffice\PhpSpreadsheet\IOFactory;

require 'vendor/autoload.php';

$spreadsheet = IOFactory::load("data.xlsx");
$sheet = $spreadsheet->getActiveSheet();
$data = $sheet->toArray();

print_r($data);
?>
```

✔ `toArray()`로 전체 데이터를 배열로 받아서 반복 처리 가능  

#### 3. 엑셀 파일 생성

```php
<?php
use PhpOffice\PhpSpreadsheet\Spreadsheet;
use PhpOffice\PhpSpreadsheet\Writer\Xlsx;

$spreadsheet = new Spreadsheet();
$sheet = $spreadsheet->getActiveSheet();

$sheet->setCellValue('A1', '이름');
$sheet->setCellValue('B1', '나이');
$sheet->setCellValue('A2', '홍길동');
$sheet->setCellValue('B2', 30);

$writer = new Xlsx($spreadsheet);
$writer->save("output.xlsx");
?>
```

✔ 코드로 직접 셀을 설정하고 저장할 수 있습니다.  
✔ `header()`를 사용하여 **브라우저에서 다운로드**도 가능  

#### 4. 브라우저 다운로드 (엑셀)

```php
header("Content-Type: application/vnd.openxmlformats-officedocument.spreadsheetml.sheet");
header("Content-Disposition: attachment; filename=\"report.xlsx\"");
$writer->save("php://output");
```

---

## 4️⃣ PDF/Excel 저장 위치 및 권한 주의

- `save()` 함수로 파일을 저장하려면 해당 디렉토리에 **쓰기 권한**이 있어야 합니다.  
- 브라우저로 전송할 경우 `php://output` 사용  
- 한글 폰트나 스타일 적용이 필요하다면 외부 `.ttf` 폰트를 추가로 설정해야 할 수도 있습니다.

---

## 5️⃣ 실무 활용 예시

| 기능                             | 방법 |
|----------------------------------|------|
| 견적서, 영수증 PDF 출력          | FPDF, DOMPDF 사용 |
| 관리자용 보고서 엑셀로 출력       | PhpSpreadsheet 사용 |
| 사용자 업로드 엑셀 분석           | PhpSpreadsheet + `toArray()` |
| PDF에 표, 이미지 삽입            | TCPDF / DOMPDF 테이블, 이미지 삽입 기능 활용 |
| HTML을 그대로 PDF로 변환         | DOMPDF 추천 |

---

## 🎯 정리

✔ PHP에서 PDF는 `FPDF`, `DOMPDF`, `TCPDF` 등을 통해 생성할 수 있습니다.  
✔ HTML 기반 PDF가 필요하면 `DOMPDF`, 코드 기반 레이아웃이면 `FPDF`가 적합합니다.  
✔ Excel 파일 읽기/생성/다운로드는 `PhpSpreadsheet` 하나로 대부분 처리 가능합니다.  
✔ 브라우저 다운로드 시에는 `php://output`과 `Content-Type` 헤더를 꼭 설정해야 합니다.

