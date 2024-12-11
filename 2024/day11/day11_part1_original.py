from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

res = 0

nums = [int(n) for n in input().split()]
print(nums)

for i in range(75):
  print(i, len(nums))
  newnums = []
  for n in nums:
    if n == 0:
      newnums.append(1)
    elif len(str(n)) % 2 == 0:
      L = len(str(n)) // 2
      newnums.append(int(str(n)[:L]))
      newnums.append(int(str(n)[L:]))
    else:
      newnums.append(n * 2024)
  nums = newnums
  # print(nums)

print(len(nums))
