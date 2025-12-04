import sys

lines = list(map(str.strip, sys.stdin.readlines()))

ROWS = len(lines)
COLS = len(lines[0])

res = 0

for y in range(ROWS):
    for x in range(COLS):

        if lines[y][x] != "@": continue
        cur = 0

        for dx in range(-1, 2):
            for dy in range(-1, 2):
                if dx == 0 and dy == 0: continue
                if not (0 <= x + dx < COLS): continue
                if not (0 <= y + dy < ROWS): continue
                if lines[y + dy][x + dx] == "@":
                    cur += 1

        if cur < 4:
            # print(y, x)
            res += 1

print(res)
