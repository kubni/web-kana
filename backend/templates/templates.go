package templates

import (
	"html/template"
)

// Our custom functions to be used in templates
var funcMap = template.FuncMap{
	// The name "inc" is what the function will be called in the template
	"inc": func(i int) int {
		return (i + 1)
	},
	"dec": func(i int) int {
		return (i + 1)
	},
	"assign": func(x int, y int) int {
		x = y
		return x
	},
}

/*
From the docs:

	Since the templates created by ParseFiles are named by the base names of the argument files,
	t should usually have the name of one of the (base) names of the files. If it does not, depending on t's contents before calling ParseFiles,
	t.Execute may fail.

	This is why we pass "layout_game.html" to our template.New() func
*/


// TODO: Potential race condition if Clone isn't used?
// https://www.calhoun.io/intro-to-templates-p3-functions/

var TmpMain = template.Must(template.ParseFiles("./templates/layout_main.html"))
var TmpGame = template.Must(template.New("layout_game.html").Funcs(funcMap).ParseFiles("./templates/layout_game.html"))
