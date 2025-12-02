import sys

ranges = "".join(sys.stdin.readlines()).split(",")
ranges = [r.strip() for r in ranges]

res = 0

def invalid(q):
    return q[len(q)//2:] == q[:len(q)//2]

for r in ranges:
    a, b = map(int, r.split("-"))

    for q in range(a, b+1):
        if invalid(str(q)):
            res += q

print(res)
