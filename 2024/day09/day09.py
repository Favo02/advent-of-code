# https://adventofcode.com/2024/day/9
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def checksum(disk):
  res = 0
  id = 0
  for file, qty, val in disk:
    if not file:
      id += qty
      continue
    for _ in range(qty):
      res += id*val
      id += 1
  return res

def part1(disk):
  expanded = []
  for file, size, val in disk:
    for _ in range(size):
      expanded.append((file, 1, val))

  l = 0
  r = len(expanded)-1

  while l < r:
    while l < r and expanded[l][0]:
      l += 1
    while l < r and not expanded[r][0]:
      r -= 1
    expanded[l], expanded[r] = expanded[r], expanded[l]

  return expanded

def part2(disk):
  r = len(disk)-1
  while r >= 0:
    r_file, r_size, r_val = disk[r]
    if not r_file:
      r -= 1
      continue

    l = 0
    while l < r:
      l_file, l_size, l_val = disk[l]
      if l_file:
        l += 1
        continue
      assert l_val == 0

      if l_size == r_size:
        disk[r], disk[l] = disk[l], disk[r]
        break
      elif l_size > r_size:
        disk[l] = disk[r]
        disk.insert(l+1, (False, l_size - r_size, 0))
        r += 1
        disk[r] = (False, r_size, 0)
        break

      l += 1

    r -= 1
  return disk

# disk encoding: (bool, int, int)
# - True: file, False: free space
# - int: size of the file/free space
# - int: value of the file (all free spaces are 0)
disk = [(i % 2 == 0, int(val), i // 2 if i % 2 == 0 else 0) for i, val in enumerate(input()) if int(val) > 0]

print("Part 1:", checksum(part1(disk.copy())))
print("Part 2:", checksum(part2(disk.copy())))
