import advent

advent.setup(2023, 3)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

def reconstructDigit(line,x):
  s = x
  if line[x-1].isdigit():
    s -= 1
    if line[x-2].isdigit():
      s -= 1
  e = x
  if line[x+1].isdigit():
    e += 1
    if line[x+2].isdigit():
      e += 1
  print(int("".join(line[s:e+1])))
  return int("".join(line[s:e+1]))


def recSearch(lines,x,y):
  digits = set()
  for dy in range(-1,2):
    if (y == 0 and dy == -1) or (y == len(lines)-1 and dy == 1): continue
    for dx in range(-1,2):
      if (x == 0 and dx == -1) or (x == len(lines[0])-1 and dx == 1): continue
      if dx == dy == 0: continue
      if lines[y+dy][x+dx].isdigit():
        digits.add(reconstructDigit(lines[y+dy],x+dx))
  return digits

part1 = 0
part2 = 0

for li, line in enumerate(lines):
  buffer = []
  for ci,c in enumerate(line):
    if c == '*':
      dig = recSearch(lines,ci,li)
      print(dig)
      if len(dig) == 2:
        dig=list(dig)
        part2 += dig[0] * dig[1]

# advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
