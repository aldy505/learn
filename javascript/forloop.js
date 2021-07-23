var data = ['😀', '😁', '🤣', '😍', '😎'];

for (var i = 0; i < data.length / 2; i++) {
  console.log(`${i}: ${data.length} ${data.length / 2}`)
  data.push('😭')
}

console.log(data)
// How many 😭's in data array?

/**
 * Answer options:
 * Error: Javascript heap out of memory
 * 0
 * 1
 * 2
 * 3
 * 4
 * 5
 */