var i = 25;
var x = (i + 5 || i == 30) ? (i + 10 || i < 5) ? i + 10 ? i - 5 : i - 3 : i - 2 : i;
console.log(x);

/**
 * Answer options:
 * 20
 * 22
 * 23
 * 25
 * 30
 * 35
 * 40
 */