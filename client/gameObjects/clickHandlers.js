export function onRectClick(x, y, width, height, onClick) {
  return new ClickableRectangle({ x, y, width, height, onClick })
}

class ClickableRectangle {
  clickHandler;

  constructor({ x, y, width, height, onClick }) {
    this.x = x;
    this.y = y;
    this.width = width;
    this.height = height;
    this.clickHandler = onClick
  }

  click() {
    !!this.clickHandler && this.clickHandler();
  }

  intersects(x, y) {
    return this.x <= x && x <= this.x + this.width && this.y <= y && y <= this.y + this.height;
  }
}
