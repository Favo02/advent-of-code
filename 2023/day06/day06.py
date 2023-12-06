# https://adventofcode.com/2023/day/6
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = 1
part2 = 0

p1times = [int(n) for n in input().split()[1:]]
p1records = [int(n) for n in input().split()[1:]]

p2time = int("".join(map(str, p1times)))
p2record = int("".join(map(str, p1records)))

def bruteforce(time, rec):
  count = 0
  for start in range(time):
    if start*(time-start) > rec:
      count += 1
  return count

for time, rec in zip(p1times, p1records):
  part1 *= bruteforce(time, rec)

part2 = bruteforce(p2time, p2record)

print("Part 1:", part1)
print("Part 2:", part2)
