from sys import stdin

field = []

lefts = []
rights = []

for r, line in enumerate(stdin):
  line = line.strip()
  field.append(line)

  r, l = line.split("  ")
  lefts.append(int(l))
  rights.append(int(r))

  for c, cell in enumerate(line):
    pass

lefts.sort()
rights.sort()

rs = 0
for l, r in zip(lefts, rights):
  print(abs(l - r))
  rs += abs(l - r)

print(rs)
