import advent

advent.setup(2023, 22)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

def overlap(a, b):
  return not (a[0] > b[1] or a[1] < b[0])

def fall(bricks):
  bsorted = sorted(bricks, key=lambda b: b[2][0])
  fallen = []
  for bri in bsorted:
    z = 1
    height = bri[2][1] - bri[2][0]
    for under in fallen:
      if overlap(bri[0], under[0]) and overlap(bri[1], under[1]):
        z = max(z, under[2][1]+1)
    fallen.append((bri[0], bri[1], (z, z+height)))
  return fallen

part1 = 0
part2 = 0

bricks = []
for line in lines:
  above, e = line.split("~")
  xs, ys, zs = [int(n) for n in above.split(",")]
  xe, ye, ze = [int(n) for n in e.split(",")]
  bricks.append(((xs,xe), (ys,ye), (zs,ze)))

fallen = sorted(fall(bricks), key=lambda b: b[2][0])

for i, brick in enumerate(fallen):
  removed = fallen[:i] + fallen[i+1:]
  simulation = fall(removed)

  for j in range(len(bricks)-1):
    if removed[j] != simulation[j]:
      part1 += 1
      break

part1 = len(bricks) - part1

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
