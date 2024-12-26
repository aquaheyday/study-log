empty_list = list()
print(len(empty_list)) # 0

fruits = ['apple', 'banana', 'orange']
print('Fruits:', fruits)
print('Number of fruits:', len(fruits))

fruits = ['apple', 'banana', 'orange']
first_fruit = fruits[0]
print(first_fruit) # apple
second_fruit = fruits[1]
print(second_fruit) # banana
last_fruit = fruits[2]
print(last_fruit) # orange
last_index = len(fruits) - 1
last_fruit = fruits[last_index]
print(last_fruit) # orange


fruits = ['apple', 'banana', 'orange']
last_fruit = fruits[-1]
print(last_fruit) # orange
second_last = fruits[-2]
print(second_last) # banana

# slicing
fruits = ['apple', 'banana', 'orange']
all_fruits = fruits[0:3]
print(all_fruits) # ['apple', 'banana', 'orange']
all_fruits = fruits[0:]
print(all_fruits) #['apple', 'banana', 'orange']
banana_orange = fruits[1:3]
print(banana_orange) # ['banana', 'orange']
orange = fruits[2:]
print(orange) #['orange']

fruits = ['apple', 'banana', 'orange']
all_fruits = fruits[-3:]
print(all_fruits) # ['apple', 'banana', 'orange']
banana = fruits[-2:-1]
print(banana) # ['banana']
banana_orange = fruits[-2:]
print(banana_orange) # ['banana', 'orange']

fruits = ['banana', 'orange']
fruits[0] = 'apple'
print(fruits) # ['apple', 'banana', 'orange']

fruits = ['apple', 'banana', 'orange']
does_exist = 'banana' in fruits
print(does_exist) # True
does_exist = 'mango' in fruits
print(does_exist) # False

fruits = ['apple', 'banana']
fruits.append('orange')
print(fruits) # ['apple', 'banana', 'orange']
fruits.append('mango')
print(fruits) # ['apple', 'banana', 'orange', 'mango']

# insert
fruits = ['apple', 'banana', 'mango']
fruits.insert(2, 'orange')
print(fruits) # ['apple', 'banana', 'orange', 'mango']
fruits.insert(3, 'lemon')
print(fruits) # ['apple', 'banana', 'orange', 'mango', 'lemon']

# remove
fruits = ['apple', 'banana', 'orange', 'mango', 'lemon']
fruits.remove('banana')
print(fruits) # ['apple', 'orange', 'mango', 'lemon']
fruits.remove('lemon')
print(fruits) # ['apple', 'orange', 'mango']

# pop
fruits = ['apple', 'banana', 'orange', 'mango', 'lemon']
fruits.pop()
print(fruits) # ['apple', 'banana', 'orange', 'mango']
fruits.pop(0)
print(fruits) # ['banana', 'orange', 'mango']

# del 
fruits = ['apple', 'banana', 'orange', 'mango', 'lemon']
del fruits[0]
print(fruits) # ['banana', 'orange', 'mango', 'lemon']
del fruits[1]
print(fruits) # ['banana', 'mango', 'lemon']

# clear
fruits = ['apple', 'banana', 'orange', 'mango', 'lemon']
fruits.clear()     
print(fruits) # []

# copying a lists
fruits = ['apple', 'banana', 'orange', 'mango', 'lemon']
fruits_copy = fruits.copy()     
print(fruits_copy) # ['apple', 'banana', 'orange', 'mango', 'lemon']

# join
fruits = ['apple', 'banana', 'orange', 'mango', 'lemon']
vegetables = ['cabbage', 'tomato', 'potato', 'onion', 'carrot']
fruits_vegetables = fruits + vegetables
print(fruits_vegetables) # ['apple', 'banana', 'orange', 'mango', 'lemon', 'cabbage', 'tomato', 'potato', 'onion', 'carrot']

# join with extend
fruits = ['apple', 'banana', 'orange', 'mango', 'lemon']
vegetables = ['cabbage', 'tomato', 'potato', 'onion', 'carrot']
fruits.extend(vegetables)
print(fruits) # ['apple', 'banana', 'orange', 'mango', 'lemon', 'cabbage', 'tomato', 'potato', 'onion', 'carrot']

# count
fruits = ['apple', 'banana', 'orange', 'mango', 'lemon', 'lemon']
print(fruits.count('banana')) # 1
print(fruits.count('lemon')) # 2

# index
fruits = ['apple', 'banana', 'orange', 'mango', 'lemon']
print(fruits.index('orange')) # 2

# Reverse
fruits = ['apple', 'banana', 'orange', 'mango', 'lemon']
fruits.reverse()
print(fruits) # ['lemon', 'mango', 'orange', 'banana', 'apple']

# sort
fruits = ['apple', 'banana', 'orange', 'mango', 'lemon']
fruits.sort()
print(fruits) # ['apple', 'banana', 'lemon', 'mango', 'orange']
fruits.sort(reverse=True)
print(fruits) # ['orange', 'mango', 'lemon', 'banana', 'apple']
 
