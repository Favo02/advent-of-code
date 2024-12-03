import sys

def sign(a, b):
  if a-b == 0: return +1
  return (a-b) / abs(a-b)

def valid(report):
  assert len(report) > 1

  s = None

  for a, b in zip(report, report[1:]):
    if s == None: s = sign(a, b)
    if sign(a, b) != s: return False
    if not (1 <= abs(a-b) <= 3): return False

  return True

part1 = part2 = 0

for line in sys.stdin:
  nums = list(map(int, line.split()))

  if valid(nums):
    part1 += 1

  for i in range(len(nums)):
    if valid(nums[:i] + nums[i+1:]):
      part2 += 1
      break

print("Part 1:", part1)
print("Part 2:", part2)
