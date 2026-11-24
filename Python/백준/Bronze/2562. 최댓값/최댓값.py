max = 0
count = 0

for i in range(1, 10):
    n = int(input())
    if n > max:
        max = n
        count = i

print(max)
print(count)