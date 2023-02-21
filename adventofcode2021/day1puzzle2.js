const fs = require('fs');

const depths = fs.readFileSync("day1puzzle1.txt", "utf-8").split("\n");
const depthsArray = depths.map(Number);
const depthsLength = depthsArray.length;

let increased = 0;

let i = 0;
let j = 1;

while (j+2 < depthsLength) {
    let firstSum = depthsArray[i] + depthsArray[i+1] + depthsArray[i+2];
    let secondSum = depthsArray[j] + depthsArray[j+1] + depthsArray[j+2];
    if (secondSum > firstSum) {
        increased++;
    }
    i++;
    j++;
}

console.log(increased)