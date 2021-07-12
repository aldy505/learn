#=
Julia enables package developers and users to document functions, 
types and other objects easily via a built-in documentation system.

The basic syntax is simple: any string appearing at the toplevel right before an object
(function, macro, type or instance) will be interpreted as documenting it 
(these are called docstrings). Note that no blank lines or comments may intervene 
between a docstring and the documented object.

Documentation is interpreted as Markdown, so you can use indentation and code fences
to delimit code examples from text. 
=#

"Tells you the meaning of life"
function meaningOfLife()
  return 42
end

println(meaningOfLife()) # you guess it, it's 42

# This might be how I'd approach Julia's function from now on.
function optional(;x=5, y=8)
  println(x)
  println(y)
end

optional(y=5, x=1)