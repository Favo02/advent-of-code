const fs = require("fs")
const path = require("path")

function parseInput(filepath) {
  return fs
    .readFileSync(path.join(__dirname, filepath), { encoding: "utf-8"})
    .split("\n")
    .slice(0, -1)
}

const formulas = new Map()

function parseFormulas(flines) {
  flines.forEach(f => {
    let tokens = f.split(" -> ")
    formulas.set(tokens[0], tokens[1])
  })
}

function executeReactionStep(pol) {
  let newPol = []

  for (let i = 0; i < pol.length-1; i++) {
    const sub = pol.slice(i, i+2)
    // console.log(sub)
    if (formulas.has(sub)) {
      newPol.push(sub[0])
      newPol.push(formulas.get(sub))
    }
    else {
      newPol.push(sub[0])
    }
  }
  newPol.push(pol.charAt(pol.length-1))
  return newPol.join("")
}

function getMax(str) {
  var max = 0,
      maxChar = '';
   str.split('').forEach(function(char){
     if(str.split(char).length > max) {
         max = str.split(char).length;
         maxChar = char;
      }
   });
   return max;
 };

 function getMin(str) {
  var max = 10198219872817,
      maxChar = '';
   str.split('').forEach(function(char){
     if(str.split(char).length < max) {
         max = str.split(char).length;
         maxChar = char;
      }
   });
   return max;
 };

const lines = parseInput(process.argv[2])

let pol = lines[0]
const flines = lines.slice(2)
parseFormulas(flines)

console.log(pol)
console.log(formulas)

for (let i = 0; i < 10; i++) {
  pol = executeReactionStep(pol)
  console.log(pol)
}

const max = getMax(pol)
const min = getMin(pol)

console.log(max, min)
console.log((max-1)-(min-1))
