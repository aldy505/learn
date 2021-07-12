i = 0

while i < 4
  println(i)
  global i += 1
end

# 0
# 1
# 2
# 3

for j = 1:3
  println(j)
end

# 1
# 2
# 3

# I love to make complicated things sometimes
fruits = ["apple", "orange", "pear", "banana"]
for k in fruits
  # find pear
  if k == "pear"
    println("found it! it's $(k)")
  end
end

# found it! it's pear