import Utils from 'utils';

Utils.GetLines().then(processLines);
/*let sampleInput = [
  'abcde',
  'fghij',
  'klmno',
  'pqrst',
  'fguij',
  'axcye',
  'wvxyz'
];
processLines(sampleInput);
*/

function processLines(lines: string[]) {
  lines.forEach((currentLine, outerIteratorIndex) => {
    for (let j = outerIteratorIndex; j < lines.length; j++) {
      let otherLine = lines[j];
      if (currentLine.length === otherLine.length) {
        let numDiffs = 0;
        let lastDiffIndex = -1;
        for (let i = 0; i < currentLine.length; i++) {
          if (currentLine.split('')[i] !== otherLine.split('')[i]) {
            ++numDiffs;
            lastDiffIndex = i;
          }
        }

        if (numDiffs === 1) {
          let currentLineExploded = currentLine.split('');
          currentLineExploded.splice(lastDiffIndex, 1);
          console.log(currentLineExploded.join(''));
        }
      }
    }
  });
}
