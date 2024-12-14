from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def move():
  global robots
  newrobots = set()
  for x, y, vx, vy in robots:
    newrobots.add(((x + vx) % COLS, (y + vy) % ROWS, vx, vy))
  robots = newrobots

def pprint():
  rrrobots = Counter(map(lambda a: a[:2], robots))
  for y in range(ROWS):
    for x in range(COLS):
      if (x,y) in rrrobots:
        print(rrrobots[(x,y)], end="")
      else:
        print(".", end="")
    print()
  print()

def check():
  rrrobots = Counter(map(lambda a: a[:2], robots))
  segmetns = 0
  for y in range(ROWS):
    cons = 0
    last = False
    for x in range(COLS):
      if (x,y) in rrrobots:
        if last: cons+=1
        last = True
      else:
        if cons >= 5:
          segmetns += 1
        cons = 0
        last = False
  return segmetns

res = 0

robots = set()

for y, line in enumerate(fin):
  pos, vel = line.strip().split()
  x, y = map(int, pos[2:].split(","))
  vx, vy = map(int, vel[2:].split(","))
  robots.add((x, y, vx, vy))

print(robots)

ROWS = 103
COLS = 101

# ROWS = 7
# COLS = 11

for time in range(10000):
  print(time+1)
  move()
  # pprint()
  if check() > 5:
    pprint()
    break


# print(robots)

res = [0] * 4

for x, y, _, _ in robots:
  if x == (COLS // 2): continue
  if y == (ROWS // 2): continue

  if x < COLS // 2:
    if y < ROWS // 2:
      res[0] += 1
    else:
      res[2] += 1
  else:
    if y < ROWS // 2:
      res[1] += 1
    else:
      res[3] += 1

p1 = 1
for r in res:
  p1 *= r

print("RES:", p1)
