const fs = require("fs")
const path = require("path")

const lines = parseInput(process.argv[2])

const matrix = []
buildGraph(lines)

expandHorizontally(4)
expandVertically(4)

// const target = "49-49"
const target = "499-499"
const dist = dijkstra(0, 0, target)
console.log(dist[target])
// console.log(dist)

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
function dijkstra(startX, startY, target) {

  const dist = {}

  let queue = []

  for (let y = 0; y < matrix.length; y++) {
    for (let x = 0; x < matrix[y].length; x++) {
      dist[`${x}-${y}`] = { dist: 100000, manDist: manhattanDistance(startX, startY, x, y) }
      queue.push(`${x}-${y}`)
    }
  }

	dist[`${startX}-${startY}`].dist = 0

  while (queue.length > 0) {

    let min = queue[0]
    for (let i = 0; i < queue.length; i++) {
      const el = queue[i]

      if (dist[el].dist === 100000) {
        continue
      }

      // dijkstra
      // if (dist[el] < dist[min]) {
      //   min = el
      // }

      // a*
      if ((dist[el].dist * (500 - dist[el].manDist)) < (dist[min].dist * (500 - dist[el].manDist))) {
        min = el
      }
    }

    queue = queue.filter(e => e !== min)

    const adj = getAdjacent(min)

    adj.forEach(ad => {
      if ((dist[min].dist + parseInt(ad.w)) < dist[ad.k].dist) {
        dist[ad.k].dist = dist[min].dist + parseInt(ad.w)
      }
    })

    if (min === target) {
      return dist
    }
	}

	return dist
}

// get adjacent elements: up, down, left, right
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

function expandHorizontally(steps) {

  const length = matrix.length

  for (let step = 0; step < steps; step++) {
    for (let i = 0; i < matrix.length; i++) {
      const source = matrix[i].slice(length*step)
      matrix[i] = [...matrix[i], ...newWeight(...source)]
    }
  }
}

function expandVertically(steps) {
  const heigth = matrix.length*steps

  for (let i = 0; i < heigth; i++) {
    matrix.push(newWeight(...matrix[i]))
  }
}

function newWeight(...arr) {
  for (let i = 0; i < arr.length; i++) {
    // process.stdout.write(arr[i] + "->")
    arr[i] = arr[i] == 8 ? 9 : ((arr[i]+1) %9) // this wont work with === :)
    // console.log(arr[i])
  }
  return arr
}

function printMatrix() {
  for (let y = 0; y < matrix.length; y++) {
    console.log(matrix[y].join(""))
  }
}

function manhattanDistance(aX, aY, bX, bY) {
  return Math.abs(bX-aX) + Math.abs(bY-bY)
}
