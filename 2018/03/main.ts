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

  let Squares:Set<Square> = new Set();
  
  lines.forEach(x => {
    let s = Square.Create(x);
    Squares.add(s);
  });

  let fabric2:Map<string, number> = new Map();
  
  function claimSquare(left: number, top: number) {
    let positionHash = getPositionHash(left, top);
    let claimsOnPosition = fabric2.get(positionHash);
    if (claimsOnPosition === undefined) {
      fabric2.set(positionHash, 1);
    } else if (claimsOnPosition >= 1) {
      fabric2.set(positionHash, ++claimsOnPosition);
    }
  }
  
  let failedClaims = 0;
  Squares.forEach(square => {
    for (var l = square.left; (l < (square.left + square.width)); l++) {
      for (var t = square.top; (t < (square.top + square.height)); t++) {
        claimSquare(l, t);
      }
    }
  });

  fabric2.forEach(x => x > 1 ? failedClaims++ : void 0);
  console.log(failedClaims);
}
