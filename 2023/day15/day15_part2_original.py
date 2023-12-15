import advent
from collections import OrderedDict

advent.setup(2023, 15)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

def hash(string):
  cur = 0
  for c in string:
    cur += ord(c)
    cur *= 17
    cur %= 256
  return cur

part1 = 0
part2 = 0

line = lines[0].strip().split(",")

boxes = [OrderedDict() for _ in range(256)]

for t in line:
  label = None
  lens = None
  remove_op = None
  if '=' in t:
    label = t.split("=")[0]
    remove_op = False
    lens = int(t.split("=")[1])
  else:
    label = t.split("-")[0]
    remove_op = True

  box = hash(label)
  print(t, label, remove_op, lens)

  if remove_op:
    if label in boxes[box]:
      del boxes[box][label]
  else:
    boxes[box][label] = lens
  # print(boxes[box])


for box, content in enumerate(boxes):
  val = 1+box
  for i, len in enumerate(content.values()):
    # print(val, i, len)
    part2 += val*(i+1)*len

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
