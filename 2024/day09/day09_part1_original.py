from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

field = []
for y, line in enumerate(fin):
  line = line.rstrip()
  field.append(line)

  for x, cell in enumerate(line):
    pass

all = "".join(field)

res = []
file = True
id = 0

for i, a in enumerate(all):
  for _ in range(int(a)):
    res.append((file, id))
  if file: id += 1
  file = not file

l = 0
r = len(res)-1

while l < r:
  while l < r and res[l][0] == True:
    l += 1
  while l < r and res[r][0] == False:
    r -= 1
  res[l], res[r] = res[r], res[l]

p1 = 0
for i, (what, who) in enumerate(res):
  if what:
    p1 += who * i

print(p1)
