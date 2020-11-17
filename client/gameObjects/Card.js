import { drawCircle, drawRect, drawText } from "../graphics.js";
import { GameObject } from "./GameObject.js";

export class Card extends GameObject {
  static SIZE_X = 50;
  static SIZE_Y = 70;

  constructor(
    {
      x,
      y,
      sizeX = Card.SIZE_X,
      sizeY = Card.SIZE_Y,
      color,
      number,
    },
  ) {
    super({ x, y, sizeX, sizeY });

    this.color = color;
    this.number = number;
  }

  shouldChooseOverride() {
    return this.number === "PLUS_FOUR" || this.number == "CHANGE_COLOR";
  }

  draw(g) {
    super.draw(g);

    const { background, text: textColor } = mapColors(this.color);
    const { text, textFunc, font = "24px Arial" } = mapNumberToText(
      this.number,
    );

    drawRect(
      g,
      this.x,
      this.y,
      Card.SIZE_X,
      Card.SIZE_Y,
      background,
    );

    textFunc
      ? textFunc(g, this.x + 5, this.y + 15, this.sizeX - 10, this.sizeY - 30)
      : drawText(
        g,
        this.x,
        this.y,
        Card.SIZE_X,
        Card.SIZE_Y,
        text,
        textColor,
        font,
      );
  }
}

function mapNumberToText(number) {
  switch (number) {
    case "REVERSE":
      return { text: "🔄" };
    case "PLUS_TWO":
      return { text: "+2" };
    case "PLUS_FOUR":
      return { text: "+4" };
    case "SKIP_TURN":
      return { text: "🚫" };
    case "CHANGE_COLOR":
      return {
        textFunc: (g, x, y, sizeX, sizeY) => {
          drawText(g, x, y, sizeX / 2, sizeY / 2, "🔴");
          drawText(g, x + sizeX / 2, y, sizeX / 2, sizeY / 2, "🟡");
          drawText(g, x, y + sizeY / 2, sizeX / 2, sizeY / 2, "🟢");
          drawText(g, x + sizeX / 2, y + sizeY / 2, sizeX / 2, sizeY / 2, "🔵");
        },
        font: "8px Arial",
      };
    default:
      return { text: number };
  }
}

function mapColors(color) {
  switch (color) {
    case "NO_COLOR":
      return { background: "black", text: "white" };
    case "YELLOW":
      return { background: "yellow", text: "black" };
    case "GREEN":
      return { background: "green", text: "white" };
    case "BLUE":
      return { background: "blue", text: "white" };
    case "RED":
      return { background: "red", text: "white" };
    case "UNO":
      return { background: "red", text: "yellow" };
    default:
      throw new Error(`Unknown color: ${color}`);
  }
}
