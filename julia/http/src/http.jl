module http
using Sockets
using HTTP
using JSON

println("Open http://localhost:3000 - send SIGINT (Ctrl + C) to terminate.")

response = JSON.json(Dict("message" => "Hello world!"))

HTTP.listen(Sockets.localhost, 3000) do http::HTTP.Stream
  @show http.message
  @show HTTP.header(http, "Content-Type")

  HTTP.setstatus(http, 200)
  HTTP.setheader(http, "Content-Type" => "application/json")
  HTTP.startwrite(http)

  write(http, response)
end

end # module
