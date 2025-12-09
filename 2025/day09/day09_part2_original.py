import sys
from shapely.geometry import Point
from shapely.geometry.polygon import Polygon

lines = list(map(lambda l: tuple(map(int, l.strip().split(","))), sys.stdin.read().strip().splitlines()))

fullpoly = Polygon(list(map(lambda x: Point(x[0], x[1]), lines)))
# print(fullpoly)

res = 0
for a, (xa, ya) in enumerate(lines):
    for b, (xb, yb) in enumerate(lines):
        if a == b: continue
        # print((xa, ya), (xb, yb))

        if xb < xa: continue

        o1x, o1y = xa, yb
        o2x, o2y = xb, ya

        # if not (polygon.touches(Point(xa, ya)) or polygon.contains(Point(xa, ya))):
        #     continue
        # if not (polygon.touches(Point(o1x, o1y)) or polygon.contains(Point(o1x, o1y))):
        #     continue
        # if not (polygon.touches(Point(xb, yb)) or polygon.contains(Point(xb, yb))):
        #     continue
        # if not (polygon.touches(Point(o2x, o2y)) or polygon.contains(Point(o2x, o2y))):
        #     continue

        poly = Polygon([(xa, ya), (o1x, o1y), (xb, yb), (o2x, o2y)])

        if poly.covered_by(fullpoly):
            # print(poly, (abs(xa-xb)+1)*(abs(ya-yb)+1))
            a = (abs(xa-xb)+1)*(abs(ya-yb)+1)
            res = max(res, a)
            # res = max(res, poly.area)
        # print("--", poly, poly.area)
        # if poly.area == 0: continue

        # for c in range(len(lines)):
        #     line = LineString([lines[c], lines[(c+1) % len(lines)]])
        #     if line.within(poly):
        #         print("within", line)
        #         break
        #     if poly.contains(line):
        #         print("contains", line)
        #         break
        # else:
        #     print("valid", (abs(xb-xa)+1)*(abs(yb-ya)+1))
        #     res = max(res, (abs(xb-xa)+1)*(abs(yb-ya)+1))

        # for c, (xc, yc) in enumerate(lines):
        #     if a == c: continue
        #     if b == c: continue

        #     # print((xc, yc))
        #     if (xa < xc < xb) and (ya < yc < yb): break
        #     if (xb < xc < xa) and (ya < yc < yb): break
        #     if (xb < xc < xa) and (yb < yc < ya): break
        #     if (xa < xc < xb) and (yb < yc < ya): break

        #     if (xa <= xc <= xb) and (ya <= yc <= yb) \
        #     or (xb <= xc <= xa) and (ya <= yc <= yb) \
        #     or (xb <= xc <= xa) and (yb <= yc <= ya) \
        #     or (xa <= xc <= xb) and (yb <= yc <= ya):
        #         d = (c+1) % len(lines)
        #         if not (d == a or d == b):
        #             xd, yd = lines[d]

        #             if (xa <= xd <= xb) and (ya <= yd <= yb): break
        #             if (xb <= xd <= xa) and (ya <= yd <= yb): break
        #             if (xb <= xd <= xa) and (yb <= yd <= ya): break
        #             if (xa <= xd <= xb) and (yb <= yd <= ya): break

        #         d = (c-1) % len(lines)
        #         if not (d == a or d == b):
        #             xd, yd = lines[d]

        #             if (xa <= xd <= xb) and (ya <= yd <= yb): break
        #             if (xb <= xd <= xa) and (ya <= yd <= yb): break
        #             if (xb <= xd <= xa) and (yb <= yd <= ya): break
        #             if (xa <= xd <= xb) and (yb <= yd <= ya): break

        # else:

print(res)

'''

..............
.......OXXX#..
.......XXXXX..
..#XXXX#XXXX..
..XXXXXXXXXX..
..#XXXXXX#XX..
.........XXX..
.........#X#..
..............


'''
