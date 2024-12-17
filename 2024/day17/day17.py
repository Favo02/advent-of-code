# https://adventofcode.com/2024/day/17
# https://github.com/Favo02/advent-of-code

import re
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def combo(registers, n):
  if n <= 3: return n
  if n == 4: return registers["a"]
  if n == 5: return registers["b"]
  if n == 6: return registers["c"]
  assert False, "invalid combo"

def simulate(a, b, c, instructions):
  registers = {"a": a, "b": b, "c": c}
  ip = 0

  res = []

  while ip < len(instructions):
    match instructions[ip]:
      case 0:
        den = 2**combo(registers, instructions[ip+1])
        registers["a"] = registers["a"] // den

      case 1:
        registers["b"] = registers["b"] ^ instructions[ip+1]

      case 2:
        registers["b"] = combo(registers, instructions[ip+1]) % 8

      case 3:
        if registers["a"] != 0:
          ip = instructions[ip+1]
          ip -= 2

      case 4:
        registers["b"] = registers["b"] ^ registers["c"]

      case 5:
        res.append(combo(registers, instructions[ip+1]) % 8)

      case 6:
        den = 2**combo(registers, instructions[ip+1])
        registers["b"] = registers["a"] // den

      case 7:
        den = 2**combo(registers, instructions[ip+1])
        registers["c"] = registers["a"] // den

    ip += 2

  return res

def reverse_simulate(a, expected):
  for ex in expected:
    # formula unique to my input!
    val = ((a % 8) ^ 4 ^ (a // 2**((a % 8) ^ 1))) % 8
    a //= 2**3
    if a < 0 or val != ex: return False
  return True

registers, instructions = fin.read().split("\n\n")
a = int(re.findall(r"A:\s(\d+)", registers)[0])
b = int(re.findall(r"B:\s(\d+)", registers)[0])
c = int(re.findall(r"C:\s(\d+)", registers)[0])
instructions = list(map(int, instructions.split(": ")[1].split(",")))

part1 = ",".join(map(str, simulate(a, b, c, instructions)))

part2 = 0
for i in range(1, len(instructions)+1):
  part2 *= 8
  while True:
    if reverse_simulate(part2, instructions[-i:]):
      break
    part2 += 1

print("Part 1:", part1)
print("Part 2:", part2)
