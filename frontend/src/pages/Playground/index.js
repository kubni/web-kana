// TODO: For some reason there is an Uncaught promise DOMException on this page, even though everything seems to work.
//
//
// FIXME: If there is no input and only Enter is pressed, the correct answer message will appear like it should, but character won't change

import "./playground.css";
import React, { useState, useEffect, useRef } from "react";

export default function Playground(props) {
  // console.log(props)
  const chosenAlphabet = props.chosenAlphabet;

  // Refs:
  const ref = useRef();

  // States:
  const [targetCharacter, setTargetCharacter] = useState("");

  const [playgroundInfo, setPlaygroundInfo] = useState({
    userAnswer: "",
    isAnswerCorrect: true,
    correctAnswerRomaji: "",
    displayCorrectAnswerRomaji: false,
    wrongAnswerMessage: "",
  });

  // AbortController for useEffect fetch cleanup
  useEffect(() => {
    const fetchHiraganaAbortController = new AbortController(); // TODO: Can 1 AbortController monitor more than 1 fetch request
    const fetchKatakanaAbortController = new AbortController();

    if (chosenAlphabet === "hiragana") {
      fetch("http://localhost:8000/game/generateHiraganaCharacter", {
        signal: fetchHiraganaAbortController.signal,
      })
        .then((res) => res.json())
        .then((jsonData) => setTargetCharacter(jsonData));
    } else {
      fetch("http://localhost:8000/game/generateKatakanaCharacter", {
        signal: fetchKatakanaAbortController.signal,
      })
        .then((res) => res.json())
        .then((jsonData) => setTargetCharacter(jsonData));
    }

    // Cleanup function
    return () => {
      fetchHiraganaAbortController.abort();
      fetchKatakanaAbortController.abort();
    };
  }, [playgroundInfo.userAnswer, props.currentPlayerScore]);
  console.log("Target Character: ", targetCharacter);

  function onAnswerFormSubmit(event) {
    event.preventDefault();

    // Answer validation: the answer can't be ""
    // FIXME: This doesn't work as intended
    // if (ref.current.value === "") {
    //   setPlaygroundInfo({
    //     ...playgroundInfo,
    //     wrongAnswerMessage: "Error: You must type something!",
    //   });
    // }

    const userAnswer = ref.current.value;
    ref.current.value = "";

    fetch("http://localhost:8000/game/checkAnswer", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        userAnswer: userAnswer,
        correctAnswerCharacter: targetCharacter,
      }),
    })
      .then((res) => res.json())
      .then((jsonData) => {
        if (jsonData.isAnswerCorrect === true) {
          /*
           * Important: Two setStates happen here, but React is using batching probably so there are no rerenders
           * between them
           */
          setPlaygroundInfo({
            ...playgroundInfo,
            userAnswer: userAnswer,
            isAnswerCorrect: true,
            correctAnswerRomaji: jsonData.correctAnswerRomaji,
           });
          props.changeCurrentPlayerScoreBy(1);
        } else {
          setPlaygroundInfo({
            userAnswer: userAnswer,
            isAnswerCorrect: false,
            correctAnswerRomaji: jsonData.correctAnswerRomaji,
            displayCorrectAnswerRomaji: userAnswer === "" ? false : true,
            wrongAnswerMessage: userAnswer === "" ? "Error: You must type something!" : "Wrong, the correct answer was "
          });
          props.changeCurrentPlayerScoreBy(-1);
        }
      });
  }
  console.log("Is answer correct: ", playgroundInfo.isAnswerCorrect);

  function onFinishGameFormSubmit(event) {
    event.preventDefault();
    props.finishGame();
  }

  return (
    <div className="playground">
      <div className="target-char">
        <h1>{targetCharacter}</h1>
      </div>

      <form onSubmit={onAnswerFormSubmit}>
        <input className="answer-input" type="text" ref={ref} autoFocus />
      </form>

      <div className="result-info">
        <h2>
          Current score:{" "}
          <span className="total-score">{props.currentPlayerScore}</span>
        </h2>
        {!playgroundInfo.isAnswerCorrect && (
          <h3 className="result-message">
            {playgroundInfo.wrongAnswerMessage}
            {playgroundInfo.displayCorrectAnswerRomaji && (
              <span className="correct-answer">
                {playgroundInfo.correctAnswerRomaji}
              </span>
            )}
          </h3>
        )}
      </div>

      <form onSubmit={onFinishGameFormSubmit}>
        <button className="finish-button">Finish</button>
      </form>
    </div>
  );
}
