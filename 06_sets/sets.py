#집합(set)은 중복되지 않은 데이터를 저장하는 자료형 (순서가 없으며 중복 원소 제거)

# 집합 선언
empty_set = set()
print(empty_set) # set()

numbers = {1, 2, 3, 4, 5}
print(numbers) # {1, 2, 3, 4, 5}
print(len(numbers)) # 5

duplicates = {1, 2, 3, 4, 5, 5}
print(duplicates) # {1, 2, 3, 4, 5}

# 원소 추가
# add(): 집합에 단일 원소 추가
# update(): 집합에 여러 원소 추가
numbers = {1, 2, 3}
numbers.add(4)
numbers.update({5, 6})
print(numbers) # {1, 2, 3, 4, 5, 6}

# 원소 삭제
# remove(): 특정 원소 제거 (원소가 없으면 오류 발생)
# discard(): 특정 원소 제거 (원소가 없어도 오류 발생 하지 않음)
numbers = {1, 2, 3, 4}
numbers.remove(3) # 숫자 3 삭제
numbers.discard(5) # 숫자 5 삭제
print(numbers) # {1, 2, 4}

# 집합
# 합집합: | 또는 union()
set1 = {1, 2, 3}
set2 = {3, 4, 5}
union_set = set1 | set2
print(union_set) # {1, 2, 3, 4, 5}
# 교집합: & 또는 intersection()
intersection_set = set1 & set2
print(intersection_set) # {3}
# 차집합: - 또는 difference()
difference_set = set1 - set2
print(difference_set) # {1, 2}
# 대칭 차집합: ^ 또는 symmetric_difference()
symmetric_difference_set = set1 ^ set2
print(symmetric_difference_set) # {1, 2, 4, 5}

# 집합 순회
# 집합은 순서가 없지만, 반복문을 사용해 모든 원소를 순회할 수 있습니다.
numbers = {1, 2, 3, 4}
for num in numbers:
    print(num)

# in 연산자: 특정 값이 집합에 포함되어 있는지 확인
# not in 연산자: 특정 값이 집합에 포함되지 않았는지 확인
fruits = {'apple', 'banana', 'cherry'}
# in 연산자
print('apple' in fruits) # True
print('orange' in fruits) # False
# not in 연산자
print('apple' not in fruits) # False
print('orange' not in fruits) # True

