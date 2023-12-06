import advent

advent.setup(2023, 6)
# fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
# lines = advent.get_lines(fin)

part1 = 1
part2 = 0

time = 47707566
record = 282107911471062

# time = 71530
# record = 940200

for t in range(time):
  if t % 10**6 == 0: print(t)
  if t*(time-t) > record:
    part2 += 1


# advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
