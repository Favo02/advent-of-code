# https://adventofcode.com/2023/day/15
# https://github.com/Favo02/advent-of-code

import sys
import re
from functools import reduce
from collections import OrderedDict
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def hash(string):
  return reduce(lambda acc, elem: ((acc + ord(elem)) * 17 % 256), string, 0)

part1 = 0
part2 = 0

tokens = fin.read().strip().split(",")
boxes = [OrderedDict() for _ in range(256)]

for token in tokens:
  part1 += hash(token)

  label, lens = re.split(r'[=-]', token)
  box = hash(label)

  # add/replace operation
  if '=' in token:
    boxes[box][label] = int(lens)
  # remove operation
  elif '-' in token and label in boxes[box]:
    del boxes[box][label]

for box, content in enumerate(boxes):
  for pos, lens in enumerate(content.values()):
    part2 += (box+1)*(pos+1)*lens

print("Part 1:", part1)
print("Part 2:", part2)
