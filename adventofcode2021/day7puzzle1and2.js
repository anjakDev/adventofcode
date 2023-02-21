const fs = require('fs');

let crabs = fs.readFileSync("day7input.txt", "utf-8").split(",").map(crab => parseInt(crab));

let min = 0;

for (let i = 0; i <= Math.max(...crabs); i++) {
    let chosenPosition = i;
    let sumFuel = 0;

    for (let j = 0; j < crabs.length; j++) {
        let distance = Math.abs(chosenPosition - crabs[j]) + 1;
        let sumDistance = ((distance) * (distance - 1 )) / 2
        sumFuel += sumDistance;
    }
    if (i === 0) {
        min = sumFuel;
    } else {
        min = min <= sumFuel ? min : sumFuel;
    }
}
 console.log(min)

 