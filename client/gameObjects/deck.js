import { drawRect } from "../graphics.js";

const MAX_DECK_SIZE = 108;
const MAX_SIZE_DECK_DRAW_HEIGH = 20;

export function drawDeck({ g, x, y, width, height, deckSize, drawTopCard }) {
  const baseY = y;
  const baseHeight = height;

  const deckHeight = MAX_SIZE_DECK_DRAW_HEIGH * deckSize / MAX_DECK_SIZE;

  height = baseHeight + deckHeight;
  y = baseY - deckHeight;

  drawRect(g, x, y, width, height, "darkgrey");
  drawTopCard({ g, x, y, width, height })
}
