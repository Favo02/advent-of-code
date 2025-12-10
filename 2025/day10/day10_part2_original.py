import sys
from collections import deque
from z3 import *

lines = list(map(lambda l: tuple(l.strip().split()), sys.stdin.read().strip().splitlines()))

# def solve(cur, buttons, target):
#     print(cur, target)
#     q = deque()
#     q.append((0, cur))
#     seen = set()
#     while q:
#         t, cur = q.popleft()
#         if cur == target:
#             print(t)
#             return t
#         if cur in seen:
#             continue
#         seen.add(cur)
#         if any(c > t for c,t in zip(cur, target)):
#             continue
#         for but in buttons:
#             next = list(cur)
#             for i in but:
#                 next[i] += 1
#             q.append((t+1, tuple(next)))
#     assert False

# b4 + b5 = 3
# b1 + b5 = 5
# b2 + b3 + b4 = 4
# b0 + b1 + b3 = 7

def solve(buttons, target):
    B = len(buttons)
    T = len(target)

    buts = [Int(f'b{i}') for i in range(B)]

    cost = Int('cost')
    opt = Optimize()

    for j in range(T):
        opt.add(Sum([buts[i] * (j in buttons[i]) for i in range(B)]) == target[j])

    for b in buts:
        opt.add(b >= 0)
    opt.add(cost == Sum(buts))
    opt.minimize(cost)

    opt.check()
    return opt.model()[cost].as_long()

res = 0
for l in lines:
    # print(l)
    buttons = [tuple(map(int, b[1:-1].split(","))) for b in l[1:-1]]
    target = tuple(map(int, l[-1][1:-1].split(",")))
    res += solve(buttons, target)
print(res)

