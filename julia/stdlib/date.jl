using Dates

# I LOVE Julia
# The dates default to ISO 8601 format, which I literally use heavily on every of my projects!

a = Dates.DateTime(2025) 
b = Dates.isleapyear(Date("2021"))
c = Dates.today()
d = Dates.now()
e = Dates.Oct # this is a constant
f = Dates.DateTime(2020, 8, 4, 12, 10)

println(a) # 2025-01-01T00:00:00
println(b) # false
println(c) # 2021-06-16
println(d) # 2021-06-16T15:16:02.247
println(e) # 10
println(f) # 2020-08-04T12:10:00