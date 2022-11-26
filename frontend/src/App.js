import { BrowserRouter, Routes, Route } from "react-router-dom";

import Home from "./pages/Home";
import Game from "./pages/Game";

export default function App() {
  // FIXME: This should dinamically change depending on the button that was clicked on main page.
  const pageTitle = "hiragana";

  /* TODO:
   * path="/" vs index?
   * GamePage should have states passed as props  (maybe an object of objects (corresponding to the component))!!
   */

  const gameInfo = {
    pageTitle: "",
    isFinished: false,
    isDisplayScoreboard: false,
    isUsernameValid: false,
  }

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route
          path="/game/hiragana"
          element={
            <Game
              {...gameInfo}
              pageTitle="hiragana"
            />
          }
        />
        <Route
          path="/game/katakana"
          element={
            <Game
              {...gameInfo}
              pageTitle="katakana"
            />
          }
        />
      </Routes>
    </BrowserRouter>
  );
}
