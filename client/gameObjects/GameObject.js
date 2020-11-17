export class GameObject {
  clickHandler;

  constructor({ x, y, sizeX, sizeY }) {
    this.x = x;
    this.y = y;
    this.sizeX = sizeX;
    this.sizeY = sizeY;
  }

  draw() {}
  addClickHandler(handler) {
    const self = this;
    this.clickHandler = () => handler(self);
  }
  click() {
    !!this.clickHandler && this.clickHandler();
  }

  intersects(x, y) {
    return this.x <= x && x <= this.x + this.sizeX && this.y <= y &&
      y <= this.y + this.sizeY;
  }
}
