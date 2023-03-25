import sys

# returns a list of strings (input)
def parseInput():
  lines = []
  for line in sys.stdin:
    lines.append(line.rstrip())
  return lines

# converts list of strings to int matrix
# returns an int matrix
def convertToMatrix(lines):
  matrix = []
  for i, line in enumerate(lines):
    matrix.append([])
    for c in line:
      matrix[i].append(int(c))
  return matrix

# runs a step (increment every point and flash) on the matrix
def runStep(matrix):
  count = 0
  toflash = []

  # increment every point and save the points that flashes
  for y, _ in enumerate(matrix):
    for x, _ in enumerate(matrix[y]):
      matrix[x][y]+=1
      if (matrix[x][y] > 9):
        toflash.append((x, y))
  
  # flash points
  for xy in toflash:
    flash(matrix, xy[0], xy[1])

  # reset points that flashed
  for y, _ in enumerate(matrix):
    for x, _ in enumerate(matrix[y]):
      if (matrix[x][y] > 9):
        count+=1
        matrix[x][y] = 0
  
  return count

# flashes a point (increment all adjacent points)
def flash(matrix, x, y):
  # print("flash", x, y)
  # printMatrix(matrix)

  for iy in range(max(0, y-1), min(len(matrix), y+2)):
    for ix in range(max(0, x-1), min(len(matrix[iy]), x+2)):
      # print("inc", ix, iy)
      matrix[ix][iy]+=1
      if (matrix[ix][iy] == 10):
        flash(matrix, ix, iy)

# prints the matrix in a decent format
def printMatrix(matrix):
  for line in matrix:
    for n in line:
      # print(str(n).ljust(2), end="")
      print(n, end="")
    print()
  print()



lines = parseInput()
matrix = convertToMatrix(lines)

STEPS = 100
count = 0
allflash = 0

# printMatrix(matrix)
i = 0
while True:
  stepCount = runStep(matrix)
  
  if (i < STEPS):
    count += stepCount
  
  if (stepCount == len(matrix) * len(matrix[0])):
    allflash = i+1
    break

  # printMatrix(matrix)
  i+=1

print("part1:", count)
print("part2:", allflash)
