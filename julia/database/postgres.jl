using UUIDs, LibPQ, Tables

conn = LibPQ.Connection("postgres://postgres:password@localhost:5432/learn_julia")

try
  result = execute(conn, """
    CREATE TABLE IF NOT EXISTS users (
      id VARCHAR(255) PRIMARY KEY,
      name VARCHAR(255),
      email VARCHAR(255) UNIQUE
    );
  """)

  println(columntable(result))

  execute(conn, "BEGIN;")
  LibPQ.load!(
    (id = [UUIDs.uuid4(), UUIDs.uuid4(), UUIDs.uuid4()], name = ["Tomas", "Steph", "Morgan"], email = ["tomasgoff@gmail.com", "stephjgilbert@gmail.com", "ayomorgan@gmail.com"]),
    conn,
    "INSERT INTO users (id, name, email) VALUES (\$1, \$2, \$3);",
  )
  execute(conn, "COMMIT;")
  
  select = execute(conn, "SELECT * FROM users")
  println(columntable(select))

  close(conn)
catch e
  println(e)
end