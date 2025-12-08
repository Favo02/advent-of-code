# https://adventofcode.com/2025/day/8
# https://github.com/Favo02/advent-of-code

import sys
from math import sqrt
from functools import reduce
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = part2 = 0

nums = list(map(lambda l: tuple(map(int, l.strip().split(","))), sys.stdin.read().strip().splitlines()))

P = [i for i in range(len(nums))]
S = [1 for _ in range(len(nums))]

def find(n):
    if P[n] == n:
        return n
    P[n] = find(P[n])
    return P[n]

def union(a, b):
    ap = find(a)
    bp = find(b)
    if ap == bp:
        return
    if S[ap] < S[bp]:
        ap, bp = bp, ap
    S[ap] += S[bp]
    S[bp] = 0
    P[bp] = ap

dists = []
for ai, (ax, ay, az) in enumerate(nums):
    for bi, (bx, by, bz) in enumerate(nums[ai+1:]):
        bi += ai+1
        dist = sqrt((ax-bx)**2 + (ay-by)**2 + (az-bz)**2)
        dists.append((dist, ai, bi))

dists.sort()

i = 0
while len([s for s in S if s == 0]) < len(nums)-1:
    if i == 1000:
        part1 = reduce(lambda a, b: a*b, sorted(S, reverse=True)[:3], 1)

    union(dists[i][1], dists[i][2])
    i+=1

part2 = nums[dists[i-1][1]][0] * nums[dists[i-1][2]][0]

print("Part 1:", part1)
print("Part 2:", part2)
