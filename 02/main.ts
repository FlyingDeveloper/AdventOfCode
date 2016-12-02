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

var positions = new Array<{ point:Point, position: string}>();
positions.push({ point: new Point(1, 1), position: '1' });
positions.push({ point: new Point(1, 2), position: '2' });
positions.push({ point: new Point(1, 3), position: '3' });
positions.push({ point: new Point(2, 1), position: '4' });
positions.push({ point: new Point(2, 2), position: '5' });
positions.push({ point: new Point(2, 3), position: '6' });
positions.push({ point: new Point(3, 1), position: '7' });
positions.push({ point: new Point(3, 2), position: '8' });
positions.push({ point: new Point(3, 3), position: '9' });

function getPointFromPosition(position: string):Point {
    var filtered = positions.filter(i=>i.position === position);
    if (filtered.length === 0) {
        throw `Unable to find point at position ${position}`;
    }
    
    return filtered[0].point;
}

function getPositionFromPoint(point: Point):string {
    var filtered = positions.filter(i=> {
        return i.point.row === point.row && i.point.col === point.col;
    });
    if (filtered.length === 0) {
        throw `Unable to find position at point ${point.row}, ${point.col}`;
    }
    
    return filtered[0].position;
}

var accumulator = getPointFromPosition('5').clone();

var positionMap = instructionSets.map(is => {
    var instructions = is.split('');
    instructions.forEach(input=> {
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
        if (accumulator.row > 3) {
            accumulator.row = 3;
        } else if (accumulator.row < 1) {
            accumulator.row = 1;
        } else if (accumulator.col > 3) {
            accumulator.col = 3;
        } else if (accumulator.col < 1) {
            accumulator.col = 1;
        }

        showSteps && console.log(`-${input} -- ${accumulator.row}, ${accumulator.col} -- ${getPositionFromPoint(accumulator)}`);
    });

    showSteps && console.log(`----\n${getPositionFromPoint(accumulator)}`);
    return getPositionFromPoint(accumulator);
});

console.log(positionMap.join(''));