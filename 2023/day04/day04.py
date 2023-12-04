# https://adventofcode.com/2023/day/4
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin
lines = []
for l in fin:
  lines.append(l.rstrip())

part1 = 0
part2 = 0

cardCopies = [1 for _ in lines]

for cardNum, card in enumerate(lines):
  card = card.rstrip()
  card = card[card.index(":")+2:]
  parts = card.split(" | ")

  winners = [int(p) for p in parts[0].split(" ") if len(p)]
  my = [int(p) for p in parts[1].split(" ") if len(p)]

  cardPoints = 0

  for m in my:
    if m in winners:
      cardPoints += 1

  if cardPoints > 0:
    part1 += 2**(cardPoints-1)

    for i in range(cardNum+1, cardNum+1+cardPoints):
      cardCopies[i] += cardCopies[cardNum]

part2 = sum(cardCopies)

print("Part 1:", part1)
print("Part 2:", part2)
