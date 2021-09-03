def decor(func):
  def inner_func():
    return func() + 10
  return inner_func

@decor
def return_five():
  return 5

print(return_five())
print(return_five)