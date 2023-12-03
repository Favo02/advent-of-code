import advent

advent.setup(2023, 3)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

def checkAdj(lines,x,y):
  # print("\nch",x,y)
  for dy in range(-1,2):
    if (y == 0 and dy == -1) or (y == len(lines)-1 and dy == 1): continue
    for dx in range(-1,2):
      if (x == 0 and dx == -1) or (x == len(lines[0])-1 and dx == 1): continue
      if dx == dy == 0: continue
      # print(lines[y+dy][x+dx])
      if (lines[y+dy][x+dx] != ".") and (not lines[y+dy][x+dx].isdigit()):
        return True
  return False

def checkBuf(lines, li, i, buffer):
  i -= len(buffer)
  for l in range(len(buffer)):
    if checkAdj(lines,i+l,li):
      return True
  return False

part1 = 0
part2 = 0

for li, line in enumerate(lines):
  buffer = []
  for ci,c in enumerate(line):
    if c.isdigit(): buffer.append(c)
    else:
      if buffer:
        res = checkBuf(lines, li, ci, buffer)
        print(int("".join(buffer)), res)
        if res:
          part1 += int("".join(buffer))
        buffer = []
  if buffer:
    res = checkBuf(lines, li, len(line), buffer)
    print(int("".join(buffer)), res)
    if res:
      part1 += int("".join(buffer))
    buffer = []
  # print(line)

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

# advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
