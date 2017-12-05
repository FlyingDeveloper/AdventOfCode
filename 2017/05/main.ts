import * as fs from 'fs';

function processFile(contents: string) {
    if (contents.endsWith('\n')) {
        contents = contents.slice(0, contents.length - 1);
    }
    const instructions = contents.split('\n').map(Number);
    let currentPosition = 0;
    let numberOfMoves = 0;
    while (currentPosition >= 0 && currentPosition < instructions.length) {
        let currentMove = instructions[currentPosition];
        ++instructions[currentPosition];
        currentPosition += currentMove;
        ++numberOfMoves;
    }

    console.log(numberOfMoves);
}

/* Test case */
//processFile('0\n3\n0\n1\n-3\n'); // 5

fs.readFile("input.txt", { encoding: 'utf8' }, (err, contents) => processFile(contents));