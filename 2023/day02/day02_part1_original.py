import advent

advent.setup(2023, 2)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
# lines = advent.get_lines(fin)

part1 = 0
part2 = 0

maxx = {'blue':14,'green':13,'red':12}

def checkPart(str):
  d = {'blue':0,'green':0,'red':0}
  tokens = str.split(" ")
  print(tokens)
  tokens = [t for t in tokens if len(t) > 0]
  for i in range(0,len(tokens),2):
    k = tokens[i+1].replace(",", "")
    k = k.replace(";", "")

    d[k] += int(tokens[i])
  print(li+1, d)

  for k,v in d.items():
    if d[k] > maxx[k]:
      return False
  return True

days = []
for li,line in enumerate(fin):
  out = line.rstrip().split(": ")
  parts = out[1].split(";")

  for p in parts:
    if not checkPart(p):
      break
  else:
    part1 += li+1

advent.print_answer(1, part1)
advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
