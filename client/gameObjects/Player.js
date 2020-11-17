import {
  drawCircle,
  drawRect,
  drawRectWithText,
  drawSemiCircle,
  drawText,
} from "../graphics.js";
import { GameObject } from "./GameObject.js";

export class Player extends GameObject {
  static SIZE_X = 60;
  static SIZE_Y = 90;

  constructor(
    {
      x,
      y,
      sizeX = Player.SIZE_X,
      sizeY = Player.SIZE_Y,
      name,
      handSize,
      turn,
    },
  ) {
    super({ x, y, sizeX, sizeY });

    this.name = name;
    this.handSize = handSize;
    this.turn = turn;
  }

  draw(g) {
    const text = `${this.name} [${this.handSize || "0"}]`;
    const backgroundColor = this.turn ? "orangered" : "orange";
    const textColor = this.turn ? "white" : "black";

    drawSemiCircle(
      g,
      this.x + this.sizeX / 2,
      this.y + this.sizeY,
      this.sizeX,
      backgroundColor,
    );

    drawCircle(
      g,
      this.x + this.sizeX / 2,
      this.y + this.sizeY - this.sizeX,
      this.sizeX,
      backgroundColor,
    );

    drawText(
      g,
      this.x,
      this.y + this.sizeY - this.sizeX / 2,
      this.sizeX,
      this.sizeY - this.sizeX,
      text,
      textColor,
    );
  }
}
