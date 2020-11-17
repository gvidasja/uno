import { drawRect } from "../graphics.js";
import { Card } from "./Card.js";
import { GameObject } from "./GameObject.js";

export class Deck extends GameObject {
  static MAX_DECK_SIZE = 108;
  static MAX_SIZE_DECK_DRAW_HEIGH = 20;

  recalculateHeight() {
    const deckHeight = Deck.MAX_SIZE_DECK_DRAW_HEIGH * this.deckSize /
      Deck.MAX_DECK_SIZE;

    this.sizeY = this.baseSizeY + deckHeight;
    this.y = this.baseY - deckHeight;

    this.topCard.x = this.x;
    this.topCard.y = this.y;
  }

  constructor(
    {
      x,
      y,
      sizeX = Card.SIZE_X,
      sizeY = Card.SIZE_Y,
      deckSize,
      topCard,
    },
  ) {
    super({ x, y, sizeX, sizeY });
    this.topCard = topCard;
    this.baseY = y;
    this.baseSizeY = sizeY;
    this.deckSize = deckSize;
  }

  draw(g) {
    this.recalculateHeight();

    super.draw(g);

    drawRect(g, this.x, this.y, this.sizeX, this.sizeY, "darkgrey");
    this.topCard.draw(g);
  }
}
