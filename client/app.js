import { openSocket } from "./io.js";

const socket = openSocket(`ws://${location.host}/ws`);

socket.subscribe((message) => console.log(message));
socket.send({ type: "PING", data: { lel: 5 } });
