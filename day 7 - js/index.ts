import input from "./puzzleInput.json";

const TypeDictionary: { [key: string]: number } = {
  HighCard: 0,
  OnePair: 1,
  TwoPair: 2,
  ThreeOfAKind: 3,
  FullHouse: 4,
  FourOfAKind: 5,
  FiveOfAKind: 6,
};

const cardStringToNumericalValueMap: { [key: string]: number } = {
  "2": 0,
  "3": 1,
  "4": 2,
  "5": 3,
  "6": 4,
  "7": 5,
  "8": 6,
  "9": 7,
  T: 8,
  J: 9,
  Q: 10,
  K: 11,
  A: 12,
};

interface ICard {
  bid: number;
  cards: number[];
  type: number;
}

const transformInputLineToCard = (
  line: string,
  shouldUseNewJokerRole: boolean
): ICard => {
  const [inputHand, bid] = line.split(" ");
  const cards = inputHand
    .split("")
    .map((c) => cardStringToNumericalValueMap[c]);
  const bidNumber = Number(bid);
  const handType = getTypeOfHand(cards, shouldUseNewJokerRole);

  const result: ICard = {
    type: handType,
    bid: bidNumber,
    cards,
  };

  return result;
};

const getTypeOfHand = (cards: number[], shouldUseNewJokerRole: boolean) => {
  const numberOfEachCardDictionary = cards.reduce<{ [key: string]: number }>(
    (previous, card) => {
      if (!previous[card]) {
        previous[card] = 0;
      }

      previous[card] += 1;
      return previous;
    },
    {}
  );

  const numberOfCards = Object.values(numberOfEachCardDictionary);
  if (!shouldUseNewJokerRole) {
    const type = transformHandToType(numberOfCards);
    return type;
  }

  const kvParisOfNumberOfCards = Object.entries(numberOfEachCardDictionary);
  console.log("kvParisOfNumberOfCards", kvParisOfNumberOfCards);
  let numberOfJokers = cards.reduce<number>((p, card) => {
    const isJoker = card === cardStringToNumericalValueMap["J"];
    if (isJoker) {
      return p + 1;
    }

    return p;
  }, 0);
  console.log("numberOfJokers", numberOfJokers);
  const indexOfJoker = kvParisOfNumberOfCards.findIndex(
    (k) => k[0] === cardStringToNumericalValueMap["J"].toString()
  );
  console.log("indexOfJoker", indexOfJoker);
  const highestPossibleHandWithJokers = kvParisOfNumberOfCards.map(
    ([k, _], i) => {
      const isJoker = i === indexOfJoker;
      const currentNumbers = [...numberOfCards];
      console.log(k, cardStringToNumericalValueMap["J"]);
      console.log(isJoker);
      if (!isJoker) {
        currentNumbers[i] += numberOfJokers;
        currentNumbers[indexOfJoker] = 0;
      }
      console.log(`currentNumbers iteration ${i}: ${currentNumbers}`);
      const t = transformHandToType(currentNumbers);
      return t;
    }
  );
  console.log("highestPossibleHandWithJokers: ", highestPossibleHandWithJokers);
  const max = Math.max(...highestPossibleHandWithJokers);
  console.log("MAX: ", max);
  return max;
};

const transformHandToType = (numberOfCards: number[]): number => {
  switch (true) {
    case numberOfCards.includes(5): {
      return TypeDictionary.FiveOfAKind;
    }
    case numberOfCards.includes(4): {
      return TypeDictionary.FourOfAKind;
    }
    case numberOfCards.includes(3) && numberOfCards.includes(2): {
      return TypeDictionary.FullHouse;
    }
    case numberOfCards.includes(3): {
      return TypeDictionary.ThreeOfAKind;
    }
    case numberOfCards.filter((n) => n === 2).length === 2: {
      return TypeDictionary.TwoPair;
    }

    case numberOfCards.includes(2): {
      return TypeDictionary.OnePair;
    }
    default:
      return TypeDictionary.HighCard;
  }
};

const groupCardsByType = (cards: ICard[]) => {
  const cardsByType: ICard[][] = [[], [], [], [], [], [], []];

  cards.forEach((card) => {
    cardsByType[card.type].push(card);
  });
  console.log(cardsByType);
  return cardsByType;
};

const sortCardsPerGroup = (cards: ICard[][]) => {
  const sortedCards = cards.map((group) => {
    return group.sort((a, b) => {
      for (let i = 0; i < group.length; i += 1) {
        const aCard = a.cards[i];
        const bCard = b.cards[i];

        if (aCard === bCard) {
          continue;
        }

        return aCard < bCard ? -1 : 1;
      }

      return 0;
    });
  });

  return sortedCards;
};

const getTotalWinnings = (cardsGroupedByType: ICard[][]) => {
  let currentRank = 1;
  const numberForLogging = 3;
  const result = cardsGroupedByType.reduce((sum, group, i) => {
    const groupResult = group.reduce((s, c) => {
      if (i === numberForLogging) {
        console.log(`Rank: ${currentRank}; ${JSON.stringify(c)} `);
      }
      const product = c.bid * currentRank;
      currentRank += 1;
      return s + product;
    }, 0);
    if (i === numberForLogging) {
      console.log(
        `Group: ${
          Object.entries(TypeDictionary).find(([_, v]) => v === i)?.[0]
        }`
      );
    }
    return sum + groupResult;
  }, 0);

  return result;
};

(() => {
  const shouldUseNewJokerRole = true;
  if (shouldUseNewJokerRole) {
    cardStringToNumericalValueMap["J"] = -1;
  }

  const cards = input.map((line) =>
    transformInputLineToCard(line, shouldUseNewJokerRole)
  );
  const cardsGroupedByType = groupCardsByType(cards);
  const sortedCards = sortCardsPerGroup(cardsGroupedByType);
  const totalWinnings = getTotalWinnings(sortedCards);
  console.log(totalWinnings);
})();
