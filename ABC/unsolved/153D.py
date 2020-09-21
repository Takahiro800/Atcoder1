H = int(input())
import pdb; pdb.set_trace()
def atk(n):
  if n==1: return 1
  return 1 + 2*(atk(n//2))

print(atk(H))