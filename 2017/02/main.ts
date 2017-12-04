const fs = require('fs');

function processLine(line:string) {
  if (line.length === 0) {
    return 0;
  }

  const fields = line.split('\t');
  let smallest = Number.MAX_VALUE;
  let largest = Number.MIN_VALUE;
  fields.map(x => Number(x)).forEach(field => {
    if (field > largest) {
      largest = field;
    }

    if (field < smallest) {
      smallest = field;
    }
  });

  return largest - smallest;
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
// console.log(processFile(undefined, '5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8')); // expect 18

fs.readFile('input.txt', { encoding: 'utf8' }, processFile);
