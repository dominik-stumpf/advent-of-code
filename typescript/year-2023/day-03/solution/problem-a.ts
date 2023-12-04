const _exampleFile = Bun.file('input/example-a');
const _inputFile = Bun.file('input/input-a');

const validDigits = '0123456789';

interface PartNumber {
  value: string;
  startIndex: number;
}

function getPartNumbers(engineSchematic: string): PartNumber[] {
  const partNumbers: PartNumber[] = [];
  let currentNumber = '';

  for (let i = 0; i < engineSchematic.length; i += 1) {
    const char = engineSchematic[i];
    if (validDigits.includes(char)) {
      currentNumber += char;
    } else if (currentNumber.length !== 0) {
      partNumbers.push({
        value: currentNumber,
        startIndex: i - currentNumber.length,
      });
      currentNumber = '';
    }
  }

  return partNumbers;
}

interface IndexedPartNumber extends PartNumber {
  adjacentIndices: number[][];
}

function getAdjacentPartNumberIndices(
  partNumbers: PartNumber[],
  engineSchematic: string,
): IndexedPartNumber[] {
  const lineLength = engineSchematic.indexOf('\n');
  // console.log(lineLength);
  const adjacentRange = 1;
  const indexedPartNumbers: IndexedPartNumber[] = partNumbers.map(
    (current) => ({ ...current, adjacentIndices: [] }),
  );

  for (const partNumber of indexedPartNumbers) {
    const adjacentIndices = [];

    for (let i = 0; i < adjacentRange * 2 + 1; i += 1) {
      const rowOverflowIndex =
        partNumber.startIndex -
        adjacentRange +
        (i - adjacentRange) * (lineLength + adjacentRange) +
        1;
      const didOverflowHorizontally =
        rowOverflowIndex < 0 || rowOverflowIndex > engineSchematic.length;

      if (didOverflowHorizontally) {
        console.log('this row is overflowing');
        continue;
      }

      const rowIndex = Math.floor(partNumber.startIndex / (lineLength + 1));
      const _overflowLeftAmount = Math.min(
        0,
        partNumber.startIndex - adjacentRange - rowIndex * (lineLength + 1),
      );

      const overflowRightAmount = Math.max(
        0,
        partNumber.startIndex +
          partNumber.value.length +
          adjacentRange -
          rowIndex * (lineLength + 1) -
          lineLength,
      );

      const row = Array.from(
        { length: partNumber.value.length + adjacentRange * 2 },
        (_, j) => {
          return (
            partNumber.startIndex +
            j -
            adjacentRange +
            (i - adjacentRange) * (lineLength + adjacentRange)
          );
        },
      );

      console.log(row, partNumber.value, overflowRightAmount);
      adjacentIndices.push(row);
    }

    console.log();

    partNumber.adjacentIndices = adjacentIndices;
  }

  return indexedPartNumbers;
}

function _visualizeAdjacentIndices(
  indexedPartNumbers: IndexedPartNumber[],
  engineSchematic: string,
) {
  const splittedSchematic = engineSchematic.split('');
  for (const indexedPartNumber of indexedPartNumbers) {
    for (const adjacentIndices of indexedPartNumber.adjacentIndices) {
      for (const adjacentIndex of adjacentIndices) {
        splittedSchematic[adjacentIndex] = 'X';
      }
    }
  }

  console.log(splittedSchematic.join(''));
}

async function main() {
  const _input = await _exampleFile.text();
  const input = `467..114..
...*......
..35..633.
.......123
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`;

  console.log(input);

  const partNumbers = getPartNumbers(input);
  // console.log(partNumbers)

  const indexedPartNumbers = getAdjacentPartNumberIndices(partNumbers, input);
  // console.log(indexedPartNumbers);

  _visualizeAdjacentIndices(indexedPartNumbers, input);
}

main();
