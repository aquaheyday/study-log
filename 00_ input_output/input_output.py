# 입력과 출력
# 입력: input()
# 출력: print()

# input(): 입력값을 문자열로 반환
text = input() # 입력: hello
print(text) # hello

x = input() # 입력: hello
y = input() # 입력: world
print(x, y) # hello world 

# type(): 데이터 타입 확인
text = input() # 입력: hello
print(text) # hello
print(type(text)) # <class 'str'>

number = input() # 입력: 1234
print(number) # 1234
print(type(number)) # <class 'str'>

# int(): int type 으로 형 변환
number = input() # 입력: 1234
number = int(text)
print(number) # 1234
print(type(number)) # <class 'int'>

number = int(input()) # 입력: 1234
print(number) # 1234
print(type(a)) # <class 'int'>

x, y = input().split() # 입력: hello world
print(x, y) # hello world
print(x) # hello
print(y) # world

x, y = input().split() # 입력: 10 5
x = int(x)
y = int(y)
print(a + b) # 15

x, y = input().split() # 입력: 12 3
print(x + y) # 123

# map(): map 함수는 여러 개의 입력값을 동시에 처리할 때 유용합니다.
x, y = map(int, input().split()) # 입력: 3 10
print(x + y) # 13

# 여러 개의 입력을 리스트로 받아 출력하고 각 요소의 데이터 타입 확인
text_list = input().split() # 입력: hello world
print(text_list) # ['hello', 'world']
print(type(text_list)) # <class 'list'>
print(type(text_list[0])) # <class 'str'>

# map 함수를 사용하여 여러 개의 정수 입력을 리스트로 변환하여 출력하고 각 요소의 데이
number_list = list(map(int, input().split())) # 입력: 1 22 333
print(number_list) # [1, 22, 333]
print(type(number_list)) # <class 'list'>
print(type(number_list[0])) # <class 'int'>

# 출력 형식 지정
# sep 파라미터 (출력값 사이에 넣을 문자열을 선언)
# 기본값(공백): sep = ''
print('apple', 'banana', 'cherry') # apple banana cherry
print('apple', 'banana', 'cherry', sep = ', ') # apple, banana, cherry

# end 파라미터
# end 는 출력의 끝에 넣을 문자열을 지정
# 기본값(줄바꿈): end = '\n'
print('Hello', end = ' ') # Hello 
print('World') # Hello World

# 두 개의 입력을 받아 정수로 변환하고 몫과 나머지를 출력
a, b = map(int, input().split())
print(a // b, a % b)


# 빠른 입력: import sys, input = sys.stdin.readline
# input() 은 표준 입력을 읽어 개행 문자를 제거한 문자열을 반환합니다.
# sys.stdin.readline() 은 표준 입력 스트림에서 한 줄을 그대로 반환합니다. (개행 문자 포함)
# 개행 문자 제거 과정이 생략되어있어 sys.stdin.readline() 이 input() 보다 속도가 빠릅니다.
import sys
input = sys.stdin.readline
text = int(input())
