# Single line comment
text = 'Hello, World!'
print(text) # Hello, World!
print(len(text)) # 13

# Multiline string
multiline_text = '''Hello,
World!'''
print(multiline_text)

# String concatenation
first_text = 'Hello,'
last_text = 'World!'
space = ' '
full_text = first_text + space + last_text
print(full_text) # Hello, World!

# Unpacking characters
text = 'Hello'
a, b, c, d, e = text
print(a) # H
print(b) # e
print(c) # l
print(d) # l
print(e) # o

# Accessing characters in strings by index
text = 'Hello'
first_letter = text[0]
second_letter = text[1]
last_index = len(text) - 1
last_letter = text[last_index]
print(first_letter) # H
print(second_letter) # e
print(last_letter)  # o

# If we want to start from right end we can use negative indexing. -1 is the last index
text = 'Hello'
last_letter = text[-1]
second_last = text[-2]
print(last_letter)  # o
print(second_last)  # l

# Slicing
text = 'Hello'
last_three = text[2:5]
print(last_three)  # llo
last_three = text[-3:]
print(last_three)  # llo
last_three = text[2:]
print(last_three)  # llo

# Skipping character
text = 'Hello, World!'
skip_text = text[0:13:2]
print(skip_text) # Hlo ol!

# Escape sequence
print('Hello, \nWorld!')
print('Hello, W\tor\tld!')
print('Hello, World!\\')
print('\'Hello\', World!')

# String methods
# capitalize(): Converts the first character the string to Capital Letter
text = 'Hello, World!'
print(text.capitalize()) # Hello, world!

# count(): returns occurrences of substring in string, count(substring, start=.., end=..)
text = 'Hello, World!'
print(text.count('o')) # 2
print(text.count('o', 7, 13)) # 1
print(text.count('llo')) # 1

# endswith(): Checks if a string ends with a specified ending
text = 'Hello, World!'
print('rld!:', text.endswith('rld!')) # True
print('rld:', text.endswith('rld')) # False

# find(): Returns the index of first occurrence of substring
text = 'Hello, World!'
print('l:', text.find('l')) # 2
print('Hel:', text.find('Hel')) # 0
print('Wl:', text.find('Wl')) # -1

# format()	formats string into nicer output    
first_text = 'He'
second_text = 'Wor'
last_text = '!'
full_text = '{}llo, {}ld{}'.format(first_text, second_text, last_text)
print('{}llo, {}ld{}:', full_text) # Hello, World!

# isalnum(): Checks alphanumeric character
text = 'HelloWorld'
print('HelloWorld:', text.isalnum()) # True
text = 'Hello, World!'
print('Hello, World!:', text.isalnum()) # False

# isdigit(): Checks Digit Characters
text = 'age25'
print('age25:', text.isdigit()) # False

# isdecimal():Checks decimal characters
text = '10'
print('10:', text.isdecimal()) # True
text = '10.5'
print('10.5', text.isdecimal()) # False

# islower():Checks if all alphabets in a string are lowercase
text = 'hello, world!'
print('hello, world!:', text.islower()) # True
text = 'Hello, World!'
print('Hello, World!:', text.islower()) # False

# isupper(): returns if all characters are uppercase characters
text = 'HELLO, WORLD!'
print('HELLO, WORLD!:', text.isupper()) # True
text = 'Hello, World!'
print('Hello, World!:', text.isupper()) # False

# isnumeric():Checks numeric characters
text = '2'
print('2:', text.isnumeric()) # True
text = 'two'
print('two:', text.isnumeric()) # False

# join(): Returns a concatenated string
web_tech = ['He', 'llo', 'Wor', 'ld']
result = ', '.join(web_tech)
print(result)  # He, llo, Wor, ld

# replace(): Replaces substring inside
text = 'Hello, World!'
print('Hello, World!:', text.replace('Hello', 'Bye'))  # Bye, World!

# split():Splits String from Left
text = 'Hello, World!'
print('Hello, World!:', text.split())  # ['Hello,', 'World!']

# swapcase(): Checks if String Starts with the Specified String
text = 'Hello, World!'
print(text.swapcase())  # hELLO, wORLD!

# startswith(): Checks if String Starts with the Specified String
text = 'Hello, World!'
print('Hello, World!:', text.startswith('Hello'))  # True
text = 'Hello, World!'
print('Hello, World!:', text.startswith('Bye'))  # False
