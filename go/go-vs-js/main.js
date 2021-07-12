const Console = require('console')
const multiply = require('./neighbour')

function main() {
  //array
  const array = [1,2,3,4,5]
  Console.log(array)
  
  // there is no slice in javascript
  const slice = [1, 2, 3, 4, 5, 6, 7, 8, 4, 5, 6]
  Console.log(slice)
  slice.push(10)
  Console.log(slice)
  
  // only 1 return values in javascript
  plus = add(5, 10)
  Console.log(plus)

  multiply(2, -2)
    .then(data => Console.log(data))
    .catch(err => Console.error(err))
}

function add(x, y) {
  return x + y
}

main()