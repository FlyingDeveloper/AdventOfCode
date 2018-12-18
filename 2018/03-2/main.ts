import Utils from 'utils';

Utils.GetLines().then(processLines);

let sampleInput = [
  '#1 @ 1,3: 4x4',
  '#2 @ 3,1: 4x4',
  '#3 @ 5,5: 2x2'
];
//processLines(sampleInput);



function processLines(lines: string[]) {
  class Square {
    private static formatRegex = /#(\d+)\s@\s(\d+),(\d+):\s(\d+)x(\d+)$/
    static Create(input: string): Square {
      var result = Square.formatRegex.exec(input);
      if (result == undefined) {
        throw 'Invalid input';
      }
  
      var id = Number(result[1]);
      var left = Number(result[2]);
      var top = Number(result[3]);
      var width = Number(result[4]);
      var height = Number(result[5]);
  
      return new Square(id, left, top, width, height);
    }
  
    private constructor(
      public id: number,
      public left: number,
      public top: number,
      public width: number,
      public height: number) {
    }
    
    public ToString() {
      return `#${this.id} @ ${this.left},${this.top}: ${this.width}x${this.height}`;
    }
  }

  function getPositionHash(left: number, top: number):string {
    return `${left},${top}`;
  }

  let Squares:Square[] = [];
  
  lines.forEach(x => {
    let s = Square.Create(x);
    Squares.push(s);
  });

  function doSquaresIntersect(s1:Square, s2:Square): boolean {
    if (s1 === s2 || s1.left > (s2.left + s2.width - 1) || s1.top > (s2.top + s2.height - 1) ||
      (s1.left + s1.width - 1) < s2.left || (s1.top + s1.height - 1) < s2.top) {
      return false;
    };

    return true;
  }

  for (var i = 0; i < Squares.length; i++) {
    let hasOverlap = false;
    for (var j = 0; j < Squares.length; j++) {
      if (doSquaresIntersect(Squares[i], Squares[j])) {
        hasOverlap = true;
        break;
      }
    }
    if (!hasOverlap) {
      console.log(Squares[i].id);
      break;
    }
  }
}
