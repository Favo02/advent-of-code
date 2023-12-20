import advent
from collections import deque
from math import lcm

advent.setup(2023, 20)
fin = advent.get_input()
lines = advent.get_lines(fin)

class Component:
  def __init__(self, name, subscribers, state=False):
    self.name = name
    self.subscribers = subscribers
    self.state = False
  def receive(self, fromm, signal):
    # print(f"{fromm} -({signal})> {self.name}")
    # COUNTER[signal] += 1
    pass

class FlipFlop(Component):
  def receive(self, fromm, signal):
    # print(f"{fromm} -({signal})> {self.name}")
    # COUNTER[signal] += 1
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
    # COUNTER[signal] += 1
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
    # COUNTER[signal] += 1
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

# ----------------------- PARSE INPUT
BUTTONS = {}
conjunctions_to_build = {}

for line in lines:
  name, subscribers = line.split(" -> ")
  typee = name[0]
  name = name[1:]
  subscribers = subscribers.split(", ")

  if typee == "%": # ff
    BUTTONS[name] = FlipFlop(name, subscribers)
  elif typee == "&": # con
    conjunctions_to_build[name] = (subscribers, [])
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

  for conj_name in conjunctions_to_build:
    if conj_name in subscribers:
      conjunctions_to_build[conj_name][1].append(name)

for name, (subs, adj) in conjunctions_to_build.items():
  BUTTONS[name] = Conjunction(name, subs, {a: False for a in adj})
# ----------------------- /PARSE INPUT

conj_to_rx = ["tq", "pf", "kx", "rj"]
cycles = { c: -1 for c in conj_to_rx }

QUEUE = deque()
time = 0
while any(c == -1 for c in cycles.values()):
  time += 1
  BUTTONS["broadcaster"].receive("HQ", False)
  while QUEUE:
    s = QUEUE.popleft()

    if s.fromm in conj_to_rx and not s.value and cycles[s.fromm] == -1:
      cycles[s.fromm] = time

    BUTTONS[s.to].receive(s.fromm, s.value)

part2 = lcm(*cycles.values())

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)

#    224602953547789 giusta :)))))))))))))))))))
#    562949953421312 too high
#    562949953421312 too high
# 288230376151711744 too high
# 392624287792451301217568538230784000 :)
