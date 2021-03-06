import fs = require('fs');
import { isUndefined } from 'util';

function findIndexOfBiggest(banks: Array<number>): number {
    let currentBiggestIndex = Number.MIN_VALUE;
    let currentBiggestValue = Number.MIN_VALUE;
    banks.forEach((x, i) => {
        if (x > currentBiggestValue && x !== currentBiggestIndex) {
            currentBiggestIndex = i;
            currentBiggestValue = x;
        }
    });

    return currentBiggestIndex;
}

function redistribute(banks: Array<number>) {
    const indexOfBiggest = findIndexOfBiggest(banks);
    const numberAtBiggestIndex = banks[indexOfBiggest];
    banks[indexOfBiggest] = 0;
    let currentIndexInBanks = indexOfBiggest;
    for (var i = 0; i < numberAtBiggestIndex; i++) {
        banks[(++currentIndexInBanks) % banks.length]++;
    }
}

function processFile(input: string) {
    const banks = input.split('\t').map(Number);
    let knownPatterns = new Map<string, { timesSeen: number, lastSeenInRedistribution: number }>();
    let numberOfRedistributions = 0;
    while (true) {
        console.log(JSON.stringify(banks));
        redistribute(banks);
        ++numberOfRedistributions;
        let banksAsString = banks.join('-');
        const timesSeenSoFar = knownPatterns.get(banksAsString) || {
            timesSeen: 0,
            lastSeenInRedistribution: Number.MIN_VALUE
        };

        if (timesSeenSoFar.timesSeen === 2) {
            break;
        } else {
            knownPatterns.set(banksAsString, { timesSeen: timesSeenSoFar.timesSeen + 1, lastSeenInRedistribution: numberOfRedistributions });
        }
    }

    console.log(JSON.stringify(banks));
    console.log(numberOfRedistributions);
}

/* Test case */
//processFile([0, 2, 7, 0].join('\t'));

fs.readFile('input.txt', { encoding: 'utf8' }, (err, contents) => processFile(contents));
//9073 is too high