#!/usr/bin/python

import random
import threading
import socket
import sys


alph = { "0": "zero", "1": "one", "2": "two", "3": "three", "4": "four", "5": "five", "6": "six", "7": "seven", "8": "eight", "9": "nine" , ".": "dot"}

count = 0

while count < 1000:
    x=random.random() * random.randint(1,1000)
    j=random.random() * random.randint(1,1000)
    x_str = ""
    j_str = ""

    for c in map(str, str(x)):
        x_str += alph[str(c)] + " "
    for c in map(str, str(j)):
        j_str += alph[str(c)] + " "

    op = ["+", "-", "*", "/", "<<", ">>"]
    op_num = random.randint(1,6)
    data = ""
    solve = 0
    op_num = 5

    if op_num == 1:
        solve = round(float(x + j), 5)
        data = ("what is %s + %s\n" % (x_str , j_str))
    elif op_num == 2:
        solve = round(float(x - j), 2)
        data = ("what is %s - %s\n" % (x_str , j_str))
    elif op_num == 3:
        solve = round(float(x * j), 1)
        data = ("what is %s * %s\n" % (x_str , j_str))
    elif op_num == 4:
        solve = round(float(x / j), 3)
        data = ("what is %s / %s\n" % (x_str, j_str))
    elif op_num == 5:
        solve = int(x) << int(j)
        data = ("what is %s << %s ? (You have to round these to ints)\n" % (x_str , j_str))
    else:
        solve = int(x) >> int(j)
        data = ("what is %s >> %s ? (You have to round these to ints)\n" % (x_str , j_str))

    try:
        r = input(data)
    except Exception as e:
        print("not sure what you are doing... but I can't handle that.\r\n")
        continue

    try:
        val = float(r)
    except ValueError:
        print("SEND ME A NUMBER!")
        continue

    if op_num in [1, 2, 3, 4]:
        if float(r) == solve:
            count = count + 1
        else:
            print("Wrong... try again", solve)
    else:
        if int(r) == solve:
            count = count + 1
        else:
            print("Wrong... try again", solve)

if count == 1000:
    print('flag{IEN3ZYcBSbAIYvp6TW8Jp3Sanv4O8tvp}')
