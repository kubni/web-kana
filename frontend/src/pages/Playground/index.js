import "./playground.css"
import React, {useState, useEffect, useRef} from "react"

// ClassName changes:
// target-char -> targetChar
// total-score -> totalScore
// correct-answer -> correctAnswer 

export default function Playground(props) {

  // console.log(props)
  const chosenAlphabet = props.chosenAlphabet

  // Refs:
  const ref = useRef()
  
  // States:
  const [targetCharacter, setTargetCharacter] = useState("")
  const [answer, setAnswer] = useState({"userAnswer": "", "isAnswerCorrect": false})
  
   useEffect(() => {
    if(chosenAlphabet === "hiragana")
    {
        fetch("http://localhost:8000/game/generateHiraganaCharacter")
          .then(res => res.json())
          .then(jsonData => setTargetCharacter(jsonData)) 
    } else {
      fetch("http://localhost:8000/game/generateKatakanaCharacter")
        .then(res => res.json())
        .then(jsonData => setTargetCharacter(jsonData))
    }

  }, [answer.userAnswer]) 
  console.log("Target Character: ", targetCharacter)


  function onAnswerFormSubmit (event) {
    event.preventDefault()
    const userAnswer = ref.current.value

    fetch("http://localhost:8000/game/checkAnswer", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        "userAnswer": userAnswer,
        "correctAnswer": targetCharacter
      })
    }).then(res => res.json()) 
      .then(jsonData => jsonData === true ? setAnswer({"userAnswer": userAnswer, "isAnswerCorrect": true}) 
                                          : setAnswer({"userAnswer": userAnswer, "isAnswerCorrect": false})
      )
  }
  console.log("Is answer correct: ", answer.isAnswerCorrect)

  return (
    <div className="playground">
      <div className="target-char">
         <h1>{targetCharacter}</h1>
      </div>

      <form onSubmit={onAnswerFormSubmit}>
        <input className="answer-input" type="text" ref={ref} autoFocus />
      </form>

      <div className="result-info">
        <h2>Current score: <span className="total-score">props.currentPlayerScore</span></h2>
        <h3 className="result-message">props.resultMessage<span className="correct-answer">props.correctAnswer</span></h3>
      </div>
        
      <form>
      </form>

    </div>
  )
}

