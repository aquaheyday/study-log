# 딕셔너리(dictionary)는 키-값 쌍(key-value pair)으로 데이터를 저장하는 자료형임
# 딕셔너리는 파이썬에서 가변적(mutable)임

# 딕셔너리 생성
# 딕셔너리는 {} 또는 dict() 함수를 사용해 생성
empty_dict = dict()
print(empty_dict) # {}

student_scores = {'Alice': 85, 'Bob': 90, 'Charlie': 78}
print(student_scores) # {'Alice': 85, 'Bob': 90, 'Charlie': 78}
print(len(student_scores))

# 딕셔너리 원소 추가/수정
# dict[key] = value 구문: 키-값 쌍 추가 및 기존 값 수정
student_scores = {'Alice': 85, 'Bob': 90}
print(student_scores) # {'Alice': 85, 'Bob': 90}
student_scores['Charlie'] = 78
print(student_scores) # {'Alice': 85, 'Bob': 90, 'Charlie': 78}
student_scores['Alice'] = 88
print(student_scores) # {'Alice': 88, 'Bob': 90, 'Charlie': 78}

# 딕셔너리 원소 제거
# del: 특정 키-값 제거
# pop(): 특정 키의 값 제거 및 반환 (키값이 없으면 오류 발생)
student_scores = {'Alice': 85, 'Bob': 90, 'Charlie': 78}
del student_scores['Bob'] # 'Bob' 삭제
print(student_scores) # {'Alice': 85, 'Charlie': 78}
removed_score = student_scores.pop('Charlie')
print(removed_score) # 78
print(student_scores) # {'Alice': 88}

# 딕셔너리 키 와 값 조회
# keys(): 모든 키 반환
# values(): 모든 값 반환
# items(): 모든 키-값 쌍을 튜플 형태로 반환
student_scores = {'Alice': 85, 'Bob': 90, 'Charlie': 78}
print(student_scores.keys()) # dict_keys(['Alice', 'Bob', 'Charlie'])
print(student_scores.values()) # dict_values([88, 90, 78])
print(student_scores.items()) # dict_items([('Alice', 85), ('Bob', 90), ('Charlie', 78)])

# 딕셔너리 순회
# 딕셔너리의 키, 값, 키-값 쌍을 각각 반복문으로 순회할 수 있음
student_scores = {'Alice': 88, 'Bob': 90, 'Charlie': 78}
for name in student_scores.keys():
    print(name)
for value in student_scores.values():
    print(value)
for item in student_scores.items():
    print(item)


student_scores = {'Alice': 88, 'Bob': 90, 'Charlie': 78}
# in 연산자 사용
print('Alice' in student_scores) # True
print('David' in student_scores) # False
# not in 연산자 사용
print('Alice' not in student_scores) # False
print('David' not in student_scores) # True
