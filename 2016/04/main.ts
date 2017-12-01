import fs = require('fs');

var input = fs.readFileSync('input.txt', { encoding: 'utf8' });
var inputLines = input.split('\n');

var inputFormat = /([a-z-]+)-(\d+)\[([a-z]{5})\]/;

function validateChecksum(encryptedString: string, possibleChecksum: string): boolean {
    var characterCounts = new Map<string, number>();
    var encryptedName = encryptedString
    for (var i = 0; i < encryptedName.length; i++) {
        var currentCharacter = encryptedName[i];
        if (!characterCounts.get(currentCharacter)) {
            characterCounts.set(currentCharacter, 1);
        } else {
            characterCounts.set(
                currentCharacter,
                characterCounts.get(currentCharacter) + 1
            );
        }
    }

    var a = new Array<{ character: string, count: number }>();
    characterCounts.forEach((count, character) => a.push({ character: character, count: count }));
    a.sort((a, b) => {
        if (a.count === b.count) {
            return (a.character < b.character) ? 1 : -1;
        }

        return (a.count > b.count) ? 1 : -1;
    });

    var checksum = a
        .reverse()
        .filter(item => item.character !== '-')
        .map(item => item.character)
        .join('')
        .substring(0, 5);
    return checksum === possibleChecksum;
}

var sectorSum = inputLines.reduce((accumulator, line) => {
    if (line === '') {
        return accumulator;
    }

    var parsed = inputFormat.exec(line);
    var encryptedString:string = parsed[1];
    var sectorId: number = parseInt(parsed[2]);
    var checksum:string = parsed[3];

    var isValidChecksum = validateChecksum(encryptedString, checksum);
    var returnValue: number;
    if (isValidChecksum) {
        returnValue = accumulator + sectorId;
    } else {
        returnValue = accumulator;
    }
    return returnValue;
}, 0);

console.log(sectorSum);