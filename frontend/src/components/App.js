import { BrowserRouter, Routes, Route} from "react-router-dom";


import MainPage from "./MainPage"
import Game from "./Game"

export default function App() {
  // TODO: path="/" vs index?
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<MainPage />} />
        <Route path="/game" element={<Game />} />
      </Routes>
    </BrowserRouter>
  )
}
