const fs = require('fs');

function peek(input, currentIndex) {
  if (currentIndex === input.length - 1) {
    return input[0];
  }

  return input[currentIndex + 1];
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
console.log(processContents('1122'));
console.log(processContents('1111'));
console.log(processContents('1234'));
console.log(processContents('91212129'));
*/

fs.readFile('input.txt', { encoding: 'utf8' }, function(err, contents) {
  console.log(processContents(contents.replace('\n', '')));
});
