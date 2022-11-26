import React, {useState} from "react"

import Playground from "../../components/Playground.js"
import FinishPage from "../../components/FinishPage.js"
import Scoreboard from "../../components/Scoreboard.js"

export default function GamePage(props) {
  /*  Class names changes 
    * gamePage -> container 
  */

  // We create this here because all of the children will need it

  // TODO: Should i implement this here or just send a request to golang. We could have something
  // like a getter for everything in golang for everything we need here?
  const [currentPlayerScore, setCurrentPlayerScore] = useState(0)  


  // const currentScoreFromGolang = getScoreFromBackend()
  function incrementPlayerScore() {
    // setCurrentPlayerScore(currentScoreFromGolang + 1)
    setCurrentPlayerScore(currentPlayerScore + 1)
  }

  return (
      // TODO: Add the {} around props placeholders when they get properly sent here
      // Sending the data for props from backend to frontend?
        // Probably as a response to the GET request, but again the problem is that our controller doesn't recognize the request 

      <div className="gamePage">  
       <header>
        <h1>{props.pageTitle}</h1>
       </header>

      {/* For now I have recreated it as it was in Golang. There is a high chance it can be made simpler in React */}
      {  !props.isFinished && <Playground />}     {/* if(!isFinished)*/} 
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
            <p color="red">The username you entered already exists! Please, choose another one.</p>
            <FinishPage />
          </>
      } 


      </div>
  )
}
