const fs = require('fs');

interface Array<T> {
  find(predicate: (search: T, index: number) => boolean) : T;
}

function processLine(line:string) {
  if (line.length === 0) {
    return 0;
  }

  const fields = line.split('\t');
  let output:number = 0;
  fields.map(Number).forEach((field, fieldIndex) => {
    fields.map(Number).find((currentDivisor, currentDivisorIndex) => {
      if (field % currentDivisor === 0 && fieldIndex !== currentDivisorIndex) {
        output = field / currentDivisor;
        return true;
      }

      return false;
    });
  });

  return output;
}

function processFile(err:any, input:string) {
  const lines = input.split('\n');
  var checksum = lines
    .map(line => processLine(line))
    .reduce((accumulator, current) => {
      return accumulator + current;
    }, 0);
  console.log(checksum);
}

/* test case */
//processFile(undefined, '5\t9\t2\t8\n9\t4\t7\t3\n3\t8\t6\t5\n'); // expect 9

fs.readFile('input.txt', { encoding: 'utf8' }, processFile);
