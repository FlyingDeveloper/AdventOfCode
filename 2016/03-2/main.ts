import fs = require('fs');

var input = fs.readFileSync('input.txt', {encoding: 'utf8'});
var possibleTriangles = input.split('\n');

var inputParser = /\s*(\d+)\s+(\d+)\s+(\d+)/;
var accumulatorA = new Array<number>();
var accumulatorB = new Array<number>();
var accumulatorC = new Array<number>();
var accumulators = [accumulatorA, accumulatorB, accumulatorC];

var possibleTriangleCount = 0;

function testForPossibleTriangle(a:number, b:number, c:number):boolean {
    var sum = a + b + c;
    if (sum - a > a && sum - b > b && sum - c > c) {
        return true;
    }
    
    return false;
}

var result = possibleTriangles.forEach(line => {
    var pttResult = inputParser.exec(line);
    if (!pttResult) {
        return;
    }

    var [full, aString, bString, cString] = pttResult;
    var a = parseInt(aString);
    var b = parseInt(bString);
    var c = parseInt(cString);

    accumulatorA.push(a);
    accumulatorB.push(b);
    accumulatorC.push(c);

    accumulators.forEach(accumulator => {
        if (accumulator.length === 3) {
            var result = testForPossibleTriangle(accumulator.pop(), accumulator.pop(), accumulator.pop());
            result && possibleTriangleCount++;
        }
    });
});

console.log(`Number of possible triangles: ${possibleTriangleCount}`);