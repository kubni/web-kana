import React, {useState, useEffect, useRef} from "react"
import {Link} from "react-router-dom"

export default function Form() {
  const ref = useRef()

  function handleSubmit(event) {
    event.preventDefault()
    console.log("Nesto")
    console.log(event.submitted)
    console.log(event.nativeEvent.submitter.className)
    const className = event.nativeEvent.submitter.className 
    const chosenAlphabet = className === "hiragana-button" ? "hiragana" : "katakana"
       

    fetch("http://localhost:8000/game", {
      method: "POST",
      headers: {
        "Content-Type": "application/json" 
      },
      body: JSON.stringify({chosenAlphabet})
    })

    console.log(ref.current)
    ref.current.click()
  }


  function handleClick(event) {
    event.preventDefault()
      }

  return (
    <>
    <form id="testId" onSubmit={handleSubmit}> 
      <div className="buttons">
        <button className="hiragana-button">Hiragana</button>
        <button className="katakana-button">Katakana</button>
        <Link to="/game" ref={ref} hidden></Link>
      </div>
    </form>
    </>
  )
}
