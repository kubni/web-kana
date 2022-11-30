import "./scoreboard.css";

import { useState, useEffect } from "react";

// ClassName changes:
// player-rank -> playerRank
// active-row -> activeRow
export default function Scoreboard(props) {
  console.log("Scoreboard props: ", props);

  // TODO: State?
  const [currentPlayerRank, setCurrentPlayerRank] = useState(0);

  // TODO: Add AbortController and a cleanup function
  useEffect(() => {
    // TODO: For some reason, there is a React warning if this is outside of useEffect even though userData wouldn't change
    const userData = {
      currentPlayerStringID: props.currentPlayerStringID,
      currentPlayerScore: props.currentPlayerScore,
    };

    console.log("userData: ", userData);

    fetch("http://localhost:8000/game/calculatePlayerRank", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ ...userData }),
    })
      .then((res) => res.json())
      .then((jsonData) => setCurrentPlayerRank(jsonData.currentPlayerRank));
  }, [props.currentPlayerStringID, props.currentPlayerScore]);

  return (
    <div className="scoreboard-page">
      <h1>Scoreboard</h1>
      <h2>
        Your rank is: <span className="player-rank">{currentPlayerRank}</span>
      </h2>
      <table className="scoreboard-table">
        <thead>
          <tr>
            <th>Rank</th>
            <th>Username</th>
            <th>Score</th>
          </tr>
        </thead>
        <tbody>
          {props.scoreboard.map((player) => {
            return (
              <tr
                key={player.ID}
                className={
                  player.ID === props.currentPlayerStringID ? "activeRow" : ""
                }
              >
                <td>{player.rank}</td>
                <td>{player.username}</td>
                <td>{player.score}</td>
              </tr>
            );
          })}
        </tbody>
      </table>

      <div className="pagination-buttons">
        {props.currentPage !== 0 && <form></form>}

        {props.currentPage + 1 < props.numberOfPages && <form></form>}
      </div>

      <h3>Page {props.currentPage + 1}</h3>
    </div>
  );
}
