import { BrowserRouter, Routes, Route } from "react-router-dom";

import Home from "./pages/Home";
import Game from "./pages/Game";

export default function App() {
  /* TODO:
   * path="/" vs index?
   * GamePage should have states passed as props  (maybe an object of objects (corresponding to the component))!!
   */

  const gameInfo = {
    chosenAlphabet: "",
    pageTitle: "",
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
              chosenAlphabet="hiragana"
              pageTitle="ひらがな"
            />
          }
        />
        <Route
          path="/game/katakana"
          element={
            <Game
              {...gameInfo}
              currentAlphabet="katakana"
              pageTitle="カタカナ"
            />
          }
        />
      </Routes>
    </BrowserRouter>
  );
}
