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

    if count < 500:
        if op_num == 1:
            solve = x + j
            data = ("what is %d + %d\n" % (x , j))
        elif op_num == 2:
            solve = x - j
            data = ("what is %d - %d\n" % (x , j))
        else:
            solve = x * j
            data = ("what is %d * %d\n" % (x , j))

    else:
        if op_num == 1:
            solve = x + j
            data = ("what is %s + %s\n" % (hex(x) , hex(j)))
        elif op_num == 2:
            solve = x - j
            data = ("what is %s - %s\n" % (hex(x) , hex(j)))
        else:
            solve = x * j
            data = ("what is %s * %s\n" % (hex(x) , hex(j)))
    try:
        r = input(data)
    except Exception as e:
        print("Not sure what you are doing, but I cannot take that...")
        continue

    try:
        val = int(r)
    except ValueError:
        print("That's not an int!")
        continue

    if int(r) == solve:
        count = count + 1
    else:
        print("Wrong... try again\r\n")

if count == 1000:
    print('flag{p9ovNL0HeWqbmDQk_JAaNgbp_QCYMhu0}\r\n')
