import advent

advent.setup(2023, 10)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
# fin = advent.get_input("easy_input2.txt")
# fin = advent.get_input("easy_input3.txt")
# fin = advent.get_input("easy_input4.txt")
# fin = advent.get_input("easy_input5.txt")
fin = advent.get_input("easy_input6.txt")
lines = advent.get_lines(fin)

PIPES = {
  "|": [(0,-1), (0,+1)],
  "-": [(-1,0), (+1,0)],
  "L": [(0,-1), (+1,0)],
  "J": [(0,-1), (-1,0)],
  "7": [(0,+1), (-1,0)],
  "F": [(0,+1), (+1,0)]
}

ADJS = [(0,-1),(0,+1),(-1,0),(+1,0)]

def pointing_to(type, pipe, target):
  for dx,dy in PIPES[type]:
    if (pipe[0]+dx, pipe[1]+dy) == target:
      return True
  return False

def main_loop(matrix, start):
  path = [start]

  for dx,dy in ADJS:
    x,y = start[0]+dx, start[1]+dy
    if matrix[y][x] in PIPES and pointing_to(matrix[y][x], (x,y), start):
      path.append((start[0]+dx, start[1]+dy))
      # print(x,y, matrix[y][x])
      break

  cx,cy = path[-1]
  while (cx,cy) != start:
    for dx,dy in PIPES[matrix[cy][cx]]:
      if (cx+dx, cy+dy) != path[-2]:
        path.append((cx+dx, cy+dy))
        break
    cx,cy = path[-1]

  return path[:-1]

def bfs(matrix,start):
  queue = [start]
  while queue:
    cx,cy = queue.pop()
    for dx,dy in ADJS:
      if 0 <= cy+dy < len(matrix) and 0 <= cx+dx < len(matrix[0]):
        if matrix[cy+dy][cx+dx] == ".":
          matrix[cy+dy][cx+dx] = "_"
          queue.append((cx+dx,cy+dy))
  return matrix

def expand(matrix):
  new_matrix = []
  HOR = ["L", "F", "-"]
  VER = ["7", "F", "|"]
  for r in matrix:
    row = []
    row2 = []
    for c in r:
      row.append(c)
      row.append("-" if c in HOR else ".")
      row2.append("|" if c in VER else ".")
      row2.append(".")
    new_matrix.append(row)
    new_matrix.append(row2)
  return new_matrix

def shrink(matrix):
  new_matrix = []
  for r in matrix[::2]:
    row = []
    for c in r[::2]:
      row.append(c)
    new_matrix.append(row)
  return new_matrix

def printm(matrix):
  for l in matrix:
    for c in l:
      print(c, end="")
    print()
  print()

part1 = 0
part2 = 0

matrix = []
# devo rimpiazzare a mano S con il tubo giusto perchè altrimenti si apre una voragine
# che mangia tutto quando espando la matrice (dato che S non so dove espanderlo)
# e quindi devo hardcodare la start
# start = (88, 128) # input vero
# start = (12, 4) # easyinput4.txt
start = (4,0) # easyinput6.txt
# start = (1,1) # easyinput3.txt
for y, line in enumerate(lines):
  if "S" in line:
    start = (line.index("S"), y)
  matrix.append(line)

# ===========================
# calcolare path che di tubi connessa ad S (parte1)
print(start)
path = main_loop(matrix, start)
part1 = len(path)//2

printm(matrix)

# ========================================
# togliere i tubi non usati (mettendoci #)
new_matrix = []
for y in range(len(matrix)):
  new_matrix.append(["#"] * (len(matrix[0])+2))
  new_matrix[y][0] = "."
  new_matrix[y][-1] = "."
  for x in range(len(matrix[y])):
    if (x,y) not in path:
      new_matrix[y][x+1] = "."
    if (x,y) in path:
      new_matrix[y][x+1] = matrix[y][x]
new_matrix.insert(0,["."] * (len(matrix[0])+2))
new_matrix.append(["."] * (len(matrix[0])+2))

printm(new_matrix)

# =======================================
# espandere matrice, ogni riga diventano 2, ogni colonna diventano 2
# per far passare la bfs anche tra i tubi appiccicati (ma non "contigui")
expanded = expand(new_matrix)

printm(expanded)

# =======================================
# bfs: controllare quali putni sono "landlocked" e quali no
# grazie all'espansione si passa anche dove i tubi sono "stretti"
# per questa operazione va rimpiazzata a mano la S nell'input
new_matrix = bfs(expanded, (0,0))
printm(new_matrix)

# =======================================
# inverso dell'expand, rimpicciolire la matrice alla size originale
# scartando le righe dispari
shrinked = shrink(new_matrix)
printm(shrinked)

# contare punti :))))
for row in shrinked:
  for c in row:
    if c == ".":
      part2 += 1

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
