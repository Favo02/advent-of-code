import advent

advent.setup(2023, 7)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

part1 = 0
part2 = 0

CARDS = ["A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"]

def fiveChecker(occs):
  for v in occs.values():
    if v == 5:
      return True
  return False

def fourChecker(occs):
  for v in occs.values():
    if v == 4:
      return True
  return False

def fullChecker(occs):
  vals = occs.values()
  if 3 in vals and 2 in vals:
    return True
  return False

def threeChecker(occs):
  for v in occs.values():
    if v == 3:
      return True
  return False

def twopairChecker(occs):
  pairs = 0
  for v in occs.values():
    if v == 2:
      pairs += 1
  return pairs == 2

def onepairChecker(occs):
  pairs = 0
  for v in occs.values():
    if v >= 2:
      pairs += 1
  return pairs == 1

def chain(occs):
  value = 0
  if fiveChecker(occs): value = 6
  elif fourChecker(occs): value = 5
  elif fullChecker(occs): value = 4
  elif threeChecker(occs): value = 3
  elif twopairChecker(occs): value = 2
  elif onepairChecker(occs): value = 1
  return value

def tieBraker(h1, h2):
  for c1, c2 in zip(h1, h2):
    # print(c1, c2, f" {CARDS.index(c1)}, {CARDS.index(c2)}")
    if CARDS.index(c1) == CARDS.index(c2):
      continue
    if CARDS.index(c1) < CARDS.index(c2):
      return True
    else: return False
  return False

def countCards(hand):
  occs = {k:hand.count(k) for k in hand}
  Js = occs.get('J', 0)

  if Js == 5:
    return {"A": 5}

  if Js: del occs['J']

  while Js:
    maxK = list(occs.keys())[0]
    for k,v in occs.items():
      if v > occs[maxK]:
        maxK = k
    occs[maxK] += 1
    Js -= 1
  return occs

def comparator(hand1, hand2):
  h1 = hand1[0]
  h2 = hand2[0]

  v1 = chain(countCards(h1))
  v2 = chain(countCards(h2))

  if v1 == v2:
    # print(f"tie {h1} {h2}, winner = {tieBraker(h1, h2)}")
    return 1 if tieBraker(h1, h2) else -1

  return v1-v2

hands = []
for line in lines:
  t = line.split()
  hands.append((t[0], int(t[1])))
  count = countCards(hands[-1][0])

from functools import cmp_to_key
hands = sorted(hands, key=cmp_to_key(comparator))

for i,(h,bet) in enumerate(hands):
  part2 += bet*(i+1)

# advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part1)
