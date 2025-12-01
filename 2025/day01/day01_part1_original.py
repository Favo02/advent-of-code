import sys

lines = sys.stdin.readlines()

cur = 50
res = 0
for l in lines:
    dir, qty = l[0], int(l[1:])
    if dir == "L":
        cur -= qty
    else:
        cur += qty
    cur = cur % 100
    if cur == 0:
        res += 1
print(res)
