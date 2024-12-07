# https://adventofcode.com/2024/day/7
# https://github.com/Favo02/advent-of-code

from itertools import chain
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def valid(TARGET, NUMS, CONCATENATION, i):
  if i == 0:
    return [ NUMS[0] ]

  rec = [v for v in valid(TARGET, NUMS, CONCATENATION, i-1) if v <= TARGET]

  add = [NUMS[i] + r for r in rec]
  mul = [NUMS[i] * r for r in rec]

  if not CONCATENATION: return add + mul

  con = [r*(10**len(str(NUMS[i]))) + NUMS[i] for r in rec]
  return add + mul + con

part1 = part2 = 0

for line in fin:
  target, nums = line.strip().split(": ")
  target = int(target)
  nums = list(map(int, nums.split(" ")))

  if target in valid(target, nums, False, len(nums)-1): part1 += target
  if target in valid(target, nums, True, len(nums)-1): part2 += target

print("Part 1:", part1)
print("Part 2:", part2)
