# https://adventofcode.com/2020/day/1
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = 0
part2 = 0

tokens = []
for line in fin:
  tokens.append(line.rstrip())

for i,t in enumerate(tokens):
  for j,tt in enumerate(tokens):
    if i == j: continue

    if int(t) + int(tt) == 2020:
      part1 = int(t) * int(tt)

    for k,ttt in enumerate(tokens):
      if k==i or k==j: continue
      t=int(t)
      tt=int(tt)
      ttt=int(ttt)

      if t + tt + ttt == 2020:
        part2 = t * tt * ttt

print("Part 1:", part1)
print("Part 2:", part2)
