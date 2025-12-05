import sys

ranges, ingrs = sys.stdin.read().strip().split("\n\n")

ran = []
for r in ranges.strip().split("\n"):
    a, b = map(int, r.split("-"))
    ran.append((a, b))

res = 0
for i in map(int, ingrs.strip().split("\n")):
    for a, b in ran:
        if a <= i <= b:
            break
    else:
        continue
    res += 1
print(res)
