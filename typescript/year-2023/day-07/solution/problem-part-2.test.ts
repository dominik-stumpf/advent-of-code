import { describe, expect, it } from 'bun:test';
import { HandType, evaluateHand, rankHands } from './problem-part-2';

// test inputs copied from: https://www.reddit.com/r/adventofcode/comments/18cubzw/2023_day_7_part_2_a_bunch_of_sample_data/

// Five of a kind
const fiveOfAKindHands = [
  'JJJJJ',
  'AAAAA',
  'JAAAA',
  'AJAAA',
  'AAJAA',
  'AAAJA',
  'AAAAJ',
];
// Four of a kind
const fourOfAKindHands = [
  'AA8AA',
  'TTTT8',
  'JTTT8',
  'TJTT8',
  'TTJT8',
  'TTTJ8',
  'TTT8J',
  'T55J5',
  'KTJJT',
  'QQQJA',
  'QJJQ2',
  'JJQJ4',
  'JJ2J9',
  'JTJ55',
];
// Full house
const fullHouseHands = [
  '23332',
  'J2233',
  '2J233',
  '22J33',
  '223J3',
  '2233J',
  '22333',
  '25J52',
];
// Three of a kind
const threeOfAKindHands = [
  'AJKJ4',
  'TTT98',
  'JTT98',
  'TJT98',
  'TTJ98',
  'TT9J8',
  'TT98J',
  'T9T8J',
  'T98TJ',
  'T98JT',
  'TQJQ8',
];
// Two pair
const twoPairHands = ['23432', 'KK677', 'KK677'];
// One pair
const onePairHands = [
  '32T3K',
  'A23A4',
  '32T3K',
  'J2345',
  '2J345',
  '23J45',
  '234J5',
  '2345J',
  '5TK4J',
];
// High card
const highCardHands = ['23456'];

const handTypeHands: Record<HandType, string[]> = {
  fiveOfAKind: fiveOfAKindHands,
  fourOfAKind: fourOfAKindHands,
  fullHouse: fullHouseHands,
  threeOfAKind: threeOfAKindHands,
  twoPair: twoPairHands,
  onePair: onePairHands,
  highCard: highCardHands,
};

describe('utility functions for day-07/part-2', () => {
  for (const [untypedHandType, rawHands] of Object.entries(handTypeHands)) {
    const handType = untypedHandType as HandType;

    it(`should recognize ${handType} hands`, () => {
      const hands = rawHands.map((hand) => ({ bid: 0, cards: hand }));
      const evaluatedHands = hands.map((hand) => evaluateHand(hand));
      const rankedHands = rankHands(evaluatedHands);

      for (const hand of rankedHands) {
        if (hand.type !== handType) {
          console.log(hand.cards);
        }
        expect(hand.type).toBe(handType);
      }
    });
  }
});
