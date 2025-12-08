import sys
from math import sqrt
from functools import reduce

nums = list(map(lambda l: tuple(map(int, l.strip().split(","))), sys.stdin.read().strip().splitlines()))

P = [i for i in range(len(nums))]
S = [1 for i in range(len(nums))]

def find(n):
    if P[n] == n:
        return n
    P[n] = find(P[n])
    return P[n]

def union(a, b):
    A = find(a)
    B = find(b)
    if A == B:
        return
    if S[A] < S[B]:
        A, B = B, A
    S[A] += S[B]
    S[B] = 0
    P[B] = A

used = set()
for t in range(1000):

    closest = (float("inf"), None, None)

    for ai, (ax, ay, az) in enumerate(nums):
        for bi, (bx, by, bz) in enumerate(nums):
            if ai == bi: continue
            if (ai, bi) in used: continue
            if (bi, ai) in used: continue
            dist = sqrt((ax-bx)**2 + (ay-by)**2 + (az-bz)**2)
            closest = min(closest, (dist, ai, bi))

    used.add((closest[1], closest[2]))
    if closest[0] == float("inf"):
        print("what")
        continue

    # print(t)
    union(closest[1], closest[2])
    # print(closest[0], nums[closest[1]], nums[closest[2]])

# print(S)
SS = sorted(S, reverse=True)
print(reduce(lambda a, b: a*b, SS[:3], 1))

