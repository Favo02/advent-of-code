import advent

advent.setup(2023, 14)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

part1 = 0
part2 = 0

matrix = []
for line in lines:
  matrix.append(line)

print(matrix)
for x in range(len(matrix[0])):
  multiplier = len(matrix)+1
  print()
  for y in range(len(matrix)):
    if matrix[y][x] == 'O':
      multiplier -= 1
      part1 += multiplier
      print("O" , x,y, multiplier)
    if matrix[y][x] == '#':
      multiplier = len(matrix) - y


advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
