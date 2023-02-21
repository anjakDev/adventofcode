const fs = require('fs');

const binaries = fs.readFileSync("day3input.txt", "utf-8").split("\n");
const numOfBits = binaries[0].length

const findOxy = (array) => {
    let oxyArray = [...array];
    for (let i = 0; i < numOfBits; i++) {

        let countOnes = 0;
        let countZeros = 0;

        for (let j = 0; j < oxyArray.length; j++) {
            parseInt(oxyArray[j][i]) === 0 ? countZeros++ : countOnes++;
        }
    
        if (countOnes > countZeros || countOnes === countZeros) {
            oxyArray = oxyArray.filter((element) => element[i] === "1");
        } else {
            oxyArray = oxyArray.filter((element) => element[i] === "0");
        }
        
        if (oxyArray.length === 1) {
            break;
        }
    }
    return oxyArray[0];
}

const findCO2 = (array) => {
    let co2Array = [...array]
    for (let i = 0; i < numOfBits; i++) {

        let countOnes = 0;
        let countZeros = 0;

        for (let j = 0; j < co2Array.length; j++) {
            parseInt(co2Array[j][i]) === 0 ? countZeros++ : countOnes++;
        }
        
        if (countOnes > countZeros || countOnes === countZeros) {
            co2Array = co2Array.filter((element) => element[i] === "0")
        } else {
            co2Array = co2Array.filter((element) => element[i] === "1");
        } 
        
        if (co2Array.length === 1) {
            break;
        } 
    }
    return co2Array[0];
}

const oxy = parseInt(findOxy(binaries), 2)
const co2 = parseInt(findCO2(binaries), 2)

console.log(oxy * co2)


