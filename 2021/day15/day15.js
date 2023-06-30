const fs = require("fs")
const path = require("path")

const lines = parseInput(process.argv[2])

const matrix = []
buildGraph(lines)

const dist = dijkstra(0, 0)
console.log(dist)

// parse input from file given as process arg
function parseInput(filepath) {
  return fs
    .readFileSync(path.join(__dirname, filepath), { encoding: "utf-8"})
    .split("\n")
    .slice(0, -1)
}

// builds the graph as a matrix (modifies matrix)
function buildGraph(lines) {
  lines.forEach(line => {
    matrix.push(line.split(""))
  })
}

// does not works with negative weights
function dijkstra(startX, startY) {

  const dist = {}

  let queue = []

  for (let y = 0; y < matrix.length; y++) {
    for (let x = 0; x < matrix[y].length; x++) {
      dist[`${x}-${y}`] = 100000
      queue.push(`${x}-${y}`)
    }
  }

	dist[`${startX}-${startY}`] = 0

  while (queue.length > 0) {

    let min = queue[0]
    queue.forEach(el => {
      if (dist[el] < dist[min]) {
        min = el
      }
    })

    queue = queue.filter(e => e !== min)

    const adj = getAdjacent(min)

    adj.forEach(ad => {
      if ((parseInt(dist[min]) + parseInt(ad.w)) < parseInt(dist[ad.k])) {
        dist[ad.k] = parseInt(dist[min]) + parseInt(ad.w)
      }
    })
	}

	return dist
}

function getAdjacent(el) {
  const [xs, ys] = el.split("-")
  const x = parseInt(xs)
  const y = parseInt(ys)

  const adj = []
  adj.push({ k: `${x-1}-${y}`, w: (matrix[x-1] ?? [])[y]})
  adj.push({ k: `${x+1}-${y}`, w: (matrix[x+1] ?? [])[y]})
  adj.push({ k: `${x}-${y-1}`, w: (matrix[x] ?? [])[y-1]})
  adj.push({ k: `${x}-${y+1}`, w: (matrix[x] ?? [])[y+1]})

  return adj.filter(i => !!i.w)
}
