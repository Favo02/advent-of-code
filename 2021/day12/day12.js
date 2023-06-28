const fs = require("fs")
const path = require("path")

function parseInput(filepath) {
  return fs
    .readFileSync(path.join(__dirname, filepath), { encoding: "utf-8"})
    .split("\n")
    .slice(0, -1)
}

function buildGraph(lines) {
  const graph = new Map()
  lines.forEach(l => {
    const tokens = l.split("-")
    const [from, to] = tokens

    graph.set(from, [...graph.get(from) ?? [], to])
    graph.set(to, [...graph.get(to) ?? [], from])
  })
  return graph
}

const tree = new Map()
var count = 0

// build tree of possible paths, number of leaves are number of possible paths to reach end
function buildTree(graph, cur) {
  if (cur.includes("end")) {
    count++
    return
  }

  const curV = cur.at(-1)

  const childs = []
  // for each adjacent add a path
  graph.get(curV).forEach(adj => {
    if (!(adj === adj.toLowerCase() && cur.includes(adj))) {
      childs.push([...cur, adj])
    }
  })
  tree.set(cur.toString(), childs)
  
  childs.forEach(next => {
    buildTree(graph, next)
  })
}

const lines = parseInput("input.txt")
// console.log(lines)

const graph = buildGraph(lines)
// console.log(graph)

buildTree(graph, ["start"])

console.log(tree)
console.log(count)

