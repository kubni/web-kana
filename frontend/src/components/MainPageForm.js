import React, {useState} from "react"
import {Link} from "react-router-dom"

export default function Form() {
  const [formData, setFormData] = useState(
    {
      chosenAlphabet: ""
    }
  )

  console.log(formData)

  function handleSubmit(event) {
    event.preventDefault()

    // Send this to /game endpoint so the backend can process the data.
    //   FIXME: GET request should not have the body.
    /* 
      * The problem is that the backend expects  chosenAlphabet to come with a GET request
      * I could send the post request and read the FormValue("chosen-alphabet") there,
      * but the problem is that moving Play_all_gamemode into the else block (which is what gets 
      * triggered on POST request) breaks the game 
    */ 
    fetch("http://localhost:8000/game", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
        // "Content-Type": "application/json",
      },
      body: JSON.stringify(formData) 
    })
  } 


  function handleClick(event) {
    const {className} = event.target
    console.log(event)
    setFormData(() => {
      return {
        chosenAlphabet: className === "hiragana-button" ? "hiragana" : "katakana"
      } 
    })

  }

  return (
    <>
    <form id="testId" onSubmit={handleSubmit}> 
      <div className="buttons">
        <Link to="/game" className="hiragana-button" onClick={handleClick}>Hiragana</Link>
        <Link to="/game" className="katakana-button" onClick={handleClick}>Katakana</Link>
      </div>
    </form>
    </>
  )
}
