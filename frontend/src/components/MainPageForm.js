import React, {useState, useEffect} from "react"
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

    // FIXME: useEffect doesn't work !!!!!!!!!!!!!!!!!!!
    // useEffect(function () {
    //   fetch("http://localhost:8000/game", {
    //     method: "POST",
    //     headers: {
    //       "Content-Type": "application/x-www-form-urlencoded",
    //       // "Content-Type": "application/json",
    //     },
    //     body: JSON.stringify(formData) 
    //   })
    // }, [])
    //
    //
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
