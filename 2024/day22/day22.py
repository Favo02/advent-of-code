# https://adventofcode.com/2024/day/22
# https://github.com/Favo02/advent-of-code

import sys
from collections import defaultdict
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

  return s

SEQS = defaultdict(int)

def buyer(start):
  secret = start
  price = secret % 10

  seq = ()
  seen_seqs = set()

  for _ in range(ITERS-1):
    secret = generate(secret)
    price, diff = secret % 10, (secret % 10) - price

    if len(seq) < 4:
      seq += (diff,)
    else:
      seq = tuple(seq[1:] + (diff,))

    if len(seq) == 4 and seq not in seen_seqs:
      seen_seqs.add(seq)
      SEQS[seq] += price

  return generate(secret)

ITERS = 2000

part1 = 0
for line in fin:
  part1 += buyer(int(line.strip()))

part2 = max(SEQS.values())

print("Part 1:", part1)
print("Part 2:", part2)
