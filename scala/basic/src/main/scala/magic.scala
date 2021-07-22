package basic

def stringInterpolation() = {
  val a = s"Their age is ${2021 - 1988}"
  println(a)
  val b = 5
  println(s"Five in number is $b") // A bit similar like PHP isnt it
  println(raw"Another raw text")
}