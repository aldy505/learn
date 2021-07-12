# in julia, array starts with 1, not 0
arr = [1,2,3,4,5,6]

println(typeof(arr))    # Vector{Int64}
println(arr[1])         # 1

multiple = [1, 2,"hello", 3, "there"]

println(typeof(multiple))   # Vector{Any}
println(multiple[3])        # hello
println(multiple[end])      # there

int = Int8[1, 3, 5]

println(typeof(int))  # Vector{Int8}

println(length(int))  # 3

# Let's modify some array

println(push!(int, 7))          # Int8[1, 3, 5, 7]
println(push!(int, 9, 11, 13))  # Int8[1, 3, 5, 7, 9, 11, 13]

println(append!(int, arr))  # Int8[1, 3, 5, 7, 9, 11, 13, 1, 2, 3, 4, 5, 6]
println(pop!(int))          # 6

# let's see...
println(int)  # Int8[1, 3, 5, 7, 9, 11, 13, 1, 2, 3, 4, 5]
println(arr)  # [1, 2, 3, 4, 5, 6]

println(sort(int))    # Int8[1, 1, 2, 3, 3, 4, 5, 5, 7, 9, 11, 13] 
# when you're not using a function not followed with !, the original value won't change
println(int)          # Int8[1, 3, 5, 7, 9, 11, 13, 1, 2, 3, 4, 5]
println(sort!(int))   # Int8[1, 1, 2, 3, 3, 4, 5, 5, 7, 9, 11, 13]

println(popfirst!(int))   # 1
println(popat!(int, 5))   # 4