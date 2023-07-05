const fs = require("fs")
const path = require("path")

const lines = parseInput(process.argv[2])

const res = sumAllAndReduce(lines)
const mag = magnitude(res)

console.log(res.join(""))
console.log("magnitude (part1):", mag)

// parse input from file given as process arg
function parseInput(filepath) {
  return fs
    .readFileSync(path.join(__dirname, filepath), { encoding: "utf-8"})
    .split("\n")
    .slice(0, -1)
    .map(l => parseToArray(l))
}

// parse string to array (parse numbers to int)
function parseToArray(expr) {
  let numBuffer = ""
  let arr = []

  for (let i = 0; i < expr.length; i++) {
    const element = expr[i]
    
    if (element >= '0' && element <= '9') {
      numBuffer += element
    }
    else {
      if (numBuffer.length > 0) {
        number = parseInt(numBuffer)
        arr.push(number)
        numBuffer = ""
      }

      arr.push(element)
    }
  }

  return arr
}

// sum two expressions and return new expression
function sum(expr1, expr2) {
  return [ '[', ...expr1, ',', ...expr2, ']' ]
}

// reduce to normal form an expression
function reduce(expr) {
  let cont = 0

  loopExplode:
  for (let i = 0; i < expr.length; i++) {

    // reset this loop
    function reset() {
      cont = 0
      i = -1
    }

    const element = expr[i]
    
    if (element == '[') cont++
    if (element == ']') cont--

    // explode
    if (cont > 4) {
      // console.log("bef explode: \n", expr.join(""))
      expr = explode(expr, i)
      // console.log("aft explode: \n", expr.join(""), "\n")
      reset()
      continue
    }

    // check for possible splits only if no explode left
    if (i == expr.length -1) {

      for (let j = 0; j < expr.length; j++) {
        // split
        if (Number.isInteger(expr[j]) && expr[j] > 9) {
          // console.log("bef split: \n", expr.join(""))
          expr = split(expr, j)
          // console.log("aft split: \n", expr.join(""), "\n")

          // reset loop and start checking for new explode or slipts
          reset()
          continue loopExplode
        }
      }
    }
  }

  return expr
}

// explode the couple starting at start index
function explode(expr, start) { 
  const end = start+5
  
  const elem = expr.slice(start, end)

  // find number before
  let befInd = -1
  for (let i = start; i >= 0; i--) {
    if (Number.isInteger(expr[i])) {
      befInd = i
      break
    }  
  }
  // console.log("befind", befInd, expr[befInd])

  // find number after
  let aftInd = -1
  for (let i = end; i < expr.length; i++) {
    if (Number.isInteger(expr[i])) {
      aftInd = i
      break
    }   
  }
  // console.log("aftind", aftInd, expr[aftInd])

  if (befInd != -1) expr[befInd] += elem[1]
  if (aftInd != -1) expr[aftInd] += elem[3]

  return [...expr.slice(0, start), 0, ...expr.slice(end)]
}

// split the couple starting at start index
function split(expr, start) {
  const a = Math.floor(expr[start]/2)
  const b = Math.ceil(expr[start]/2)

  return [ ...expr.slice(0, start), '[', a, ',', b, ']', ...expr.slice(start+1) ]
}

// sums all expressions and return reduced final expression
function sumAllAndReduce(exprs) {
  let expr = reduce(exprs[0])
  
  for (let i = 1; i < exprs.length; i++) {
    expr = sum(expr, exprs[i])
    expr = reduce(expr)
  }

  return expr
}

// return magnitude of expression
function magnitude(expr) {
  for (let i = 0; i < expr.length; i++) {
    const element = expr[i]

    if (element == ']') {
      const magnitude = expr[i-1]*2 + expr[i-3]*3
      expr = [...expr.slice(0, i-4), magnitude, ...expr.slice(i+1)]
      i = -1
    }
  }

  return expr[0]
}
