import sys
from collections import deque

lines = list(map(lambda l: tuple(l.strip().split()), sys.stdin.read().strip().splitlines()))

def solve(cur, buttons, target):
    q = deque()
    q.append((0, cur))
    seen = set()
    while q:
        t, cur = q.popleft()
        # print(t, bin(cur))
        if cur == target:
            # print(t)
            return t
        if cur in seen:
            continue
        seen.add(cur)
        for but in buttons:
            next = cur
            for i in but:
                next ^= (1 << i)
            q.append((t+1, next))
    assert False

res = 0
for l in lines:
    cur = 0
    for i, c in enumerate(l[0][1:-1]):
        if c == "#":
            cur |= (1 << i)
    buttons = [tuple(map(int, b[1:-1].split(","))) for b in l[1:-1]]
    res += solve(0, buttons, cur)
print(res)
