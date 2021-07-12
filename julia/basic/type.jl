a = 1
b = 10
c = 10000
d = 10.2
e = 100.505

println(typeof(a)) # Int64
println(typeof(b)) # Int64
println(typeof(c)) # Int64
println(typeof(d)) # Float64
println(typeof(d)) # Float64

f = "hi there"
g = 'h'

println(typeof(f)) # String
println(typeof(g)) # Char

h = true

println(typeof(h)) # Bool

# These below would return an error
# syntax: type declarations on global variables are not yet supported

i::Int8 = 2
j::Int16 = 20
k::Int32 = 200

println(typeof(i)) #
println(typeof(j)) #
println(typeof(k)) #