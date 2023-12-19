import advent
import re

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

lines[1] = lines[1][:-1]
for mac in lines[1].split("\n"):
  match = re.match(r"\{x=(\d+),m=(\d+),a=(\d+),s=(\d+)\}", mac)
  vals = tuple(map(int, match.groups()))
  vals = {'x': vals[0], 'm': vals[1], 'a': vals[2], 's':vals[3]}

  cur_rule = "in"
  accepted = None
  while True:
    for rule in workflows[cur_rule]:

      # single token: R, A or new rule
      if len(rule) == 1:
        if rule[0] == 'R':
          accepted = False
          break
        elif rule[0] == 'A':
          accepted = True
          break
        else:
          cur_rule = rule[0]
          break
      # more tokens: expression to parse, redirection
      else:
        var = rule[0][0]
        sign = rule[0][1]
        num = int(rule[0][2:])
        this_rule = False

        if sign == '>':
          if vals[var] > num:
            this_rule = True
        elif sign == '<':
          if vals[var] < num:
            this_rule = True
        else:
          assert False, "unknown sign"

        if this_rule:
          if rule[1] == 'R':
            accepted = False
            break
          elif rule[1] == 'A':
            accepted = True
            break
          else:
            cur_rule = rule[1]
            break

    if accepted != None:
      if accepted:
        print("A", mac)
        part1 += sum(vals.values())
      else:
        print("R", mac)
      break

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
