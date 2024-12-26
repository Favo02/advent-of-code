# https://adventofcode.com/2024/day/21
# https://github.com/Favo02/advent-of-code

import sys
from itertools import permutations
from functools import cache
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

MOVE = {
  ">": (1,0),
  "<": (-1,0),
  "v": (0,1),
  "^": (0,-1),
  "A": (0,0),
}

NUMERICAL = {
  "7": (0,0),
  "8": (1,0),
  "9": (2,0),
  "4": (0,1),
  "5": (1,1),
  "6": (2,1),
  "1": (0,2),
  "2": (1,2),
  "3": (2,2),
  "X": (0,3), # gap
  "0": (1,3),
  "A": (2,3),
}

DIRECTIONAL = {
  'X': (0,0), # gap
  '^': (1,0),
  'A': (2,0),
  '<': (0,1),
  'v': (1,1),
  '>': (2,1),
}

def generate_moves(x, y, tx, ty):
  return "<" * max(0, x-tx) +\
         ">" * max(0, tx-x) +\
         "^" * max(0, y-ty) +\
         "v" * max(0, ty-y)

def is_valid_path(x, y, path, gap, target):
  if (x, y) == gap: return False
  for p in path:
    dx, dy = MOVE[p]
    x += dx
    y += dy
    if (x, y) == gap: return False
  assert (x, y) == target
  return True

NUM_PATHS = {}
DIR_PATHS = {}

def precalc_paths():
  for dict, res in [(NUMERICAL, NUM_PATHS), (DIRECTIONAL, DIR_PATHS)]:
    for fr, fcoords in dict.items():
      if fr == "X": continue
      for to, tcoords in dict.items():
        if to == "X": continue
        if fr == to:
          res[(fr, to)] = [("A",)]
          continue

        all_moves = permutations(generate_moves(*fcoords, *tcoords))
        res[(fr, to)] = {a + ("A",) for a in all_moves if is_valid_path(*fcoords, a, dict["X"], tcoords)}

@cache
def length(fr, to, depth):

  min_l = float("inf")

  if (fr, to) in DIR_PATHS:
    DICT = DIR_PATHS
  else:
    DICT = NUM_PATHS

  for path in DICT[(fr, to)]:
    if depth == 1:
      return len(path)

    l = 0
    last = "A"
    for elem in path:
      l += length(last, elem, depth-1)
      last = elem

    min_l = min(min_l, l)

  return min_l

def encode(code, depth):
  l = 0
  last = "A"
  for elem in code:
    l += length(last, elem, depth)
    last = elem
  return l

codes = [c.strip() for c in fin]

precalc_paths()

part1 = part2 = 0

for code in codes:
  p1 = encode(code, 3)
  p2 = encode(code, 26)
  part1 += int(code[:-1]) * p1
  part2 += int(code[:-1]) * p2

print("Part 1:", part1)
print("Part 2:", part2)
