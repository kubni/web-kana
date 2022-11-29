import "./game.css";
import React, { useState } from "react";

import Playground from "../../pages/Playground";
import FinishPage from "../../pages/Finish";
import Scoreboard from "../../pages/Scoreboard";

export default function GamePage(props) {
  
  const [currentPlayerScore, setCurrentPlayerScore] = useState(0)

  return (
    <div className="game-page">
      <header>
        <h1 className="game-header">{props.pageTitle}</h1>
      </header>

      {!props.isFinished && (
        <Playground
          chosenAlphabet={props.chosenAlphabet}
          setCurrentPlayerScore = 
          { 
            (changeScoreBy) => {
              setCurrentPlayerScore( (oldPlayerScore) => {
                const newScore = oldPlayerScore + changeScoreBy 
                return newScore > 0 ? newScore : 0
              })
            }
          }
        />
      )}
      {props.isFinished && !props.isDisplayScoreboard && <FinishPage />}
      {props.isFinished && props.isDisplayScoreboard && props.isUsernameValid && (
        <Scoreboard
          scoreboard={[
            { ID: "1", rank: 1, username: "test", score: "4" },
            { ID: "2", rank: 2, username: "test2", score: "2" },
          ]}
          currentPlayerStringID="1"
          currentPage={0}
        />
      )}
      {props.isFinished && props.isDisplayScoreboard && !props.isUsernameValid && (
        <>
          <p className="invalid-username-msg">
            The username you entered already exists! Please, choose another one.
          </p>
          <FinishPage />
        </>
      )}
    </div>
  );
}
