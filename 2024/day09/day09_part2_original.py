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
  res.append((file, int(a), id if file else -1))
  if file: id += 1
  file = not file


def pprint():
  for r in res:
    print(r)
  print("-" * 10)

tried = set()
while True:
  tomove = None
  r = None
  for rr, tm in enumerate(res[::-1]):
    if tm in tried: continue
    if tm[0]:
      tomove = tm
      r = len(res) -1 - rr
      tried.add(tm)
      break

  if not tomove: break

  # print("MOVE", tomove, r)

  free = None
  for l, fr in enumerate(res):
    if l >= r: break
    if fr[0]: continue
    if fr[1] == tomove[1]:
      res[l], res[r] = res[r], res[l]
      break
    elif fr[1] > tomove[1]:
      lwhat, lqty, lval = res[l]
      rwhat, rqty, rval = res[r]
      res[l] = (rwhat, rqty, rval)
      res[r] = (False, rqty, -1)
      res.insert(l+1, (lwhat, lqty - rqty, lval))
      break

pprint()
p2 = 0
id = 0
for what, qty, val in res:
  for _ in range(qty):
    if what:
      p2 += id * val
    id += 1


print(p2)
