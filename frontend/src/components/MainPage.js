import Form from "./Form.js"
export default function MainPage() {
  // Za sad je prakticno MainPage == App ali nekako mora i Game da se ubaci u App
  // Takodje ne moze samo Form.js da se zove komponenta za formu jer nije svaka forma ista
  return (
    <div className="container">
      <header>
        <h1>Welcome! Choose which kana alphabet you want to practice: </h1>
      </header>
      <Form />
    </div>
  )
}
