from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
from operator import itemgetter
import sys
import re
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

res = 0

matchessss = []
field = []

for line in sys.stdin:
  field.append(line.strip())


for r, line in enumerate(field):
  pattern = re.compile(r'mul\(\d+,\d+\)')
  for match in pattern.finditer(line.strip()):
    start_index = match.start()
    matchessss.append(("MUL", (r, start_index), match))

for r, line in enumerate(field):
  pattern = re.compile(r'do\(\)')
  for match in pattern.finditer(line.strip()):
    print("DO")
    start_index = match.start()
    matchessss.append(("DO", (r, start_index), match))

for r, line in enumerate(field):
  pattern = re.compile(r'don\'t\(\)')
  for match in pattern.finditer(line.strip()):
    start_index = match.start()
    matchessss.append(("DONT", (r, start_index), match))

matchessss.sort(key=itemgetter(1))
valid = True
for (what, aa, match) in matchessss:
  print(what, aa, match, valid)
  if what == "MUL" and valid:
    nums = re.findall(r'\d+', match.group())
    if nums:
      num1, num2 = map(int, nums)
      if not (1 <= len(str(num1)) <= 3): continue
      if not (1 <= len(str(num2)) <= 3): continue
      res += num1 * num2
  elif what == "DO":
    valid = True
  elif what == "DONT":
    valid = False

print("RES:", res)
