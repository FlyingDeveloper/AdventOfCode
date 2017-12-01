import fs = require('fs');

var input = fs.readFileSync("input.txt", "utf8");

var hypernetSeeker = /\[([^\]]*)\]/g;
var supernetSeeker = /([^\[\]]+)(\[|$)/g
var abaSeeker = /([^\[\]])([^\[\]])\1/g;

var output = input.split('\n').filter(line => {
    var supernetSequences = new Array<string>();
    var hypernetSequences = new Array<string>();

    var result: RegExpExecArray;
    while (result = hypernetSeeker.exec(line)) {
        hypernetSequences.push(result[1]);
    }

    while (result = supernetSeeker.exec(line)) {
        supernetSequences.push(result[1]);
    }

    var abaCollection = new Array<{ a: string, b: string }>();
    supernetSequences.forEach(current => {
        var abaResult: RegExpExecArray;
        while (abaResult = abaSeeker.exec(current)) {
            if (abaResult !== null) {
                abaSeeker.lastIndex = abaResult.index + 1;
                abaCollection.push({ a: abaResult[1], b: abaResult[2] });
            }
        }
    });

    var z = new Array<string>();

    return abaCollection.some(aba => {
        var babSeeker = new RegExp(`${aba.b}${aba.a}${aba.b}`);
        var hasBABInHypernetSequences = hypernetSequences.some(current => {
            return babSeeker.test(current);
        });
        return hasBABInHypernetSequences;
    });
});

console.log(output.length);