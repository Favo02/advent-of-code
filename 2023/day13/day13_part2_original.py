import advent

advent.setup(2023, 13)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
# lines = advent.get_lines(fin)

def printm(matrix):
  for r in matrix:
    print("".join(r))
  print()

def smudge(matrix):
  before = check(matrix)
  for row in matrix:
    for x,c in enumerate(row):
      if c == "#": row[x] = "."
      else: row[x] = "#"

      res = check(matrix, before)
      # print(res)
      # printm(matrix)

      if res > 0 and res != before: return res
      else: row[x] = c

  assert False, f"no smudge {before} {matrix}"

def check(matrix, ignore=-1):
  # print(len(matrix), len(matrix[0]), len(matrix) * len(matrix[0]))
  for p in range(1, len(matrix[0])):
    if is_specular_h(matrix, p):
      if p == ignore: continue
      return p

  for p in range(1, len(matrix)):
    if is_specular_v(matrix, p):
      if p*100 == ignore: continue
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
    part2 += smudge(matrix)
    matrix = []
  else:
    matrix.append(list(line.rstrip()))
part2 += smudge(matrix)


# advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
