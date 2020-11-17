import { Player } from "./Player.js";
import { Deck } from "./Deck.js";
import { Card } from "./Card.js";
import { GameObject } from "./GameObject.js";

export class GameScene {
  /** @type {UnoState} */
  state = [];

  /**
   * @param {UnoState} newState
   * @param {CanvasRenderingContext2D} g
   * */
  update(newState) {
    if (!newState) return;

    this.state = newState;
  }

  click(x, y) {
    for (let i = this.objects.length - 1; i >= 0; i--) {
      if (this.objects[i].intersects(x, y)) {
        this.objects[i].click();
        break;
      }
    }
  }

  draw(g, onAction) {
    requestAnimationFrame(() => {
      const state = this.state;

      if (!state.globalState) {
        return;
      }

      const players = createPlayers(
        state.globalState.players,
        state.me,
        0,
        0,
        g.canvas.width / 2,
        g.canvas.height,
      );

      const cardZoneX = g.canvas.width / 2 + 10;
      const cardZoneY = 10;
      const cardZoneWidth = g.canvas.width - cardZoneX - 10;
      const cardPerRow = Math.floor(cardZoneWidth / Card.SIZE_X);

      const hand = state.hand
        ? state.hand.map((x, i) => {
          const col = i % cardPerRow;
          const row = Math.floor(i / cardPerRow);

          const card = new Card({
            x: cardZoneX + col * Card.SIZE_X,
            y: cardZoneY + row * Card.SIZE_Y,
            color: x.color,
            number: x.number,
          });

          state.me.turn && card.addClickHandler((card) =>
            onAction(
              "PLAY_CARD",
              {
                CARD_COLOR: card.color,
                CARD_NUMBER: card.number,
                COLOR_OVERRIDE: card.shouldChooseOverride()
                  ? prompt(
                    "Enter 'RED', 'GREEN', 'YELLOW', 'BLUE'",
                  )
                  : null,
              },
            )
          );

          return card;
        })
        : [];

      const middleX = g.canvas.width / 4;
      const middleY = g.canvas.height / 2;

      const deck = new Deck(
        {
          deckSize: state.globalState.deckSize,
          x: middleX - Card.SIZE_X - 10,
          y: middleY - Card.SIZE_Y / 2,
          topCard: new Card({ color: "UNO", number: "UNO" }),
        },
      );

      deck.addClickHandler(() => onAction("DRAW_CARD"));

      const pile = state.globalState.pileSize
        ? [
          new Deck({
            deckSize: state.globalState.pileSize,
            x: middleX + 10,
            y: middleY - Card.SIZE_Y / 2,
            topCard: new Card(
              {
                color: state.globalState.topCard.color,
                number: state.globalState.topCard.number,
              },
            ),
          }),
        ]
        : [];

      /** @type {GameObject[]} */
      this.objects = [...players, deck, ...pile, ...hand].sort((o1, o2) =>
        (o1.y - o2.y) || (o1.x - o2.x)
      );

      g.fillStyle = "lightgrey";
      g.fillRect(0, 0, g.canvas.width, g.canvas.height);
      this.objects.forEach((x) => x.draw(g));
    });
  }
}

function createPlayers(statePlayers, me, areaX, areaY, areaSizeX, areaSizeY) {
  if (statePlayers && statePlayers.length > 0) {
    const middleX = areaX + areaSizeX / 2;
    const middleY = areaY + areaSizeY / 2;

    const playerEveryRad = Math.PI * 2 / statePlayers.length;

    while (statePlayers[0].name !== me.name) {
      statePlayers.push(statePlayers.shift());
    }

    const players = statePlayers.map((x, i) =>
      new Player({
        x: middleX +
          Math.cos(Math.PI / 2 + i * playerEveryRad) * middleX * 0.8 -
          Player.SIZE_X / 2,
        y: middleY +
          Math.sin(Math.PI / 2 + i * playerEveryRad) * middleY * 0.8 -
          Player.SIZE_Y / 2,
        handSize: x.handSize,
        name: x.name,
        turn: x.turn,
      })
    );
    return players;
  } else {
    return [];
  }
}
