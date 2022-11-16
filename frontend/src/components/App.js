import { BrowserRouter, Routes, Route} from "react-router-dom";

import MainPage from "./MainPage"
import GamePage from "./GamePage"

export default function App() {

  // FIXME: This should dinamically change depending on the button that was clicked on main page.
  const pageTitle = "hiragana"



  // TODO: path="/" vs index?
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<MainPage />} />
        <Route path="/game" element={
          <GamePage 
           pageTitle={pageTitle}
          />}
        />
      </Routes>
    </BrowserRouter>
  )
}
