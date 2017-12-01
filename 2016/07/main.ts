import fs = require('fs');
import process = require('process');

var input = fs.readFileSync("input.txt", { encoding: 'utf8' });

var result = input.split('\n').filter(line => {
    var stageOneTest = /\[[^\]]*(.)(.)(\2)(\1).*\]/;
    var stageOneResult = stageOneTest.exec(line);
    if (!stageOneResult) {
        var stageTwoTest = /(.)(.)(\2)(\1)/;
        var stageTwoResult = stageTwoTest.exec(line);
        if (stageTwoResult) {
            var innerCharsAreDifferentFromOuterChars = stageTwoResult[1] !== stageTwoResult[2];
            if (innerCharsAreDifferentFromOuterChars) {
                // We have a match!
                return true;
            }
        }
    }
});

console.log(result.length);