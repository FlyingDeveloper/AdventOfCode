const fs = require('fs');

function isValidPassphrase(line:string): boolean {
  if (line.length === 0) {
    return false;
  }

  var words = line.split(' ');
  var foundWords:Map<string, boolean> = new Map();
  var isPassphraseValid = true;
  words.find((word:string) => {
    if (foundWords.get(word)) {
      isPassphraseValid = false;
      return true; // to break out of loop
    }

    foundWords.set(word, true);
    return false;
  });

  return isPassphraseValid;
}

function processInput(err:any, input:string) {
  const lines = input.split('\n');
  let numberOfValidPassphrases = lines.map(isValidPassphrase).filter(p => p).length;
  console.log(numberOfValidPassphrases);
}

/* test cases */
/*
console.log(isValidPassphrase('aa bb cc dd ee')); // true
console.log(isValidPassphrase('aa bb cc dd aa')); //false
console.log(isValidPassphrase('aa bb cc dd aaa')); //true
*/

fs.readFile('input.txt', { encoding: 'utf8' }, processInput);
