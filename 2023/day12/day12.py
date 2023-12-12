# https://adventofcode.com/2023/day/12
# https://github.com/Favo02/advent-of-code

import sys
from functools import lru_cache
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def reduce_input(line):
  new = [line[0]]
  for i in range(1, len(line)):
    if line[i] == line[i-1] == ".":
      continue
    new.append(line[i])
  return "".join(new)

def expand_input(string, tokens):
  string = "?".join([string]*5)
  tokens = tokens * 5
  return string, tokens

@lru_cache(maxsize=None)
def solve(string, tokens):
  if not tokens:
    return 1 if ('#' not in string) else 0

  pivot = max(tokens)
  pivot_i = tokens.index(pivot)

  LEN = len(string)

  res = 0
  for i in range(0, LEN-pivot +1):
    if "." in string[i:i+pivot]:
      continue

    bef = aft = 0

    if i == 0 or string[i-1] != "#":
      bef = solve(string[:max(0,i-1)], tokens[:pivot_i])

    if i+pivot == LEN or string[pivot+i] != "#":
      aft = solve(string[i+pivot+1:], tokens[pivot_i+1:])

    res += bef * aft
  return res

part1 = 0
part2 = 0

for line in fin:
  string, tokens = line.strip().split()
  string = reduce_input(string)
  tokens = [int(n) for n in tokens.split(",")]

  part1 += solve(string, tuple(tokens))
  string, tokens = expand_input(string, tokens)
  part2 += solve(string, tuple(tokens))

print("Part 1:", part1)
print("Part 2:", part2)
