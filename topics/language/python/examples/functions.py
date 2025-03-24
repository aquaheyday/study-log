# 함수

# 함수 생성
def add(x, y):
    return x + y
# 함수 호출
result = add(3, 4)
print(result) # 7

def fruit_info(name, price):
    print(name, price)
fruit_info('apple', 300) # apple 300

def square(x):
    return x * x
result = square(5)
print(result) # 25

def add_subtract(a, b):
    return a + b, a - b
sum_result, diff_result = add_subtract(10, 3)
print(sum_result, diff_result) # 13 7

# 전역 변수
def show_count():
    print(count) # 전역 변수 호출

count = 10 # 전역 변수 생성
show_count()

# global: 함수 내부에서 전역 변수를 수정할 때 호출
def increment():
    global count
    count += 1 # 전역 변수 수정

count = 10 # 전역 변수 생성
increment()
print(count) # 11

# 리스트나 딕셔너리는 global 호출 없이 수정 가능
def append_number():
    numbers.append(4)

numbers = [1, 2, 3]
append_number()
print(numbers) # [1, 2, 3, 4]

# 재귀 함수
def countdown(n):
    if n <= 0:
        print('end')
        return
    else:
        print(n)
        countdown(n - 1)

countdown(5)
