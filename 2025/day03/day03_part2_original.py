import sys

lines = sys.stdin.readlines()

res = 0

ans = 0
for l in lines:
    cur = 0
    l = list(map(int, list(l.strip())))

    res = []
    start = 0
    for _ in range(12):
        cur = (0, 0)
        for j, n in enumerate(l):
            if j < start: continue
            if j >= len(l)-11+len(res):
                break
            cur = max(cur, (n, -j))
        print(l[start:], cur)
        res.append(cur[0])
        start = -cur[1] + 1


    # print(res)
    print("".join(map(str, res)))
    ans += int("".join(map(str, res)))

print(ans)
