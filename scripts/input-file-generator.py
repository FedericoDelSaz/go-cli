import random

# Define the number of lines and the range of unsigned 64-bit integers
num_lines = 30_000_000
min_val = 0
max_val = 18_446_744_073_709_551_615

# File path
file_path = "input-max.txt"

# Generate and write to the file
with open(file_path, "w") as f:
    for _ in range(num_lines):
        f.write(f"{random.randint(min_val, max_val)}\n")

file_path
