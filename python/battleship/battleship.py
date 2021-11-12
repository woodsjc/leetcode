from typing import List

def getHitProbability(R: int, C: int, G: List[List[int]]) -> float:
  # count total ships & r*c for total cells
  total = 0
  for line in G:
    total += sum(line)
  return total / (R*C)
