# a few ways to concatenate String

hello = "henlo"
there = "there"

println(string(hello, " ", there))  # henlo there
println(hello * " " * there)        # henlo there
println("$hello $there")            # henlo there

println(hello ^ 3)