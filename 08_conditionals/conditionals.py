# if 문
num = 10
if num > 5:
    print('num 은 5보다 크다')

age = int(input())
if age >= 18:
    print('age 는 18세 이상')

# else 문
num = 3
if num > 5:
    print('num 은 5보다 큽다')
else:
    print('num 은 5보다 작거나 같다')

score = int(input())
if score >= 60:
    print('합격')
else:
    print('불합격')

# elif 문
num = 5
if num > 5:
    print('num 은 5보다 크다')
elif num == 5:
    print('num 은 5와 같다')
else:
    print('num 은 5보다 작다')

score = int(input())
if score >= 90:
    print('A 학점')
elif score >= 80:
    print('B 학점')
elif score >= 70:
    print('C 학점')
elif score >= 60:
    print('D 학점')
else:
    print('F 학점')

# 중첩 조건문
x = 10
y = 20
if x > 5:
    if y > 5:
        print('x 와 y 모두 5보다 크다')

score = int(input())
if score >= 60:
    if score >= 80:
        print('80점 이상 합격')
    else:
        print('60점 이상 80점 미만 합격')
else:
    print('불합격')

# 복합 조건문 (여러 조건을 결합하여 복합적인 조건을 처리)
x = 10
y = 20
if x > 5 and y > 5:
    print('x 와 y 모두 5보다 크다')
else:
    print('x 와 y 중 하나 또는 둘 다 5보다 작거나 같다')

x = 10
y = 5
if x > 5 or y > 5:
    print('x 또는 y 중 하나는 5보다 크다')
else:
    print('x 와 y 모두 5보다 작거나 같다')
