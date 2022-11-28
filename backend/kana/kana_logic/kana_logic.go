package kana_logic

import (
  // "fmt"
	"math/rand"
	"time"
	"web_kana_v1/kana/tables"
)

func get_keys(mymap map[string][]string) ([]string, int) {
	keys := make([]string, len(mymap))

	i := 0
	for key := range mymap {
		keys[i] = key
		i++
	}

	return keys, i
}

func Check_answer(answer string, target string) (bool, string) {
  // fmt.Println("Inside check_answer: answer: ", answer, ", target: ", tables.Romaji_table[target])
  correctAnswerRomaji := tables.Romaji_table[target]
	if answer != correctAnswerRomaji {
		return false, correctAnswerRomaji
	} else {
		return true, correctAnswerRomaji
	}
}

func Play_all_gamemode(table map[string][]string) string {
	// We need the map keys in order to begin the process of random selection
	keys, keys_length := get_keys(table)
	rand.Seed(time.Now().Unix())

	// We randomly select a key from the map. This corresponds to randomly selecting a table column.
	target_column := keys[rand.Intn(keys_length)]

	// Now we need to randomly select a character from that previously selected column
	column_length := len(table[target_column])
	target_character := table[target_column][rand.Intn(column_length)]

	return target_character
}
