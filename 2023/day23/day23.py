# https://adventofcode.com/2023/day/23
# https://github.com/Favo02/advent-of-code

import sys
from collections import defaultdict, deque
sys.setrecursionlimit(10**5)
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def matrix_dfs(matrix, cur, cur_dist, seen=set(), dist=defaultdict(int)):
  if cur in seen:
    return dist
  seen.add(cur)

  if dist[cur] < cur_dist:
    dist[cur] = cur_dist

  x, y = cur
  for dx, dy in DIRS:
    ax, ay = x+dx, y+dy
    if not ((0 <= ax < W) and (0 <= ay < H)):
      continue
    if matrix[ay][ax] == "#":
      continue
    if matrix[ay][ax] in ALLOWED_DIR:
      if (dx, dy) != ALLOWED_DIR[matrix[ay][ax]]:
        continue
    matrix_dfs(matrix, (ax, ay), cur_dist+1, seen, dist)
  seen.remove(cur)
  return dist

def compress(cur):
  compressed_graph = defaultdict(list)
  queue = deque([ (cur, cur, 0) ])
  seen = set()

  while queue:
    last_split, (x, y), dist = queue.popleft()
    if (last_split, (x, y)) in seen:
      continue

    adjs = []
    for dx, dy in DIRS:
      ax, ay = x+dx, y+dy
      if not ((0 <= ax < W) and (0 <= ay < H)):
        continue
      if matrix[ay][ax] == "#":
        continue
      if (last_split, (ax, ay)) in seen:
        continue
      adjs.append((ax, ay))

    # dead end (likely start or end)
    if len(adjs) == 0:
      compressed_graph[last_split].append(((x, y), dist))
      seen.add((last_split, (x, y)))
    # normal road (only one possible path to take)
    elif len(adjs) == 1:
      next = adjs[0]
      queue.append((last_split, next, dist+1))
      seen.add((last_split, (x, y)))
    # split (multiple possible paths)
    else:
      compressed_graph[last_split].append(((x, y), dist+1))
      seen.add(((x, y), (x, y)))
      for next in adjs:
        queue.append(((x, y), next, 0))
  return compressed_graph

def graph_dfs(graph, cur, cur_dist, seen=set(), dist=defaultdict(int)):
  if cur in seen:
    return dist
  seen.add(cur)

  if dist[cur] < cur_dist:
    dist[cur] = cur_dist

  for (ax, ay), weight in graph[cur]:
    dist = graph_dfs(graph, (ax, ay), cur_dist + weight, seen, dist)
  seen.remove(cur)
  return dist

part1 = 0
part2 = 0

matrix = []
start = None
end = None
for y, line in enumerate(fin):
  matrix.append(line.strip())
  if start == None:
    start = (line.index("."), y)
  end = (line.index("."), y)

assert start
assert end

W, H = len(matrix[0]), len(matrix)
DIRS = [(0,1), (0,-1), (1,0), (-1,0)]
ALLOWED_DIR = {
  ">": (+1, 0),
  "<": (-1, 0),
  "^": (0, -1),
  "v": (0, +1),
}

part1 = matrix_dfs(matrix, start, 0)[end]

graph = compress(start)
part2 = graph_dfs(graph, start, 0)[end]

print("Part 1:", part1)
print("Part 2:", part2)
