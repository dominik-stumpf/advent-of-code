const _exampleFile = Bun.file('input/example-b');
const inputFile = Bun.file('input/input-b');

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

  console.log(splittedSchematic.join(''));
}

// biome-ignore lint: spaghetti is my speciality
function getPartNumberProducts(
  indexedPartNumbers: IndexedPartNumber[],
  engineSchematic: string,
): number[] {
  const partNumberProducts: number[] = [];

  for (let i = 0; i < engineSchematic.length; i += 1) {
    const char = engineSchematic[i];
    if (char === '*') {
      const collectedPartNumbers: IndexedPartNumber[] = [];

      partNumberLoop: for (const indexedPartNumber of indexedPartNumbers) {
        for (const adjacentIndices of indexedPartNumber.adjacentIndices) {
          for (const adjacentIndex of adjacentIndices) {
            if (adjacentIndex === i) {
              collectedPartNumbers.push(indexedPartNumber);
            }

            if (collectedPartNumbers.length > 2) {
              break partNumberLoop;
            }
          }
        }
      }

      if (collectedPartNumbers.length === 2) {
        partNumberProducts.push(
          collectedPartNumbers.reduce(
            (accumulator, partNumber) => accumulator * Number(partNumber.value),
            1,
          ),
        );
      }
    }
  }

  return partNumberProducts;
}

async function main() {
  const input = await inputFile.text();

  const partNumbers = getPartNumbers(input);
  const indexedPartNumbers = getAdjacentPartNumberIndices(partNumbers, input);
  const partNumberProducts = getPartNumberProducts(indexedPartNumbers, input);
  console.log('partNumberProducts', partNumberProducts);

  const sum = partNumberProducts.reduce(
    (accumulator, product) => accumulator + product,
    0,
  );
  console.log('sum:', sum);
}

main();
