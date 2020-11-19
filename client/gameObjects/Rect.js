import { drawRect } from "../graphics.js";
import { GameObject } from "./GameObject.js";

export class Rect extends GameObject {
  constructor({ x, y, sizeX, sizeY, color }) {
    super({ x, y, sizeX, sizeY });
    this.color = color;
  }

  draw(g) {
    super.draw(g);
    drawRect(g, this.x, this.y, this.sizeX, this.sizeY, this.color);
  }
}
