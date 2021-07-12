using Genie
using Genie.Router
using Genie.Requests
using Genie.Renderer.Json

Genie.config.run_as_server = true

route("/", method = "GET") do
  json(Dict("message" => "hi there"))
end

route("/", method = "POST") do 
  @show Requests.jsonpayload()

  json(Dict("name" => "$(jsonpayload()["name"])"))
end

route("/", method="DELETE") do 
  name = @params(:name)
  json(Dict("name" => "$name"))
end

Genie.startup()