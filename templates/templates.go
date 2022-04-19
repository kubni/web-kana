package templates

import (
  "html/template"
)


var TmpMain  = template.Must(template.ParseFiles("./templates/layout_main.html"))
var TmpGame = template.Must(template.ParseFiles("./templates/layout_game.html"))
 
 

