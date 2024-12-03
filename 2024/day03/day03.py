# https://adventofcode.com/2024/day/3
# https://github.com/Favo02/advent-of-code

import re
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

lines = []
for line in sys.stdin:
  lines.append(line.strip())

matches = []
MUL = re.compile(r'mul\((\d+),(\d+)\)')
DO = re.compile(r'do\(\)')
DONT = re.compile(r'don\'t\(\)')

for y, line in enumerate(lines):
  for match in MUL.finditer(line):
    i = match.start()
    a, b = map(int, (match.group(1), match.group(2)))
    if not (1 <= a <= 999): continue
    if not (1 <= b <= 999): continue
    matches.append(((y, i), "MUL", (a, b)))

for y, line in enumerate(lines):
  for match in DO.finditer(line):
    i = match.start()
    matches.append(((y, i), "DO", (0, 0)))

for y, line in enumerate(lines):
  for match in DONT.finditer(line):
    i = match.start()
    matches.append(((y, i), "DONT", (0, 0)))

matches.sort()

part1 = part2 = 0
valid = True

for (_, what, (a, b)) in matches:
  if what == "DO": valid = True
  elif what == "DONT": valid = False

  if what == "MUL":
    part1 += a * b
    if valid: part2 += a * b

print("Part 1:", part1)
print("Part 2:", part2)
