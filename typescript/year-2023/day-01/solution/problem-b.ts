const _exampleFile = Bun.file('input/example-b');
const inputFile = Bun.file('input/input-b');

const tokenMap = {
  one: '1',
  two: '2',
  three: '3',
  four: '4',
  five: '5',
  six: '6',
  seven: '7',
  eight: '8',
  nine: '9',
  '1': '1',
  '2': '2',
  '3': '3',
  '4': '4',
  '5': '5',
  '6': '6',
  '7': '7',
  '8': '8',
  '9': '9',
};

const tokenKeys = Object.keys(tokenMap);

interface TokenMapSet {
  key: keyof typeof tokenMap;
  set: string;
}

function getTokenKeySetPair() {
  return Array.from({ length: tokenKeys.length }, (_, i) =>
    Object.fromEntries([
      ['key', tokenKeys[i]],
      ['set', ''],
    ]),
  ) as unknown as TokenMapSet[];
}

function calcCalibrationFromLeft(calibrationLine: string): string {
  const tokenKeySetPair = getTokenKeySetPair();

  for (let i = 0; i < calibrationLine.length; i += 1) {
    for (let j = 0; j < tokenKeySetPair.length; j += 1) {
      const map = tokenKeySetPair[j];
      if (calibrationLine[i] === map.key[map.set.length]) {
        map.set += calibrationLine[i];
        if (map.set === map.key) {
          return tokenMap[map.key];
        }
      }
    }
  }

  throw new Error('no calibration found');
}

function calcCalibrationFromRight(calibrationLine: string): string {
  const tokenKeySetPair = getTokenKeySetPair();

  for (let i = calibrationLine.length - 1; i >= 0; i -= 1) {
    for (let j = 0; j < tokenKeySetPair.length; j += 1) {
      const map = tokenKeySetPair[j];
      if (calibrationLine[i] === map.key[map.key.length - 1 - map.set.length]) {
        map.set = calibrationLine[i] + map.set;
        if (map.set === map.key) {
          return tokenMap[map.key];
        }
      }
    }
  }

  throw new Error('no calibration found');
}

function calcCalibrationValue(calibrationLine: string): number {
  return Number(
    calcCalibrationFromLeft(calibrationLine) +
    calcCalibrationFromRight(calibrationLine),
  );
}

async function main() {
  const input = await inputFile.text();
  const inputLines = input.split('\n');

  inputLines.pop(); // omit last trailing empty line

  const sum = inputLines.reduce((accumulator, current) => {
    const calibrationValue = calcCalibrationValue(current);
    console.log(current, calibrationValue);
    return accumulator + calibrationValue;
  }, 0);

  console.log('\nsum:', sum);
}

main();
