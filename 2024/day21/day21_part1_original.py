from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, cache
from itertools import permutations
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

MOVE = {
  ">": (1, 0),
  "<": (-1, 0),
  "v": (0, 1),
  "^": (0, -1),
}

grid_nums = {
  '7': (0,0),
  '8': (1,0),
  '9': (2,0),
  '4': (0,1),
  '5': (1,1),
  '6': (2,1),
  '1': (0,2),
  '2': (1,2),
  '3': (2,2),
  '0': (1,3),
  'A': (2,3),
}

# +---+---+---+
# | 7 | 8 | 9 |
# +---+---+---+
# | 4 | 5 | 6 |
# +---+---+---+
# | 1 | 2 | 3 |
# +---+---+---+
#     | 0 | A |
#     +---+---+

grid_robot = {
  '^': (1,0),
  'A': (2,0),
  '<': (0,1),
  'v': (1,1),
  '>': (2,1),
}

#     +---+---+
#     | ^ | A |
# +---+---+---+
# | < | v | > |
# +---+---+---+

def moves(sx, sy, ex, ey):
  for _ in range(sx, ex): yield ">"
  for _ in range(ex, sx): yield "<"
  for _ in range(sy, ey): yield "v"
  for _ in range(ey, sy): yield "^"
  yield "A"

def encode_nums(code, x, y):
  res = []
  for c in code:
    res.extend(moves(x, y, *grid_nums[c]))
    x, y = grid_nums[c]
  return res

def encode_robot(code, x, y):
  res = []
  for c in code:
    res.extend(moves(x, y, *grid_robot[c]))
    x, y = grid_robot[c]
  return res

def all_combs(enc):
  parts = enc.split('A')

  results = set(map(lambda t: t + ('A',), permutations(parts[0])))

  for p in parts[1:]:
    newres = set()
    if not p: continue
    for r in results:
      for perm in set(permutations(p)):
        newres.append(r + perm + ('A',))
    results = newres

  return results

def valid_nums(x, y, path):
  for p in path:
    x += MOVE[p][0]
    y += MOVE[p][1]
    if not (0 <= x < 3): return False
    if not (0 <= y < 4): return False
    if (x, y) == (0, 3): return False
  return True

def valid_robots(x, y, path):
  for p in path:
    x += MOVE[p][0]
    y += MOVE[p][1]
    if not (0 <= x < 3): return False
    if not (0 <= y < 2): return False
    if (x, y) == (0, 0): return False
  return True

def precalc_nums():
  for c1, coords1 in grid_nums.items():
    for c2, coords2 in grid_nums.items():

      all_paths = set(permutations(list(moves(*coords1, *coords2))[:-1]))
      all_paths = {p + ('A',) for p in all_paths if valid_nums(*coords1, p)}
      PRE_N[(c1, c2)] = all_paths

      all_paths = set(permutations(list(moves(*coords2, *coords1))[:-1]))
      all_paths = {p + ('A',) for p in all_paths if valid_nums(*coords2, p)}
      PRE_N[(c2, c1)] = all_paths

def precalc_robot():
  for c1, coords1 in grid_robot.items():
    for c2, coords2 in grid_robot.items():

      all_paths = set(permutations(list(moves(*coords1, *coords2))[:-1]))
      all_paths = {p + ('A',) for p in all_paths if valid_robots(*coords1, p)}
      PRE_R[(c1, c2)] = all_paths

      all_paths = set(permutations(list(moves(*coords2, *coords1))[:-1]))
      all_paths = {p + ('A',) for p in all_paths if valid_robots(*coords2, p)}
      PRE_R[(c2, c1)] = all_paths

def chain_robots(last_path, times):
  for _ in range(times):
    new_path = set()
    for path in last_path:
      cur = set()
      last = 'A'
      for char in path:
        newcur = set()

        if not cur:
          for mapping in PRE_R[(last, char)]:
            cur.add(mapping)
        else:
          for c in cur:
            for mapping in PRE_R[(last, char)]:
              newcur.add(c + mapping)
          cur = newcur

        last = char
      new_path.update(cur)
    last_path = new_path
  return last_path

def super_precalc():
  for f1 in grid_nums.keys():
    for f2 in grid_nums.keys():
      n_paths = PRE_N[(f1, f2)]
      PREC[(f1, f2)] = chain_robots(n_paths, 2)

def prec_best():
  for k, v in PREC.items():
    best = None
    for path in v:
      if best is None or len(path) < len(best):
        best = path
      elif len(path) == len(best):
        best = min(best, path)
    BEST[k] = "".join(best)

res = 0

PRE_N = {}
PRE_R = {}

PREC = {}
BEST = {}

precalc_nums()
# for k, v in PRE_N.items():
#   print(k, v)
# print()

precalc_robot()
# for k, v in PRE_R.items():
#   print(k, v)
# print()

super_precalc()
# for k, v in PREC.items():
#   print(k, v)

prec_best()

print("fine prec")

part1 = 0

for y, line in enumerate(fin):
  line = line.strip()

  res = []
  last = 'A'
  for l in line:
    res.extend(BEST[(last, l)])
    print(f"{last}->{l}: {BEST[(last, l)]}")
    last = l
    # print(BEST[('A', l)])
  print("RES:", "".join(res))
  part1 += len(res) * int(line[:-1])

  # enc4, (x4, y4) = encode_robot(enc3, x4, y4)

print("RES:", part1)
