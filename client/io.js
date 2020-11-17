export function openSocket(url, auth) {
  const ws = new WebSocket(url);

  const openWs = new Promise((resolve, reject) => {
    ws.addEventListener("open", () => resolve());
    ws.addEventListener("error", (e) => reject(e));
  });

  return {
    async send(message) {
      await openWs;
      ws.send(JSON.stringify({ ...message, auth }));
    },

    subscribe(messageHandler) {
      ws.addEventListener("message", (target) => {
        messageHandler(JSON.parse(target.data));
      });
    },
  };
}
