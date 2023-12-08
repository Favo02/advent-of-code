import advent

advent.setup(2023, 8)
fin = advent.get_input()
fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

part1 = 0
part2 = 0

dirs = lines[0]
graph = {}

starts = []

ends = []

for line in lines[2:]:
  fromm, to = line.split(" = ")
  l,r = to.split(", ")
  l = l[1:]
  r = r[:-1]

  if fromm[2] == "A": starts.append(fromm)
  if fromm[2] == "Z": ends.append(fromm)
  graph[fromm] = (l, r)

time = 0

curs = starts

wait_cycle = {}
cycles = {}

def calcRepet(rep, time):
  print("=" * 20)
  print(rep)
  adsa = [t[1] for t in rep]
  jump_size = max(adsa)

  stop = False
  while not stop:
    stop = True
    for fromm, cycle in rep:
      if (time-fromm) % cycle != 0:
        stop = False
    if not stop:
      time += jump_size

  return time

while True:
  dir = 0 if (dirs[time % len(dirs)] == "L") else 1

  for i,cur in enumerate(curs):
    if cur in cycles: continue
    curs[i] = graph[cur][dir]

    if curs[i][2] == "Z":

      if curs[i] in wait_cycle:
        cycles[curs[i]] = (time+1, time - wait_cycle[curs[i]])
      else:
        wait_cycle[curs[i]] = time

  time += 1

  if len(cycles) == len(curs):
    part2 = calcRepet(list(cycles.values()), time)
    break

print(part2)

# advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

# advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
