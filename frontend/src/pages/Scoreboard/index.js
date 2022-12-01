import "./scoreboard.css";

import { useState, useEffect } from "react";

// ClassName changes:
// player-rank -> playerRank
// active-row -> activeRow
export default function Scoreboard(props) {
  // console.log("Scoreboard props: ", props);

  const [currentPlayerRank, setCurrentPlayerRank] = useState(0);

  // TODO: Maybe pagination, setPagination as an object that can keep more things like numOfPages
  const [paginationData, setPaginationData] = useState({
    scoreboard: [],
    currentPage: 0,
    numOfPages: 1,
  }); // TODO: Remove currentPage from props in Game

  // TODO: Add AbortController and a cleanup function to both useEffects !!!
  useEffect(() => {
    const userData = {
      currentPlayerStringID: props.currentPlayerStringID,
      currentPlayerScore: props.currentPlayerScore,
    };

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


  // Here we write a useEffect that will POST the scoreboard for the current page 
  useEffect(() => {
    
    fetch("http://localhost:8000/game/getScoreboard", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({currentPage: paginationData.currentPage})
    })
      .then(res => res.json())
      .then(jsonData => setPaginationData({
        ...paginationData, 
        scoreboard: jsonData.scoreboard,
        numOfPages: jsonData.numOfPages
      }))
  }, [paginationData.currentPage]) 


  // TODO: onPageChange  -> currentPage + 1

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
          {paginationData.scoreboard.map((player) => {
            // console.log("Player: ", player)
            return (
              <tr
                key={player.ID}
                className={
                  player.ID === props.currentPlayerStringID ? "active-row" : ""
                }
              >
                <td>{player.Rank}</td>
                <td>{player.Username}</td>
                <td>{player.Score}</td>
              </tr>
            );
          })}
        </tbody>
      </table>

      <div className="pagination-buttons">
        {props.currentPage !== 0 && <form></form>}

        {props.currentPage + 1 < props.numberOfPages && <form></form>}
      </div>

      <h3>Page {paginationData.currentPage + 1}</h3>
    </div>
  );
}
