const fs = require("fs")
const path = require("path")

const lines = parseInput(process.argv[2])

const pol = lines[0]
const flines = lines.slice(2)

const formulas = {}
parseFormulas(flines)

let couples = {}
convertStringToCouples(pol)

let part1
for (let i = 0; i < 40; i++) {
  if (i == 10) {
    const { count, lenght } = countValues(pol)
    const { max, min } = getMaxMin(count)
    part1 = max-min  
  }
  executeReactionStep()
}

const { count, lenght } = countValues(pol)
const { max, min } = getMaxMin(count)

console.log("part1", part1)
console.log("part2", max-min)

// parse input from file given as process arg
function parseInput(filepath) {
  return fs
    .readFileSync(path.join(__dirname, filepath), { encoding: "utf-8"})
    .split("\n")
    .slice(0, -1)
}

// parse formula lines into map object (modifies formulas)
function parseFormulas(flines) {
  flines.forEach(f => {
    let tokens = f.split(" -> ")
    formulas[tokens[0]] = tokens[1]
  })
}

// convert starting string to couples (modifies couples)
function convertStringToCouples(pol) {
  for (let i = 0; i < pol.length-1; i++) {
    const sub = pol.slice(i, i+2)
    couples[sub] = (couples[sub] ?? 0) +1
  }
}

// execute one reaction step (modifies couples)
function executeReactionStep() {
  const newCouples = {}

  // console.log("old", couples)

  for (const k in couples) {
    if (formulas[k]) {
      const value = couples[k]
      
      const first = k.charAt(0) + formulas[k]
      const second = formulas[k] + k.charAt(1)
      
      // console.log(k, "=>", first, second, value)

      newCouples[first] = (newCouples[first] ?? 0) + value
      newCouples[second] = (newCouples[second] ?? 0) + value
    }
  }

  couples = newCouples
  // console.log("new", couples)
  // console.log(countValues(pol))
  // console.log()
}

// count the number of each element (and length)
function countValues(pol) {
  const count = {}
  let length = 0

  for (const k in couples) {
    const first = k.charAt(0)
    const second = k.charAt(1)

    count[first] = (count[first] ?? 0) + couples[k]
    count[second] = (count[second] ?? 0) + couples[k]

    length += couples[k]
  }
  
  count[pol.charAt(0)]++
  count[pol.charAt(pol.length-1)]++

  for (const k in count) {
    count[k] /= 2
  }

  return { count, lenght: length+1 }
}

// get most common and least common element
function getMaxMin(count) {
  let max = 0
  let min = Number.MAX_SAFE_INTEGER // porcata ma ok
  
  for (const key in count) {
    if (count[key] > max) {
      max = count[key]
    }

    if (count[key] < min) {
      min = count[key]
    }
  }

  return { max, min }
}
