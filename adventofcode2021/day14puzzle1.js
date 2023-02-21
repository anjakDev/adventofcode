const fs = require('fs');

let input = fs.readFileSync('day14input.txt', 'utf-8').split('\n');

let template = input.shift();

input.shift();

let instructions = {};

input.forEach(instruction => {
    let [pair,insert] = instruction.split(" -> ");

    instructions[pair] = insert;
})

let steps = 0;

while (steps < 10) {
    let modifiedTemplate = "";
    for (let i = 0; i < template.length - 1; i++) {
        let firstPair = template[i]+template[i+1];
        let insert = instructions[firstPair];
        if (i === template.length - 2) {
            modifiedTemplate += template[i] + insert + template[i+1];
        } else {
            modifiedTemplate += template[i] + insert;
        }
    }
    template = modifiedTemplate;
    steps++;
}

let counts = {};

let uniqueCharacters = new Set();

[...template].forEach(character => uniqueCharacters.add(character));

for (let i = 0; i < template.length; i++) {
    let count = counts[template[i]]

    counts[template[i]] = count ? count + 1 : 1;
}

let max = Math.max(...Object.values(counts));
let min = Math.min(...Object.values(counts));

console.log(max - min);
