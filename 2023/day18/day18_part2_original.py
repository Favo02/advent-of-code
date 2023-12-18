import advent
from shapely.geometry import Polygon

advent.setup(2023, 18)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
# fin = advent.get_input("easy_input2.txt")
lines = advent.get_lines(fin)

part1 = 0
part2 = 0

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

cur = (0,0)
points = [ cur ]
extra = 0

for line in lines:
  hex = line.split()[2][2:-1]
  amt = int(hex[:5], 16)
  dir = HEX_TO_DIR[hex[5]]

  curx, cury = cur
  dirx, diry = DIRS[dir]
  next = (curx + amt*dirx, cury + amt*diry)
  extra += (abs(curx - next[0]) + abs(cury - next[1]))
  points.append(next)
  cur = next

pgon = Polygon(points)
part2 = int(pgon.area + (extra / 2) + 1)

# advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# 952_408_144_115 correct
# 952_411_346_717 area + extra -2
# 952_411_346_731 area + extra -1
# 952_404_941_483 area

# advent.submit_answer(2, part2)
