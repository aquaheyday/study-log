name = 'HongGilDong'
age = 25
cm = 173.5
is_men = True
skills = ['HTML', 'CSS', 'JS', 'Python']
person_info = {
    'name':'HongGilDong',
    'age': 25,
    'is_men': True,
    'skills': ['HTML', 'CSS', 'JS', 'Python']
    }

print('Name:', name)
print('Name type:', type(name))
print('Name id:', id(name))
print('Name length:', len(name))

print('Age:', age)
print('Age type:', type(age))
print('Age id:', id(age))

print('Cm:', cm)
print('Cm type:', type(cm))
print('Cm id:', id(cm))

print('Is men: ', is_men)
print('Is men type:', type(is_men))
print('Is men id:', id(is_men))

print('Skills: ', skills)
print('Skills type:', type(skills))
print('Skills id:', id(skills))
print('Skills length:', len(skills))

print('Person Info:', person_info)
print('Person Info type:', type(person_info))
print('Person Info id:', id(person_info))
print('Person Info length:', len(person_info))


name, age, is_men = 'HongGilDong', 25, True

print(name, age, is_men)
