import fs = require("fs");

var showSteps = process.env.SHOW_STEPS === '1';

class Instruction {
    turnDirection:string;
    distance:number;
    constructor(instruction:string) {
        var validRegex = /^([LR]{1})(\d+)$/;
        var regexResult = validRegex.exec(instruction);

        if (!regexResult) {
            throw `Invalid instruction - '${instruction}'`;
        }

        this.turnDirection = regexResult[1];
        this.distance = parseInt(regexResult[2]);
    }

    toString() {
        return `${this.turnDirection}${this.distance}`;
    }
}

class Point {
    x:number;
    y:number;

    constructor(x:number, y:number) {
        this.x = x;
        this.y = y;
    }

    toString() : string {
        return `${this.x}, ${this.y}`;
    }
}

enum Direction {
    North,
    South,
    East,
    West
}

var input = fs.readFileSync('input.txt', {encoding:'utf8'});

var instructions = new Array<Instruction>();

input.split(',').forEach(i=>instructions.push(new Instruction(i.trim())));

var currentPosition = new Point(0, 0);
var currentDirection = Direction.North;

function makeMove(direction:Direction, distance:number) {
    switch (direction) {
        case Direction.North:
          currentPosition.y = currentPosition.y + distance;
          break;
        case Direction.East:
          currentPosition.x = currentPosition.x + distance;
          break;
        case Direction.South:
          currentPosition.y = currentPosition.y - distance;
          break;
        case Direction.West:
          currentPosition.x = currentPosition.x - distance;
          break;
    }
}

instructions.forEach(i=>{
    switch (currentDirection) {
        case Direction.North:
            switch (i.turnDirection) {
                case "L":
                    currentDirection = Direction.West;
                    break;
                case "R":
                    currentDirection = Direction.East;
                    break;
            }
            break;
        case Direction.East:
            switch(i.turnDirection) {
                case "L":
                    currentDirection = Direction.North;
                    break;
                case "R":
                    currentDirection = Direction.South;
                    break;
            }
            break;
        case Direction.South:
            switch(i.turnDirection) {
                case "L":
                    currentDirection = Direction.East;
                    break;
                case "R":
                    currentDirection = Direction.West;
                    break;
            }
            break;
        case Direction.West:
            switch(i.turnDirection) {
                case "L":
                    currentDirection = Direction.South;
                    break;
                case "R":
                    currentDirection = Direction.North;
                    break;
            }
    }
    makeMove(currentDirection, i.distance);
    showSteps && console.log(currentPosition.toString());
});

console.log(currentPosition.toString());

var TotalDistance = currentPosition.x + currentPosition.y;
console.log(`Total distance: ${TotalDistance}`);