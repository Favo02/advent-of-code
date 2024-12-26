from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def generate(s):

  r1 = s*64
  s ^= r1
  s %= 16777216

  r2 = s//32
  s ^= r2
  s %= 16777216

  r3 = s * 2048
  s ^= r3
  s %= 16777216

  return s, s%10

K = 2000
# K = 10

SEQS = defaultdict(int)

for y, line in enumerate(fin):
  line = line.strip()
  n = int(line)
  price = n % 10

  seen = set()
  seq = ()

  last_price = price
  for i in range(K-1):
    n, price = generate(n)
    diff = price - last_price
    last_price = price

    # print("---", price, diff)

    if len(seq) < 4:
      seq += (diff,)
    else:
      seq = tuple(seq[1:] + (diff,))

    if len(seq) == 4:
      if seq in seen: continue
      seen.add(seq)
      # print(seq, price)
      SEQS[seq] += price

    # print(SEQS)
    # print(ld, diff, seq, ld)

res = max(SEQS.values())
print(res)

# 2096
# 2045
# 2036
