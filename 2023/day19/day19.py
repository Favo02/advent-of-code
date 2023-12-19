# https://adventofcode.com/2023/day/19
# https://github.com/Favo02/advent-of-code

import sys
import re
from math import prod
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def parse_workflows(workflows):
  result = {}
  for wf in workflows.split("\n"):
    name, rules = wf.split("{")
    rules = rules[:-1].split(",")
    result[name] = []

    for r in rules:
      tokens = r.split(":")
      result[name].append(tuple(tokens))
  return result

# split an interval into ok and no parts:
# part that is valid (to the sign, limit)
# and part that is valid
def split_interval(inter, sign, limit):
  s, e = inter
  if sign == "<":
    ok = (s, limit)
    no = (limit, e)
  elif sign == ">":
    no = (s, limit+1)
    ok = (limit+1, e)
  else:
    assert False, "unknown sign"
  return [ok, no]

# runs a workflow, modifying global variable ACCEPTED
def run_workflow(inters, workflow):
  rules = WORKFLOWS.get(workflow, [ workflow ])

  forward = inters
  for rule in rules:
    # accepted, foprward to next rule
    forward = run_rule(forward, rule)
    if not forward:
      break

# runs a workflow, modifying global variable ACCEPTED
# returns the ranges to forward to next rule
def run_rule(ranges, rule):
  # single token: R, A or new rule
  if len(rule) == 1:
    if rule[0] == 'A':
      ACCEPTED_RANGES.append(ranges)
    elif rule[0] != 'R':
      run_workflow(ranges, rule[0])
    return None

  # more tokens: expression to parse, redirection
  else:
    criteria, target = rule
    var = criteria[0]
    sign = criteria[1]
    num = int(criteria[2:])

    pos = VAR_TO_INDEX[var]
    to_split = ranges[pos]
    accepted, to_forward = split_interval(to_split, sign, num)

    ranges_ok = [ranges[i] if i != pos else accepted for i in range(4)]
    run_workflow(ranges_ok, target)

    ranges_no = [ranges[i] if i != pos else to_forward for i in range(4)]
    return ranges_no

workflows, parts = fin.read().split("\n\n")
WORKFLOWS = parse_workflows(workflows)

part1 = 0
part2 = 0

VAR_TO_INDEX = { 'x': 0, 'm': 1, 'a': 2, 's': 3 }
ACCEPTED_RANGES = []

start = [(1, 4001), (1, 4001), (1, 4001), (1, 4001)]
run_workflow(start, "in")

# part1
parts = [p for p in parts.split("\n") if p]
for part in parts:
  match = re.match(r"\{x=(\d+),m=(\d+),a=(\d+),s=(\d+)\}", part)
  assert match
  vals = tuple(map(int, match.groups()))

  for acc in ACCEPTED_RANGES:
    if all(vals[i] >= s and vals[i] < e for i,(s,e) in enumerate(acc)):
      part1 += sum(vals)

# part2
part2 = sum(prod(s-e for s,e in interval) for interval in ACCEPTED_RANGES)

print("Part 1:", part1)
print("Part 2:", part2)
