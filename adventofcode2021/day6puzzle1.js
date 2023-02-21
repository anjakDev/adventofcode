const fs = require('fs');

let initialState = fs.readFileSync("day6input.txt", "utf-8").split(",");

let state = initialState.map(fish => parseInt(fish));

let newFish = 0;

let days = 0;

while (days <= 80) {
    state.push(...Array(newFish).fill(8))
    newFish = 0;
    state = state.map(fish => {
        if (fish === 0) {
            newFish++;
            return 6
        } else {
            return fish - 1
        }
    })
    days++;
}

console.log(state.length)