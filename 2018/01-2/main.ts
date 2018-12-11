import { default as fs } from 'fs';
import Utils from 'utils';

Utils.GetLinesAsNumber().then(splitInput => {
  let visitedNumbers:{[num:number]: boolean} = {};
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
