N,K = map(int,input().split())
H = list(map(int, input().split()))

monster_list = sorted(H, reverse=True)[:K]
ans = sum(H) - sum(monster_list)

print(ans)