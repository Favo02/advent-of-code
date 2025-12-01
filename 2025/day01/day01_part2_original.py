import sys

lines = sys.stdin.readlines()

cur = 50
res = 0

for l in lines:
    dir, qty = l[0], int(l[1:])
    prev = cur

    for _ in range(qty):
        if dir == "L":
            cur -= 1
        else:
            cur += 1

        cur = cur % 100
        if cur == 0:
            res += 1

print(res)
