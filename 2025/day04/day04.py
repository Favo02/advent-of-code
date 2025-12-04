# https://adventofcode.com/2025/day/4
# https://github.com/Favo02/advent-of-code

import sys
from collections import deque
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = part2 = 0

grid = list(map(list, map(str.strip, fin.readlines())))

ROWS = len(grid)
COLS = len(grid[0])

def check(x, y, mut):
    if grid[y][x] != "@":
        return False, set()

    adjs = set()
    for dx, dy in ((dx, dy) for dx in range(-1, 2) for dy in range(-1, 2)):
        if dx == 0 and dy == 0: continue
        if not (0 <= x + dx < COLS): continue
        if not (0 <= y + dy < ROWS): continue
        if grid[y + dy][x + dx] == "@":
            adjs.add((x+dx, y+dy))

    if len(adjs) < 4:
        if mut: grid[y][x] = "."
        return True, adjs
    return False, set()

part1 = sum(check(x, y, False)[0] for x in range(COLS) for y in range(COLS))

q = deque()
for y in range(ROWS):
    for x in range(COLS):
        q.append((x, y))
while q:
    b, new = check(*q.popleft(), True)
    part2 += b
    q += list(new)

print("Part 1:", part1)
print("Part 2:", part2)
