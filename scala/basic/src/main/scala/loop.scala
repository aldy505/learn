package basic

def doWhile() = {
  var i = 0
  while (i < 5) {
    println(i)
    i += 1
  }
}

def forLoop() = {
  var i = 1 to 5
  // this is soooo cool
  i.foreach(println)
}