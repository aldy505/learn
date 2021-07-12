using Random

x = Random.MersenneTwister()

println(x)

a = Random.rand(Int16, 10)
b = Random.rand(Random.MersenneTwister(), "therefore")
c = Random.rand(Random.MersenneTwister(), ("where", "have", "you", "been", "all", "this", "time"))

println(a)
println(b)
println(c)
