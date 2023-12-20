import advent
from collections import deque

advent.setup(2023, 20)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
# fin = advent.get_input("easy_input2.txt")
lines = advent.get_lines(fin)

class Component:
  def __init__(self, name, subscribers, state=False):
    self.name = name
    self.subscribers = subscribers
    self.state = False
  def receive(self, fromm, signal):
    # print(f"{fromm} -({signal})> {self.name}")
    COUNTER[signal] += 1

class FlipFlop(Component):
  def receive(self, fromm, signal):
    # print(f"{fromm} -({signal})> {self.name}")
    COUNTER[signal] += 1
    if signal == False:
      self.state = not self.state
      for sub in self.subscribers:
        QUEUE.append(Signal(fromm=self.name, to=sub, value=self.state))

class Conjunction(Component):
  def __init__(self, name, subscribers, state):
    self.name = name
    self.subscribers = subscribers
    self.state = state
  def receive(self, fromm, signal):
    # print(f"{fromm} -({signal})> {self.name}")
    COUNTER[signal] += 1
    self.state[fromm] = signal
    if all(self.state.values()):
      for sub in self.subscribers:
        QUEUE.append(Signal(fromm=self.name, to=sub, value=False))
    else:
      for sub in self.subscribers:
        QUEUE.append(Signal(fromm=self.name, to=sub, value=True))

class Broadcaster(Component):
  def receive(self, fromm, signal):
    # print(f"{fromm} -({signal})> {self.name}")
    COUNTER[signal] += 1
    for sub in self.subscribers:
      QUEUE.append(Signal(fromm=self.name, to=sub, value=signal))

class Signal:
  def __init__(self, fromm, to, value):
    self.fromm = fromm
    self.to = to
    self.value = value

def printb():
  print("\n---")
  for b in BUTTONS.values():
    print(f"{b.name} ({type(b).__name__}): {b.state}")
  print("---\n")

part1 = 0
part2 = 0

BUTTONS = {}
temp_conj = {}

for line in lines:
  name, subscribers = line.split(" -> ")
  typee = name[0]
  name = name[1:]
  subscribers = subscribers.split(", ")

  if typee == "%": # ff
    BUTTONS[name] = FlipFlop(name, subscribers)
  elif typee == "&": # con
    temp_conj[name] = (subscribers, [])
  else: # broad
    BUTTONS["broadcaster"] = Broadcaster("broadcaster", subscribers)

  for sub in subscribers:
    if sub not in BUTTONS:
      BUTTONS[sub] = Component(sub, [], False)

for line in lines:
  name, subscribers = line.split(" -> ")
  typee = name[0]
  name = name[1:]
  subscribers = subscribers.split(", ")

  for con_name in temp_conj:
    if con_name in subscribers:
      temp_conj[con_name][1].append(name)

for name, (subs, adj) in temp_conj.items():
  BUTTONS[name] = Conjunction(name, subs, {a: False for a in adj})

# printb()

QUEUE = deque()
COUNTER = { False:0, True:0 }
PRESSES = 1000
for i in range(PRESSES):
  BUTTONS["broadcaster"].receive("HQ", False)
  while QUEUE:
    s = QUEUE.popleft()
    BUTTONS[s.to].receive(s.fromm, s.value)
  # printb()

print(COUNTER)
part1 = COUNTER[False] * COUNTER[True]

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
