import advent

advent.setup(2023, 9)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

def solve(nums):
  if all([i == 0 for i in nums]):
    return 0

  diffs = []
  for i in range(1, len(nums)):
    diffs.append(nums[i] - nums[i-1])

  sol = solve(diffs)

  # return nums[-1] + sol # part1
  return nums[0] - sol # part2

part1 = 0
part2 = 0

for line in lines:
  nums = [int(n) for n in line.split()]
  res = solve(nums)
  # part1 += res
  part2 += res

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
