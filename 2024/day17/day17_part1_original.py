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

print(instructions)

REG = {"a": a, "b": b, "c": c}

IP = 0

res = []

while IP < len(instructions):
  match instructions[IP]:
    case 0:
      den = 2**combo(instructions[IP+1])
      REG["a"] = REG["a"] // den
      IP += 2

    case 1:
      REG["b"] = REG["b"] ^ instructions[IP+1]
      IP += 2

    case 2:
      REG["b"] = combo(instructions[IP+1]) % 8
      IP += 2

    case 3:
      if REG["a"] == 0:
        IP += 2
        continue
      IP = instructions[IP+1]

    case 4:
      REG["b"] = REG["b"] ^ REG["c"]
      IP += 2

    case 5:
      res.append(combo(instructions[IP+1]) % 8)
      print(",".join(map(str, res)))
      IP += 2

    case 6:
      den = 2**combo(instructions[IP+1])
      REG["b"] = REG["a"] // den
      IP += 2

    case 7:
      den = 2**combo(instructions[IP+1])
      REG["c"] = REG["a"] // den
      IP += 2
