const fs = require('fs');

const binaries = fs.readFileSync("day3input.txt", "utf-8").split("\n");

let gammaRate = "";
let epsilonRate = "";

const numOfBits = binaries[0].length

for (let i = 0; i < numOfBits; i++) {
    let countOnes = 0;
    let countZeros = 0;
    for (let j = 0; j < binaries.length; j++) {
        parseInt(binaries[j][i]) === 0 ? countZeros++ : countOnes++;
    }
    if (countOnes > countZeros) {
        gammaRate += "1";
        epsilonRate += "0";
    } else {
        gammaRate += "0";
        epsilonRate += "1";
    }
}

console.log(parseInt(gammaRate, 2) * parseInt(epsilonRate, 2))
