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

  // TODO: Is it good to use one big object instead of dealing with multiple setters (and rerenders)
  const [playgroundInfo, setPlaygroundInfo] = useState({
    "userAnswer": "",
    "isAnswerCorrect": true,
    "correctAnswerRomaji": "",
    "currentPlayerScore": 0
  })

  // AbortController for useEffect fetch cleanup
  const fetchHiraganaAbortController = new AbortController(); // TODO: Can 1 AbortController monitor more than 1 fetch request
  const fetchKatakanaAbortController = new AbortController();
  useEffect(() => {
    if(chosenAlphabet === "hiragana")
    {
      fetch
      (
        "http://localhost:8000/game/generateHiraganaCharacter",
        {
          signal: fetchHiraganaAbortController.signal
        }  
      ).then(res => res.json())
       .then(jsonData => setTargetCharacter(jsonData)) 
    }
    else 
    {
      fetch
      (
        "http://localhost:8000/game/generateKatakanaCharacter",
        {
          signal: fetchKatakanaAbortController.signal
        }
      ).then(res => res.json())
       .then(jsonData => setTargetCharacter(jsonData))
    }

    // Cleanup function
    return () => {
      fetchHiraganaAbortController.abort()
      fetchKatakanaAbortController.abort()
    }

  }, [playgroundInfo.userAnswer]) 
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
        "correctAnswerCharacter": targetCharacter
      })
    }).then(res => res.json()) 
      .then
      (
        jsonData => jsonData.isAnswerCorrect === true 
          ? setPlaygroundInfo({
            "userAnswer": userAnswer,
            "isAnswerCorrect": true,
            "correctAnswerRomaji": jsonData.correctAnswerRomaji,
            "currentPlayerScore": playgroundInfo.currentPlayerScore + 1
          }) 
          : setPlaygroundInfo({
            "userAnswer": userAnswer,
            "isAnswerCorrect": false,
            "correctAnswerRomaji": jsonData.correctAnswerRomaji,
            "currentPlayerScore": playgroundInfo.currentPlayerScore > 0 
                                    ? playgroundInfo.currentPlayerScore - 1 
                                    : playgroundInfo.currentPlayerScore
          })
      )
  }
  console.log("Is answer correct: ", playgroundInfo.isAnswerCorrect)

  return (
    <div className="playground">
      <div className="target-char">
         <h1>{targetCharacter}</h1>
      </div>

      <form onSubmit={onAnswerFormSubmit}>
        <input className="answer-input" type="text" ref={ref} autoFocus />
      </form>

      <div className="result-info">
        <h2>Current score: <span className="total-score">{playgroundInfo.currentPlayerScore}</span></h2>
        { !playgroundInfo.isAnswerCorrect && 
          <h3 className="result-message">
            Wrong, the correct answer was <span className="correct-answer">{playgroundInfo.correctAnswerRomaji}</span>
          </h3>
        }
      </div>
        
      <form>
      </form>

    </div>
  )
}

