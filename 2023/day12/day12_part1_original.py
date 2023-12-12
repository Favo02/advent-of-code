import advent

advent.setup(2023, 12)
fin = advent.get_input()
fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

def count(spring):
  cur = 0
  found = []
  for s in spring:
    if s == '#' or s == '?':
      cur += 1
    elif cur != 0:
      found.append(cur)
      cur = 0
  if cur != 0:
    found.append(cur)
  return found

def is_valid(spring, nums):
  return count(spring) == nums

def invalid(spring):
  alll = ["#" if s=="#" else "." for s in spring]
  count_all = count(alll)
  if count_all and max(count_all) > max(nums):
    return True

  none = ["." if s=="." else "#" for s in spring]
  count_none = count(none)
  if count_none and max(count_none) < max(nums):
    return True
  return False

used = set()
def solve(spring, spaces):
  if spaces > 2 and invalid(spring):
    # print("inv")
    return 0
  if spaces == 0:
    res = 1 if is_valid(spring, nums) else 0
    if res == 0:
      # print(nums, "".join(spring))
      pass
    if "".join(spring) in used:
      res = 0
    else:
      used.add("".join(spring))
    return res

  to_fill = spring.count("?")

  res = 0
  for i in range(0, len(spring)):
    if to_fill+1 == spaces:
      break
  # for i,s in enumerate(spring):
    if spring[i] != "?":
      continue
    to_fill -= 1
    spring[i] = "."
    call = solve(spring, spaces-1)
    res += call
    spring[i] = "?"

  return res

part1 = 0
part2 = 0

# 1 4 1 1 4 10 = 21
for i, line in enumerate(lines):
  spring, nums = line.split()
  nums = [int(n) for n in nums.split(",")]
  cou = count(spring)
  spaces = sum(cou) - sum(nums)
  print(spring, cou, nums, spaces)
  print(f"{i} / {len(lines)}")
  spring = list(spring)

  used.clear()
  if spaces == 0:
    print(1)
    part1 += 1
  else:
    res = solve(spring, spaces)
    print(res)
    part1 += res

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
