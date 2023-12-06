# https://adventofcode.com/2023/day/5
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def split(toSplit, where):
  [ts, te] = toSplit
  [ws, we] = where

  # where:   [               )

  #    [ )                              = 1 range (= old)
  #                               [  )  = 1 range (= old)
  #    [         )                      = 2 ranges
  #                       [        )    = 2 ranges
  #               [   )                 = 1 range (= old)
  #      [                         )    = 3 ranges

  before =        ts, min(ws, te)
  intersection =  max(ws, ts), min(we, te)
  after =         max(we, ts), te

  # filter out only valid
  valid = []
  if before[0] < before[1]: valid.append(before)
  if intersection[0] < intersection[1]: valid.append(intersection)
  if after[0] < after[1]: valid.append(after)
  return valid

def shred(seeds, operations):
  for operation in operations:
    done = []
    for seed in seeds:
      done += split(seed, operation)
    seeds = done
  return seeds

def solvePart(steps, seeds):
  for step in steps:
    seeds = shred(seeds, step)

    for i, (seedS, seedE) in enumerate(seeds):
      for (operS, operE), offset in step.items():
        if operS <= seedS < seedE <= operE:
          seeds[i] = seedS+offset, seedE+offset
  return seeds

part1 = 0
part2 = 0

seeds = [int(n) for n in fin.readline()[6:].split() if len(n)]

# trasfrom to ranges in format [start, end) closed on start, open on end
p1seeds = [(s, s+1) for s in seeds]
p2seeds = [(seeds[i], seeds[i] + seeds[i+1]) for i in range(0, len(seeds), 2)]

steps = []

for line in fin:
  line = line.lstrip()

  if "map" in line: # separator between steps
    steps.append({})

  elif len(line): # line not empty
    tokens = [int(n) for n in line.split()]
    operS, operE, offset = tokens[1], tokens[1] + tokens[2], tokens[0]-tokens[1]
    steps[-1][(operS,operE)] = offset

part1 = min(solvePart(steps, p1seeds))[0]
part2 = min(solvePart(steps, p2seeds))[0]

print("Part 1:", part1)
print("Part 2:", part2)
