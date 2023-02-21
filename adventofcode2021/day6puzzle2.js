const fs = require('fs');

let initialState = fs.readFileSync("day6input.txt", "utf-8").split(",");

let state = initialState.map(fish => parseInt(fish));

// let agesMatrix = Array.from(Array(9), () => new Array(9).fill(0));
let ages = Array(9).fill(0);

state.forEach(fish => ages[fish] += 1);

let days = 1;

while (days <= 256) {
    let ages0 = ages[0];
    ages[0] = ages[1];
    ages[1] = ages[2];
    ages[2] = ages[3];
    ages[3] = ages[4];
    ages[4] = ages[5];
    ages[5] = ages[6];
    ages[6] = ages[7] + ages0;
    ages[7] = ages[8];
    ages[8] = ages0;
    days++;
}

console.log(ages.reduce((a,b) => a+b))