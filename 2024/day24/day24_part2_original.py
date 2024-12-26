from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

OP = {
  "AND": lambda a, b: "1" if (a == b == "1") else "0",
  "OR": lambda a, b: "1" if (a == "1" or b == "1") else "0",
  "XOR": lambda a, b: "1" if (a != b) else "0",
}

def todec(vals, who):
  ress = [(k,v) for k,v in vals.items() if k[0] == who]
  ress.sort(reverse=True)
  # print(vals, ress, who)
  return int("".join(r[1] for r in ress), 2)

def calc(vals, gates):
  queue = deque(gates)
  while queue:
    i1, op, i2, out = queue.popleft()
    if not (i1 in vals and i2 in vals):
      queue.append((i1, op, i2, out))
    else:
      vals[out] = OP[op](vals[i1], vals[i2])
  return todec(vals, "z")

def path(who):
  pat = [who]
  i = 0
  while i < len(pat):
    if pat[i] in back:
      i1, op, i2 = back[pat[i]]
      print(pat[i], "=", back[pat[i]])
      pat.append(i1)
      pat.append(i2)
    i += 1
  return pat

res = 0

wires, raw_gates = fin.read().split("\n\n")
wires = wires.strip().split("\n")
raw_gates = raw_gates.strip().split("\n")

vals = {}
for w in wires:
  label, val = w.split(": ")
  vals[label] = val

gates = []
back = {}
for g in raw_gates:
  input, output = g.split(" -> ")
  i1, op, i2 = input.split()
  gates.append([i1, op, i2, output])
  back[output] = (i1, op, i2)

X = todec(vals, "x")
Y = todec(vals, "y")

def check(a,b, c,d, e,f, g,h):
  gates[a][3], gates[b][3] = gates[b][3], gates[a][3]
  gates[c][3], gates[d][3] = gates[d][3], gates[c][3]
  gates[e][3], gates[f][3] = gates[f][3], gates[e][3]
  gates[g][3], gates[h][3] = gates[h][3], gates[g][3]
  Z = calc(vals.copy(), gates)
  gates[a][3], gates[b][3] = gates[b][3], gates[a][3]
  gates[c][3], gates[d][3] = gates[d][3], gates[c][3]
  gates[e][3], gates[f][3] = gates[f][3], gates[e][3]
  gates[g][3], gates[h][3] = gates[h][3], gates[g][3]
  return Z

print(X, Y)

# for i in range(46):
#   print("---")
#   if i < 10:
#     print(path(f"z0{i}"))
#   else:
#     print(path(f"z{i}"))


# CARRY = {}
# for g1, op, g2, out in gates:
#   if op == "AND" and g1[0] in ["x", "y"] and g2[0] in ["x", "y"]:
#     CARRY[g1[1:]] = out
#     assert g1[1:] == g2[1:]

# full adder:
# (xi XOR yi) = XOR1
# (xi AND yi) = AND1
# (xor1 XOR cin) = XOR2
# (xor1 AND cin) = AND2
# (and1 OR and2) = Cout

def find(op, left):
  for g1, opp, g2, out in gates:
    if op == opp and (g1 == left or g2 == left):
      return out

print("-"*10)

for i in range(46):
  whoi = f"{i:02}"

  xor1 = find("XOR", "x"+whoi)
  assert xor1 == find("XOR", "y"+whoi)

  and1 = find("AND", "x"+whoi)
  assert and1 == find("AND", "y"+whoi)

  xor2 = find("XOR", xor1)
  and2 = find("XOR", xor1)

  cout = find("OR", and1)

  print(f"{i}: {xor1=} {and1=} {xor2=} {and2=} {cout=}")

res = ["..."] # hand picked :)
res.sort()
print(",".join(res))
