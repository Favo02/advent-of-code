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

// true => valid for part2, false => not valid for part2
function checkPart2Valid(arr) {
  var dupFound

  for (let i = 0; i < arr.length; i++) {
    const e = arr[i]
    
    if (e === e.toLowerCase()) {
      const occ = arr.filter(el => el === e).length
      if (occ >= 2) {
        if (dupFound && dupFound !== e) {
          // console.log("fail for", e)
          return false
        } else {
          dupFound = e
        }
      }
    }
  }
  return true
}

// const tree = new Map()
var count = 0

// build tree of possible paths, number of leaves are number of possible paths to reach end
function buildTree(graph, cur) {
  // console.log(cur.toString())
  if (cur.includes("end")) {
    if (checkPart2Valid(cur)) {
      // console.log("accepted", cur.toString())
      count++
    } else {
      // console.log("rejected", cur.toString())
    }
    // tree.delete(cur.toString())
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
  // tree.set(cur.toString(), childs)

  childs.forEach((next) => {
    buildTree(graph, next)
  })
}

const lines = parseInput(process.argv[2])
// console.log(lines)

const graph = buildGraph(lines)
// console.log(graph)

buildTree(graph, ["start"])

// console.log(tree)
console.log(count)

