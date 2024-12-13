from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
import re
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def addd(p1, p2):
  a1, b1 = p1
  a2, b2 = p2
  return (a1+a2), (b1+b2)

def solve(a, b, target):
  poss = {}
  poss[(0, 0)] = (0, 0)

  for i in range(1, 201):
    newposs = poss.copy()
    for p, (aa, bb) in poss.items():

      newa = addd(p,a)
      if newa not in newposs:
        newposs[newa] = (aa+1, bb)
      else:
        oldcost = newposs[newa][0]*3 + newposs[newa][1]
        newcost = (aa+1)*3 + bb
        if newcost < oldcost:
          newposs[newa] = (aa+1, bb)

      newb = addd(p,b)
      if newb not in newposs:
        newposs[newb] = (aa, bb+1)
      else:
        oldcost = newposs[newb][0]*3 + newposs[newb][1]
        newcost = aa*3 + (bb+1)
        if newcost < oldcost:
          newposs[newb] = (aa, bb+1)

    poss = newposs
    if target in poss:
      return poss[target]

  return poss.get(target, (-1,-1))


res = 0
macs = list(fin.read().strip().split("\n\n"))

for i, mac in enumerate(macs):
  a,b,target = mac.split("\n")

  a = tuple(map(int, re.findall(r'\d+', a)))
  b = tuple(map(int, re.findall(r'\d+', b)))
  target = tuple(map(int, re.findall(r'\d+', target)))
  # print(target)

  aa, bb = solve(a,b,(target))
  print(f"--- {i}/{len(macs)} ---")

  if aa != -1:
    print(f"{aa=} {bb=}")
    res += aa*3 + bb


print("RES:", res)
