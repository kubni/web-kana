import MainPageForm from "./MainPageForm.js"

// CSS
/* TODO:
 * This doesn't actually limit the css only to this page,
 * its just that I don't style anything that encompasses other components
 * in the mainPage.css
 */
import "../stylesheets/mainPage.css"


export default function MainPage() {
  return (
    <div className="mainPage">
      <div className="container">
        <header>
          <h1>Welcome! Choose which kana alphabet you want to practice: </h1>
        </header>
        <MainPageForm />
      </div>
    </div>
  )
}
