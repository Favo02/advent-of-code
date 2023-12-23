import advent
import sys
from collections import deque, defaultdict

sys.setrecursionlimit(10**5)
advent.setup(2023, 23)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

part1 = 0
part2 = 0

matrix = []
start = None
end = None
for i,line in enumerate(lines):
  if i == 0:
    start = (line.index("."), 0)
  if i == len(lines)-1:
    end = (line.index("."), len(lines)-1)
  matrix.append(list(line))

assert start
assert end

W, H = len(matrix[0]), len(matrix)
DIRS = [(0,1),(0,-1),(1,0),(-1,0)]
# print(end)

def compress(cur):
  compressed_graph = defaultdict(list)
  queue = deque([ (cur,0,cur) ])
  seen = set()

  while queue:
    last_incr, dist, (x,y) = queue.popleft()
    # print(x,y)
    if (last_incr,x,y) in seen:
      continue
    seen.add((last_incr,x,y))
    adjs = []

    for dx,dy in DIRS:
      if not ((0 <= x+dx < W) and (0 <= y+dy < H)):
        continue
      if matrix[y+dy][x+dx] == "#":
        continue
      adjs.append((x+dx,y+dy))

    # print(x,y, adjs, queue)

    if len(adjs) > 2:
      # print("incrocio", x,y)
      if last_incr != (x,y):
        compressed_graph[last_incr].append((x,y, dist+1))
      for a in adjs:
        queue.append(((x,y), 0, a))
    elif len(adjs) == 1:
      # print("deadend", x,y)
      if last_incr != (x,y):
        compressed_graph[last_incr].append((x,y, dist+1))
      queue.append(((x,y), 0, adjs[0]))
    else:
      for a in adjs:
        queue.append((last_incr, dist+1, a))
    # print(queue)
  return compressed_graph

graph = compress(start)
# print(graph)

dist = defaultdict(int)
visited = set()
def brutesolve(cur, cur_sum):
  # print(cur)
  if cur in visited:
    return
  visited.add(cur)
  if dist[cur] < cur_sum:
    dist[cur] = cur_sum

  for ax, ay, w in graph[cur]:
    brutesolve((ax, ay), cur_sum + w)
  visited.remove(cur)

brutesolve(start, 0)
# print(dist)
part2 = dist[end]

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
