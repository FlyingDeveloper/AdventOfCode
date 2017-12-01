import c = require('crypto');

const relevantHashRegex = /^00000/;
const passwordLength = 8;
const input = 'cxdnnyjw';

function findRelevantHash(input: string, startingIndex?: number): { index: number, hash: string } {
    for (var i = startingIndex || 0; i < Number.MAX_VALUE; i++) {
        var stringToHash = input + i.toString();
        var h = c.createHash('md5');
        h.update(stringToHash);
        var hash: string = h.digest('hex');
        var isRelevantHash = relevantHashRegex.test(hash);
        if (isRelevantHash) {
            return {
                index: i,
                hash: hash
            };
        }
    }

    throw 'No relevant hash found';
}

var currentIndex = 0;
var output = '________'.split('');
while (output.indexOf('_') !== -1) {
    var result = findRelevantHash(input, currentIndex);
    if (isNaN(parseInt(result.hash[5])) || parseInt(result.hash[5]) >= passwordLength) {
        currentIndex = result.index + 1;
        continue;
    }
    var position = parseInt(result.hash[5]);
    if (output[position] === '_') {
        output[position] = result.hash[6];
        console.log(output.map(a => a || '_').join(''));
    }
    currentIndex = result.index + 1;
}