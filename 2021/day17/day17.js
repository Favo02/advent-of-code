const fs = require("fs")
const path = require("path")

const target = parseInput(process.argv[2])
const { my, cont } = findAll(target)
console.log("max y:", my)
console.log("cont:", cont)

// parse input from file given as process arg
function parseInput(filepath) {
  const inp = fs
    .readFileSync(path.join(__dirname, filepath), { encoding: "utf-8"})
    .split("\n")[0]
    .split("target area:")[1]
    .split(", ")

  const [ x1, x2 ] = inp[0].substring(3).split("..")
  const [ y1, y2 ] = inp[1].substring(2).split("..")

  return {
    x1: parseInt(x1), x2: parseInt(x2),
    y1: parseInt(y1), y2: parseInt(y2)
  }
}

// return true if point (x,y) is inside target
function isInsideTarget(target, x, y) {
  if (x >= target.x1 && x <= target.x2) {
    if (y >= target.y1 && y <= target.y2) {
      return true
    }
  }
  return false
}

// returns new values for x, y, xvel, yvel after 1 time
function newPosition(x, y, xvel, yvel) {

  x += xvel
  y += yvel

  if (xvel > 0) xvel--
  if (xvel < 0) xvel++

  yvel--

  return { x, y, xvel, yvel }
}

// returns true if the point cannot reach the target
function isOver(target, x, y, xvel, yvel) {
  if (y < target.y1) return true // y already under targer
  if (xvel <= 0 && x < target.x1) return true // x before target and x velocity not climbing up
  return false
}

// returns true if the starting velocities make the point fall into the target
function canFall(gx, gy, gxvel, gyvel, target) {
  let maxY = 0
  while(!isOver(target, gx, gy, gxvel, gyvel)) {
    const { x, y, xvel, yvel } = newPosition(gx, gy, gxvel, gyvel)
    gx = x
    gy = y
    gxvel = xvel
    gyvel = yvel
    if (y > maxY) maxY = y
    const fall = isInsideTarget(target, gx, gy)
    // console.log(gx, gy, fall, xvel, yvel)
    if (fall) return { res: true, maxY }
  }
  return { res: false }
}

// finds all initial velocities that fall into target. returns number of velocities (cont) and max y reached (my)
function findAll(target) {
  let my = 0
  let cont = 0

  for (let y = (target.y1)-1; y < 1000; y++) {
    for (let x = 1; x < 1000; x++) {     
      const { res, maxY } = canFall(0, 0, x, y, target)
      if (res) {
        cont++
        // console.log(x,y)
        if (maxY > my) my = maxY
      }
    }
  }

  return {my, cont}
}
