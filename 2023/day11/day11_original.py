import advent

advent.setup(2023, 11)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

def man_dist(g1, g2):
  x1, y1 = g1
  x2, y2 = g2
  return abs(x1 - x2) + abs(y1 - y2)

def dists(galaxs):
  dists = {}

  for g1 in galaxs:
    for g2 in galaxs:
      if g1 == g2: continue
      d = man_dist(g1, g2)
      if (g1,g2) not in dists:
        dists[(g1,g2)] = d
        dists[(g2,g1)] = d
      else:
        dists[(g1,g2)] = max(dists[(g1,g2)], d)
        dists[(g2,g1)] = max(dists[(g2,g1)], d)

  return dists

part1 = 0
part2 = 0

P2_DIST = 1000000
# P2_DIST = 100 # example

galaxs = []
graph = {}
empty_rows = []
empty_cols = []

for y,line in enumerate(lines):
  empty = True
  for x,cell in enumerate(line):
    if cell == "#":
      galaxs.append((x,y))
      empty = False
  if empty:
    empty_rows.append(y)
for x in range(len(lines[0])):
  empty = True
  for row in lines:
    if row[x] == "#":
      empty = False
      break
  if empty:
    empty_cols.append(x)

# print(galaxs)

# print(empty_rows)
er_index = 0
for i, (gx,gy) in enumerate(galaxs):
  while er_index < len(empty_rows) and gy > empty_rows[er_index]:
    er_index += 1

  galaxs[i] = (gx,(gy-er_index)+(er_index*P2_DIST))

# print(galaxs)
galaxs.sort()
# print(galaxs)

rows_index = 0
for i, (gx,gy) in enumerate(galaxs):
  while rows_index < len(empty_cols) and gx > empty_cols[rows_index]:
    rows_index += 1

  galaxs[i] = ((gx-rows_index)+(rows_index*P2_DIST),gy)

res = dists(galaxs)
# print(res)

used = set()
for (f,t), d in res.items():
  if ((f,t) not in used) and ((t,f) not in used):
    part2 += d
    used.add((f,t))

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
