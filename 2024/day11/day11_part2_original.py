from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

res = 0

nums = [int(n) for n in input().split()]

nums = Counter(nums)

for i in range(75):
  newnums = defaultdict(int)
  for n, qty in nums.items():
    if n == 0:
      newnums[1] += qty
    elif len(str(n)) % 2 == 0:
      L = len(str(n)) // 2
      newnums[int(str(n)[:L])] += qty
      newnums[int(str(n)[L:])] += qty
    else:
      newnums[n * 2024] += qty
  nums = newnums

print(sum(nums.values()))
