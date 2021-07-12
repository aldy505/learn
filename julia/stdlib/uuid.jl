using UUIDs

v4 = UUIDs.uuid4()
println(v4)

v1 = UUIDs.uuid1()
println(v1)

println(UUIDs.uuid_version(v4)) # 4
println(UUIDs.uuid_version(v1)) # 1