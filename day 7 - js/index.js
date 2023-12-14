import input from "./puzzleInput.json" assert { type: "json" };

const typeDictionary = {
  FiveOfAKind: 0,
  FourOfAKind: 1,
  FullHouse: 2,
  ThreeOfAKind: 3,
  TwoPair: 4,
  OnePair: 5,
  HighCard: 6,
};

const cardStringToNumericalValueMap = {
  2: 0,
  3: 1,
  4: 2,
  5: 3,
  6: 4,
  7: 5,
  8: 6,
  9: 7,
  T: 8,
  J: 9,
  Q: 10,
  K: 11,
  A: 12,
};

/**
 *
 * @param {string} line
 * @returns {type: typeDictionary, bid: number, cards: string[]}
 */
const transformInputLineToCard = (line) => {
  const [inputHand, bid] = line.split(" ");
  const cards = inputHand
    .split("")
    .map((c) => cardStringToNumericalValueMap[c]);
  const bidNumber = Number(bid);
  const handType = getTypeOfHand(cards);

  return {
    type: handType,
    bid: bidNumber,
    cards,
  };
};

/**
 * @param {number[]} cards
 * @returns {typeDictionary}
 */
const getTypeOfHand = (cards) => {
  const numberOfEachCardDictionary = cards.reduce((previous, card) => {
    if (!previous[card]) {
      previous[card] = 0;
    }

    previous[card] += 1;
    return previous;
  }, {});

  const numberOfEachCard = Object.values(numberOfEachCardDictionary);

  if (numberOfEachCard.includes(5)) {
    return typeDictionary.FiveOfAKind;
  }

  if (numberOfEachCard.includes(4)) {
    return typeDictionary.FourOfAKind;
  }

  if (numberOfEachCard.includes(3) && numberOfEachCard.includes(2)) {
    return typeDictionary.FullHouse;
  }

  if (numberOfEachCard.includes(3)) {
    return typeDictionary.ThreeOfAKind;
  }

  if (numberOfEachCard.filter((n) => n === 2).length === 2) {
    return typeDictionary.TwoPair;
  }

  if (numberOfEachCard.includes(2)) {
    return typeDictionary.OnePair;
  }

  return typeDictionary.HighCard;
};

/**
 *
 * @param {{type: typeDictionary, bid: number, cards: string[]}[]} cards
 * @returns {{type: typeDictionary, bid: number, cards: string[]}[]}
 */
const groupCardsByType = (cards) => {
  const cardsByType = [[], [], [], [], [], [], []];

  cards.forEach((card) => {
    cardsByType[card.type].push(card);
  });

  return cardsByType;
};

/**
 *
 * @param {{type: typeDictionary, bid: number, cards: string[]}[]} cards
 * @returns {{type: typeDictionary, bid: number, cards: string[]}[]}
 */
const sortCardsPerGroup = (cards) => {
  const sortedCards = cards.map((group) => {
    return group.sort((a, b) => {
      for (let i = 0; i < group.length; i += 1) {
        const aCard = a.cards[i];
        const bCard = b.cards[i];

        if (aCard === bCard) {
          continue;
        }

        return aCard > bCard ? -1 : 1;
      }

      return 0;
    });
  });

  return sortedCards;
};

/**
 *
 * @param {{type: typeDictionary, bid: number, cards: string[]}[]} cardsGroupedByType
 * @param {number} numberOfElements
 * @returns {number}
 */
const getTotalWinnings = (cardsGroupedByType, numberOfElements) => {
  let currentRank = numberOfElements;
  const result = cardsGroupedByType.reduce((sum, group) => {
    const groupResult = group.reduce((s, c) => {
      const product = c.bid * currentRank;
      currentRank -= 1;
      return s + product;
    }, 0);
    return sum + groupResult;
  }, 0);

  return result;
};

(() => {
  const cards = input.map(transformInputLineToCard);
  console.log("cards", cards);
  const cardsGroupedByType = groupCardsByType(cards);
  console.log("cardsGroupedByType", cardsGroupedByType);
  const sortedCards = sortCardsPerGroup(cardsGroupedByType);
  console.log("sortedCards", sortedCards);
  const totalWinnings = getTotalWinnings(sortedCards, cards.length);
  console.log(totalWinnings);
})();
