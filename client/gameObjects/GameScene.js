import "../types.js";
import { drawPlayer } from "./player.js";
import { drawDeck } from "./deck.js";
import { drawCard, shouldChooseOverride } from "./card.js";
import { drawRect, drawEllipse, drawRectWithText } from "../graphics.js";
import { onRectClick } from './clickHandlers.js'

const CARD_WIDTH = 50
const CARD_HEIGHT = 70
const PLAYER_WIDTH = 60
const PLAYER_HEIGHT = 90

const clickHandlers = new Set()

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
    for (let clickHandler of clickHandlers) {
      if (clickHandler.intersects(x, y)) {
        clickHandler.click(x, y)
        break;
      }
    }
  }

  draw(g, onAction) {
    requestAnimationFrame(() => {
      clickHandlers.clear()

      const state = this.state;

      if (!state.globalState) {
        return;
      }

      const { players: playersZone, hand: handZone } = getZones(g);

      drawRect(g, 0, 0, g.canvas.width, g.canvas.height, "lightgrey");

      drawPlayers({
        g,
        ...playersZone,
        players: state.globalState.players,
        me: state.me,
        winner: state.globalState.winner,
      });

      drawHand({
        g,
        ...handZone,
        hand: state.hand,
        onAction,
      });

      drawTable({
        g,
        ...playersZone,
        deckSize: state.globalState.deckSize,
        pileSize: state.globalState.pileSize,
        topCard: state.globalState.topCard,
        colorOverride: state.globalState.colorOverride,
        onAction,
      });
    });
  }
}

function drawPlayers(
  { g, players: statePlayers, me, x, y, width, height, winner },
) {
  if (!statePlayers.length) {
    return
  }

  const middleX = x + width / 2;
  const middleY = y + height / 2;

  const playerEveryRad = Math.PI * 2 / statePlayers.length;

  while (statePlayers[0].id !== me.id) {
    statePlayers.push(statePlayers.shift());
  }

  const distanceFromCenterX = Math.min(width / 2 * 0.8, 250);
  const distanceFromCenterY = Math.min(height / 2 * 0.8, 200);

  statePlayers.forEach((x, i) =>
    drawPlayer({
      g,
      x: middleX +
        Math.cos(Math.PI / 2 + i * playerEveryRad) * distanceFromCenterX -
        PLAYER_WIDTH / 2,
      y: middleY +
        Math.sin(Math.PI / 2 + i * playerEveryRad) * distanceFromCenterY -
        PLAYER_HEIGHT / 2,
      width: 60,
      height: 90,
      handSize: x.handSize,
      name: x.id,
      turn: x.turn,
      winner: x.winner,
    })
  );
}

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

function drawHand({ g, hand, x, y, width, height, onAction }) {
  drawRect(g, x, y, width, height, "#aac");

  const MARGIN = 10;
  const cardsX = x + MARGIN;
  const cardsY = y + MARGIN;
  const cardsWidth = (width - 2 * MARGIN);
  const cardPerRow = Math.floor(cardsWidth / CARD_WIDTH);

  (hand || []).forEach((card, i) => {
    const col = i % cardPerRow;
    const row = Math.floor(i / cardPerRow);

    const cardX = cardsX + col * CARD_WIDTH
    const cardY = cardsY + row * CARD_HEIGHT

    drawCard({
      g,
      x: cardX,
      y: cardY,
      width: CARD_WIDTH,
      height: CARD_HEIGHT,
      color: card.color,
      number: card.number,
    });

    clickHandlers.add(onRectClick(cardX, cardsY, CARD_WIDTH, CARD_HEIGHT, () => 
      onAction(
        "PLAY_CARD",
        {
          CARD_COLOR: card.color,
          CARD_NUMBER: card.number,
          COLOR_OVERRIDE: shouldChooseOverride(card.number)
            ? toColor(prompt(
              "R G Ž M",
            ))
            : null,
        },
      )))
  });

  if (hand && hand.length == 2) {
    const callUnoX = cardsX
    const callUnoY = cardsY + Math.floor(hand.length / cardPerRow + 1) * CARD_HEIGHT

    drawRectWithText(g, callUnoX, callUnoY, CARD_WIDTH, CARD_HEIGHT, 'UNO!')

    clickHandlers.add(onRectClick(callUnoX, callUnoY, CARD_WIDTH, CARD_HEIGHT, () => onAction('CALL_UNO')))
  }
}

function drawTable({
  g,
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
  const centerX = x + width / 2;
  const centerY = y + height / 2;

  drawEllipse(g, centerX, centerY, 300, 240, "#987");

  const deckX = centerX - CARD_WIDTH - 10
  const deckY = centerY - CARD_HEIGHT / 2

  drawDeck({
    g,
    deckSize: deckSize,
    x: deckX,
    y: deckY,
    width: CARD_WIDTH,
    height: CARD_HEIGHT,
    drawTopCard: ({ g, x, y, width, height }) => drawCard({ g, x, y, width, height, color: "UNO", number: "UNO" }),
  });

  clickHandlers.add(onRectClick(deckX, deckY, CARD_WIDTH, CARD_HEIGHT, () => onAction("DRAW_CARD")));

  if (pileSize > 0) {
    const pileX = centerX + 10 
    const pileY = centerY - CARD_HEIGHT / 2

    drawDeck({
      g,
      deckSize: pileSize,
      x: pileX,
      y: pileY,
      width: CARD_WIDTH,
      height: CARD_HEIGHT,
      drawTopCard: ({ g, x, y, width, height }) => drawCard({ g, x, y, width, height, 
        color: colorOverride || topCard.color,
        number: topCard.number,
      }),
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
