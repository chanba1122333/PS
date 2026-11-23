a = int(input())
b = int(input())

for i in range(3):
    print(a * (b // (10 ** i) % 10))

print(a * b)