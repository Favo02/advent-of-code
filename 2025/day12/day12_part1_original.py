import sys
from copy import deepcopy

*shapes, areas = sys.stdin.read().strip().split("\n\n")
shapes = list(map(lambda s: s.split("\n")[1:], shapes))
sa = [len([True for line in shape for c in line if c == "#"]) for shape in shapes]

def pprint(grid):
    for l in grid:
        print("".join(l))
    print()

def place(grid, x, y, shape):
    grid = deepcopy(grid)
    for sy in range(len(shape)):
        for sx in range(len(shape[0])):
            if shape[sy][sx] == ".": continue
            if not (0 <= y+sy < len(grid)): return False
            if not (0 <= x+sx < len(grid[0])): return False
            if grid[y+sy][x+sx] == "#": return False
            grid[y+sy][x+sx] = "#"
    return grid

def solve(grid, rem):
    for i, r in enumerate(rem):
        if r != 0: break
    else:
        return True

    # print(sum(rem))
    # pprint(grid)

    ROWS, COLS = len(grid), len(grid[0])
    shape = shapes[i]
    for rot in range(4):
        for y in range(ROWS):
            for x in range(COLS):
                res = place(grid, x, y, shape)
                if res is False: continue
                rem[i] -= 1
                if solve(res, rem): return True
                rem[i] += 1

        shape = [list(reversed(col)) for col in zip(*shape)]


# print(shapes)

res = 0
for a in areas.split("\n"):
    size, qties = a.split(": ")
    x, y = tuple(map(int, size.split("x")))
    qties = list(map(int, qties.split()))

    if sum([q * a for q, a in zip(qties, sa)]) > x * y:
        continue

    res += 1
    # print(size, qties)

    # if solve([["." for _ in range(x)] for _ in range(y)], qties):
    #     res += 1
    #     print("VALID")
    # else:
    #     print("INVALID")

print(res)
