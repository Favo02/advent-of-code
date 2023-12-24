import advent

advent.setup(2023, 24)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

def printt(t):
  for tt in t:
    print(tt)

def line_formula(point1, point2):
  (x1, y1), (x2, y2) = point1, point2
  if x1 == x2:
    return None, None
  else:
    m = (y2 - y1) / (x2 - x1)
    b = y1 - m * x1
    return m, b

def line_intersection(m1, b1, m2, b2):
  if m1 == m2:
    return None, None
  else:
    x = (b2 - b1) / (m1 - m2)
    y = m1 * x + b1
    return x, y

def past_checker(i, intx,inty):
  t = things[i]
  if t[SPEEDS][X] < 0:
    if not (intx < t[COORDS][X]):
      return False
  else:
    if not (intx > t[COORDS][X]):
      return False
  if t[SPEEDS][Y] < 0:
    if not (inty < t[COORDS][Y]):
      return False
  else:
    if not (inty > t[COORDS][Y]):
      return False
  return True

part1 = 0
part2 = 0

COORDS, SPEEDS = 0, 1
X, Y, Z = 0, 1, 2

things = []
formulas = []
for line in lines:
  coords, speed = line.split(" @ ")
  x,y,z = map(int, coords.split(", "))
  vx,vy,vz = map(int, speed.split(", "))
  things.append(((x,y,z), (vx,vy,vz)))
  formulas.append(line_formula((x,y),(x+vx,y+vy)))

# printt(things)
# printt(formulas)
# print("="*10)
BOUNDS = 7, 27
BOUNDS = 200000000000000, 400000000000000

for i,f1 in enumerate(formulas):
  for j, f2 in list(enumerate(formulas))[i+1:]:
    intx,inty = line_intersection(*f1, *f2)
    if intx != None and inty != None:
      if BOUNDS[0] <= intx <= BOUNDS[1]:
        if BOUNDS[0] <= inty <= BOUNDS[1]:
          if past_checker(i,intx,inty):
            if past_checker(j,intx,inty):
              # print("int", i,j,"-", f1,f2, ">", (intx,inty))
              part1 += 1

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
