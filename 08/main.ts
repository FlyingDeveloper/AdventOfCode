import fs = require('fs');

var input = fs.readFileSync("input.txt", "utf8");
var inputLines = input.split('\n');

const onCharacter = '#';
const offCharacter = '·';

class Pixel {
    isOn: boolean;

    constructor() {
        this.isOn = false;
    }

    turnOn() {
        this.isOn = true;
    }

    turnOff() {
        this.isOn = false;
    }

    toggle() {
        this.isOn = !this.isOn;
    }
}

class Row {
    Pixels: Array<Pixel>;
}

class Screen {
    rows: Array<Row>;
    constructor(width: number, height: number) {
        this.rows = new Array<Row>();
        this.rows.length = height;
        for (var i = 0; i < height; i++) {
            this.rows[i] = new Row();
            this.rows[i].Pixels = new Array<Pixel>();
            this.rows[i].Pixels.length = width;
            for (var j = 0; j < width; j++) {
                this.rows[i].Pixels[j] = new Pixel();
            }
        }
    }

    displayScreen() {
        for (var i = 0; i < this.rows.length; i++) {
            var output = '';
            for (var j = 0; j < this.rows[i].Pixels.length; j++) {
                output += this.rows[i].Pixels[j].isOn ? onCharacter : offCharacter;
            }
            console.log(output);
        }
    }

    makeRectangle(width: number, height: number) {
        for (var i = 0; i < height; i++) {
            for (var j = 0; j < width; j++) {
                this.rows[i].Pixels[j].turnOn();
            }
        }
    }

    rotateRow(row: number, distance: number) {
        var rowToRotate = this.rows[row];
        for (var i = 0; i < distance; i++) {
            rowToRotate.Pixels.unshift(rowToRotate.Pixels.pop());
        }
    }

    rotateCol(col: number, distance: number) {
        var column = new Array<Pixel>();
        for (var i = 0; i < this.rows.length; i++) {
            column.push(this.rows[i].Pixels[col]);
        }

        for (var i = 0; i < distance; i++) {
            column.unshift(column.pop());
        }

        var output = ''
        column.forEach(c => output += c.isOn ? 1 : 0);
        for (var i = 0; i < this.rows.length; i++) {
            if (output[i] === '1') {
                this.rows[i].Pixels[col].turnOn();
            } else {
                this.rows[i].Pixels[col].turnOff();
            }
        }
    }

    get numberOfPixelsTurnedOn() :number {
      return this.rows.reduce((a,r)=> a + r.Pixels.reduce((b, p) => b + (p.isOn ? 1 : 0), 0), 0)
    }
}

var s = new Screen(50, 6);

var rectCommand = /^rect (\d+)x(\d+)$/;
var rotateRowCommand = /^rotate row y=(\d+) by (\d+)$/;
var rotateColumnCommand = /^rotate column x=(\d+) by (\d+)$/;

inputLines.forEach(line => {
  var regexResult:RegExpExecArray;
  if (regexResult = rectCommand.exec(line)) {
    var width = parseInt(regexResult[1]);
    var height = parseInt(regexResult[2]);
    s.makeRectangle(width, height);
  } else if (regexResult = rotateRowCommand.exec(line)) {
    var row = parseInt(regexResult[1]);
    var distance = parseInt(regexResult[2]);
    s.rotateRow(row, distance);
  } else if (regexResult = rotateColumnCommand.exec(line)) {
    var col = parseInt(regexResult[1]);
    var distance = parseInt(regexResult[2]);
    s.rotateCol(col, distance);
  }
});

s.displayScreen();
console.log(s.numberOfPixelsTurnedOn);