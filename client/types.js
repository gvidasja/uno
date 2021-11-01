/**
 * @typedef {{
 *    color: string
 *    number: string
 *  }} UnoCard
 * 
 * @typedef {{
 *    id: string
 *    turn: boolean
 *    handSize: number
 *    winner?: boolean,
 *  }} UnoPlayer
 * 
 * @typedef {{
 *    hand?: UnoCard[]
 *    me: UnoPlayer,
 *    globalState: {
 *      state: 'PREPARATION' | 'PLAYING' | 'FINISHED',
 *      players: UnoPlayer[],
 *      pileSize: number,
 *      deckSize: number,
 *      topCard?: UnoCard,
 *      colorOverride?: string,
 *      errors?: string[]
 *  }} UnoState
 * */
