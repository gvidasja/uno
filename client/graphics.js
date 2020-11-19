/** @param {CanvasRenderingContext2D} g */
export function drawRectWithText(
  g,
  x,
  y,
  sizeX,
  sizeY,
  text,
  backgroundColor = "white",
  textColor = "black",
) {
  drawRect(g, x, y, sizeX, sizeY, backgroundColor);
  drawText(g, x, y, sizeX, sizeY, text, textColor);
}

/** @param {CanvasRenderingContext2D} g */
export function drawSemiCircle(g, xCenter, yCenter, diameter, backgrounColor) {
  g.fillStyle = backgrounColor;
  g.beginPath();
  g.arc(xCenter, yCenter, diameter / 2, Math.PI, Math.PI * 2);
  g.closePath();
  g.fill();
}

/** @param {CanvasRenderingContext2D} g */
export function drawCircle(g, xCenter, yCenter, diameter, backgrounColor) {
  g.fillStyle = backgrounColor;
  g.beginPath();
  g.arc(xCenter, yCenter, diameter / 2, 0, Math.PI * 2);
  g.closePath();
  g.fill();
}

/** @param {CanvasRenderingContext2D} g */
export function drawEllipse(g, xCenter, yCenter, sizeX, sizeY, backgrounColor) {
  g.fillStyle = backgrounColor;
  g.beginPath();
  g.ellipse(xCenter, yCenter, sizeX / 2, sizeY / 2, 0, 0, Math.PI * 2);
  g.closePath();
  g.fill();
}

/** @param {CanvasRenderingContext2D} g */
export function drawRect(g, x, y, sizeX, sizeY, backgroundColor = "white") {
  g.fillStyle = backgroundColor;
  g.fillRect(x, y, sizeX, sizeY);
}

/** @param {CanvasRenderingContext2D} g */
export function drawText(
  g,
  x,
  y,
  sizeX,
  sizeY,
  text,
  textColor = "black",
  font = "14px Arial",
) {
  g.textAlign = "center";
  g.textBaseline = "middle";
  g.font = font;
  g.fillStyle = textColor;
  g.fillText(
    text,
    x + sizeX / 2,
    y + sizeY / 2,
  );
}
