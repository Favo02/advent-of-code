# https://adventofcode.com/2021/day/21
# https://github.com/Favo02/advent-of-code

pa = int(input().split(": ")[1])-1
pb = int(input().split(": ")[1])-1

sa = sb = 0

dice_rolls = 0
dice_val = -1

def dice():
  global dice_val
  global dice_rolls

  dice_val = (dice_val + 1) % 100
  dice_rolls += 1
  return dice_val+1

while True:
  pa = (pa + dice()) % 10
  pa = (pa + dice()) % 10
  pa = (pa + dice()) % 10
  sa += pa+1
  if sa >= 1000:
    print(dice_rolls * sb)
    break

  pb = (pb + dice()) % 10
  pb = (pb + dice()) % 10
  pb = (pb + dice()) % 10
  sb += pb+1
  if sa >= 1000:
    print("Part 1:", dice_rolls * sa)
    break
