enum Direction {
  Up = 0,
  Down = 1,
  Left = 2,
  Right = 4
}

const memoryOffset = 1000;

let rows:Array<Array<number>> = [];
for (var i = 0; i < memoryOffset * 10; i++) {
  rows[i] = [];
}

const destination = 368078;

let minX = 0;
let minY = 0;
let maxX = 0;
let maxY = 0;
let currentX = 0;
let currentY = 0;
let current = 1;
let currentDirection = Direction.Right;

function getValueAtSquare(x:number, y:number) {
  return rows[y + memoryOffset][x + memoryOffset];
}

function getAdjacentValues() {
  var topLeft = getValueAtSquare(currentX - 1, currentY + 1) || 0;
  var top = getValueAtSquare(currentX, currentY + 1) || 0;
  var topRight = getValueAtSquare(currentX + 1, currentY + 1) || 0;
  var right = getValueAtSquare(currentX + 1, currentY) || 0;
  var bottomRight = getValueAtSquare(currentX + 1, currentY - 1) || 0;
  var bottom = getValueAtSquare(currentX, currentY - 1) || 0;
  var bottomLeft = getValueAtSquare(currentX - 1, currentY - 1) || 0;
  debugger;
  var left = getValueAtSquare(currentX - 1, currentY) || 0;

  return topLeft + top + topRight + right + bottomRight + bottom + bottomLeft + left;
}

function turn() {
  switch(currentDirection) {
    case Direction.Right:
      currentDirection = Direction.Up;
      break;
    case Direction.Up:
      currentDirection = Direction.Left;
      break;
    case Direction.Left:
      currentDirection = Direction.Down;
      break;
    case Direction.Down:
      currentDirection = Direction.Right;
      break;
    default:
      throw "Not implemented";
  }
}

let turnOnNextMove = false;
rows[currentY + memoryOffset][currentX + memoryOffset] = 1;
let lastValueWritten = 1;

while (current <= destination && lastValueWritten <= destination) {
  if (current !== 1) {
    let currentValue = getAdjacentValues();
    rows[currentY + memoryOffset][currentX+memoryOffset] = currentValue;
    lastValueWritten = currentValue;
  }
  if (turnOnNextMove === true) {
    turn();
    turnOnNextMove = false;
  }

  switch (currentDirection) {
    case Direction.Right:
      ++currentX;
      if (currentX > maxX) {
        maxX = currentX;
        turnOnNextMove = true;
      }
      break;
    case Direction.Up:
      ++currentY;
      if (currentY > maxY) {
        maxY = currentY;
        turnOnNextMove = true;
      }
      break;
    case Direction.Left:
      --currentX;
      if (currentX < minX) {
        minX = currentX;
        turnOnNextMove = true;
      }
      break;
    case Direction.Down:
      --currentY;
      if (currentY < minY) {
        minY = currentY;
        turnOnNextMove = true;
      }
      break;
    default:
      throw "Not implemented";
  }

  ++current;
}

console.log(lastValueWritten);

/*for (var i = 5; i > -5; i--) {
  let output = [`${i}\t`];
  for (var j = -5; j < 5; j++) {
    output.push(`${rows[i + memoryOffset][j+memoryOffset] || 'u'}\t`);
  }
  console.log(output.join(''));
}*/
