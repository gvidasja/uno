import {
  drawCircle,
  drawSemiCircle,
  drawText,
} from "../graphics.js";

export function drawPlayer({ g, x, y, width, height, name, handSize, turn, winner }) {
  const text = `${name} [${handSize || "0"}]`;
  const backgroundColor = turn ? "orangered" : "orange";
  const textColor = turn ? "white" : "black";

  drawSemiCircle(g, x + width / 2, y + height, width, backgroundColor);
  drawCircle(g, x + width / 2, y + height - width, width, backgroundColor);

  drawText(g, x, y + height - width / 2, width, height - width, text, textColor);

  winner && drawText(g, x, y, width, width, "WINNER", "black", "30px Arial");
}
