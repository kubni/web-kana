import "./scoreboard.css"
// ClassName changes:
// player-rank -> playerRank
// active-row -> activeRow 
export default function Scoreboard(props) {
  console.log(props)
  return (
    <div className="scoreboard">
      <h1>Scoreboard</h1>
      <h2>Your rank is: <span className="player-rank">props.currentPlayerRank</span></h2>
      <table className="scoreboard-table">
        <thead>
          <tr>
            <th>Rank</th>
            <th>Username</th>
            <th>Score</th>
          </tr>
        </thead>
        <tbody>
          {
            props.scoreboard.map(player => {
              return (
                <tr key={player.ID} className={player.ID === props.currentPlayerStringID ? "activeRow" : ""}>
                  <td>{player.rank}</td>
                  <td>{player.username}</td>
                  <td>{player.score}</td>
                </tr>
              )
            })
          }
        </tbody>
      </table>

      <div className="pagination-buttons">
      { 
        props.currentPage !== 0 && 
          <form>    
          </form>
      }

      {
        (props.currentPage+1 < props.numberOfPages) &&
          <form>
          </form>
      }
      </div>

      <h3>Page {props.currentPage + 1}</h3>
    </div>
  )
}
