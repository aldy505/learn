x = 10
y = 30

if x < y
  println("x is smaller than y")
elseif x > y
  println("x is bigger than y")
else 
  println("they're the same")
end

# x is smaller than y

if x >= 0 && y >= 0 
  println("they both exists")
elseif x >= 0 || y >= 0
  println("only one of them exists")
else
  println("none of them exists")
end

# they both exists