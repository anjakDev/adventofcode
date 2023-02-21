const fs = require('fs');
const _ = require('lodash');

let input = fs.readFileSync('day14input.txt', 'utf-8').split('\n');

let template = input.shift();
input.shift();

let instructions = {};
let counts = {};

input.forEach(instruction => {
    let [pair,insert] = instruction.split(" -> ");
    instructions[pair] = [pair[0] + insert, insert + pair[1]];
})

let pairsAndIdxs = {};
for (let i = 0; i < Object.keys(instructions).length; i++) {
    pairsAndIdxs[Object.keys(instructions)[i]] = 0;
}

for (let j = 0; j < template.length - 1; j++) {
    let pair = template[j] + template[j+1];
    let count = counts[template[j]];
    counts[template[j]] = count ? count + 1 : 1;
    if (j === template.length - 2) {
        let countLast = counts[template[j+1]];
        counts[template[j+1]] = countLast ? countLast + 1 : 1;
    }
    pairsAndIdxs[pair] += 1;
}

let max = 0;

let steps = 0;
while (steps < 40) {
    let copyPairs = {...pairsAndIdxs};
    let filteredByPairCount = Object.keys(pairsAndIdxs).filter(key => pairsAndIdxs[key] > 0);
    filteredByPairCount.forEach(key => {
        let pairCount = pairsAndIdxs[key];
        let [newPairOne, newPairTwo] = instructions[key];
        copyPairs[key] -= pairCount;
        copyPairs[newPairOne] += pairCount;
        copyPairs[newPairTwo] += pairCount;
        let c = counts[newPairOne[1]];
        counts[newPairOne[1]] = c ? c + pairCount : pairCount;
        max = max > counts[newPairOne[1]] ? max : counts[newPairOne[1]];
    })
    pairsAndIdxs = {...copyPairs};
    steps++;
}

let min = Math.min(...Object.values(counts));

console.log(max - min);
