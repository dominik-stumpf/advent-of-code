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
        continue;
      }

      const rowIndex = Math.floor(partNumber.startIndex / (lineLength + 1));
      const overflowLeftAmount = Math.min(
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
        {
          length:
            (i === adjacentRange ? 0 : partNumber.value.length) +
            adjacentRange * 2 -
            overflowRightAmount +
            overflowLeftAmount,
        },
        (_, j) => {
          return (
            (i === adjacentRange && j >= adjacentRange + overflowLeftAmount
              ? partNumber.value.length
              : 0) +
            partNumber.startIndex +
            j -
            adjacentRange +
            (i - adjacentRange) * (lineLength + adjacentRange) -
            overflowLeftAmount
          );
        },
      );

      adjacentIndices.push(row);
    }

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
}

function sumValidPartNumbers(
  indexedPartNumbers: IndexedPartNumber[],
  engineSchematic: string,
): number {
  let sum = 0;

  for (const indexedPartNumber of indexedPartNumbers) {
    for (const adjacentIndices of indexedPartNumber.adjacentIndices) {
      for (const adjacentIndex of adjacentIndices) {
        if (
          engineSchematic[adjacentIndex] !== '.' &&
          !validDigits.includes(engineSchematic[adjacentIndex])
        ) {
          sum += Number(indexedPartNumber.value);
          break;
        }
      }
    }
  }

  return sum;
}

export function solveProblem(input: string): number {
  const partNumbers = getPartNumbers(input);
  const indexedPartNumbers = getAdjacentPartNumberIndices(partNumbers, input);
  const sum = sumValidPartNumbers(indexedPartNumbers, input);

  return sum;
}
