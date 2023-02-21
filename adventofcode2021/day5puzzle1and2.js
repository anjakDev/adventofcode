const fs = require('fs');

let lines = fs.readFileSync("day5input.txt", "utf-8").split("\n");
lines = lines.map(line => line.split(" -> "));

let x1Array = []
let x2Array = []
let y1Array = []
let y2Array = []

lines.forEach(entry => {
    let first = entry[0].split(",");
    let second = entry[1].split(",");

    x1Array.push(parseInt(first[0]));
    x2Array.push(parseInt(second[0]));
    y1Array.push(parseInt(first[1]));
    y2Array.push(parseInt(second[1]));
})

const maxX1 = Math.max(...x1Array);
const maxX2 = Math.max(...x2Array);
const maxY1 = Math.max(...y1Array);
const maxY2 = Math.max(...y2Array);

let diagram = Array.from(Array(Math.max(maxY1, maxY2) + 1), () => new Array(Math.max(maxX1, maxX2) + 1).fill(0));

lines.forEach(entry => {
    let [x1, y1] = entry[0].split(",").map(elem => parseInt(elem))
    let [x2, y2] = entry[1].split(",").map(elem => parseInt(elem))

    // HORIZONTAL LINES
    if (x1 === x2) {
        for (let i = Math.min(y1,y2); i <= Math.max(y1,y2); i++) {
            diagram[i][x1] += 1;
        }
    } 
    // VERTICAL LINES
    else if (y1 === y2) {
        for (let j = Math.min(x1,x2); j <= Math.max(x1,x2); j++) {
            diagram[y1][j] += 1;
        }
    } 
    // DIAGONAL LINES 45 DEGREE
    else if (Math.abs(x1 - x2) === Math.abs(y1 - y2)) {

        for (let m = 0; m <= Math.abs(x1 - x2); m++) {
            if (x1 < x2 && y1 < y2) { // case 4
                diagram[y1 + m][x1 + m] += 1;
            } else if (x1 < x2 && y1 > y2) { // case 1
                diagram[y1 - m][x1 + m] += 1;
            } else if (x1 > x2 && y1 < y2) { // case 2
                diagram[y1 + m][x1 - m] += 1;
            } else if (x1 > x2 && y1 > y2) { // case 3
                diagram[y1 - m][x1 - m] += 1;
            }
        }
    }
})

let flattenedDiagram = [].concat(...diagram);
let filtered = flattenedDiagram.filter(elem => elem >= 2);
console.log(filtered.length)