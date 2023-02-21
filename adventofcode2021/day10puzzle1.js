const fs = require('fs');
const readline = require('readline');

const rl = readline.createInterface({
    input: fs.createReadStream('day10input.txt', 'utf-8')
})

let sumSyntaxScore = 0;

let openingBrackets = ["(", "[", "{", "<"];

let bracketPairs = {
    ")": "(",
    "]": "[",
    "}": "{",
    ">": "<"
}

const getBracketSyntaxScore = (bracket) => {
    switch (bracket) {
        case ")":
            return 3;
        case "]": 
            return 57;
        case "}":
            return 1197;
        case ">":
            return 25137;
    }
}

rl.on('line', (line) => {
    let bracketqueue = [];
    for (let i = 0; i < line.length; i++) {
        let bracket = line[i];
        if (bracketqueue.length === 0 || openingBrackets.includes(bracket)) {
            bracketqueue.push(bracket);
        } else {
            let latestOpenBracket = bracketqueue[bracketqueue.length - 1];
            if (bracketPairs[bracket] !== latestOpenBracket) {
                sumSyntaxScore += getBracketSyntaxScore(bracket);
                break;
            } else {
                bracketqueue.splice(bracketqueue.length-1,1);
            }
        }
    }
})

rl.on('close', () => {
    console.log(sumSyntaxScore);
})