// TODO: Add Context here to replace passing down state to children


import "./game.css";
import React, { useState } from "react";

import Playground from "../../pages/Playground";
import FinishPage from "../../pages/Finish";
import Scoreboard from "../../pages/Scoreboard";

export default function GamePage(props) {
  
  const [currentPlayerScore, setCurrentPlayerScore] = useState(0)
  const [isFinished, setIsFinished] = useState(false) 
  const [scoreboardInfo, setScoreboardInfo] = useState({
    isDisplayScoreboard: false,
    isUsernameValid: false,
    currentPlayerStringID: "",
  })


  // TODO: Do we really need to pass isDisplayScoreboard and isUsernameValid to finishPage or just their setter

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
      {isFinished && !scoreboardInfo.isDisplayScoreboard &&
        <FinishPage 
          currentPlayerScore = {currentPlayerScore}
          chosenAlphabet = {props.chosenAlphabet}
          scoreboardInfo = {{...scoreboardInfo}}
          setScoreboardInfo = {(isValid, stringID) => setScoreboardInfo({
            isDisplayScoreboard: true,
            isUsernameValid: isValid,
            currentPlayerStringID: stringID
          })}

        />
      }
      {isFinished && scoreboardInfo.isDisplayScoreboard && scoreboardInfo.isUsernameValid && (
        <Scoreboard
          currentPlayerScore = {currentPlayerScore}
          currentPlayerStringID = {scoreboardInfo.currentPlayerStringID}
          scoreboard={[
            { ID: "1", rank: 1, username: "test", score: "4" },
            { ID: "2", rank: 2, username: "test2", score: "2" },
          ]}
          currentPage={0}
        />
      )}
      {isFinished && scoreboardInfo.isDisplayScoreboard && !scoreboardInfo.isUsernameValid && (
        <>
          <p className="invalid-username-msg">
            The username you entered already exists! Please, choose another one.
          </p>
          <FinishPage
            currentPlayerScore = {currentPlayerScore}
            chosenAlphabet = {props.chosenAlphabet}
            scoreboardInfo = {{...scoreboardInfo}}
            setScoreboardInfo = {(isValid, stringObjectID) => setScoreboardInfo({
              isDisplayScoreboard: true,
              isUsernameValid: isValid,
              currentPlayerStringObjectID: stringObjectID
            })}
           
          />
        </>
      )}
    </div>
  );
}

// The second rendering of FinishPage when isUsernameValid is false can be done in a better way, probably together with the first one
