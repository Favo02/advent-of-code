const fs = require("fs")
const path = require("path")

const hex = parseInput(process.argv[2])
const bin = hex2bin(hex)

const package = analyzePackage(bin)
// console.log(JSON.stringify(package, null, 2))

const version = sumVersions(package)
console.log("version", version)

const value = calculatePackageValue(package)
console.log("value", value)

// parse input from file given as process arg
function parseInput(filepath) {
  return fs
    .readFileSync(path.join(__dirname, filepath), { encoding: "utf-8"})
}

// hexadecimal to binary
function hex2bin(hex){
  hex = hex.toLowerCase()
  let out = ""

  for(let c of hex) {
    switch(c) {
      case '0': out += "0000"; break
      case '1': out += "0001"; break
      case '2': out += "0010"; break
      case '3': out += "0011"; break
      case '4': out += "0100"; break
      case '5': out += "0101"; break
      case '6': out += "0110"; break
      case '7': out += "0111"; break
      case '8': out += "1000"; break
      case '9': out += "1001"; break
      case 'a': out += "1010"; break
      case 'b': out += "1011"; break
      case 'c': out += "1100"; break
      case 'd': out += "1101"; break
      case 'e': out += "1110"; break
      case 'f': out += "1111"; break
    }
  }

  return out
}

function analyzePackage(p) {
  let info = {}

  info.version = p.substring(0, 3)
  info.type = parseInt(p.substring(3, 6), 2)
  const payload = p.substring(6)

  // literal type
  if (info.type == 4) {
    const parsed = parseLiteralPackage(payload)
    info.literal = parsed.literal
    info.remaining = parsed.remaining
  }
  // operator type
  else {
    const parsed = parseOperatorPackage(payload)
    info.packages = parsed.packages
    info.remaining = parsed.remaining
  }

  return info
}

function parseLiteralPackage(payload) {
  const literalGroups = chunkSubstr(payload, 5)
  let literal = ""

  for (let i = 0; i < literalGroups.length; i++) {
    const group = literalGroups[i]
    
    literal += group.substr(1)

    if (group[0] == '0') {
      return { literal, remaining: literalGroups.slice(i+1).join("") }
    }

  }
}

// split string in chunks by size
function chunkSubstr(str, size) {
  const numChunks = Math.ceil(str.length / size)
  const chunks = new Array(numChunks)

  for (let i = 0, o = 0; i < numChunks; ++i, o += size) {
    chunks[i] = str.substr(o, size)
  }

  return chunks
}

function parseOperatorPackage(payload) {
  const lengthType = payload[0]

  // the next 15 bits are a number that represents the total length in bits of the sub-packets contained by this packet
  if (lengthType == '0') {
    const subLengthStr = payload.substring(1, 16)
    const subLength = parseInt(subLengthStr, 2)

    let subPayload = payload.substring(16, (16+subLength))

    const packages = []
    while(subPayload.length > 0) {
      // console.log("sp", subPayload, typeof subPayload)
      const analyzed = analyzePackage(subPayload)
      // console.log(analyzed)
      subPayload = analyzed.remaining
      packages.push(analyzed)
    }
    
    return { packages, remaining: payload.substring(16+subLength)}
  }

  // the next 11 bits are a number that represents the number of sub-packets immediately contained by this packet
  else {
    const subPackageNumberStr = payload.substring(1, 12)
    const subPackageNumber = parseInt(subPackageNumberStr, 2)
    // console.log("spn", subPackageNumber)

    const res = analyzePackage(payload.substring(12))

    let subPayload = payload.substring(12)

    // console.log(res)
    const packages = []
    let analyzedN = 0
    while(subPayload.length > 0 && analyzedN < subPackageNumber) {
      // console.log("sp", subPayload, typeof subPayload)
      const analyzed = analyzePackage(subPayload)
      // console.log(analyzed)
      subPayload = analyzed.remaining
      packages.push(analyzed)
      analyzedN++
    }

    return { packages, remaining: subPayload }
  }
}

// sum versions of package recursively
function sumVersions(package) {
  let sum = 0

  for (let key in package) {
    if (typeof package[key] === 'object') {
      sum += sumVersions(package[key])
    }
    else if (key === 'version') {
      const ver = parseInt(package[key], 2)
      sum += ver
    }
  }

  return sum
}

function calculatePackageValue(package) {
  if (package.type == 4) {
    return parseInt(package.literal, 2)
  }

  const values = [...package.packages].map(p => calculatePackageValue(p))

  // console.log(packagesValues)

  switch(package.type) {
    
    case 0: // sum
      return values.reduce((sum, a) => sum + a, 0)

    case 1: // product
      return values.reduce((prod, a) => prod * a, 1)

    case 2: // min
      return Math.min(...values)

    case 3: // max
      return Math.max(...values)

    case 5: // gt
      return values[0] > values[1] ? 1 : 0

    case 6: // lt
      return values[0] < values[1] ? 1 : 0

    case 7: // eq
      return values[0] == values[1] ? 1 : 0

    default:
      console.log("error")
  }
}
