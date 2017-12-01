import fs = require("fs");

var showSteps = process.env.SHOW_STEPS === '1';

var input = fs.readFileSync("input.txt", { encoding:'utf8' });
var instructionSets:Array<string> = input.split('\n');

class Point {
    row: number;
    col: number;

    constructor(row:number, col:number) {
        this.row = row;
        this.col = col;
    }

    clone() {
        return new Point(this.row, this.col);
    }
}

enum MovesAllowed {
    U = 1,
    D = 2,
    L = 4,
    R = 8
}

function getEnumFromString(direction: string): MovesAllowed {
    switch(direction) {
        case 'U':
            return MovesAllowed.U;
        case 'D':
            return MovesAllowed.D;
        case 'L':
            return MovesAllowed.L;
        case 'R':
            return MovesAllowed.R;
    }
}

class Position {
    point: Point;
    position: string;
    movesAllowed: MovesAllowed;
}

var positions = new Array<Position>();
positions.push({ point: new Point(1, 3), position: '1', movesAllowed: MovesAllowed.D });
positions.push({ point: new Point(2, 2), position: '2', movesAllowed: MovesAllowed.R | MovesAllowed.D });
positions.push({ point: new Point(2, 3), position: '3', movesAllowed: 15 /* all */ });
positions.push({ point: new Point(2, 4), position: '4', movesAllowed: MovesAllowed.L | MovesAllowed.D });
positions.push({ point: new Point(3, 1), position: '5', movesAllowed: MovesAllowed.R });
positions.push({ point: new Point(3, 2), position: '6', movesAllowed: 15 });
positions.push({ point: new Point(3, 3), position: '7', movesAllowed: 15 });
positions.push({ point: new Point(3, 4), position: '8', movesAllowed: 15 });
positions.push({ point: new Point(3, 5), position: '9', movesAllowed: MovesAllowed.L });
positions.push({ point: new Point(4, 2), position: 'A', movesAllowed: MovesAllowed.U | MovesAllowed.R });
positions.push({ point: new Point(4, 3), position: 'B', movesAllowed: 15 });
positions.push({ point: new Point(4, 4), position: 'C', movesAllowed: MovesAllowed.U | MovesAllowed.L });
positions.push({ point: new Point(5, 3), position: 'D', movesAllowed: MovesAllowed.U });

function getPointFromPosition(position: string):Point {
    var filtered = positions.filter(i=>i.position === position);
    if (filtered.length === 0) {
        throw `Unable to find point at position ${position}`;
    }
    
    return filtered[0].point;
}

function getPositionFromPoint(point: Point):Position {
    var filtered = positions.filter(i=> {
        return i.point.row === point.row && i.point.col === point.col;
    });
    if (filtered.length === 0) {
        throw `Unable to find position at point ${point.row}, ${point.col}`;
    }
    
    return filtered[0];
}

var accumulator = getPointFromPosition('5').clone();

var positionMap = instructionSets.map(instructionSet => {
    var instructions = instructionSet.split('');
    instructions.forEach(input=> {
        var currentPosition = getPositionFromPoint(accumulator);
        if (currentPosition.movesAllowed & getEnumFromString(input)) {
            switch (input) {
                case 'U':
                    accumulator.row--;
                    break;
                case 'D':
                    accumulator.row++;
                    break;
                case 'L':
                    accumulator.col--;
                    break;
                case 'R':
                    accumulator.col++;
                    break;
            }
        }

        showSteps && console.log(`-${input} -- ${accumulator.row}, ${accumulator.col} -- ${getPositionFromPoint(accumulator).position}`);
    });

    showSteps && console.log(`----\n${getPositionFromPoint(accumulator).position}`);
    return getPositionFromPoint(accumulator).position;
});

console.log(positionMap.join(''));