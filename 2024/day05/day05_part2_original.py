from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

res = 0

ordering = defaultdict(set)

for y, line in enumerate(fin):
  if line == "\n": break
  a, b = map(int, line.split("|"))
  ordering[b].add(a)

updates = []
for y, line in enumerate(fin):
  updates.append(list(map(int, line.split(","))))

def isvalid(update):
  banned = set()

  for elem in update:
    if elem in banned:
      return False
    banned.update(ordering[elem])

  return True

incorrect = []
for update in updates:
  if not isvalid(update):
    incorrect.append(update)

def fix(update):
  seen = set()
  for e in update:
    if e in seen:
      continue
    seen.add(e)
    yield e

res = []
for update in incorrect:
  # print(update)
  queue = deque()
  sett = set(update)
  final = []

  for elem in update:
    queue.append(elem)

  while queue:
    cur = queue.popleft()
    # print(queue)
    valid = True
    for prec in ordering[cur]:
      if prec in sett and prec not in final:
        queue.append(prec)
        valid = False

    if valid:
      final.append(cur)
    else:
      queue.append(cur)

  res.append(list(fix(final)))
  print("__________", final)

ans = 0
for r in res:
  ans += r[len(r) // 2]

print(ans)
