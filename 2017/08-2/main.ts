import fs = require('fs');

const instruction_regex = /([\S]+)\s(\S{3})\s(-?\d+)\sif\s([\S]+)\s([\!=><]+)\s(-?\d+)/;
const registers: Map<string, number> = new Map();

class Condition {
    register: string;
    operator: string;
    value: number;

    constructor(register: string, operator: string, value: number) {
        this.register = register;
        this.operator = operator;
        this.value = value;
    }

    public evaluate(): boolean {
        var registerValue = registers.get(this.register);
        if (registerValue === undefined) {
            registerValue = 0;
        }

        switch (this.operator) {
            case '==':
                return registerValue == this.value;
            case '>':
                return registerValue > this.value;
            case '<':
                return registerValue < this.value;
            case '>=':
                return registerValue >= this.value;
            case '<=':
                return registerValue <= this.value;
            case '!=':
                return registerValue != this.value;
            default:
                throw 'Unknow operator specified: ' + this.operator;
        }
    }
}

function processInstruction(instruction: string) {
    const parsed = instruction_regex.exec(instruction);
    if (parsed === null) {
        throw `Invalid instruction: ${instruction}`;
    }

    const registerToModify = parsed[1];
    const direction = parsed[2];
    const amount = Number(parsed[3]);
    const condition = new Condition(parsed[4], parsed[5], Number(parsed[6]));

    if (condition.evaluate()) {
        const register = registers.get(registerToModify);
        let currentValue: number;
        if (register === undefined) {
            currentValue = 0;
        } else {
            currentValue = register;
        }

        if (direction === 'inc') {
            registers.set(registerToModify, currentValue + amount);
        } else {
            registers.set(registerToModify, currentValue - amount);
        }
    }
}

function findLargestRegister(): [string, number] {
    let max: [string, number] = ['unknown', Number.MIN_VALUE];
    registers.forEach((value, key) => {
        if (value > max[1]) {
            max[0] = key;
            max[1] = value;
        }
    });

    return max;
}

/*
Test case
const instructions = [
    'b inc 5 if a > 1',
    'a inc 1 if b < 5',
    'c dec -10 if a >= 1',
    'c inc -20 if c == 10'
];

instructions.forEach(processInstruction);
registers.forEach((value, key) => console.log(`${key}: ${value}`));
*/

fs.readFile('input.txt', { encoding: 'utf8' }, (err, data) => {
    let largestValueSeen = Number.MIN_VALUE;
    data.split('\n').filter(s => s.length > 0).reduce((accumulator, current) => {
        processInstruction(current);
        let currentLargest = findLargestRegister();
        if (currentLargest[1] > largestValueSeen) {
            largestValueSeen = currentLargest[1];
        }
        return currentLargest
    }, ['unknown', Number.MIN_VALUE]);
    console.log(largestValueSeen);
});