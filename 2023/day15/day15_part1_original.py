import advent

advent.setup(2023, 15)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

part1 = 0
part2 = 0

line = lines[0].strip().split(",")
print(line)

for t in line:
  cur = 0
  for c in t:
    cur += ord(c)
    cur *= 17
    cur %= 256
  part1 += cur

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
