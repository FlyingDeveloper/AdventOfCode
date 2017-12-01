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
                hash: hash,
            };
        }
    }

    throw 'No relevant hash found';
}

var currentIndex = 0;
for (var i = 0; i < passwordLength; i++) {
    var result = findRelevantHash(input, currentIndex);
    console.log(result.hash[5]);
    currentIndex = result.index + 1;
}