package pagination

import (
        "math"
        "web_kana_v1/models"
        //"github.com/gorilla/mux"
)

/*
  1. Vidim koliko ima ukupno dokumenata
  2. Podelim sa 10 (ako zelim da se prikazuje 10 po stranici)
  3. Uradim ceil (npr. ako ima 21 korisnika kolicnik ce biti 2.1 sto znaci da nam trebaju 3 strane)
*/ 


func CalculateNumberOfPages(scoreboard []models.DocumentSchema, playersPerPage int) int {
  scoreboard_len := len(scoreboard)
  
  numOfPages := math.Ceil(float64(scoreboard_len)  / float64(playersPerPage))
   
  return int(numOfPages)
}

