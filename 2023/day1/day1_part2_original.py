# https://adventofcode.com/2023/day/1
# https://github.com/Favo02/advent-of-code

import sys

def parseInput():
  lines = []
  for line in sys.stdin:
    lines.append(line.rstrip())
  return lines

lines = parseInput()
count = 0

rep = {"1":1,"2":2,"3":3,"4":4,"5":5,"6":6,"7":7,"8":8,"9":9,"0":0,"one":1, "two":2, "three":3, "four":4, "five":5, "six":6, "seven":7, "eight":8, "nine":9}

summ = 0
for l in lines:

  count = 0

  i,i2 = None,None
  f,f2 = 0,0
  for r in rep.keys():
    newi = l.find(r)
    newi2 = l.rfind(r)
    # print(l,r,newi,newi2)
    if (newi != -1) and (i == None or newi < i):
      i = newi
      f = r
    if (newi2 != -1) and (i2 == None or newi2 > i2):
      i2 = newi2
      f2 = r

  summ += int(f"{rep[f]}{rep[f2]}")

  print(i,i2,f,f2)
  # if len(d) >= 1:

print(summ)
