using Sockets
using HTTP

include("handler.jl")
include("routes.jl")

println("Open http://localhost:3000 - send SIGINT (Ctrl + C) to terminate.")

HTTP.serve(Routes, Sockets.localhost, 3000)
