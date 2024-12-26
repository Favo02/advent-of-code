# https://adventofcode.com/2024/day/24
# https://github.com/Favo02/advent-of-code

from collections import deque
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

OP = {
  "AND": lambda a, b: "1" if (a == b == "1") else "0",
  "OR": lambda a, b: "1" if (a == "1" or b == "1") else "0",
  "XOR": lambda a, b: "1" if (a != b) else "0",
}

def todec(vals, who):
  elems = [(k,v) for k,v in vals.items() if k[0] == who]
  elems.sort(reverse=True)
  return int("".join(e[1] for e in elems), 2)

def calc(vals, gates):
  queue = deque(gates)
  while queue:
    i1, op, i2, out = queue.popleft()
    if not (i1 in vals and i2 in vals):
      queue.append((i1, op, i2, out))
    else:
      vals[out] = OP[op](vals[i1], vals[i2])
  return todec(vals, "z")

def find(op, elem):
  for g1, opp, g2, out in gates:
    if op == opp and (g1 == elem or g2 == elem):
      return out

wires, raw_gates = fin.read().split("\n\n")
wires = wires.strip().split("\n")
raw_gates = raw_gates.strip().split("\n")

vals = {}
for w in wires:
  label, val = w.split(": ")
  vals[label] = val

gates = []
for rg in raw_gates:
  input, output = rg.split(" -> ")
  i1, op, i2 = input.split()
  gates.append([i1, op, i2, output])

part1 = calc(vals, gates)

# --- full adder ---
# (xi XOR yi) = XOR1
# (xi AND yi) = AND1
# (xor1 XOR cin) = XOR2
# (xor1 AND cin) = AND2
# (and1 OR and2) = Cout

swapped = []

for i in range(46):
  whoi = f"{i:02}"

  xor1 = find("XOR", "x"+whoi)
  assert xor1 == find("XOR", "y"+whoi)

  and1 = find("AND", "x"+whoi)
  assert and1 == find("AND", "y"+whoi)

  xor2 = find("XOR", xor1)
  and2 = find("XOR", xor1)
  assert xor2 == and2

  cout = find("OR", and1)

  if i == 0 or i == 45: continue

  if xor2 is None:
    swapped.append(xor1)
    swapped.append(and1)
  elif xor2 != "z"+whoi:
    swapped.append(xor2)
    swapped.append("z"+whoi)

part2 = ",".join(sorted(swapped))

print("Part 1:", part1)
print("Part 2:", part2)
