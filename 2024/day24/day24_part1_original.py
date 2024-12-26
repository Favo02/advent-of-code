from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

res = 0

wires, gates = fin.read().split("\n\n")
wires = wires.strip().split("\n")
gates = gates.strip().split("\n")

vals = {}
for w in wires:
  label, val = w.split(": ")
  vals[label] = val

queue = deque()
for g in gates:
  input, output = g.split(" -> ")
  i1, op, i2 = input.split()
  queue.append((i1, op, i2, output))

OP = {
  "AND": lambda a, b: "1" if (a == b == "1") else "0",
  "OR": lambda a, b: "1" if (a == "1" or b == "1") else "0",
  "XOR": lambda a, b: "1" if (a != b) else "0",
}

print(queue)
while queue:
  i1, op, i2, out = queue.popleft()
  if not (i1 in vals and i2 in vals):
    queue.append((i1, op, i2, out))
    continue

  vals[out] = OP[op](vals[i1], vals[i2])

print(vals)

ress = []
for k,v in vals.items():
  if k[0] == "z":
    ress.append((k, v))
    print(k, v)

ress.sort(reverse=True)
res = int("".join(r[1] for r in ress), 2)

print("RES:", res)
