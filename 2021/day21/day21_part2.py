# https://adventofcode.com/2021/day/21
# https://github.com/Favo02/advent-of-code

from collections import defaultdict

start_a = int(input().split(": ")[1])-1
start_b = int(input().split(": ")[1])-1

dice_comb = defaultdict(int)
for r1 in [1,2,3]:
  for r2 in [1,2,3]:
    for r3 in [1,2,3]:
      dice_comb[r1+r2+r3] += 1
assert sum(dice_comb.values()) == 27

TURN_A = 0
TURN_B = 1

POSITIONS = 10
SCORE = 21

# 5d dp: dp[turn][scoreA][posA][scoreB][posB] (2 * 21 * 10 * 21 * 10 = 88200 states)

dp = [[[[[0 for _ in range(POSITIONS)] for _ in range(SCORE)] for _ in range(POSITIONS)] for _ in range(SCORE)] for _ in range(2)]
dp[TURN_A][0][start_a][0][start_b] = 1

resA = resB = 0

for scoreA in range(SCORE):
  for posA in range(POSITIONS):
    for scoreB in range(SCORE):
      for posB in range(POSITIONS):
        for turn in range(2):

          if turn == TURN_A:
            # A turn
            for val, qty in dice_comb.items():
              newpos = (posA+val) % 10
              newscore = scoreA + (newpos+1)
              if newscore >= SCORE:
                resA += dp[TURN_A][scoreA][posA][scoreB][posB] * qty
              else:
                dp[TURN_B][newscore][newpos][scoreB][posB] += dp[TURN_A][scoreA][posA][scoreB][posB] * qty

          if turn == TURN_B:
            # B turn
            for val, qty in dice_comb.items():
              newpos = (posB+val) % 10
              newscore = scoreB + (newpos+1)
              if newscore >= SCORE:
                resB += dp[TURN_B][scoreA][posA][scoreB][posB] * qty
              else:
                dp[TURN_A][scoreA][posA][newscore][newpos] += dp[TURN_B][scoreA][posA][scoreB][posB] * qty

print("Part 2:", max(resA, resB))


dp[turn][scoreA][scoreA][scoreB][posB]
