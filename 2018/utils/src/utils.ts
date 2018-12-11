import {default as fs} from 'fs';

export default class Utils {
  static GetFileContents(filename:string = "input.txt") {
    return new Promise<string>((res, rej) => {
      fs.readFile(filename, {encoding: 'utf8'}, (err, data) => {
        if (err) {
          rej(err);
        } else {
          res(data);
        }
      });
    });
  }

  static GetLines(filename: string = "input.txt") {
    return Utils.GetFileContents(filename).then(data => data.split('\n').filter(x => x !== ''));
  }

  static GetLinesAsNumber(filename: string = "input.txt") {
    return Utils.GetLines(filename).then(data => data.map(Number));
  }
}
