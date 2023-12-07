# https://adventofcode.com/2023/day/7
# https://github.com/Favo02/advent-of-code

import sys
from enum import Enum
from typing import Dict
from collections import Counter
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = 0
part2 = 0

CARDS_P1 = "AKQJT98765432"
CARDS_P2 = "AKQT98765432J"

class Hands(Enum):
  FIVE_OF_A_KIND =  6
  FOUR_OF_A_KIND =  5
  FULL_HOUSE =      4
  THREE_OF_A_KIND = 3
  TWO_PAIR =        2
  ONE_PAIR =        1
  HIGH_CARD =       0

class Checkers:
  def __init__(self, occs: Dict[str, int]):
    self.occs = list(occs.values())

  fiveChecker =   lambda self: 5 in self.occs
  fourChecker =   lambda self: 4 in self.occs
  fullChecker =   lambda self: 3 in self.occs and 2 in self.occs
  threeChecker =  lambda self: 3 in self.occs
  twoChecker =    lambda self: self.occs.count(2) == 2
  oneChecker =    lambda self: 2 in self.occs

def chain(hand, part2):
  checker = Checkers(countCards(hand, part2))

  if   checker.fiveChecker():   return Hands.FIVE_OF_A_KIND
  elif checker.fourChecker():   return Hands.FOUR_OF_A_KIND
  elif checker.fullChecker():   return Hands.FULL_HOUSE
  elif checker.threeChecker():  return Hands.THREE_OF_A_KIND
  elif checker.twoChecker():    return Hands.TWO_PAIR
  elif checker.oneChecker():    return Hands.ONE_PAIR
  return                               Hands.HIGH_CARD

def countCards(hand, part2):
  occs = Counter(hand)
  jokers = occs.get("J", 0)

  if part2 and jokers == 5:
    return {"A": 5}

  if part2 and jokers > 0:
    del occs["J"]
    mostOccsCard = max(occs.items(), key=lambda x: x[1])[0]
    occs[mostOccsCard] += jokers
  return occs

hands = []
for line in fin:
  hand, bet = line.rstrip().split()
  hands.append((hand, int(bet)))

p1_hands = sorted(hands, key=lambda h: [
  chain(h[0], False).value,
  [-CARDS_P1.index(c) for c in h[0]]
])

p2_hands = sorted(hands, key=lambda h: [
  chain(h[0], True).value,
  [-CARDS_P2.index(c) for c in h[0]]
])

for i, ((_, bet1), (_, bet2)) in enumerate(zip(p1_hands, p2_hands)):
  part1 += bet1*(i+1)
  part2 += bet2*(i+1)

print("Part 1:", part1)
print("Part 2:", part2)
