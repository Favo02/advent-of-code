# https://adventofcode.com/2023/day/17
# https://github.com/Favo02/advent-of-code

import sys
import heapq
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def dijkstra(matrix, min_straight=1, max_straight=float("inf")):
  W, H = len(matrix[0]), len(matrix)
  START = (0, 0)
  END = (W-1, H-1)
  DIRS = [(-1, 0), (0, -1), (0, 1), (1, 0)]

  distances = {(START, 0, (0, 0)): 0}
  queue = [(0, START, 0, (0, 0))]

  while queue:
    cur = heapq.heappop(queue)
    cur_dist, (cur_x, cur_y), length, (cur_dirx, cur_diry) = cur

    if (cur_x, cur_y) == END:
      return cur_dist

    for delta_x, delta_y in DIRS:
      # limit min straight (ignoring start point with dir 0,0)
      if (cur_dirx, cur_diry) != (0,0) and\
         ((cur_dirx != delta_x or cur_diry != delta_y) and length < min_straight):
        continue
      # limit max straight
      if cur_dirx == delta_x and cur_diry == delta_y and length == max_straight:
        continue
      # do not turn 180°
      if cur_dirx == -delta_x and cur_diry == -delta_y:
        continue

      adj_x, adj_y = cur_x + delta_x, cur_y + delta_y

      # out of bounds
      if not ((0 <= adj_x < W) and (0 <= adj_y < H)):
        continue

      new_dist = cur_dist + matrix[adj_x][adj_y]

      new_length = length + 1 if (cur_dirx == delta_x and cur_diry == delta_y) else 1
      adj = (adj_x, adj_y), new_length, (delta_x, delta_y)

      if adj not in distances or new_dist < distances[adj]:
        distances[adj] = new_dist
        heapq.heappush(queue, (new_dist, *adj))
  assert False, "Cannot build path to END with these restraints"

part1 = 0
part2 = 0

matrix = []
for line in fin:
  matrix.append([int(n) for n in line.rstrip()])

part1 = dijkstra(matrix, 1, 3)
part2 = dijkstra(matrix, 4, 10)

print("Part 1:", part1)
print("Part 2:", part2)
