import advent

advent.setup(2023, 5)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

part1 = 0

seeds = [int(n) for n in (lines[0][6:]).split() if len(n)]

categories = {}

CATS = ["seed", "soil", "fertilizer", "water", "light", "temperature", "humidity"]

key = None
for line in lines[2:]:
  if "map" in line:
    key = line.split("-")[0]
    categories[key] = {}
  elif len(line):
    tokens = [int(n) for n in line.split()]
    start, end = tokens[1], tokens[1] + tokens[2] -1
    result = tokens[0]-tokens[1]

    categories[key][(start,end)] = result

newseeds = []
for i in range(0, len(seeds), 2):
  start = seeds[i]
  end = seeds[i+1]
  newseeds.append((start, start+end-1))

seeds = newseeds

def split(seed, mapp):
  [sS, sE] = seed
  [mS, mE] = mapp
  res = []

  if sS < mS < sE:
    res.append((sS, mS-1))
    res.append((mS, sE))

  elif sS < mE < sE:
    res.append((sS, mE))
    res.append((mE+1, sE))

  elif mS <= sS < sE <= mE:
    res.append((sS, sE))

  elif sS < mS < mE < sE:
    res.append((sS, mS-1))
    res.append((mS, mE))
    res.append((mE+1, sE))

  else:
    res.append((sS, sE))

  return res

def norm(seeds, mappings):

  for m in mappings:
    i = 0
    while i < len(seeds):
      print(f"check {m} {seeds[i]}")
      splitRes = split(seeds[i], m)
      if len(splitRes) == 1:
        i += 1
      else:
        seeds = seeds[:i] + seeds[i+1:] + splitRes
      print(seeds)
  return seeds

print(seeds)
for step in CATS:
  mappings = categories[step]
  seeds = norm(seeds, mappings)

  for j,(seedS, seedE) in enumerate(seeds):
    for (start, end), res in mappings.items():
      if start <= seedS <= seedE <= end:
        seeds[j] = seedS+res, seedE+res

print(seeds)
part2 = min(seeds)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
