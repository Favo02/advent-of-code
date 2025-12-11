import sys
from collections import deque, defaultdict
from functools import cache

lines = list(map(lambda l: l.strip().split(": "), sys.stdin.read().strip().splitlines()))

# def bfs(start, end):
#     res = 0
#     q = deque()
#     q.append((start, False, False, ""))
#     while q:
#         cur, fft, dac = q.popleft()
#         print(cur)
#         for adj in graph[cur]:
#             if adj == end:
#                 if fft and dac:
#                     res += 1
#             else:
#                 q.append((d+1, adj, fft or adj == "fft", dac or adj == "dac"))
#     return res

@cache
def dfs(cur, fft, dac):
    res = 0
    for adj in graph[cur]:
        if adj == "out":
            if fft and dac:
                res += 1
        else:
            next = (adj, fft or adj == "fft", dac or adj == "dac")
            res += dfs(*next)
    return res

graph = {}

for f, t in lines:
    t = t.split()
    graph[f] = t

# print(graph)
# print(bfs("svr", "out"))
print(dfs("svr", False, False))
