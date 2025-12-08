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

dists = []
for ai, (ax, ay, az) in enumerate(nums):
    for bi, (bx, by, bz) in enumerate(nums):
        if ai == bi: continue
        dist = sqrt((ax-bx)**2 + (ay-by)**2 + (az-bz)**2)
        dists.append((dist, ai, bi))

dists.sort()
# print(dists)

i = 0
while len([s for s in S if s == 0]) < len(nums)-1:

    # print(len([s for s in S if s == 0]), len(nums)-1)
    union(dists[i][1], dists[i][2])
    i+=1

print(nums[dists[i][1]][0] * nums[dists[i][2]][0])
