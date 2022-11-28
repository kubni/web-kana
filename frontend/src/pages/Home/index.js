import "./home.css"
import {Link} from "react-router-dom"

export default function MainPage() {
  return (
    <div className="home-page">
        <div className="home-container">
          <header>
            <h1>Welcome! Choose which kana alphabet you want to practice: </h1>
          </header>
        <div className="home-buttons">
          <Link to="/game/hiragana" className="hiragana-button">Hiragana</Link>
          <Link to="/game/katakana" className="katakana-button">Katakana</Link>
        </div>
      </div>
    </div>
  )
}


