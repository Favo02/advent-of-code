import sys

lines = sys.stdin.read().strip().splitlines()

tokens = []
ops = []

for x in range(max((len(l) for l in lines))):
    cur = []
    for y in range(len(lines)):
        c = "" if x >= len(lines[y]) else lines[y][x]
        if c != "" and c != " ":
            if c in "*+":
                ops.append(c)
            else:
                cur.append(int(c))
    tokens.append(tuple(cur))

# print(tokens)
# print(ops)

opss = {
    "*": (lambda a, b: a * b, 1),
    "+": (lambda a, b: a + b, 0)
}

res = 0
i = 0
for op in ops:
    o, cur = opss[op]
    # print("---", o, cur)
    while i < len(tokens) and len(tokens[i]) > 0:
        # print(int("".join(map(str, tokens[i]))))
        cur = o(cur, int("".join(map(str, tokens[i]))))
        i += 1
    # print(cur)
    res += cur
    i += 1

print(res)
