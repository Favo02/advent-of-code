from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
import re
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

res = 0

for line in sys.stdin:

  pattern = re.compile(r'mul\(\d+,\d+\)')
  matches = pattern.findall(line.strip())

  for match in matches:
    nums = re.findall(r'\d+', match)
    print(match)
    if nums:
      num1, num2 = map(int, nums)
      if not (1 <= len(str(num1)) <= 3): continue
      if not (1 <= len(str(num2)) <= 3): continue
      res += num1 * num2

print("RES:", res)
