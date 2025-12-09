# https://adventofcode.com/2025/day/9
# https://github.com/Favo02/advent-of-code

import sys
from shapely.geometry.polygon import Polygon
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = part2 = 0

points = list(map(lambda l: tuple(map(int, l.strip().split(","))), sys.stdin.read().strip().splitlines()))

polygon = Polygon(points)

part1 = 0
for a, (xa, ya) in enumerate(points):
    for xb, yb in points[a+1:]:
        part1 = max(part1, (abs(xb-xa)+1)*(abs(yb-ya)+1))

        rectangle = Polygon([(xa, ya), (xa, yb), (xb, yb), (xb, ya)])
        if polygon.covers(rectangle):
            part2 = max(part2, (abs(xb-xa)+1)*(abs(yb-ya)+1))

print("Part 1:", part1)
print("Part 2:", part2)
