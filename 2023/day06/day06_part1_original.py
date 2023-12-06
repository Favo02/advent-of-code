import advent

advent.setup(2023, 6)
# fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
# lines = advent.get_lines(fin)

part1 = 1
part2 = 0

time = [47, 70, 75, 66]
record = [282, 1079, 1147, 1062]

# time = [7, 15, 30]
# record = [9, 40, 200]

for i, t in enumerate(time):
  count = 0
  for inc in range(t):
    if inc*(t-inc) > record[i]:
      count += 1
  part1 *= count


advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

# advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
