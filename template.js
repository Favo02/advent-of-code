const fs = require("fs")
const path = require("path")

function parseInput(filepath) {
  return fs
    .readFileSync(path.join(__dirname, filepath), { encoding: "utf-8"})
    .split("\n")
    .slice(0, -1)
}

const lines = parseInput("input.txt")
console.log(lines)
