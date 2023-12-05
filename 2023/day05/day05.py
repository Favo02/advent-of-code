# https://adventofcode.com/2023/day/5
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def split(seed, operation):
  [sS, sE] = seed
  [mS, mE] = operation

  res = []

  if sS < mS < sE:
    res.append((sS, mS-1))
    res.append((mS, sE))

  elif sS < mE < sE:
    res.append((sS, mE))
    res.append((mE+1, sE))

  elif mS <= sS < sE <= mE:
    res.append((sS, sE))

  elif sS < mS < mE < sE:
    res.append((sS, mS-1))
    res.append((mS, mE))
    res.append((mE+1, sE))

  else:
    res.append((sS, sE))

  return res

def shred(seeds, operations):
  for m in operations:
    i = 0
    while i < len(seeds):
      splitRes = split(seeds[i], m)
      if len(splitRes) == 1:
        i += 1
      else:
        seeds = seeds[:i] + seeds[i+1:] + splitRes
  return seeds

part1 = 0
part2 = 0

seeds = [int(n) for n in fin.readline()[6:].split() if len(n)]

# trasfrom to ranges
p1seeds = [(s, s) for s in seeds]
p2seeds = [(seeds[i], seeds[i] + seeds[i+1]) for i in range(0, len(seeds), 2)]

steps = []

for line in fin:
  line = line.lstrip()

  if "map" in line: # separator between steps
    steps.append({})

  elif len(line): # line not empty
    tokens = [int(n) for n in line.split()]
    operS, operE, offset = tokens[1], tokens[1] + tokens[2]-1, tokens[0]-tokens[1]
    steps[-1][(operS,operE)] = offset

for step in steps:

  p1seeds = shred(p1seeds, step)

  for i, (seedS, seedE) in enumerate(p1seeds):
    for (operS, operE), offset in step.items():
      if operS <= seedS <= seedE <= operE:
        p1seeds[i] = seedS+offset, seedE+offset

  p2seeds = shred(p2seeds, step)

  for i, (seedS, seedE) in enumerate(p2seeds):
    for (operS, operE), offset in step.items():
      if operS <= seedS <= seedE <= operE:
        p2seeds[i] = seedS+offset, seedE+offset

part1 = min(p1seeds)[0]
part2 = min(p2seeds)[0]

print("Part 1:", part1)
print("Part 2:", part2)
