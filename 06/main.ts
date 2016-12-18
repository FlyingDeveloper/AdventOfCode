import fs = require("fs");
import _ = require('lodash');

var input = fs.readFileSync("input.txt", { encoding: "utf8" });

class Position {
    characterCounts = {};
    constructor() { };

    incrementCharacter(character:string) {
        if (character.length !== 1) {
            throw "Input character string must have length of 1";
        }

        if (!this.characterCounts.hasOwnProperty(character)) {
            this.characterCounts[character] = 0;
        }

        this.characterCounts[character]++;
    }

    getMostCommonCharacter(): string {
        var currentMax = -1;
        var currentMaxCharacter = '';
        _.forOwn(this.characterCounts, (value:number, key, object) => {
            if (value > currentMax) {
                currentMaxCharacter = key;
                currentMax = value;
            }
        });

        return currentMaxCharacter;
    }
}

var outputPositions = input.split('\n').reduce((previousValue: Array<Position>, currentValue, currentIndex, array) => {
    if (currentValue !== '') {
        currentValue.split('').forEach((character, currentIndex) => {
            if (previousValue.length <= currentIndex) {
                previousValue.push(new Position());
            }
            previousValue[currentIndex].incrementCharacter(character)
        });
    }

    return previousValue;
}, new Array<Position>());

var output = outputPositions.reduce((accumulator, current, index) => accumulator += current.getMostCommonCharacter(), '');

console.log(output);