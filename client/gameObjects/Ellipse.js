import { drawEllipse } from "../graphics.js";
import { GameObject } from "./GameObject.js";

export class Ellipse extends GameObject {
  constructor({ x, y, sizeX, sizeY, color }) {
    super({ x: x - sizeX / 2, y: y - sizeY / 2, sizeX, sizeY });
    this.color = color;
  }

  draw(g) {
    super.draw(g);
    drawEllipse(
      g,
      this.x + this.sizeX / 2,
      this.y + this.sizeY / 2,
      this.sizeX,
      this.sizeY,
      this.color,
    );
  }
}
