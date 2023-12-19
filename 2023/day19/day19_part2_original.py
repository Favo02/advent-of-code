import advent

advent.setup(2023, 19)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = fin.read().split("\n\n")

part1 = 0
part2 = 0

workflows = {}
accepted = []
for line in lines[0].split("\n"):
  name, rules = line.split("{")
  rules = rules[:-1].split(",")
  workflows[name] = []

  for r in rules:
    tokens = r.split(":")
    workflows[name].append(tuple(tokens))
  print(name, workflows[name])

to_pos = { 'x': 0, 'm': 1, 'a': 2, 's': 3 }

def split(range, sign, pivot):
  s, e = range

  if sign == "<":
    ok = (s, pivot)
    no = (pivot, e)
  elif sign == ">":
    ok = (pivot+1, e)
    no = (s, pivot+1)
  else:
    assert False, "unknown sign"

  if ok[0] >= ok[1]:
    ok = (0,0)
  if no[0] >= no[1]:
    no = (0,0)
  return [ok, no]

ACCEPTED = []

def workflow(ranges, cur_workflow):
  # print(f"call work {ranges=} {cur_workflow}")

  if cur_workflow == 'A':
    if ranges:
      ACCEPTED.append(ranges)
    return
  elif cur_workflow == 'R':
    return

  keep = ranges
  for r in workflows[cur_workflow]:
    acc, rej, ke = rule(keep, r)
    # print(f"work {cur_workflow}, applied rule {r} -> {acc}, {rej}, {ke}")
    if acc:
      ACCEPTED.append(acc)
    if ke:
      keep = ke
    else:
      # print(f"breaking after {r}")
      break

def rule(ranges, r):
  # print(f"call rule {ranges=} {r=}")

  # single token: R, A or new rule
  if len(r) == 1:
    if r[0] == 'A':
      return [ ranges, None, None ]
    elif r[0] == 'R':
      return [ None, ranges, None ]
    else:
      workflow(ranges, r[0])
      return [ None, None, None ]

  # more tokens: expression to parse, redirection
  else:
    var = r[0][0]
    sign = r[0][1]
    num = int(r[0][2:])

    pos = to_pos[var]
    to_split = ranges[pos]
    rule_ok, rule_no = split(to_split, sign, num)

    newt = []
    for i in range(4):
      if i == pos:
        newt.append(rule_ok)
      else:
        newt.append(ranges[i])
    ranges_ok = tuple(newt)
    workflow(ranges_ok, r[1])

    newt = []
    for i in range(4):
      if i == pos:
        newt.append(rule_no)
      else:
        newt.append(ranges[i])
    ranges_no = tuple(newt)
    # print(f"call rule {ranges=} {r=} return {[ None, None, ranges_no ]}")
    return [ None, None, ranges_no ]

start = ((1, 4001), (1, 4001), (1, 4001), (1, 4001))
workflow(start, "in")
# print(ACCEPTED)

for acc in ACCEPTED:
  parz = 1
  for s,e in acc:
    parz *= (e-s)
  print(acc, "=", parz)
  part2 += parz

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)

# 256000000000000 <- max possible
# 167409079868000 <- correct
# 15320205000000  <- ding dong my algorithm is wrong
