const fs = require("fs")
const path = require("path")

const lines = parseInput(process.argv[2]).map(l => parseToArray(l))

const res = sumAndReduce(lines)

console.log(res.join(""))

console.log(magnitude(res))

// parse input from file given as process arg
function parseInput(filepath) {
  return fs
    .readFileSync(path.join(__dirname, filepath), { encoding: "utf-8"})
    .split("\n")
    .slice(0, -1)
}

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

  // console.log(arr)
  return arr
}

function sum(arr1, arr2) {
  return [ '[', ...arr1, ',', ...arr2, ']' ]
}

function reduce(expr) {
  let cont = 0

  loopExplode:
  for (let i = 0; i < expr.length; i++) {
    const element = expr[i]
    
    if (element == '[') cont++
    if (element == ']') cont--

    // explode
    if (cont > 4) {
      // console.log("bef explode: \n", expr.join(""))
      expr = explode(expr, i)
      // console.log("aft explode: \n", expr.join(""), "\n")
      i = -1
      cont = 0
      continue
    }

    if (i == expr.length -1) {

      for (let j = 0; j < expr.length; j++) {
        // split
        if (Number.isInteger(expr[j]) && expr[j] > 9) {
          // console.log("bef split: \n", expr.join(""))
          expr = split(expr, j)
          // console.log("aft split: \n", expr.join(""), "\n")
          i = -1
          j = -1
          cont = 0
          continue loopExplode
        }
      }
    }
  }

  return expr
}

function explode(expr, start) { 
  const end = start+5
  
  const elem = expr.slice(start, end)
  // console.log(elem)

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

  if (befInd != -1) {
    // console.log("be", expr)
    expr[befInd] += elem[1]
    // console.log("ae", expr)
  }
  if (aftInd != -1) {
    expr[aftInd] += elem[3]
  }

  return [...expr.slice(0, start), 0, ...expr.slice(end)]
}

function split(expr, start) {
  // console.log(expr[start])

  const a = Math.floor(expr[start]/2)
  const b = Math.ceil(expr[start]/2)

  // console.log("-----", expr[start], a, b)
  // console.log(a,b)

  return [ ...expr.slice(0, start), '[', a, ',', b, ']', ...expr.slice(start+1) ]
}

function sumAndReduce(lines) {
  let expr = reduce(lines[0])
  
  for (let i = 1; i < lines.length; i++) {
    // expr = reduce(expr)
    expr = sum(expr, lines[i])
    expr = reduce(expr)
  }

  return expr
}

function magnitude(arr) {
  for (let i = 0; i < arr.length; i++) {
    const element = arr[i]

    if (element == ']') {
      const magnitude = arr[i-1]*2 + arr[i-3]*3
      arr = [...arr.slice(0, i-4), magnitude, ...arr.slice(i+1)]
      i = -1
    }
  }

  return arr[0]
}
