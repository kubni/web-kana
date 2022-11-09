import React, {useState} from "react"

export default function Form() {
  const [formData, setFormData] = useState(
    {
      chosenAlphabet: ""
    }
  )

  console.log(formData)


  function handleSubmit(event) {
    event.preventDefault()
    console.log("Zdravo iz handlesubmita")
    console.log("formData u handleSubmitu", formData)
    // TODO: axios
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
        <button className="hiragana-button" onClick={handleClick}>Hiragana</button>
        <button className="katakana-button" onClick={handleClick}>Katakana</button>
      </div>
    </form>
    </>
  )
}
