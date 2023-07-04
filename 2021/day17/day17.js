const fs = require("fs")
const path = require("path")

const target = parseInput(process.argv[2])
// console.log(target)

// const { res, maxY } = canFall(0, 0, 6, 9, target)

// if (res) {
//   console.log("maxY:", maxY)
// }

const {my, cont} = tryAll(target)
console.log(my, cont)

const debug = debugParseInput("out corretto")
// console.log(debug)

for (let i = 0; i < debug.length; i++) {
  const { res } = canFall(0, 0, debug[i].x, debug[i].y, target)
  if (!res) {
    console.log(debug[i].x, debug[i].y, res)
  }
}

console.log(canFall(0,0,7,-1,target))

function debugParseInput(filepath) {
  return fs
    .readFileSync(path.join(__dirname, filepath), { encoding: "utf-8"})
    .split("\n")
    .slice(0, -1)
    .map(l => {
      const [x,y] = l.split(",")
      return {x: parseInt(x), y: parseInt(y)}
    })
}

// parse input from file given as process arg
function parseInput(filepath) {
  const inp = fs
    .readFileSync(path.join(__dirname, filepath), { encoding: "utf-8"})
    .split("\n")[0]
    .split("target area:")[1]
    .split(", ")

  const [ xs, xe ] = inp[0].substring(3).split("..")
  const [ ys, ye ] = inp[1].substring(2).split("..")

  return {
    xs: parseInt(xs), xe: parseInt(xe),
    ys: parseInt(ys), ye: parseInt(ye)
  }
}

// return true if point (x,y) is inside target
function isInsideTarget(target, x, y) {
  if (x >= target.xs && x <= target.xe) {
    if (y >= target.ys && y <= target.ye) {
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
function isOver(target, x, y) {
  return y < (target.ye-1)
}

// returns true if the starting velocities make the point fall into the target
function canFall(gx, gy, gxvel, gyvel, target) {
  let maxY = 0
  while(true) {
    const { x, y, xvel, yvel } = newPosition(gx, gy, gxvel, gyvel)
    gx = x
    gy = y
    gxvel = xvel
    gyvel = yvel
    if (y > maxY) maxY = y
    const fall = isInsideTarget(target, gx, gy)
    // console.log(gx, gy, fall, xvel, yvel)
    if (fall) return { res: true, maxY }

    if (isOver(target, gx, gy)) break
  }
  return { res: false }
}

function tryAll(target) {
  let my = 0
  let cont = 0

  for (let y = -1000; y < 1000; y++) {
    for (let x = -1000; x < 1000; x++) {     
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
