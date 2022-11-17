// ClassName changes:
// total-score -> totalScore 
//

// Forms with "hidden" hack are probably not needed in react and can be done in a better way


export default function FinishPage() {
  return (
    <div className="finishPage">
      <h2>The game is finished. Your total score is: <span className="totalScore">props.currentPlayerScore</span></h2>

      <form>
      </form>

      <h2>If you want to save your result, type in your username: </h2>

      <form>
      </form>
    </div>
  )
}
