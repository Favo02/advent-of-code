import sys

def valid(rep):
  if not ((sorted(rep)) == rep or sorted(rep, reverse=True) == rep):
    return False

  s = rep[0]
  for a in rep[1:]:
    if a == s : return False
    if abs(a-s) > 3: return False
    s = a
  return True

field = []
for y, line in enumerate(sys.stdin):
  line = line.strip()
  field.append(list(map(int, line.split())))
  for c, cell in enumerate(line):
    pass

res = 0
for report in field:
  print(report)

  if (valid(report)): res+=1

  print(report, valid(report))

print(res)
