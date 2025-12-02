import sys

ranges = "".join(sys.stdin.readlines()).split(",")
ranges = [r.strip() for r in ranges]

res = 0

def invalid(q):
    for size in range(1, len(q)):
        if len(q) % size != 0: continue
        first = q[:size]
        valid = True
        for starts in range(0, len(q), size):
            if q[starts:starts+size] != first:
                valid = False
                break
        if valid:
            return True
    return False

for r in ranges:
    a, b = map(int, r.split("-"))

    for q in range(a, b+1):
        if invalid(str(q)):
            res += q

print(res)
