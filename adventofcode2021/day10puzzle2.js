const fs = require('fs');
const readline = require('readline');

const rl = readline.createInterface({
    input: fs.createReadStream('day10input.txt', 'utf-8')
})

let allScores = [];

let openingBrackets = ["(", "[", "{", "<"];

let bracketPairs = {
    ")": "(",
    "]": "[",
    "}": "{",
    ">": "<"
}

let bracketPairsSwapped = {
    "(": ")",
    "[": "]",
    "{": "}",
    "<": ">"
}

const getBracketScore = (bracket) => {
    switch (bracket) {
        case ")":
            return 1;
        case "]": 
            return 2;
        case "}":
            return 3;
        case ">":
            return 4;
    }
}

rl.on('line', (line) => {
    let bracketqueue = [];
    let isCorrupted = false;
    for (let i = 0; i < line.length; i++) {
        let bracket = line[i];
        if (bracketqueue.length === 0 || openingBrackets.includes(bracket)) {
            bracketqueue.push(bracket);
        } else {
            let latestOpenBracket = bracketqueue[bracketqueue.length - 1];
            if (bracketPairs[bracket] !== latestOpenBracket) {
                isCorrupted = true;
                break;
            } else {
                bracketqueue.splice(bracketqueue.length-1,1);
            }
        }
    }
    if (!isCorrupted) {
        let completionString = ""
        let lineScore = 0
        for (let j = bracketqueue.length - 1; j >= 0; j--) {
            let lastOpenBracket = bracketqueue[j];
            let closingBracket = bracketPairsSwapped[lastOpenBracket];
            lineScore = (lineScore * 5) + getBracketScore(closingBracket);
            completionString += closingBracket;
        }
        allScores.push(lineScore)
    }
})

rl.on('close', () => {
    allScores.sort((a,b) => a - b);
    console.log(allScores.at((allScores.length-1)/2))
})