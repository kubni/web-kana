import "./finish.css"
// ClassName changes:
// total-score -> totalScore 
//
import {Link} from "react-router-dom"

export default function FinishPage(props) {

  function onFormSubmit() {

  }

  return (
    <div className="finish-page">
      <h2>The game is finished. Your total score is: <span className="total-score">{props.currentPlayerScore}</span></h2>

      <div className="play-again-button">
        <Link to={`/game/${props.chosenAlphabet}`} className="go-back-button">Play Again</Link>
      </div>

      <h2>If you want to save your result, type in your username: </h2>

      <form onSubmit={onFormSubmit}>
        <input className="username-input" type="text"/>
      </form>
    </div>
  )
}
