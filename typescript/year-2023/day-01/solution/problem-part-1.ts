const validDigits = '0123456789';

function calculateCalibrationValue(calibrationLine: string): number {
  let firstDigitIndex = -1;
  let lastDigitIndex = -1;

  for (let i = 0; i < calibrationLine.length; i += 1) {
    const character = calibrationLine[i];
    if (validDigits.includes(character)) {
      firstDigitIndex = i;
      break;
    }
  }

  if (firstDigitIndex === -1) {
    throw new Error('no calibration found');
  }

  for (let i = calibrationLine.length - 1; i >= 0; i -= 1) {
    const character = calibrationLine[i];
    if (validDigits.includes(character)) {
      lastDigitIndex = i;
      break;
    }
  }

  return Number(
    calibrationLine[firstDigitIndex] + calibrationLine[lastDigitIndex],
  );
}

export function solveProblem(input: string): number {
  const inputLines = input.split('\n');
  inputLines.pop(); // omit last trailing empty line

  const sum = inputLines.reduce((accumulator, current) => {
    const calibrationValue = calculateCalibrationValue(current);
    return accumulator + calibrationValue;
  }, 0);

  return sum;
}

