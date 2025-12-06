import sys
from functools import reduce

lines = sys.stdin.read().strip().splitlines()

tokens = []
ops = None

for l in lines:
    l = list(filter(lambda e: len(e) > 0, l.split(" ")))
    # print(l)
    if l[0] in "+-*/":
        ops = l
    else:
        tokens.append(list(map(int, l)))

# print(tokens)
# print(ops)

opss = {
    "*": (lambda a, b: a * b, 1),
    "+": (lambda a, b: a + b, 0)
}

res = 0
for i, op in enumerate(ops):
    lam, start = opss[op]

    res += reduce(lam, [t[i] for t in tokens], start)
    # print(i, reduce(lam, [t[i] for t in tokens], start))

print(res)
