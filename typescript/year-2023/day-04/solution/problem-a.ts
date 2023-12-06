const _exampleFile = Bun.file('input/example-a');
const inputFile = Bun.file('input/input-a');

interface ScratchCard {
  winningNumbers: Set<number>;
  receivedNumbers: Set<number>;
}

function parseCard(cardLine: string): ScratchCard {
  const [winningNumbers, receivedNumbers] = cardLine
    .slice(cardLine.indexOf(': ') + 1)
    .split('|')
    .map((numbers) => {
      return new Set(
        numbers
          .trim()
          .split(/ +/)
          .map((number) => Number(number)),
      );
    });

  return { winningNumbers, receivedNumbers };
}

function evaluateScratchCardPoint({
  winningNumbers,
  receivedNumbers,
}: ScratchCard): number {
  let matchCount = 0;

  for (const winningNumber of winningNumbers) {
    if (receivedNumbers.has(winningNumber)) {
      matchCount += 1;
    }
  }

  return matchCount && 2 ** (matchCount - 1);
}

async function main() {
  const input = await inputFile.text();
  const inputLines = input.split('\n');

  inputLines.pop(); // omit last trailing empty line

  let sum = 0;

  for (const line of inputLines) {
    const parsedCard = parseCard(line);
    const cardPoint = evaluateScratchCardPoint(parsedCard);

    sum += cardPoint;
  }

  console.log('sum:', sum);
}

main();
