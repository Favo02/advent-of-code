# https://adventofcode.com/2023/day/20
# https://github.com/Favo02/advent-of-code

import sys
from collections import deque
from math import lcm
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

class Component:
  def __init__(self, name, subs, state):
    self.name = name
    self.subs = subs
    self.state = state
  def receive(self, _, signal):
    COUNTER[signal] += 1
  def notify_subs(self, value):
    for sub in self.subs:
      QUEUE.append(Signal(fromm=self.name, to=sub, value=value))

class FlipFlop(Component):
  def receive(self, _, signal):
    COUNTER[signal] += 1
    if not signal:
      self.state = not self.state
      self.notify_subs(self.state)

class Conjunction(Component):
  def receive(self, fromm, signal):
    COUNTER[signal] += 1
    self.state[fromm] = signal
    if all(self.state.values()):
      self.notify_subs(False)
    else:
      self.notify_subs(True)

class Broadcaster(Component):
  def receive(self, _, signal):
    COUNTER[signal] += 1
    self.notify_subs(signal)

class Signal:
  def __init__(self, fromm, to, value):
    self.fromm = fromm
    self.to = to
    self.value = value

def parse_input():
  parsed = []
  for line in fin:
    line = line.rstrip()
    name, subs = line.split(" -> ")
    typee = name[0]
    name = name[1:]
    subs = subs.split(", ")
    parsed.append((typee, name, subs))

  buttons = {}
  conjunctions_to_build = {}

  for typee, name, subs in parsed:
    if typee == "%": # flipflop
      buttons[name] = FlipFlop(name, subs, False)
    elif typee == "&": # conjunction
      conjunctions_to_build[name] = (subs, [])
    else: # broadcast
      buttons["broadcaster"] = Broadcaster("broadcaster", subs, False)

  for _, name, subs in parsed:
    # add subscriptions to conjunctions
    for conj_name in conjunctions_to_build:
      if conj_name in subs:
        conjunctions_to_build[conj_name][1].append(name)
    # add generic components (only if a specific component doesnt exists)
    for sub in subs:
      if sub not in buttons:
        buttons[sub] = Component(sub, [], False)

  # build conjunctions
  for name, (subs, adj) in conjunctions_to_build.items():
    buttons[name] = Conjunction(name, subs, {a: False for a in adj})
  return buttons

def previous(button):
  prevs = []
  for but in BUTTONS.values():
    if button in but.subs:
      prevs.append(but.name)
  assert prevs, "Cannot find previous of {button}"
  return prevs

part1 = 0
part2 = 0

BUTTONS = parse_input()

conjunctions_to_rx = [name for name, con in BUTTONS.items()
                      if type(con) == Conjunction and
                      all(type(BUTTONS[p]) == FlipFlop for p in previous(name))]

cycles = { c: 0 for c in conjunctions_to_rx }

QUEUE = deque()
COUNTER = {True: 0, False: 0}

time = 0
while any(c == 0 for c in cycles.values()):
  time += 1
  BUTTONS["broadcaster"].receive("button", False)

  while QUEUE:
    s = QUEUE.popleft()
    BUTTONS[s.to].receive(s.fromm, s.value)
    if (s.fromm in conjunctions_to_rx) and (not s.value) and (not cycles[s.fromm]):
      cycles[s.fromm] = time

  if time == 1000:
    part1 = COUNTER[True] * COUNTER[False]
    print("Part 1:", part1)
    assert "rx" in BUTTONS, "Part2 cannot be calculated: rx does not exist"

part2 = lcm(*cycles.values())
print("Part 2:", part2)
