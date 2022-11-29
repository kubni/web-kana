// TODO: Add Context here to replace passing down state to children


import "./game.css";
import React, { useState } from "react";

import Playground from "../../pages/Playground";
import FinishPage from "../../pages/Finish";
import Scoreboard from "../../pages/Scoreboard";

export default function GamePage(props) {
  
  const [currentPlayerScore, setCurrentPlayerScore] = useState(0)
  const [isFinished, setIsFinished] = useState(false) 

  return (
    <div className="game-page">
      <header>
        <h1 className="game-header">{props.pageTitle}</h1>
      </header>

      {!isFinished && (
        <Playground
          chosenAlphabet = {props.chosenAlphabet}
          currentPlayerScore = {currentPlayerScore}
          changeCurrentPlayerScoreBy = 
          { 
            (changeScoreBy) => {
              setCurrentPlayerScore( (oldPlayerScore) => {
                const newScore = oldPlayerScore + changeScoreBy 
                return newScore > 0 ? newScore : 0
              })
            }
          }
          isFinished = {isFinished}
          finishGame = {() => setIsFinished(true)}
        />
      )}
      {isFinished && !props.isDisplayScoreboard && <FinishPage />}
      {isFinished && props.isDisplayScoreboard && props.isUsernameValid && (
        <Scoreboard
          scoreboard={[
            { ID: "1", rank: 1, username: "test", score: "4" },
            { ID: "2", rank: 2, username: "test2", score: "2" },
          ]}
          currentPlayerStringID="1"
          currentPage={0}
        />
      )}
      {isFinished && props.isDisplayScoreboard && !props.isUsernameValid && (
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
