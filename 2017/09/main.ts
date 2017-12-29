import fs = require('fs');

class Group {
  public children:Array<Group>;
  public parent: Group|undefined;
  public myScore:number;

  constructor(parent?:Group) {
    this.children = [];
    this.parent = parent;
    this.myScore = this.parent ? this.parent.myScore + 1 : 1;
  }

  calculateScore():number {
    let childrenScore = this.children.reduce((accumulator, current) => {
      return accumulator + current.calculateScore();
    }, 0);

    return this.myScore + childrenScore;
  }
}

function processInput(input:string):number {
  let len = input.length;
  let i = 0;
  let currentGroup:Group|undefined;
  let currentlyInGarbage = false;

  while (i < len) {
    const currentCharacter = input[i];
    if (currentlyInGarbage && currentCharacter !== '>' && currentCharacter !== '!') {
      ++i;
      continue;
    }

    switch (currentCharacter) {
      case '!':
        ++i;
        break;
      case '<':
        currentlyInGarbage = true;
        break;
      case '>':
        currentlyInGarbage = false;
        break;
      case '{':
        let newGroup = new Group(currentGroup);
        if (currentGroup) {
          currentGroup.children.push(newGroup);
        }
        currentGroup = newGroup;
        break;
      case '}':
        if (!currentGroup) {
          throw 'currentGroup is undefined but should not be at this point in the code';
        }

        if (currentGroup.parent === undefined) {
          // stop iterating
          i = Number.MAX_VALUE;
        } else {
          currentGroup = currentGroup.parent;
        }
        break;
    }

    ++i;
  }

  if (!currentGroup) {
    throw 'currentGroup is undefined when trying to calculate the final score';
  }

  return currentGroup.calculateScore();
}

/* test cases */
/*function runTest(input:string, expectedOutput: number) {
  var output = processInput(input);
  if (output === expectedOutput) {
    console.log(true);
  } else {
    console.log(`Failed: ${input} produced ${output} when ${expectedOutput} was expected`);
  }
}

runTest('{}', 1);
runTest('{{{}}}', 6);
runTest('{{},{}}', 5);
runTest('{{{},{},{{}}}}', 16);
runTest('{<a>,<a>,<a>,<a>}', 1);
runTest('{{<ab>},{<ab>},{<ab>},{<ab>}}', 9);
runTest('{{<!!>},{<!!>},{<!!>},{<!!>}}', 9);
runTest('{{<a!>},{<a!>},{<a!>},{<ab>}}', 3);*/

fs.readFile('input.txt', {encoding:'utf8'}, (err, data) => console.log(processInput(data)));