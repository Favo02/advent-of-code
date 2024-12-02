from sys import stdin
from collections import Counter

field = []

lefts = []
rights = []

for r, line in enumerate(stdin):
  line = line.strip()
  field.append(line)

  l, r = line.split("  ")
  lefts.append(int(l))
  rights.append(int(r))

  for c, cell in enumerate(line):
    pass

count = Counter(rights)

print(count)

rs = 0
for l in lefts:
  rs += l * count[l]
  print(l, count[l])

print(rs)
