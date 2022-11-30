import "./finish.css";
// ClassName changes:
// total-score -> totalScore
//
import { useRef } from "react";
import { Link } from "react-router-dom";

export default function FinishPage(props) {
  const ref = useRef();

  // document = bson.M{
  // 				"ID":       gc.data.CurrentPlayerStringID,
  // 				"Username": gc.data.CurrentPlayer,
  // 				"Score":    gc.data.CurrentPlayerScore,
  // 				"Rank":     gc.data.CurrentPlayerRank,
  // 			}

  function onFormSubmit(event) {
    event.preventDefault();

    const userData = {
      username: ref.current.value,
      score: props.currentPlayerScore,
    };

    // Make the request for the user to be entered into the database
    fetch("http://localhost:8000/game/insertUserIntoDatabase", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ ...userData }),
    })
      .then((res) => res.json())
      // .then((jsonData) => {
      //   // TODO: Better error handling
      //   if (!jsonData.isInserted)
      //     console.log("Insert error: ", jsonData.error);
      // });
      .then((jsonData) => {
        // TODO: Is better way for error handling needed ? 
        let isValid = jsonData.isInserted 
        console.log(jsonData.error) // If there is no error then jsonData.error === ""
         
        // Call the scoreboard info setter wrapper:
        props.setScoreboardInfo(isValid)
      })
  }

  return (
    <div className="finish-page">
      <h2>
        The game is finished. Your total score is:{" "}
        <span className="total-score">{props.currentPlayerScore}</span>
      </h2>

      <div className="play-again-button">
        <Link to={`/game/${props.chosenAlphabet}`} className="go-back-button">
          Play Again
        </Link>
      </div>

      <h2>If you want to save your result, type in your username: </h2>

      <form onSubmit={onFormSubmit}>
        <input className="username-input" ref={ref} type="text" />
      </form>
    </div>
  );
}
