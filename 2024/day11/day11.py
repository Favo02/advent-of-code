# https://adventofcode.com/2024/day/11
# https://github.com/Favo02/advent-of-code

from collections import Counter, defaultdict
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

nums = Counter(list(map(int, fin.read().strip().split())))

for time in range(75):
  if time == 25:
    part1 = sum(nums.values())

  newnums = defaultdict(int)

  for n, qty in nums.items():
    if n == 0:
      newnums[1] += qty
    elif (L := len(str(n))) % 2 == 0:
      newnums[int(str(n)[:L//2])] += qty
      newnums[int(str(n)[L//2:])] += qty
    else:
      newnums[n * 2024] += qty

  nums = newnums

print("Part 1:", part1)
print("Part 2:", sum(nums.values()))
