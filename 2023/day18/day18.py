# https://adventofcode.com/2023/day/18
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

DIRS = {
  "U": (0, -1),
  "D": (0, +1),
  "L": (-1, 0),
  "R": (+1, 0)
}
HEX_TO_DIR = {
  "0": "R",
  "1": "D",
  "2": "L",
  "3": "U"
}

def parse_points(lines, part2=False):
  points = []
  cur = (0, 0)
  border_length = 0

  for line in lines:
    if part2:
      hex = line.split()[2][2:-1]
      amt = int(hex[:5], 16)
      dir = HEX_TO_DIR[hex[5]]
    else:
      dir, amt, _ = line.split()
      amt = int(amt)

    curx, cury = cur
    dirx, diry = DIRS[dir]
    next = (curx + amt*dirx, cury + amt*diry)
    border_length += (abs(curx - next[0]) + abs(cury - next[1]))
    points.append(next)
    cur = next
  return points, border_length

# credits: https://stackoverflow.com/a/24468019/14853184
def area(points):
  n = len(points)
  area = 0
  for i in range(n):
    j = (i + 1) % n
    area += points[i][0] * points[j][1]
    area -= points[j][0] * points[i][1]
  area = abs(area) // 2
  return area

part1 = 0
part2 = 0

lines = []
for line in fin:
  lines.append(line.rstrip())

points, border_length = parse_points(lines, False)
part1 = area(points) + (border_length // 2) + 1

points, border_length = parse_points(lines, True)
part2 = area(points) + (border_length // 2) + 1

print("Part 1:", part1)
print("Part 2:", part2)
