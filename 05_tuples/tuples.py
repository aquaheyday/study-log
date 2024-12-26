# 빈 튜플 생성
empty_tuple = ()
print(empty_tuple) # ()

# 초기 값이 있는 튜플 생성
coordinates = (37.5665, 126.9780)
print(coordinates) # (37.5665, 126.9780)
print(len(coordinates)) # 2

numbers = 1, 2, 3
print(numbers) # (1, 2, 3)

coordinates = (37.5665, 126.9780)
print(coordinates[0]) # 37.5665
print(coordinates[-1]) # 126.9780

numbers = 1, 2, 3, 4, 5
print(numbers[1:3]) # (2, 3)
print(numbers[:3]) # (1, 2, 3)
print(numbers[2:]) # (3, 4, 5)
print(numbers[-2:]) # (4, 5)
print(numbers[::2]) # (1, 3, 5)
print(numbers[::-1])  # (5, 4, 3, 2, 1)

coordinates = (37.5665, 126.9780)
x, y = coordinates
print(x, y) # 37.5665 126.978

a, b = 5, 10
print(a, b) # 5, 10
b, a = a, b
print(a, b) # 10, 5
