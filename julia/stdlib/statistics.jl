using Statistics

data = [29426, -27734, -16902, -16480, -26822, -21778, -1025, 7901, 11263, 16772]

a = Statistics.mean(data)
b = Statistics.median(data)
c = Statistics.middle(data)

println(a) # -4537.9
println(b) # -8752.5
println(c) # 846.0