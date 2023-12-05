import advent

advent.setup(2023, 5)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

part1 = 0
part2 = 0

seeds = [int(n) for n in (lines[0][6:]).split() if len(n)]

categories = {}

cats = ["seed", "soil", "fertilizer", "water", "light", "temperature", "humidity"]

key = None
for line in lines[2:]:
  if "map" in line:
    key = line.split("-")[0]
    categories[key] = {}
  elif len(line):
    tokens = [int(n) for n in line.split()]
    start, end = tokens[1], tokens[2]
    result = tokens[0]

    categories[key][(start,end)] = result

for i,c in enumerate(cats):
  for j,s in enumerate(seeds):
    for (start, end), res in categories[c].items():
      if start <= s < start+end:
        seeds[j] = res + (s - start)

part1 = min(seeds)

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

# advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
