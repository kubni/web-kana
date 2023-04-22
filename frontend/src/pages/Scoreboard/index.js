/* FIXME:
 * 1) 10th ranked player is shown on the 2nd page again, when it shouldn't be so.
 * 2) After switching first and pages back and forth, SOMETIMES (??) rankings under the currrent one get messed up. // IMPORTANT
 *    // It might be a backend problem?
 *    // How it works:
 *        1) The fetch request to /game/calculatePlayerRank calls the controller.CalculatePlayerRank from the backend
 *           // Maybe the rerender happens and calls it multiple times, because we have props.currentPlayerScore in dep. array,
 *              and CalculatePlayerScore in backend is changing
 *
 *
 * */



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
    const calculatePlayerRankAbortController = new AbortController();

    const userData = {
      currentPlayerStringID: props.currentPlayerStringID,
      currentPlayerScore: props.currentPlayerScore,
    };



    console.log("Fetching to /game/calculatePlayerRank...");

    fetch("http://localhost:8000/game/calculatePlayerRank", {
      signal: calculatePlayerRankAbortController.signal,
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ ...userData }),
    })
      .then((res) => res.json())
      .then((jsonData) => setCurrentPlayerRank(jsonData.currentPlayerRank));

    // Cleanup
    return () => {
      calculatePlayerRankAbortController.abort();
    };
  }, [props.currentPlayerStringID, props.currentPlayerScore]);

  useEffect(() => {
    const getScoreboardAbortController = new AbortController();

    fetch("http://localhost:8000/game/getScoreboard", {
      signal: getScoreboardAbortController.signal,
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ currentPage: paginationData.currentPage }),
    })
      .then((res) => res.json())
      .then((jsonData) =>
        setPaginationData(() => {
          return {
            ...paginationData,
            scoreboard: jsonData.scoreboard,
            numOfPages: jsonData.numOfPages,
          };
        })
      );

    // Cleanup
    return () => {
      getScoreboardAbortController.abort();
    };
  }, [paginationData.currentPage, currentPlayerRank]);

  function handlePreviousPageButtonClick() {
    setPaginationData({
      ...paginationData,
      currentPage: paginationData.currentPage - 1,
    });
  }

  function handleNextPageButtonClick() {
    setPaginationData({
      ...paginationData,
      currentPage: paginationData.currentPage + 1,
    });
  }

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
        {paginationData.currentPage !== 0 && (
          <button
            className="previous-page-button"
            onClick={handlePreviousPageButtonClick}
          >
            Previous Page
          </button>
        )}

        {paginationData.currentPage + 1 < paginationData.numOfPages && (
          <button
            className="next-page-button"
            onClick={handleNextPageButtonClick}
          >
            Next Page
          </button>
        )}
      </div>

      <h3>Page {paginationData.currentPage + 1}</h3>
    </div>
  );
}
