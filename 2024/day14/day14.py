# https://adventofcode.com/2024/day/14
# https://github.com/Favo02/advent-of-code

from collections import Counter
from functools import reduce
import re
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def quadrants(robots):
  count = [0] * 4
  for x, y, _, _ in robots:
    if x == COLS // 2: continue
    if y == ROWS // 2: continue
    match (x < COLS // 2, y < ROWS // 2):
      case True, True: count[0] += 1
      case False, True: count[1] += 1
      case True, False: count[2] += 1
      case False, False: count[3] += 1
  return count

def check_tree(robots):
  # counts the number of horizontal segments long at least 5 robots
  # if the segments are more than 5 then there is a tree

  robots_set = set(map(lambda a: a[:2], robots))
  segments = 0

  for y in range(ROWS):
    consecutive = 0
    last = False

    for x in range(COLS):

      if (x,y) in robots_set:
        if last: consecutive += 1
        last = True

      else:
        if consecutive >= 5:
          segments += 1

        consecutive = 0
        last = False

  return segments >= 5

def pretty_print():
  robots_count = Counter(map(lambda a: a[:2], robots))
  for y in range(ROWS):
    for x in range(COLS):
      if (x,y) in robots_count:
        print(robots_count[(x,y)], end="")
      else:
        print(".", end="")
    print()
  print()

robots = []
for y, line in enumerate(fin):
  pos, vel = line.strip().split()
  x, y = map(int, re.findall(r"-?\d+", pos))
  vx, vy = map(int, re.findall(r"-?\d+", vel))
  robots.append((x, y, vx, vy))

ROWS, COLS = 103, 101
# ROWS, COLS = 7, 11 # example

part1 = part2 = 0

time = 0
while not (part1 and part2) and not (part1 and ROWS == 7):
  time += 1
  robots = [((x + vx) % COLS, (y + vy) % ROWS, vx, vy) for x, y, vx, vy in robots]
  if time == 100:
    part1 = reduce(lambda a, b: a * b, quadrants(robots), 1)
  if check_tree(robots):
    part2 = time

print("Part 1:", part1)
print("Part 2:", part2)
