const fs = require('fs');

let octofield  =  fs.readFileSync('day11input.txt', 'utf-8').split('\n').map(row => [...row].map(elem => parseInt(elem)));

let steps = 0;

let allFlashed = false;

while (allFlashed === false) {
    let flashMap = octofield.map(row => row.map(() => false));

    octofield = octofield.map(row => row.map(octo => octo+1));
    let fullEnergyIdxs = [];
    for (let row = 0; row < octofield.length; row++) {
        for (let octo = 0; octo < octofield[row].length; octo++) {
            if (octofield[row][octo] > 9) {
                fullEnergyIdxs.push([row,octo]);
            }
        }
    }

    fullEnergyIdxs.forEach(point => {
        let queue = [];
        queue.push([point[0],point[1]]);
        while(queue.length > 0) {
            let p = queue.shift();

            let isFirstColumn = p[1] === 0;
            let isFirstRow = p[0] === 0;
            let isLastColumn = p[1] === octofield[p[0]].length - 1;
            let isLastRow = p[0] === octofield.length - 1;

            let up = !isFirstRow ? [p[0]-1,p[1]] : [];
            let upLeft = !isFirstRow && !isFirstColumn ? [p[0]-1, p[1]-1] : [];
            let upRight = !isFirstRow && !isLastColumn ? [p[0]-1, p[1]+1] : [];
            let down = !isLastRow ? [p[0]+1,p[1]] : [];
            let downLeft = !isLastRow && !isFirstColumn ? [p[0]+1,p[1]-1] : [];
            let downRight = !isLastRow && !isLastColumn ? [p[0]+1,p[1]+1] : [];
            let left = !isFirstColumn ? [p[0],p[1]-1] : [];
            let right = !isLastColumn ? [p[0],p[1]+1] : [];

            if (octofield[p[0]][p[1]] > 9 && !flashMap[p[0]][p[1]]) {
                flashMap[p[0]][p[1]] = true;
                octofield[p[0]][p[1]] = 0;

                if (up.length > 0 && !flashMap[up[0]][up[1]]) {
                    octofield[up[0]][up[1]] += 1;
                    queue.push(up);
                }
                if (upLeft.length > 0 && !flashMap[upLeft[0]][upLeft[1]]) {
                    octofield[upLeft[0]][upLeft[1]] += 1;
                    queue.push(upLeft);
                }
                if (upRight.length > 0 && !flashMap[upRight[0]][upRight[1]]) {
                    octofield[upRight[0]][upRight[1]] += 1;
                    queue.push(upRight);
                }
                if (down.length > 0 && !flashMap[down[0]][down[1]]) {
                    octofield[down[0]][down[1]] += 1;
                    queue.push(down);
                }
                if (downLeft.length > 0 && !flashMap[downLeft[0]][downLeft[1]]) {
                    octofield[downLeft[0]][downLeft[1]] += 1;
                    queue.push(downLeft);
                }
                if (downRight.length > 0 && !flashMap[downRight[0]][downRight[1]]) {
                    octofield[downRight[0]][downRight[1]] += 1;
                    queue.push(downRight);
                }
                if (left.length > 0 && !flashMap[left[0]][left[1]]) {
                    octofield[left[0]][left[1]] += 1;
                    queue.push(left);
                }
                if (right.length > 0 && !flashMap[right[0]][right[1]]) {
                    octofield[right[0]][right[1]] += 1;
                    queue.push(right);
                }
            }
        }
    });
    if (flashMap.filter(row => row.filter(elem => elem === false).length > 0).length === 0) {
        allFlashed = true;
    } 
    steps += 1;
}

console.log(steps)
