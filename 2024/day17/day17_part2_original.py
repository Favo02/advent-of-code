from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
import re
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def combo(n):
  if n <= 3: return n
  if n == 4: return REG["a"]
  if n == 5: return REG["b"]
  if n == 6: return REG["c"]
  assert False, "invalid combo"

registers, instructions = fin.read().split("\n\n")
a = int(re.findall(r"A:\s(\d+)", registers)[0])
b = int(re.findall(r"B:\s(\d+)", registers)[0])
c = int(re.findall(r"C:\s(\d+)", registers)[0])
instructions = list(map(int, instructions.split(": ")[1].split(",")))

REG = {"a": a, "b": b, "c": c}

def sim_math(a, expected):
  for i in range(len(expected)):
    if a <= 0:
      return False
    val = ((a % 8) ^ 4 ^ (a // 2**((a % 8) ^ 1))) % 8 # unique to my input!
    # print(val, expected[i])
    if val != expected[i]:
      return False
    a //= 2**3
  return True

expected = tuple(instructions)
print(expected)

# print(simulate(a))
# sim_math(a, expected)


cur = 0
for i in range(1, len(expected)+1):
  cur *= 8
  print(cur, expected[-i:])
  while True:
    if sim_math(cur, expected[-i:]):
      break
    cur += 1

print()
print(sim_math(cur, expected))
