function multiply(x, y) {
  return new Promise((resolve, reject) => {
    if (y <= 0) {
      return reject('can\'t multiply by zero')
    }
    return resolve(x * y)
  })
}

module.exports = multiply