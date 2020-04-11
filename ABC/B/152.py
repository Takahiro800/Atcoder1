a,b = map(int, input().split())
x = str(a)*b
y = str(b)*a

print(min(x,y))