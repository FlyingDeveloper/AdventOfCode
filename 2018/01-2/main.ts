import { default as fs } from 'fs';

fs.readFile('input.txt', { encoding: 'utf8' }, (err:any, input:string) => {
  let visitedNumbers:{[num:number]: boolean} = {};
  let splitInput = input.split('\n').filter(x => x !== '').map(Number)
  let i = 0;
  let accumulator = 0;
  
  while(true) {
    let current = splitInput[i % splitInput.length];
    if (current === undefined || isNaN(current)) {
      continue;
    }

    accumulator = accumulator + current;

    if (visitedNumbers[accumulator] === true) {
      console.log(accumulator);
      break;
    }

    visitedNumbers[accumulator] = true;
    ++i;
  }
});
