const fs = require('fs');
const Tree = require('./Tree.js');

let caveConnections = fs.readFileSync('day12input.txt', 'utf-8').split('\n');

let smallCaves = new Set();

let uniqueCaves = new Set();

caveConnections.forEach(caveConnection => {
    let [first, second] = caveConnection.split('-');
    
    uniqueCaves.add(first);
    if (first !== "start" && first !== "end" && first.toLowerCase() === first) {
        smallCaves.add(first);
    }
    
    uniqueCaves.add(second);
    if (second !== "start" && second !== "end" && second.toLowerCase() === second) {
        smallCaves.add(second);
    }
})

let caveNeighbours = {}

uniqueCaves.forEach(cave => {
    let connections = caveConnections.filter((caveconnection => {
        let [first,second] = caveconnection.split('-');
        return first === cave || second === cave;
    }));

    if (!caveNeighbours[cave]) {
        caveNeighbours[cave] = new Set();
    }
    connections.forEach(conn => {
        let [first,second] = conn.split('-');
        let other = first === cave ? second : first;
        caveNeighbours[cave].add(other);
    })
})

console.log(caveNeighbours);
console.log(smallCaves);

let key = 1;
let tree = new Tree(key,'start');

let queue = [["start",key]];
while (queue.length > 0) {
    let [value,k] = queue.shift();

    // am Ende sollen keine Nachbarn angehängt werden.
    if (value !== "end") {
        let neighbours = caveNeighbours[value];

        neighbours.forEach(neighbour => {
            // Start soll nicht mehr angehängt werden
            if (neighbour !== "start") {
                key = parseInt(`${k}` + ((tree.find(k).children.length + 1)));
                if (smallCaves.has(neighbour)) {
                    let wasVisited = false;
                    let parent = tree.find(k);
                    while (parent) {
                        if (parent.value === neighbour) {
                            wasVisited = true;
                            break;
                        } else {
                            parent = parent.parent;
                        }
                    }
                    if (!wasVisited) {
                        tree.insert(k,key,neighbour)
                        queue.unshift([neighbour,key]);
                    }
                } else {
                    tree.insert(k,key,neighbour)
                    if (neighbour !== "end") {
                        queue.unshift([neighbour,key])
                    }
                }
            }
        })
    }
}

// console.log([...tree.preOrderTraversal()].map(x => x.value));

console.log([...tree.preOrderTraversal()].map(x => x.value).filter(elem => elem === "end").length)
// console.log([...tree.preOrderTraversal()])

