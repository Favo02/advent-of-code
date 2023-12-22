# https://adventofcode.com/2023/day/22
# https://github.com/Favo02/advent-of-code

import sys
from collections import deque
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def overlap(a, b):
  return not (a[0] > b[1] or a[1] < b[0])

def fall(bricks):
  bricks.sort(key=lambda b: b[Z][START])
  fallen = []
  for b in bricks:
    z = 1
    height = b[Z][END] - b[Z][START]
    for under in fallen:
      if overlap(b[X], under[X]) and overlap(b[Y], under[Y]):
        z = max(z, under[Z][END]+1)

    assert b[Z][START] >= z
    assert b[Z][END] >= z + height
    assert z <= z+height

    fallen.append((b[X], b[Y], (z, z+height)))
  return fallen

def get_falls(start, supp, base):
  will_fall = set([ base ])
  queue = deque(start)
  while queue:
    cur = queue.popleft()
    if len(supp[cur][BELOW] - will_fall) == 0:
      will_fall.add(cur)
      queue += supp[cur][ABOVE]
  return will_fall

part1 = 0
part2 = 0

bricks = []
for line in fin:
  above, e = line.rstrip().split("~")
  xs, ys, zs = [int(n) for n in above.split(",")]
  xe, ye, ze = [int(n) for n in e.split(",")]

  assert xs <= xe
  assert ys <= ye
  assert zs <= ze
  assert (1 if xs != xe else 0) + (1 if ys != ye else 0) + (1 if zs != ze else 0) <= 1 # rectangle

  bricks.append(((xs,xe), (ys,ye), (zs,ze)))

N = len(bricks)

X, Y, Z = 0, 1, 2 # indexes to access tuples
START, END = 0, 1 # indexes to access tuples
ABOVE, BELOW = 0, 1 # index to access support dictionary tuple

fallen = fall(bricks)

support = {}
for i, b1 in enumerate(fallen):
  above = set()
  below = set()
  for j, b2 in enumerate(fallen):
    if i == j:
      continue
    if b2[Z][START] == (b1[Z][END]+1):
      if overlap(b1[X], b2[X]) and overlap(b1[Y], b2[Y]):
        above.add(j)
    if b2[Z][END] == (b1[Z][START]-1):
      if overlap(b1[X], b2[X]) and overlap(b1[Y], b2[Y]):
        below.add(j)
  support[i] = (above, below)

unremovable = set()
for k, (above, below) in support.items():
  if len(below) == 1:
    unremovable |= below
  part2 += len(get_falls(support[k][ABOVE], support, k))-1
part1 = N - len(unremovable)

print("Part 1:", part1)
print("Part 2:", part2)
