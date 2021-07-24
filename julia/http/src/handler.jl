using HTTP
using JSON

function getTodo(req::HTTP.Request)
  response = JSON.json(Dict("message" => "Hello world!", "target" => req.target))
  return HTTP.Response(200, response)
end

function getTodoByID(req::HTTP.Request)
  response = JSON.json(Dict("message" => "Hello world!", "target" => req.target))
  return HTTP.Response(200, response)
end

function addTodo(req::HTTP.Request)
  body = req.body
  response = JSON.json(Dict("message" => "got it", "body" => JSON.parse(String(body)), "headers" => req.headers))
  return HTTP.Response(200, response)
end

function editTodo(req::HTTP.Request)
  body = req.body
  response = JSON.json(Dict("message" => "got it", "body" => JSON.parse(String(body))))
  return HTTP.Response(200)
end

function deleteTodo(req::HTTP.Request)
  body = req.body
  response = JSON.json(Dict("message" => "got it", "body" => JSON.parse(String(body)), "target" => req.target))
  return HTTP.Response(200)
end