enum Direction {
  Up = 0,
  Down = 1,
  Left = 2,
  Right = 4
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

while (current <= destination) {
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

console.log(`${Math.abs(currentX) + Math.abs(currentY)}`);
