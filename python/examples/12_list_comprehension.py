# list comprehension 은 리스트를 생성하는 방법입니다.

# 제곱수 리스트 생성
squares = []
for x in range(1, 11):
    squares.append(x ** 2)
print(squares) # [1, 4, 9, 16, 25, 36, 49, 64, 81, 100]
# list comprehension 을 사용한 제곱수 리스트 생성
squares = [x ** 2 for x in range(1, 11)]
print(squares) # [1, 4, 9, 16, 25, 36, 49, 64, 81, 100]

# 짝수 리스트 생성
even_numbers = []
for x in range(1, 11):
    if x % 2 == 0:
        even_numbers.append(x)
print(even_numbers) # [2, 4, 6, 8, 10]
# list comprehension 을 사용한 짝수 리스트 생성
even_numbers = [x for x in range(1, 11) if x % 2 == 0]
print(even_numbers) # [2, 4, 6, 8, 10]
