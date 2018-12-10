import { default as fs } from 'fs';

fs.readFile('input.txt', { encoding: 'utf8' }, (err:any, input:string) => {
  let result = input.split('\n').map(Number).reduce((current:number, accumulator:number) => {
    if (isNaN(current)) {
      return accumulator;
    }

    return accumulator + current;
  }, 0);

  console.log(result);
});
