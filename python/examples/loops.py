# for 문
fruits = ['apple', 'banana', 'orange']
for fruit in fruits:
    print(fruit)

text = 'Python'
for char in text:
    print(char)

# range(): 함수는 특정 범위의 숫자를 생성
for i in range(6):
    print(i) # 0 부터 5 까지 출력

for i in range(1, 6):
    print(i) # 1 부터 5 까지 출력

for i in range(0, 10, 2):
    print(i) # 0 부터 9 까지 짝수만 출력

for i in range(10, 0, -1):
    print(i) # 10 부터 1 까지 출력

# break: 반복문을 종료
# continue: 현재 반복을 건너뛰고 다음 반복을 계속함
for i in range(1, 10):
    if i == 5:
        break # 5 일때 반복문 종료
    print(i) # 1 부터 4 까지 출력

# 짝수만 출력
for i in range(1, 10):
    if i % 2 != 0:
        continue # 2로 나눠 나머지가 있을 경우 반복문 건너뜀
    print(i) # 1 부터 10 중에 짝수만 출력

# while 문
i = 1
while i <= 5:
    print(i) # 1 부터 5까지 출력
    i += 1

total = 0
i = 1
while i <= 100:
    total += i
    if total > 50:
        break
    i += 1
print(total) # 55 출력

i = 0
while i < 10:
    i += 1
    if i % 2 == 0:
        continue
    print(i) # 1 부터 10 중 홀수만 출력

# 무한 루프
# while True 처럼 조건이 항상 참일 경우, 무한 루프에 빠질 수 있음
# 이를 제어하기 위해 break를 사용할 수 있음
# break를 사용하여 무한 루프 종료
count = 0
while True:
    print(count)
    count += 1
    if count == 3:
        break

# 입력된 숫자의 합을 구하다가 0이 입력되면 종료
total = 0
while True:
    num = int(input())
    if num == 0:
        break # 0이 입력되면 종료
    total += num # 입력된 숫자를 더함
print(total)

# 중첩된 while
# 구구단 출력
i = 2
while i <= 9: # 2단부터 9단까지 반복
    j = 1
    while j <= 9: # 각 단에 대해 1부터 9까지 곱셈
        print(i, '*', j, '=', i * j)
        j += 1
    print() # 각 단이 끝날 때마다 줄바꿈
    i += 1

i = 0
while i < 3:
    j = 0
    while j < 5:
        if j == 3:
            break # 가장 안쪽의 while 문 종료
        print(i, j)
        j += 1
    print()
    i += 1

# 중첩된 while 문에서 continue 사용
i = 0
while i < 3:
    j = 0
    while j < 5:
        j += 1
        if j % 2 == 0:
            continue # 가장 안쪽의 for 문에서 j % 2 == 0일 때 건너뜀
        print(i, j)
    print()
    i += 1
