export default function GamePage(props) {
  /*  New      ==    Old     class names 
    * gamePage == container 
  */
  return (
      <div className="gamePage">  
       <header>
        <h1>{props.pageTitle}</h1>
       </header>
      </div>
  )
}
