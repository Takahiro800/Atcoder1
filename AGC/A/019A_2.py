Q,H,S,D = map(int,input().split())
N = int(input())
h = min(2*Q,H)
s = min(2*h,S)
d = min(2*s,D)
print((N//2)*d + (N%2)*s)