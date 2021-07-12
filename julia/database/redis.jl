using Redis

conn = Redis.RedisConnection(host="localhost", port=6379)

try
  insert = Redis.set(conn, "name", "Tomas")
  println(insert) # true
  fetch = Redis.get(conn, "name")
  println(fetch) # Tomas

  Redis.disconnect(conn)
catch e
  println(e)
end