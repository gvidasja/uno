import { drawRect, drawText } from "../graphics.js";

export function shouldChooseOverride(number) {
  return number === "PLUS_FOUR" || number == "CHANGE_COLOR";
}

export function drawCard({ g, x, y, width, height, color, number }) {
  const { background: backgroundColor, text: textColor } = mapColors(color);
  const { text, textFunc, font = "24px Arial" } = mapNumberToText(number);

  drawRect(g, x, y, width, height, backgroundColor);

  textFunc
    ? textFunc(g, x + 5, y + 15, width - 10, height - 30)
    : drawText(g, x, y, width, height, text, textColor, font);
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
        textFunc: (g, x, y, width, height) => {
          drawText(g, x, y, width / 2, height / 2, "🔴");
          drawText(g, x + width / 2, y, width / 2, height / 2, "🟡");
          drawText(g, x, y + height / 2, width / 2, height / 2, "🟢");
          drawText(g, x + width / 2, y + height / 2, width / 2, height / 2, "🔵");
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
