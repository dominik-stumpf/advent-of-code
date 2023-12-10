export function solveProblem(input: string): number {
  const hands = parseHands(input);
  const evaluatedHands = hands.map((hand) => evaluateHand(hand));
  const rankedHands = rankHands(evaluatedHands);
  return rankedHands.reduce(
    (accumulator, hand) => accumulator + hand.bid * hand.rank,
    0,
  );
}

const cardLabels = [
  'A',
  'K',
  'Q',
  'T',
  '9',
  '8',
  '7',
  '6',
  '5',
  '4',
  '3',
  '2',
  'J',
] as const;

export type CardLabel = (typeof cardLabels)[number];

interface Hand {
  bid: number;
  cards: string;
}

function parseHands(rawHands: string): Hand[] {
  const hands = rawHands
    .split('\n')
    .slice(0, -1)
    .map((hand) => hand.split(' '))
    .map(([cards, bid]) => ({ cards, bid: Number(bid) }));
  return hands;
}

const handTypes = [
  'fiveOfAKind',
  'fourOfAKind',
  'fullHouse',
  'threeOfAKind',
  'twoPair',
  'onePair',
  'highCard',
] as const;

export type HandType = (typeof handTypes)[number];
type LabelCount = Record<CardLabel, number>;
type HandTypeMapCallback = (cards: string, labelCount: LabelCount) => boolean;

interface HandStrength {
  type: HandType;
}

interface EvaluatedHand extends HandStrength, Hand {}

const handTypeMap: Record<HandType, HandTypeMapCallback> = {
  fiveOfAKind: (_, labelCount) => {
    const jokerCount = labelCount.J;
    for (const [label, count] of Object.entries(labelCount)) {
      if (label === 'J') {
        continue;
      }
      if (count + jokerCount === 5) {
        return true;
      }
    }

    return false;
  },
  fourOfAKind: (_, labelCount) => {
    const jokerCount = labelCount.J;
    for (const [label, count] of Object.entries(labelCount)) {
      if (label === 'J') {
        continue;
      }
      if (jokerCount + count === 4) {
        return true;
      }
    }

    return false;
  },
  fullHouse: (_, labelCount) => {
    // 23332
    const jokerCount = labelCount.J;
    const labelValues = Object.entries(labelCount)
      .filter(([key]) => key !== 'J')
      .map(([, value]) => value);

    if (
      jokerCount === 1 &&
      labelValues.filter((label) => label === 2).length === 2
    ) {
      return true;
    }
    if (
      jokerCount === 0 &&
      labelValues.includes(3) &&
      labelValues.includes(2)
    ) {
      return true;
    }
    return false;
  },
  threeOfAKind: (_, labelCount) => {
    const jokerCount = labelCount.J;
    for (const [label, count] of Object.entries(labelCount)) {
      if (label === 'J') {
        continue;
      }
      if (count + jokerCount === 3) {
        return true;
      }
    }
    return false;
  },
  twoPair: (_, labelCount) => {
    const labelValues = Object.entries(labelCount)
      .filter(([key]) => key !== 'J')
      .map(([, value]) => value);

    if (labelValues.filter((label) => label === 2).length === 2) {
      return true;
    }

    return false;
  },
  onePair: (_, labelCount) => {
    const jokerCount = labelCount.J;
    for (const [label, count] of Object.entries(labelCount)) {
      if (label === 'J') {
        continue;
      }
      if (count + jokerCount === 2) {
        return true;
      }
    }
    return false;
  },
  highCard: () => {
    return true;
  },
};

export function evaluateHand(hand: Hand): EvaluatedHand {
  const labelCount = Object.fromEntries(
    cardLabels.map((label) => [label, 0]),
  ) as LabelCount;

  for (let i = 0; i < hand.cards.length; i += 1) {
    labelCount[hand.cards[i] as CardLabel] += 1;
  }

  for (const handType of handTypes) {
    const evaluatedHand = handTypeMap[handType](hand.cards, labelCount);
    if (evaluatedHand) {
      return { ...hand, type: handType };
    }
  }

  throw new Error('cards could not be evaluated');
}

interface RankedHand extends EvaluatedHand {
  rank: number;
}

export function rankHands(evaluatedHands: EvaluatedHand[]): RankedHand[] {
  const result: RankedHand[] = [];

  let rankCount = evaluatedHands.length;

  for (const handType of handTypes) {
    const hands = evaluatedHands.filter((hand) => hand.type === handType);
    hands.sort((a, b) => {
      for (let i = 0; i < a.cards.length; i += 1) {
        const [labelA, labelB] = [a.cards[i], b.cards[i]].map((label) =>
          cardLabels.indexOf(label as CardLabel),
        );
        if (labelA === labelB) {
          continue;
        }
        return labelA - labelB;
      }

      return 0;
    });
    result.push(...hands.map((hand) => ({ ...hand, rank: rankCount-- })));
  }

  return result;
}
