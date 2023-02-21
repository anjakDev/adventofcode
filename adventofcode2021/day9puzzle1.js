const fs = require('fs');

let heightmap = fs.readFileSync('day9input.txt', 'utf-8').split('\n');

heightmap = heightmap.map(line => [...line].map(element => parseInt(element)));

let riskLevelSum = 0;

const increaseRiskLevelSum = (height) => {
    riskLevelSum += 1 + height
}

for (let row = 0; row < heightmap.length; row++) {
    for (let column = 0; column < heightmap[row].length; column++) {
        let current = heightmap[row][column];

        // 9 is highest digit therefore cannot be lowest point
        if (current === 9) {
            continue;
        }
        if (current === 0) {
            increaseRiskLevelSum(current);
            continue;
        }

        let isFirstColumn = column === 0;
        let isFirstRow = row === 0;
        let isLastColumn = column === heightmap[row].length - 1;
        let isLastRow = row === heightmap.length - 1;

        let up = !isFirstRow ? heightmap[row-1][column] : 0;
        let down = !isLastRow ? heightmap[row+1][column] : 0;
        let left = !isFirstColumn ? heightmap[row][column-1] : 0;
        let right = !isLastColumn ? heightmap[row][column+1] : 0;

        if (isFirstRow) { 
            if (isFirstColumn) {
                if (current < down && current < right) {
                    increaseRiskLevelSum(current);
                }
            } else if (isLastColumn) {
                if (current < down && current < left) {
                    increaseRiskLevelSum(current);
                }
            } else {
                if (current < down && current < left && current < right) {
                    increaseRiskLevelSum(current);
                } 
            }
        } else if (isLastRow) {
            if (isFirstColumn) {
                if (current < up && current < right) {
                    increaseRiskLevelSum(current);
                }
            } else if (isLastColumn) {
                if (current < up && current < left) {
                    increaseRiskLevelSum(current);
                }
            } else {
                if (current < up && current < left && current < right) {
                    increaseRiskLevelSum(current);
                }
            }
        } else {
            if (isFirstColumn) {
                if (current < up && current < down && current < right) {
                    increaseRiskLevelSum(current);
                }
            } else if (isLastColumn) {
                if (current < up && current < down && current < left) {
                    increaseRiskLevelSum(current);
                }
            } else {
                if (current < up && current < down && current < left && current < right) {
                    increaseRiskLevelSum(current);
                }
            }
        }
    }
}

console.log(riskLevelSum)
