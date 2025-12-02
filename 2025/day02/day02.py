# https://adventofcode.com/2025/day/2
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = part2 = 0

def invalid(q):
    p1 = q[:len(q)//2] == q[len(q)//2:]
    for size in range(1, len(q)):
        if len(q) % size != 0: continue
        for starts in range(size, len(q), size):
            if q[starts:starts+size] != q[:size]:
                break
        else:
            return p1, True
    return p1, False

for l, r in (map(int, range) for range in map(lambda r: r.split("-"), fin.read().split(","))):
    for n in range(l, r+1):
        p1, p2 = invalid(str(n))
        if p1: part1 += n
        if p2: part2 += n

print("Part 1:", part1)
print("Part 2:", part2)
