# https://adventofcode.com/2024/day/5
# https://github.com/Favo02/advent-of-code

from collections import defaultdict, deque
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

ordering = defaultdict(set)
updates = []

for line in fin:
  if line == "\n": break
  a, b = map(int, line.split("|"))
  ordering[b].add(a)

for line in fin:
  updates.append(list(map(int, line.split(","))))

def is_valid(update):
  banned = set()
  for elem in update:
    if elem in banned:
      return False
    banned.update(ordering[elem])
  return True

def fix_order(update):
  final = []
  all_elems = set(update)

  queue = deque(update)
  while queue:
    cur = queue.popleft()
    if cur not in all_elems: continue

    valid = True
    for bef in (o for o in ordering[cur] if o in all_elems):
      queue.append(bef)
      valid = False

    if valid:
      all_elems.discard(cur)
      final.append(cur)
    else:
      queue.append(cur)

  return list(final)

part1 = sum(u[len(u)//2] for u in updates if is_valid(u))
part2 = sum(fix_order(u)[len(u)//2] for u in updates if not is_valid(u))

print("Part 1:", part1)
print("Part 2:", part2)
