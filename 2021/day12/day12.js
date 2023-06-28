// https://adventofcode.com/2021/day/12
// https://github.com/Favo02/advent-of-code

const fs = require("fs")
const path = require("path")

function parseInput(filepath) {
  return fs
    .readFileSync(path.join(__dirname, filepath), { encoding: "utf-8" })
    .split("\n")
    .slice(0, -1)
}

function buildGraph(lines) {
  const graph = new Map()
  lines.forEach((l) => {
    const tokens = l.split("-")
    const [from, to] = tokens

    graph.set(from, [...(graph.get(from) ?? []), to])
    graph.set(to, [...(graph.get(to) ?? []), from])
  })
  return graph
}

// true: valid for part1, false: not valid for part1
// not best way to calculate part1, but already calculating every path with 2 visits to small caves for part2
// old commits contains better way to calculate part1
function checkPart1Valid(arr) {
  for (let i = 0; i < arr.length; i++) {
    const e = arr[i]

    if (e === e.toLowerCase()) {
      if (arr.filter(el => el === e).length >= 2) {
        return false
      }
    }
  }
  return true
}

// true: valid for part2, false: not valid for part2
function checkPart2Valid(arr) {
  var dupFound

  for (let i = 0; i < arr.length; i++) {
    const e = arr[i]
    
    if (e === e.toLowerCase()) {
      if (arr.filter(el => el === e).length >= 2) {
        if (dupFound && dupFound !== e) {
          return false
        } else {
          dupFound = e
        }
      }
    }
  }
  return true
}

var countPart1 = 0
var countPart2 = 0

// build tree of possible paths, number of leaves are number of possible paths to reach end
function buildTree(graph, cur) {

  if (cur.includes("end")) {
    // not best way to calculate part1, but already calculating every path with 2 visits to small caves for part2
    // old commits contains better way to calculate part1
    if (checkPart1Valid(cur)) {
      countPart1++
    }
    if (checkPart2Valid(cur)) {
      countPart2++
    }
    return
  }

  const curV = cur.at(-1)

  const childs = []
  // for each adjacent add a path
  graph.get(curV).forEach((adj) => {
    if (adj !== "start") {
      if (!(adj === adj.toLowerCase() && (cur.filter((e) => e === adj).length >= 2))) {
        childs.push([...cur, adj])
      }
    }
  })

  childs.forEach((next) => {
    buildTree(graph, next)
  })
}

const lines = parseInput(process.argv[2])

const graph = buildGraph(lines)

buildTree(graph, ["start"])

console.log("part1", countPart1)
console.log("part2", countPart2)

