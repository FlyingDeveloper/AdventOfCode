import Utils from 'utils';

Utils.GetLines().then(processLines);

function processLines(lines:string[]) {
  let NumberOfIDsWithCharacterAppearingTwice = 0;
  let NumberOfIDsWithCharactersAppearingThreeTimes = 0;

  lines.forEach(line => {
    var characterAppearances = getCharacterAppearancesInString(line);
    var alreadyAppearsTwice = false;
    var alreadyAppearsThrice = false;
    Object.keys(characterAppearances).forEach(character => {
     if (characterAppearances[character] === 2 && !alreadyAppearsTwice) {
        ++NumberOfIDsWithCharacterAppearingTwice;
        alreadyAppearsTwice = true;
      } else if (characterAppearances[character] === 3 && !alreadyAppearsThrice) {
        ++NumberOfIDsWithCharactersAppearingThreeTimes;
        alreadyAppearsThrice = true;
      }
    });
  });

  console.log(`${NumberOfIDsWithCharacterAppearingTwice} * ${NumberOfIDsWithCharactersAppearingThreeTimes} = ${NumberOfIDsWithCharacterAppearingTwice * NumberOfIDsWithCharactersAppearingThreeTimes}`);
}

function getCharacterAppearancesInString(input: string): {[character: string]: number} {
  let characterAppearances:{[character:string]: number} = {};
  input.split('').forEach(character => {
    if (characterAppearances[character] === undefined) {
      characterAppearances[character] = 1;
    } else {
      characterAppearances[character]++;
    }
  });

  return characterAppearances;
}
