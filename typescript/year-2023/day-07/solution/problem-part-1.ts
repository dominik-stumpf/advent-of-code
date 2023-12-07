export function solveProblem(input: string): number {
  const hands = parseHands(input);

  const evaluatedHands = hands.map((hand) => evaluateHand(hand));

  const rankedHands = rankHands(evaluatedHands);
  console.log(rankedHands);

  return rankedHands.reduce(
    (accumulator, hand) => accumulator + hand.bid * hand.rank,
    0,
  );
}
const cardLabels = [
  'A',
  'K',
  'Q',
  'J',
  'T',
  '9',
  '8',
  '7',
  '6',
  '5',
  '4',
  '3',
  '2',
] as const;
type CardLabel = (typeof cardLabels)[number];

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

type HandType = (typeof handTypes)[number];
type LabelCount = Record<CardLabel, number>;
type HandTypeMapCallback = (
  cards: string,
  labelCount: LabelCount,
) => HandStrength | undefined;

interface HandStrength {
  type: HandType;
}

interface EvaluatedHand extends HandStrength, Hand {}

const handTypeMap: Record<HandType, HandTypeMapCallback> = {
  fiveOfAKind: (_, labelCount) => {
    for (const [, count] of Object.entries(labelCount)) {
      if (count === 5) {
        return { type: 'fiveOfAKind' };
      }
    }

    return undefined;
  },
  fourOfAKind: (_, labelCount) => {
    for (const [, count] of Object.entries(labelCount)) {
      if (count === 4) {
        return {
          type: 'fourOfAKind',
        };
      }
    }

    return undefined;
  },
  fullHouse: (_, labelCount) => {
    let threeLabel: undefined | CardLabel;
    let twoLabel: undefined | CardLabel;
    for (const [label, count] of Object.entries(labelCount)) {
      if (!threeLabel && count === 3) {
        threeLabel = label as CardLabel;
      } else if (!twoLabel && count === 2) {
        twoLabel = label as CardLabel;
      }
    }

    if (threeLabel && twoLabel) {
      return { type: 'fullHouse' };
    }

    return undefined;
  },
  threeOfAKind: (_, labelCount) => {
    for (const [, count] of Object.entries(labelCount)) {
      if (count === 3) {
        return {
          type: 'threeOfAKind',
        };
      }
    }
    return undefined;
  },
  twoPair: (_, labelCount) => {
    const twoLabels: CardLabel[] = [];
    for (const [label, count] of Object.entries(labelCount)) {
      if (count === 2) {
        twoLabels.push(label as CardLabel);
      }
      if (twoLabels.length === 2) {
        return {
          type: 'twoPair',
        };
      }
    }
    return undefined;
  },
  onePair: (_, labelCount) => {
    for (const [, count] of Object.entries(labelCount)) {
      if (count === 2) {
        return { type: 'onePair' };
      }
    }
    return undefined;
  },
  highCard: () => {
    return {
      type: 'highCard',
    };
  },
};

function evaluateHand(hand: Hand): EvaluatedHand {
  const labelCount = Object.fromEntries(
    cardLabels.map((label) => [label, 0]),
  ) as LabelCount;

  for (let i = 0; i < hand.cards.length; i += 1) {
    labelCount[hand.cards[i] as CardLabel] += 1;
  }

  for (const handType of handTypes) {
    const evaluatedHand = handTypeMap[handType](hand.cards, labelCount);
    if (evaluatedHand) {
      return { ...hand, ...evaluatedHand };
    }
  }

  throw new Error('cards could not be evaluated');
}

interface RankedHand extends EvaluatedHand {
  rank: number;
}

function rankHands(evaluatedHands: EvaluatedHand[]): RankedHand[] {
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

      throw new Error('could not sort card');
    });
    result.push(...hands.map((hand) => ({ ...hand, rank: rankCount-- })));
  }

  return result;
}
