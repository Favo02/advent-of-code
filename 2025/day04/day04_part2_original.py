import sys
from collections import deque

lines = list(map(list, map(str.strip, sys.stdin.readlines())))

ROWS = len(lines)
COLS = len(lines[0])

def check(x, y):
    if lines[y][x] != "@":
        return False, set()

    cur = 0
    adjs = set()
    for dx in range(-1, 2):
        for dy in range(-1, 2):
            if dx == 0 and dy == 0: continue
            if not (0 <= x + dx < COLS): continue
            if not (0 <= y + dy < ROWS): continue
            if lines[y + dy][x + dx] == "@":
                adjs.add((x+dx, y+dy))
                cur += 1

    if cur < 4:
        lines[y][x] = "."
        return True, adjs
    return False, set()



q = deque()

for y in range(ROWS):
    for x in range(COLS):
        q.append((x, y))

res = 0
while q:
    b, neww = check(*q.popleft())
    res += b
    for n in neww:
        q.append(n)


print(res)
