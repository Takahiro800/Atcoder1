X = int(input())

def is_prime(n):
  if n == 2:
    return True
  elif n == 1 or n%2 ==0:
    return False
  else:
    m = 3
    while m*m <= n:
      if n % m == 0:
        return False
      else:
        m += 2
    return True

while 1:
  if is_prime(X):
    print(X)
    break
  X += 1