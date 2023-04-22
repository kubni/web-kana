import "./finish.css";
// ClassName changes:
// total-score -> totalScore
//
import { useRef } from "react";
import { Link } from "react-router-dom";

export default function FinishPage(props) {
  const ref = useRef();

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
      .then((jsonData) => {
        console.log("jsonData.isInserted", jsonData.isInserted)
        const isValid = jsonData.isInserted;
        const stringID = jsonData.stringID;
        console.log(jsonData.error); // If there is no error then jsonData.error === ""

        // Call the scoreboard info setter wrapper:
        props.setScoreboardInfo(
          isValid,
          stringID,
          userData.username,
          userData.score
        );
      });
  }

  return (
    <div className="finish-page">
      <h2>
        The game is finished. Your total score is:{" "}
        <span className="total-score">{props.currentPlayerScore}</span>
      </h2>

      <div className="play-again-button">
        <Link to="/" className="go-back-button">
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
