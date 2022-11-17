
import Playground from "./Playground.js"
import FinishPage from "./FinishPage.js"
import Scoreboard from "./Scoreboard.js"

export default function GamePage(props) {
  /*  Class names changes 
    * gamePage -> container 
  */

  return (
      // TODO: Add the {} around props placeholders
      // How do i send the character from backend to frontend?
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
