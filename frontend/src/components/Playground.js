// ClassName changes:
// target-char -> targetChar
// total-score -> totalScore
// correct-answer -> correctAnswer 


export default function Playground(props) {

  return (
    <div className="playground">
      <div className="targetChar">
         <h1>props.character</h1>
      </div>

      {// Is this form needed since we now have access to onClick functions, props, states, etc. whereas in golang we had to send a request and after page refresh 
       // read the value to know that the button was clicked or some other action happened.
      }
      <form>  
      </form>

      <div className="resultInfo">
        <h2>Current score: <span className="totalScore">props.currentPlayerScore</span></h2>
        <h3 className="resultMessage">props.resultMessage<span className="correctAnswer">props.correctAnswer</span></h3>
      </div>
        
      <form>
      </form>

    </div>
  )
}
