import advent
import math

advent.setup(2023, 4)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

part1 = 0
part2 = 0

for i,line in enumerate(lines):
  tokens = line.split(": ")
  parts = tokens[1].split(" | ")
  win = [p for p in parts[0].split(" ") if len(p)]
  my = [p for p in parts[1].split(" ") if len(p)]
  # print(win,my)
  points = 0

  for m in my:
    if m in win: points += 1

  print(i+1, points, 2**(points-1))
  part1 += math.floor(2**(points-1))

# advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
