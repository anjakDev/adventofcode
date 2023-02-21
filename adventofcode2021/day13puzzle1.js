const fs = require('fs');

let input = fs.readFileSync('day13input.txt', 'utf-8').split('\n');

let idxOfEmpty = input.findIndex((elem) => elem === "");

let instructions = [];

for (let i = idxOfEmpty + 1; i < input.length; i++) {
    instructions.push(input[i]);
}

input.splice(idxOfEmpty);

let maxX = 0;
let maxY = 0;
input.forEach(coordinate => {
    let [x,y] = coordinate.split(',').map(elem => parseInt(elem));

    maxX = x > maxX ? x : maxX;
    maxY = y > maxY ? y : maxY;
})

let dotsOnPaper = Array.from(Array(maxY+1), () => new Array(maxX + 1).fill('.'));

input.forEach(coordinate => {
    let [column,row] = coordinate.split(',').map(elem => parseInt(elem));

    dotsOnPaper[row][column] = '#';
})

let foldLocation = instructions[0].split(' ')[2];
foldLocation = foldLocation.split('=');

let axis = foldLocation[0];
let idxOfFold = parseInt(foldLocation[1]);

if (axis === 'y') {
    for (let row = idxOfFold + 1; row < dotsOnPaper.length; row++) {
        for (let column = 0; column < dotsOnPaper[row].length; column++) {
            if (dotsOnPaper[row][column] === '#') {
                let newRowIdx = idxOfFold - (row - idxOfFold);
                dotsOnPaper[newRowIdx][column] = '#';
            }
        } 
    }
    dotsOnPaper.splice(idxOfFold);
} else {
    for (let row = 0; row < dotsOnPaper.length; row++) {
        for (let column = idxOfFold + 1; column < dotsOnPaper[row].length; column++) {
            if (dotsOnPaper[row][column] === '#') {
                let newColumnIdx = idxOfFold - (column - idxOfFold);
                dotsOnPaper[row][newColumnIdx] = '#';
            }
        }
    }
    dotsOnPaper.forEach(row => row.splice(idxOfFold));
}
// dotsOnPaper.forEach(row => console.log(row.join("")))

console.log([].concat(...dotsOnPaper).filter(elem => elem === '#').length)
