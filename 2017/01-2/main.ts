import fs = require('fs');

function peek(input, currentIndex) {
  debugger;
  const peekIndex = (currentIndex + (input.length / 2)) % input.length;

  return input[peekIndex];
}

function processContents(contents) {
  const input = contents.split('');
  let rollingSum = 0;

  for (let currentIndex = 0; currentIndex < input.length; currentIndex++) {
    if (input[currentIndex] === peek(input, currentIndex)) {
      rollingSum = rollingSum + Number(input[currentIndex]);
    }
  }

  return rollingSum;
}

/* Test cases */
/*
console.log(processContents('1212'));
console.log(processContents('1221'));
console.log(processContents('123425'));
console.log(processContents('123123'));
console.log(processContents('12131415'));
*/

fs.readFile('input.txt', { encoding: 'utf8' }, function(err, contents) {
  console.log(processContents(contents.replace('\n', '')));
});
