const fs = require('fs');

let heightmap = fs.readFileSync('day9input.txt', 'utf-8').split('\n');

heightmap = heightmap.map(line => [...line].map(element => parseInt(element)));

let basinSizes = [];

let markedIdxs = heightmap.map(row => row.map(() => false));

let lowPoints = []

for (let row = 0; row < heightmap.length; row++) {
    for (let column = 0; column < heightmap[row].length; column++) {
        let current = heightmap[row][column];
        // 9 is highest digit therefore cannot be lowest point
        if (current === 9) {
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
                    lowPoints.push([row,column])
                }
            } else if (isLastColumn) {
                if (current < down && current < left) {
                    lowPoints.push([row,column])
                }
            } else {
                if (current < down && current < left && current < right) {
                    lowPoints.push([row,column])
                } 
            }
        } else if (isLastRow) {
            if (isFirstColumn) {
                if (current < up && current < right) {
                    lowPoints.push([row,column])
                }
            } else if (isLastColumn) {
                if (current < up && current < left) {
                    lowPoints.push([row,column])
                }
            } else {
                if (current < up && current < left && current < right) {
                    lowPoints.push([row,column])
                }
            }
        } else {
            if (isFirstColumn) {
                if (current < up && current < down && current < right) {
                    lowPoints.push([row,column])
                }
            } else if (isLastColumn) {
                if (current < up && current < down && current < left) {
                    lowPoints.push([row,column])
                }
            } else {
                if (current < up && current < down && current < left && current < right) {
                    lowPoints.push([row,column])
                }
            }
        }
    }
}

// console.log(lowPoints)

lowPoints.forEach(point => {
    let basinSize = 0;
    let queue = [];
    queue.push(point);
    while (queue.length > 0) {
        let p = queue.shift();
        if (heightmap[p[0]][p[1]] === 9) {
            continue;
        }

        if (!markedIdxs[p[0]][p[1]]) {
            markedIdxs[p[0]][p[1]] = true;
            basinSize += 1;
        }
                
        let isFirstColumn = p[1] === 0;
        let isFirstRow = p[0] === 0;
        let isLastColumn = p[1] === heightmap[p[0]].length - 1;
        let isLastRow = p[0] === heightmap.length - 1;

        let up = !isFirstRow ? [p[0]-1,p[1]] : [];
        let down = !isLastRow ? [p[0]+1,p[1]] : [];
        let left = !isFirstColumn ? [p[0],p[1]-1] : [];
        let right = !isLastColumn ? [p[0],p[1]+1] : [];

        if (up.length > 0 && !markedIdxs[up[0]][up[1]]) {
            queue.push(up);
        }
        if (down.length > 0 && !markedIdxs[down[0]][down[1]]) {
            queue.push(down);
        }
        if (left.length > 0 && !markedIdxs[left[0]][left[1]]) {
            queue.push(left);
        }
        if (right.length > 0 && !markedIdxs[right[0]][right[1]]) {
            queue.push(right);
        }
    }
    basinSizes.push(basinSize);
})

console.log(basinSizes)
const sortedBasinSizes = basinSizes.sort((a,b) => b-a);
console.log(sortedBasinSizes[0] * sortedBasinSizes[1] * sortedBasinSizes[2])
