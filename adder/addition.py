#!/usr/bin/python

import random
import sys

count = 0
while count < 1000:
    x=random.randint(1,100)
    j=random.randint(1,200)
    op = ["+", "-", "*"]
    op_num = random.randint(1,3)
    data = ""
    solve = 0

    solve = x + j
    data = ("what is %d + %d\n" % (x , j))

    r = input(data)
    if int(r) == solve:
        count = count + 1
    else:
        print("Wrong... try again\r\n")

if count == 1000:
    print('flag{U7bA-uyNUYw-l1WTwOLWXF4HlyQWke5M}\r\n')
