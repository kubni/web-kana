package hiragana

import (
  "fmt"
  "math/rand"
  "time"
)
var hiragana_table = map[string][]string {
    "vowels_column":    {"あ", "い", "う", "え", "お"},
    "k_column":         {"か", "き", "く", "け", "こ"},
    "s_column":         {"さ", "し", "す", "せ", "そ"},
    "t_column":         {"た", "ち", "つ", "て", "と"},
    "n_column":         {"な", "に", "ぬ", "ね", "の"},
    "h_column":         {"は", "ひ", "ふ", "へ", "ほ"},
    "m_column":         {"ま", "み", "む", "め", "も"},
    "y_column":         {"や", "ゆ", "よ"},
    "r_column":         {"ら", "り", "る", "れ", "ろ"},
    "w_column":         {"わ", "を"},
    "special_column":   {"ん"},
}

var romaji_table = map[string]string {

 hiragana_table["vowels_column"][0]: "a",
 hiragana_table["vowels_column"][1]: "i",
 hiragana_table["vowels_column"][2]: "u",
 hiragana_table["vowels_column"][3]: "e",
 hiragana_table["vowels_column"][4]: "o",
                                          
 hiragana_table["k_column"][0]: "ka",
 hiragana_table["k_column"][1]: "ki",
 hiragana_table["k_column"][2]: "ku",
 hiragana_table["k_column"][3]: "ke",
 hiragana_table["k_column"][4]: "ko",
                                   
 hiragana_table["s_column"][0]: "sa",
 hiragana_table["s_column"][1]: "si",
 hiragana_table["s_column"][2]: "su",
 hiragana_table["s_column"][3]: "se",
 hiragana_table["s_column"][4]: "so",
                                     
 hiragana_table["t_column"][0]: "ta",
 hiragana_table["t_column"][1]: "chi",
 hiragana_table["t_column"][2]: "tsu",
 hiragana_table["t_column"][3]: "te",
 hiragana_table["t_column"][4]: "to",
                                   
 hiragana_table["n_column"][0]: "na",
 hiragana_table["n_column"][1]: "ni",
 hiragana_table["n_column"][2]: "nu",
 hiragana_table["n_column"][3]: "ne",
 hiragana_table["n_column"][4]: "no",
                                   
 hiragana_table["h_column"][0]: "ha",
 hiragana_table["h_column"][1]: "hi",
 hiragana_table["h_column"][2]: "fu",
 hiragana_table["h_column"][3]: "he",
 hiragana_table["h_column"][4]: "ho",
                                 
 hiragana_table["m_column"][0]: "ma",
 hiragana_table["m_column"][1]: "mi",
 hiragana_table["m_column"][2]: "mu",
 hiragana_table["m_column"][3]: "me",
 hiragana_table["m_column"][4]: "mo",
                                  
 hiragana_table["y_column"][0]: "ya",
 hiragana_table["y_column"][1]: "yu",
 hiragana_table["y_column"][2]: "yo",
                                  
 hiragana_table["r_column"][0]: "ra",
 hiragana_table["r_column"][1]: "ri",
 hiragana_table["r_column"][2]: "ru",
 hiragana_table["r_column"][3]: "re",
 hiragana_table["r_column"][4]: "ro",
                                              
 hiragana_table["w_column"][0]: "wa",
 hiragana_table["w_column"][1]: "wo",

 hiragana_table["special_column"][0]: "n",
}

func get_keys(mymap map[string][]string) ([]string, int) {
  keys := make([]string, len(mymap))

  i := 0
  for key := range mymap {
    keys[i] = key 
    i++
  }

  return keys, i 
}


func Check_answer(answer string, target string) bool {
  fmt.Println("Answer: ", answer, ", target: ", target)
  if answer != romaji_table[target] {
    return false
  } else { 
    return true 
  }
}


func Play_all_gamemode() string {
  // We need the map keys in order to begin the process of random selection
  keys, keys_length := get_keys(hiragana_table)
  
  rand.Seed(time.Now().Unix())

  // We randomly select a key from the map. This corresponds to randomly selecting a table column. 
  target_column := keys[rand.Intn(keys_length)]
    
  // Now we need to randomly select a character from that previously selected column
  column_length := len(hiragana_table[target_column])

  target_character := hiragana_table[target_column][rand.Intn(column_length)] 

  //answer_and_check(target_character)
  return target_character 
}


