import "./types.js";
import { openSocket } from "./io.js";
import { GameScene } from "./gameObjects/GameScene.js";

const roomId = location.pathname.replace(/[^\d]/, "");

const myName = localStorage.getItem("name") || prompt("Enter name");
localStorage.setItem("name", myName);
const socket = openSocket(`ws://${location.host}/ws/${roomId}`, myName);

const onAction = (action, data) => socket.send({ action, data });

/** @type {HTMLCanvasElement} */
const canvas = document.getElementById("game");
const g = canvas.getContext("2d");

const gameScene = new GameScene();

canvas.addEventListener("click", (e) => {
  const rect = e.target.getBoundingClientRect();
  const x = e.clientX - rect.left;
  const y = e.clientY - rect.top;

  gameScene.click(x, y);
});

function refitCanvas() {
  canvas.width = window.innerWidth;
  canvas.height = window.innerHeight;

  gameScene.draw(g, onAction);
}

refitCanvas();

window.addEventListener("resize", refitCanvas);
socket.subscribe((message) => {
  gameScene.update(message);
  gameScene.draw(g, onAction);
});

socket.send(
  { action: "ADD_PLAYER", data: { PLAYER_NAME: myName } },
);

document.getElementById("start-game").addEventListener(
  "click",
  () => socket.send({ action: "START_GAME" }),
);
