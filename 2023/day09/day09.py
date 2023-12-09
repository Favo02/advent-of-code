# https://adventofcode.com/2023/day/9
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def solve(nums, part2):
  if all([i == 0 for i in nums]):
    return 0

  diffs = [nums[i] - nums[i-1] for i in range(1, len(nums))]

  if part2:
    return nums[0] - solve(diffs, part2)
  else:
    return nums[-1] + solve(diffs, part2)

part1 = 0
part2 = 0

for line in fin:
  nums = [int(n) for n in line.rstrip().split()]
  part1 += solve(nums, False)
  part2 += solve(nums, True)

print("Part 1:", part1)
print("Part 2:", part2)
