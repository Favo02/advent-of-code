import advent

advent.setup(2023, 13)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
# lines = advent.get_lines(fin)

def check(matrix):
  print()
  for p in range(1, len(matrix[0])):
    print("p",p)
    if is_specular_h(matrix, p):
      print("SPECX", p)
      return p

  for p in range(1, len(matrix)):
    if is_specular_v(matrix, p):
      print("SPECY", p)
      return p*100

  return 0

def is_specular_h(matrix, pivot):
  for row in matrix:
    left = row[:pivot]
    right = row[pivot:]
    L = min(len(left), len(right))

    if left[::-1][:L] != right[:L]:
      return False

  return True

def is_specular_v(matrix, pivot):
  matrix = [[row[x] for row in matrix] for x in range(len(matrix[0]))]
  return is_specular_h(matrix, pivot)

part1 = 0
part2 = 0

matrix = []
for line in fin:
  if line == "\n":
    part1 += check(matrix)
    matrix = []
  else:
    matrix.append(line.rstrip())
part1 += check(matrix)


advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
