const fs = require('fs');
const readline = require('readline');

const rl = readline.createInterface({
    input: fs.createReadStream('day8input.txt')
});

let outputSum = 0;

const checkPatternIncludes = (pattern, array) => {
    let allIncluded = 0;

    for (let i = 0; i < array.length; i++) {
        allIncluded += pattern.includes(array[i]);
    }
    return allIncluded;
}

rl.on('line', line => {
    let [patterns, output] = line.split('|');
    patterns = patterns.trimEnd();
    output = output.trimStart();

    let patternArray = patterns.split(" ");
    let outputArray = output.split(" ");

    let digits = {
        "zero": [],
        "one": [],
        "two": [],
        "three": [],
        "four": [],
        "five": [],
        "six": [],
        "seven": [],
        "eight": [],
        "nine": []
    }

    // 1, 4, 7, 8
    patternArray.forEach(pattern => {
        switch (pattern.length) {
            case 2: 
                digits.one = [...pattern];
                break;
            case 4: 
                digits.four = [...pattern];
                break;
            case 3:
                digits.seven = [...pattern];
                break;
            case 7:
                digits.eight = [...pattern];
                break;
        }
    })
    // length 5: 2, 3, 5
    let lengthFiveDigits = patternArray.filter(elem => elem.length === 5);
    // length 6: 0, 6, 9
    let lengthSixDigits = patternArray.filter(elem => elem.length === 6);

    let threeIdx = lengthFiveDigits.findIndex(pattern => checkPatternIncludes(pattern, digits.one) === digits.one.length);
    digits.three = [...lengthFiveDigits.splice(threeIdx,1)[0]];
    let nineIdx = lengthSixDigits.findIndex(pattern => checkPatternIncludes(pattern, digits.three) === digits.three.length);
    digits.nine = [...lengthSixDigits.splice(nineIdx,1)[0]];

    let zeroIdx = lengthSixDigits.findIndex(pattern => checkPatternIncludes(pattern, digits.seven) === digits.seven.length);
    digits.zero = [...lengthSixDigits.splice(zeroIdx, 1)[0]];
    digits.six = [...lengthSixDigits[0]];

    let fiveIdx = lengthFiveDigits.findIndex(pattern => checkPatternIncludes(pattern, digits.six) === digits.six.length - 1);
    digits.five = [...lengthFiveDigits.splice(fiveIdx, 1)[0]];
    digits.two = [...lengthFiveDigits[0]];

    let outputValue = "";
    outputArray.forEach(pattern => {
        switch (pattern.length) {
            case 2:
                outputValue += "1";
                break;
            case 3:
                outputValue += "7";
                break;
            case 4:
                outputValue += "4";
                break;
            case 5:
                if (checkPatternIncludes(pattern, digits.two) === digits.two.length) {
                    outputValue += "2";
                } else if (checkPatternIncludes(pattern, digits.three) === digits.three.length) {
                    outputValue += "3";
                } else if (checkPatternIncludes(pattern, digits.five) === digits.five.length) {
                    outputValue += "5";
                }
                break;
            case 6:
                if (checkPatternIncludes(pattern, digits.zero) === digits.zero.length) {
                    outputValue += "0";
                } else if (checkPatternIncludes(pattern, digits.six) === digits.six.length) {
                    outputValue += "6";
                } else if (checkPatternIncludes(pattern, digits.nine) === digits.nine.length) {
                    outputValue += "9";
                }
                break;
            case 7:
                outputValue += "8";
                break;
        }
    })

    outputSum += parseInt(outputValue)
})

rl.on('close', () => {
    console.log(outputSum);
})