# https://adventofcode.com/2025/day/3
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = part2 = 0

def maxrange(l, start, end):
    res = 0, 0
    for i, n in enumerate(l[start:end]):
        res = max(res, (n, -i))
    return (res[0], -res[1] + start + 1)

def maxn(l, qty):
    cur = []
    start = 0
    for _ in range(qty):
        m, start = maxrange(l, start, len(l)-qty+len(cur)+1)
        cur.append(m)
    return cur

for l in (list(map(int, list(l))) for l in map(str.strip, fin.readlines())):
    part1 += int("".join(map(str, maxn(l, 2))))
    part2 += int("".join(map(str, maxn(l, 12))))

print("Part 1:", part1)
print("Part 2:", part2)
