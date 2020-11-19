import "../types.js";
import { Player } from "./Player.js";
import { Deck } from "./Deck.js";
import { Card } from "./Card.js";
import { GameObject } from "./GameObject.js";
import { Rect } from "./Rect.js";
import { Ellipse } from "./Ellipse.js";
import { drawRect } from "../graphics.js";

export class GameScene {
  /** @type {UnoState} */
  state = [];

  /**
   * @param {UnoState} newState
   * @param {CanvasRenderingContext2D} g
   * */
  update(newState) {
    if (!newState) return;

    this.state = newState;
  }

  click(x, y) {
    for (let i = this.objects.length - 1; i >= 0; i--) {
      if (this.objects[i].intersects(x, y)) {
        this.objects[i].click();
        break;
      }
    }
  }

  draw(g, onAction) {
    requestAnimationFrame(() => {
      const state = this.state;

      if (!state.globalState) {
        return;
      }

      const { players: playersZone, hand: handZone } = getZones(g);

      const players = createPlayers({
        ...playersZone,
        players: state.globalState.players,
        me: state.me,
        winner: state.globalState.winner,
      });

      const hand = createHand({
        ...handZone,
        hand: state.hand,
        onAction,
      });

      const table = createTable({
        ...playersZone,
        deckSize: state.globalState.deckSize,
        pileSize: state.globalState.pileSize,
        topCard: state.globalState.topCard,
        colorOverride: state.globalState.colorOverride,
        onAction,
      });

      /** @type {GameObject[]} */
      this.objects = [...players, ...hand, ...table].sort((o1, o2) =>
        (o1.y - o2.y) || (o1.x - o2.x)
      );

      drawRect(g, 0, 0, g.canvas.width, g.canvas.height, "lightgrey");
      this.objects.forEach((x) => x.draw(g));
    });
  }
}

function createPlayers(
  { players: statePlayers, me, x, y, width, height, winner },
) {
  if (statePlayers && statePlayers.length > 0) {
    const middleX = x + width / 2;
    const middleY = y + height / 2;

    const playerEveryRad = Math.PI * 2 / statePlayers.length;

    while (statePlayers[0].name !== me.name) {
      statePlayers.push(statePlayers.shift());
    }

    const distanceFromCenterX = Math.min(width / 2 * 0.8, 250);
    const distanceFromCenterY = Math.min(height / 2 * 0.8, 200);

    const players = statePlayers.flatMap((x, i) =>
      new Player({
        x: middleX +
          Math.cos(Math.PI / 2 + i * playerEveryRad) * distanceFromCenterX -
          Player.SIZE_X / 2,
        y: middleY +
          Math.sin(Math.PI / 2 + i * playerEveryRad) * distanceFromCenterY -
          Player.SIZE_Y / 2,
        handSize: x.handSize,
        name: x.name,
        turn: x.turn,
        winner: x.winner,
      })
    );
    return players;
  } else {
    return [];
  }
}

function createHand({ hand, x, y, width, height, onAction }) {
  const background = new Rect(
    { x, y, sizeX: width, sizeY: height, color: "#aac" },
  );

  const MARGIN = 10;
  const cardsX = x + MARGIN;
  const cardsY = y + MARGIN;
  const cardsWidth = (width - 2 * MARGIN);
  const cardPerRow = Math.floor(cardsWidth / Card.SIZE_X);

  const cards = (hand || []).map((x, i) => {
    const col = i % cardPerRow;
    const row = Math.floor(i / cardPerRow);

    const card = new Card({
      x: cardsX + col * Card.SIZE_X,
      y: cardsY + row * Card.SIZE_Y,
      color: x.color,
      number: x.number,
    });

    const toColor = (letter = "R") => {
      switch (letter.toUpperCase()) {
        case "R":
          return "RED";
        case "Ž":
          return "GREEN";
        case "G":
          return "YELLOW";
        case "M":
          return "BLUE";
        default:
          return "RED";
      }
    };

    card.addClickHandler((card) =>
      onAction(
        "PLAY_CARD",
        {
          CARD_COLOR: card.color,
          CARD_NUMBER: card.number,
          COLOR_OVERRIDE: card.shouldChooseOverride()
            ? toColor(prompt(
              "R G Ž M",
            ))
            : null,
        },
      )
    );

    return card;
  });

  return [background, ...cards];
}

function* createTable({
  deckSize,
  pileSize,
  topCard,
  colorOverride,
  x,
  y,
  width,
  height,
  onAction,
}) {
  const middleX = x + width / 2;
  const middleY = y + height / 2;

  yield new Ellipse(
    {
      x: middleX,
      y: middleY,
      sizeX: 300,
      sizeY: 240,
      color: "#987",
    },
  );

  const deck = new Deck(
    {
      deckSize: deckSize,
      x: middleX - Card.SIZE_X - 10,
      y: middleY - Card.SIZE_Y / 2,
      topCard: new Card({ color: "UNO", number: "UNO" }),
    },
  );

  deck.addClickHandler(() => onAction("DRAW_CARD"));

  yield deck;

  if (pileSize > 0) {
    yield new Deck({
      deckSize: pileSize,
      x: middleX + 10,
      y: middleY - Card.SIZE_Y / 2,
      topCard: new Card(
        {
          color: colorOverride || topCard.color,
          number: topCard.number,
        },
      ),
    });
  }
}

const ratio = (k) => [k / (1 + k), 1 / (1 + k)];

function getZones(g) {
  const width = g.canvas.width, height = g.canvas.height;

  if (width > height) {
    const [a, b] = ratio(5 / 3);

    return {
      players: { x: 0, y: 0, width: width * a, height },
      hand: { x: width * a, y: 0, width: width * b, height },
    };
  } else {
    const [a, b] = ratio(5 / 3);

    return {
      players: { x: 0, y: 0, width, height: height * a },
      hand: { x: 0, y: height * a, width, height: height * b },
    };
  }
}
