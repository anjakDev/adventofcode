const readline = require('readline');
const fs = require('fs');

const rl = readline.createInterface({
    input: fs.createReadStream('day1puzzle1.txt')
});

let previous = 0

let count = 0


rl.on('line', line => {
    const number = parseInt(line);
    if (previous !== 0) {
        if (number > previous) {
            count++;
        }
    }
    previous = number
});

rl.on('close', () => {
    console.log(count);
});