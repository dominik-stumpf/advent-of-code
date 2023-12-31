interface ScratchCard {
  winningNumbers: Set<number>;
  receivedNumbers: Set<number>;
}

interface EvaluatedScratchCard {
  match: number;
  instance: number;
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

function countScratchCardMatch({
  winningNumbers,
  receivedNumbers,
}: ScratchCard): number {
  let matchCount = 0;

  for (const winningNumber of winningNumbers) {
    if (receivedNumbers.has(winningNumber)) {
      matchCount += 1;
    }
  }

  return matchCount;
}

function evaluateScratchCards(
  scratchCards: ScratchCard[],
): EvaluatedScratchCard[] {
  const evaluatedCards: EvaluatedScratchCard[] = scratchCards.map(
    (scratchCard) => ({
      match: countScratchCardMatch(scratchCard),
      instance: 0,
    }),
  );

  for (let i = 0; i < evaluatedCards.length; i += 1) {
    const card = evaluatedCards[i];
    card.instance += 1;

    for (let j = 1; j < card.match + 1; j += 1) {
      evaluatedCards[i + j].instance += card.instance;
    }
  }

  return evaluatedCards;
}

export function solveProblem(input: string): number {
  const inputLines = input.split('\n');

  inputLines.pop(); // omit last trailing empty line

  const parsedCards = inputLines.map((line) => parseCard(line));
  const evaluatedCards = evaluateScratchCards(parsedCards);
  const sum = evaluatedCards.reduce(
    (accumulator, card) => accumulator + card.instance,
    0,
  );

  return sum;
}
