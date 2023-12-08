import advent

advent.setup(2023, 8)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

# def bfs(graph, start, dirs):
#   dist = {start : 0}
#   queue = [start]

#   time = 0

#   while queue:
#     cur = queue.pop()

#     if dirs[time] == "L":

#     for adj in cur:



part1 = 0
part2 = 0

dirs = lines[0]
graph = {}

for line in lines[2:]:
  fromm, to = line.split(" = ")
  l,r = to.split(", ")
  l = l[1:]
  r = r[:-1]
  print(fromm, l, r)
  graph[fromm] = (l, r)


cur = "AAA"
dist = 0

while cur != "ZZZ":
  print(cur)
  dir = dirs[dist % len(dirs)]
  if dir == "L":
    cur = graph[cur][0]
  else:
    cur = graph[cur][1]
  dist += 1


part1 = dist
print(dist)


print(graph)


advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
