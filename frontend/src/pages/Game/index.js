import "./game.css"
import React, {useState} from "react"

import Playground from "../../pages/Playground"
import FinishPage from "../../pages/Finish"
import Scoreboard from "../../pages/Scoreboard"

export default function GamePage(props) {
  /*  Class names changes 
    * gamePage -> container 
  */

  // We create this here because all of the children will need it
  const [currentPlayerScore, setCurrentPlayerScore] = useState(0)  


  // const currentScoreFromGolang = getScoreFromBackend()
  function incrementPlayerScore() {
    // setCurrentPlayerScore(currentScoreFromGolang + 1)
    setCurrentPlayerScore(currentPlayerScore + 1)
  }

  return (
      <div className="game-page">  
       <header>
        <h1 className="game-header">{props.pageTitle}</h1>
       </header>

      {/* For now I have recreated it as it was in Golang. There is a high chance it can be made simpler in React */}
      { !props.isFinished &&
        <Playground
          chosenAlphabet={props.chosenAlphabet}      
        />
      }     {/* if(!isFinished)*/} 
      {  props.isFinished && !props.isDisplayScoreboard && <FinishPage />} {/* else if !isDisplayScoreboard*/}     
      {  props.isFinished && props.isDisplayScoreboard && props.isUsernameValid && 
          <Scoreboard  
            scoreboard={[{ID:"1", rank:1, username:"test", score:"4"}, {ID:"2", rank:2, username:"test2", score:"2"}]}
            currentPlayerStringID="1"
            currentPage={0}
          />
      } {/* else {if .isUsernameValid} */} 
      {  props.isFinished && props.isDisplayScoreboard && !props.isUsernameValid && 
          <>
            <p>The username you entered already exists! Please, choose another one.</p>
            <FinishPage />
          </>
      } 

      </div>
  )
}
