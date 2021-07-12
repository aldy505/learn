# Dictionary is like associative array in PHP, object in Javascript, or interface in Go.

store = Dict(
  "papaya" => "\$3",
  "apple" => "\$2"
)

println(store)            # Dict("apple" => "\$2", "papaya" => "\$3")
println(store["apple"])   # $2

# there's one thing that made me curious
# is it possible to make Dict inside Dict? like a nested one?

appleContent = Dict(
  "price" => "\$2",
  "quantity" => 15
)

papayaContent = Dict(
  "price" => "\$3",
  "quantity" => 8
)

store = Dict(
  "apple" => appleContent,
  "papaya" => papayaContent
)

println(store)        # Dict{String, Dict{String, Any}}("apple" => Dict("price" => "\$2", "quantity" => 15), "papaya" => Dict("price" => "\$3", "quantity" => 8))
# now how would you access it?
applePrice = store["apple"]["price"]
println(applePrice)   # $2

# can we make it more simple?

function dictContent(price=1, quantity=1)
  Dict(
    "price" => "\$$price",
    "quantity" => quantity
  )
end

store = Dict(
  "banana" => dictContent(2.5, 9),
  "pear" => dictContent(3, 6),
  "apple" => dictContent(2, 19)
)

println(store) # Dict{String, Dict{String, Any}}("pear" => Dict("price" => "\$3", "quantity" => 6), "banana" => Dict("price" => "\$2.5", "quantity" => 9), "apple" => Dict("price" => "\$2", "quantity" => 19))
# welp it works