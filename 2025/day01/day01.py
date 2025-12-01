# https://adventofcode.com/2025/day/1
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = part2 = 0
cur = 50

for line in map(str.strip, fin):
    dir, qty = line[0], int(line[1:])

    prev = cur

    if dir == "L":
        cur -= qty
        part2 += abs(cur // 100) + (cur % 100 == 0) - (prev == 0)

    if dir == "R":
        cur += qty
        part2 += cur // 100

    cur %= 100

    if cur == 0:
        part1 += 1

print("Part 1:", part1)
print("Part 2:", part2)
