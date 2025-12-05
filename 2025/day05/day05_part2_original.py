import sys

ranges, ingrs = sys.stdin.read().strip().split("\n\n")

ran = []
for r in ranges.strip().split("\n"):
    a, b = map(int, r.split("-"))
    ran.append((a, b+1))

ran.sort()

l = []
for a, b in ran:
    l.append((a, +1, b))
    l.append((b, -1, a))

l.sort()

res = 0
open = float("inf")
close = 0
for me, what, other in l:
    # print("->", me, what, other)

    if what == 1:
        open = min(open, me)
        close = max(close, other)

    # print(open, close)

    if what == -1 and close == me and open != float("inf"):
        # print("eheh", open, close)
        res += close - open
        open = float("inf")
        close = 0

print(res)

