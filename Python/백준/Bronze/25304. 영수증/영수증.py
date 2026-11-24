x = int(input())
N = int(input())
t = 0
for _ in range(N):
    a,b = map(int,input().split())
    t += a*b
if t == x :
    print("Yes")
else :
    print("No")