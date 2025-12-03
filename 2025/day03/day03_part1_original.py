import sys

lines = sys.stdin.readlines()

res = 0
for l in lines:
    cur = 0
    l = list(map(int, list(l.strip())))
    a = 0
    b = 0
    for i, n in enumerate(l):
        if n > a and i != len(l)-1:
            a = n
            b = l[i+1]
        elif n > b:
            b = n
    print(a, b)
    res += a*10 + b

print(res)
