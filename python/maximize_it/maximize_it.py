from itertools import product

lines, modulus = list(map(int, input().split()))
num_list = []

def square_int(x: str) -> int:
    return int(x)**2

for _ in range(lines):
    num_list.append(list(map(square_int, input().split()[1:])))
    
total = 0
for combo in product(*num_list):
    local_max = sum(combo) % modulus
    if local_max > total:
        total = local_max

print(total)
