import fs = require('fs');

var input = fs.readFileSync('input.txt', { encoding: 'utf8' });

const alphabet: Array<string> = ["a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"];

function shiftCharacter(character: string, distance: number) {
    if (character.length !== 1) {
        throw 'character input should be a string of length 1';
    }

    if (character === '-') {
        return ' ';
    }

    var currentIndex = alphabet.indexOf(character);
    var newIndex = (currentIndex + distance) % alphabet.length;

    return alphabet[newIndex];
}

function shiftString(stringToShift: string, distance: number) {
    var inputArray = stringToShift.split('');
    var outputArray = inputArray.map(i => shiftCharacter(i, distance));
    return outputArray.join('');
}

var results = input.split('\n').map(a => {
    if (a.length === 0) {
        return '';
    }
    var [whole, string, distance] = /([a-z-]+)-(\d+)/.exec(a);
    return `${shiftString(string, parseInt(distance))} -- ${distance}`;
});//.forEach(console.log);

console.log(results.length);
results.forEach(i => console.log(i));