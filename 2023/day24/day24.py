# https://adventofcode.com/2023/day/24
# https://github.com/Favo02/advent-of-code

import sys
from sympy import symbols, nonlinsolve
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def line_formula(point1, point2):
  (x1, y1), (x2, y2) = point1, point2
  if x1 == x2:
    return None, None
  m = (y2 - y1) / (x2 - x1)
  b = y1 - m * x1
  return m, b

def line_intersection(m1, b1, m2, b2):
  if m1 == m2: # parallel
    return None
  x = (b2 - b1) / (m1 - m2)
  y = m1 * x + b1
  return x, y

def valid_bounds(intersection, bounds):
  assert intersection != None
  x, y = intersection
  return (bounds[0] <= x <= bounds[1]) and (bounds[0] <= y <= bounds[1])

def same_sign(a, b):
  return a*b >= 0

def valid_direction(stone, intersection):
  x, y = stone[COORDS][X], stone[COORDS][Y]
  sx, sy = stone[SPEEDS][X], stone[SPEEDS][Y]
  ix, iy = intersection
  return same_sign(ix-x, sx) and same_sign(iy-y, sy)

def solve_system(vectors):
  # solution unknowns
  sx, sy, sz, svx, svy, svz, = symbols("sx sy sz svx svy svz")

  equations = []
  times = []
  for i, (point, velocity) in enumerate(vectors):
    x, y, z = point
    vx, vy, vz = velocity
    time = symbols(f"time{i}")
    times.append(time)
    equations.append(time >= 0)
    equations.append(x + time*vx - sx - time*svx)
    equations.append(y + time*vy - sy - time*svy)
    equations.append(z + time*vz - sz - time*svz)
  unknowns = [sx, sy, sz, svx, svy, svz] + times

  try:
    [ result ] = nonlinsolve(equations, tuple(unknowns))
  except:
    assert False, "cannot calculate system"
  assert result, f"no valid solutions to system: {result}"

  return result[:3]

part1 = 0
part2 = 0

# to access tuples
COORDS, SPEEDS = 0, 1
X, Y, Z = 0, 1, 2

# BOUNDS = 7, 27 # example bounds
BOUNDS = 200000000000000, 400000000000000

hailstones = []
formulas = []
for line in fin:
  line = line.strip()
  coords, speed = line.split(" @ ")
  x,y,z = map(int, coords.split(", "))
  vx,vy,vz = map(int, speed.split(", "))
  hailstones.append(((x, y, z), (vx, vy, vz)))
  formulas.append(line_formula((x, y), (x+vx, y+vy)))

for i, formula1 in enumerate(formulas):
  for j, formula2 in enumerate(formulas[i+1:]):
    intersection = line_intersection(*formula1, *formula2)

    if intersection == None:
      continue
    if not valid_bounds(intersection, BOUNDS):
      continue
    if not valid_direction(hailstones[i], intersection):
      continue
    if not valid_direction(hailstones[i+1+j], intersection):
      continue

    part1 += 1

part2 = sum(solve_system(hailstones[:3]))

print("Part 1:", part1)
print("Part 2:", part2)
