import sys
from collections import deque

lines = list(map(lambda l: l.strip().split(": "), sys.stdin.read().strip().splitlines()))

def bfs(start, end):
    res = 0
    q = deque()
    q.append(start)
    while q:
        cur = q.popleft()
        for adj in graph[cur]:
            if adj == end:
                res += 1
            else:
                q.append(adj)
    return res


graph = {}

for f, t in lines:
    t = t.split()
    graph[f] = t

# print(graph)
print(bfs("you", "out"))
