println(2 + 2 * 31 / (9 * 2)) # 5.444444444444445

# Bitwise operators https://en.wikipedia.org/wiki/Bitwise_operation#Bitwise_operators
# Documentations: https://docs.julialang.org/en/v1/manual/mathematical-operations/#Bitwise-Operators

println(150 & 120) # 16

println(30 | 3) # 31

# Numeric comparison

println(isequal(20, 30)) # false
println(isequal(1123141, 1123141)) # true
println(isfinite(930)) # true
println(isinf(930)) # false
println(isnan(30)) # false

# advanced maths

println(round(30.13)) # 30.0
println(floor(30.13)) # 30.0
println(ceil(30.13)) # 31.0
println(trunc(30.13)) # 30.0

println(abs(-30)) # 30

println(sqrt(144)) # 12.0
println(hypot(12,13)) # 17.69180601295413
println(hypot(3,4)) # 5.0
println(log(8)) # 2.0794415416798357
println(log2(16)) # 4.0
println(log10(800)) # 2.9030899869919438
println(exponent(105.7)) # 6
