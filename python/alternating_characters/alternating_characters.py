import math
import os
import random
import re
import sys


def alternatingCharacters(s: str) -> int:
    total = 0
    prior_char = None
    for c in s:
        if prior_char is None:
            prior_char = c
        elif c == prior_char:
            total += 1
        else:
            prior_char = c
        
    return total


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input().strip())

    for q_itr in range(q):
        s = input()

        result = alternatingCharacters(s)

        fptr.write(str(result) + '\n')

    fptr.close()

