const fs = require("fs")
const path = require("path")

const lines = parseInput(process.argv[2])
console.log(lines)

// parse input from file given as process arg
function parseInput(filepath) {
  return fs
    .readFileSync(path.join(__dirname, filepath), { encoding: "utf-8"})
    .split("\n")
    .slice(0, -1)
}
