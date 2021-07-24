# So in Julia, there are a few ways to import a package/file: include, using and import.

# Say for example we have this code:

#   module MyModule

#   export x, y

#   x() = "x"
#   y() = "y"
#   p() = "p"

#   end

# using MyModule
#   in-scope: All exported names (x and y), MyModule.x, MyModule.y, and MyModule.p
#   extensible: MyModule.x, MyModule.y, and MyModule.p
  
# using MyModule: x, p
#   in-scope: x and p
#   extensible: (nothing)

# import MyModule
#   in-scope: MyModule.x, MyModule.y, and MyModule.p
#   extensible: MyModule.x, MyModule.y, and MyModule.p

# import MyModule.x, MyModule.p
#   in-scope: x and p
#   extensible: x and p

# import MyModule: x, p
#   in-scope: x and p
#   extensible: x and p

# Imports could also be renamed
#   import MyModule: x as b
#   import MyModule as MM

# When multiple using or import statements of any of the forms above are used,
# their effect is combined in the order they appear. For example,

#   using NiceStuff         # exported names and the module name
#   import NiceStuff: nice  # allows adding methods to unqualified functions

# would bring all the exported names of NiceStuff and the module name itself into scope, 
# and also allow adding methods to nice without prefixing it with a module name.