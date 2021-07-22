package basic

// Available types:
// Double
// Float
// Int
// Short
// Byte
// Unit
// Boolean
// Nothing
// Null

def types() = {
  // Declare a variable with val
  val a: Int = 10
  // Types could be infered
  val b = 10.0

  println(a)
  println(b)
}

def array() = {
  var array = Array(1,2,3,4,5,6)
  println(array)
  // to look a value
  println(array(5))
}

def map() = {
  var map = Map(
    "name" -> "Andrew",
    "age" -> 21
  )
  println(map)
  println(map("name"))
}