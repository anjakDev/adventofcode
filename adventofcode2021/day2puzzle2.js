const readline = require('readline');
const fs = require('fs');

const rl = readline.createInterface({
    input: fs.createReadStream('day2input.txt')
});

let horizontal = 0;
let depth = 0;
let aim = 0;

const FORWARD = "forward";
const DOWN = "down";
const UP = "up";

rl.on('line', line => {
    const commandParts = line.split(" ");
    const command = commandParts[0];
    const number = parseInt(commandParts[1]);

    switch(command) {
        case FORWARD: 
            horizontal += number;
            depth += aim * number;
            break;
        case DOWN: 
            aim += number;
            break;
        case UP:
            aim -= number;
            break;
    }
})

rl.on('close', () => {
    console.log(horizontal * depth);
})