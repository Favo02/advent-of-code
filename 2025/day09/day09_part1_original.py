import sys

lines = list(map(lambda l: tuple(map(int, l.strip().split(","))), sys.stdin.read().strip().splitlines()))

lines.sort()

res = 0
for a, (xa, ya) in enumerate(lines):
    for b, (xb, yb) in enumerate(lines[a+1:]):
        res = max(res, (abs(xb-xa)+1)*(abs(yb-ya)+1))

print(res)
