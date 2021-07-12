using SHA

a = SHA.sha1("lorem ipsum")
b = SHA.sha2_256("lorem ipsum")
c = SHA.sha3_256("lorem ipsum")

# These lines below should return human-unreadable bytes
println(a)
println(b)
println(c)

# let's make it readable
println(bytes2hex(a))
println(bytes2hex(b))
println(bytes2hex(c))