using HTTP

include("handler.jl")

const Routes = HTTP.Router()

HTTP.@register(Routes, "GET", "/", getTodo)
HTTP.@register(Routes, "GET", "/*", getTodoByID)
HTTP.@register(Routes, "POST", "/", addTodo)
HTTP.@register(Routes, "PATCH", "/", editTodo)
HTTP.@register(Routes, "DELETE", "/*", deleteTodo)