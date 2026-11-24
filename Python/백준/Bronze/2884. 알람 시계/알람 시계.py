a, b = map(int, input().split())

if b >= 45:
    print(a, b - 45)
else:
    print((a - 1) % 24, b + 15)