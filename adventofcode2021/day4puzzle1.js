const fs = require('fs');
const readline = require('readline');

const rl = readline.createInterface({
    input: fs.createReadStream('day4input.txt')
});

let callNumbers = [];
let boards = [];
let marked = []

let currentBoard = []
let zeros = []


let lineCount = 0;

rl.on('line', line => {
    if (callNumbers.length === 0) {
        callNumbers = line.split(",").map(Number);
    } else {
        if (line.length === 0) {
            if (currentBoard.length === 5) {
                boards.push(currentBoard);
                marked.push(zeros)
                currentBoard = [];
                zeros = [];
                lineCount = 0;
            }
        } else {
            currentBoard[lineCount] = line.split(" ").filter((elem) => elem !== "").map(Number);
            zeros[lineCount] = new Array(currentBoard[lineCount].length).fill(0)
            lineCount++;
        }
    }
})

rl.on('close', () => {
    boards.push(currentBoard);
    marked.push(zeros);

    let breakOutterLoop = false;
    let finalBoard;
    let finalNumber; 
    
    for (let i = 0; i < callNumbers.length; i++) {
        let callNumber = callNumbers[i];

        for (let j = 0; j < boards.length; j++) {
            let board = boards[j];
            let idxRow = board.findIndex((x) => x.includes(callNumber))
            if (idxRow !== -1) {
                let idxColumn = board[idxRow].findIndex((x) => x === callNumber);
                marked[j][idxRow][idxColumn] = 1;
                let sumRow = 0;
                let sumColumn = 0;
                for (let k = 0; k < 5; k++) {
                    sumRow += marked[j][idxRow][k] 
                    sumColumn += marked[j][k][idxColumn];
                }
                if (sumRow === 5 || sumColumn === 5) {
                    finalBoard = j;
                    finalNumber = callNumber;
                    breakOutterLoop = true
                    break;
                }
            }
        }

        if (breakOutterLoop) {
            break;
        }  
    } 

    let sumUnmarked = 0;
    let finalMarked = marked[finalBoard];
    for (let i = 0; i < finalMarked.length; i++) {
        for (let j = 0; j < finalMarked[i].length; j++) {
            if (finalMarked[i][j] === 0) {
                sumUnmarked += boards[finalBoard][i][j]
            }
        }
    }

    console.log(finalNumber);
    console.log(sumUnmarked);
    console.log(finalNumber * sumUnmarked)
})
