import fs = require('fs');

var input = fs.readFileSync('input.txt', {encoding: 'utf8'});
var possibleTriangles = input.split('\n');

var possibleTriangleTester = /\s*(\d+)\s+(\d+)\s+(\d+)/

var result = possibleTriangles.filter(pt => {
    var pttResult = possibleTriangleTester.exec(pt);
    if (!pttResult) {
        return false
    }

    var [full, aString, bString, cString] = pttResult;
    var a = parseInt(aString);
    var b = parseInt(bString);
    var c = parseInt(cString);

    var sum = a + b + c;
    if (sum - a > a && sum - b > b && sum - c > c) {
        return true;
    }

    return false;
});

console.log(`Number of possible triangles: ${result.length}`);