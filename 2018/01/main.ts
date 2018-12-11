import { default as fs } from 'fs';
import Utils from 'utils';

Utils.GetLinesAsNumber().then(lines => {
  let result = lines.reduce((current:number, accumulator:number) => {
    if (isNaN(current)) {
      return accumulator;
    }

    return accumulator + current;
  }, 0);

  console.log(result);
});
