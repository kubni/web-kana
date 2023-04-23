package tables

var Hiragana_table = map[string][]string{
	"vowels_column":  {"あ", "い", "う", "え", "お"},
	"k_column":       {"か", "き", "く", "け", "こ"},
	"s_column":       {"さ", "し", "す", "せ", "そ"},
	"t_column":       {"た", "ち", "つ", "て", "と"},
	"n_column":       {"な", "に", "ぬ", "ね", "の"},
	"h_column":       {"は", "ひ", "ふ", "へ", "ほ"},
	"m_column":       {"ま", "み", "む", "め", "も"},
	"y_column":       {"や", "ゆ", "よ"},
	"r_column":       {"ら", "り", "る", "れ", "ろ"},
	"w_column":       {"わ", "を"},
	"special_column": {"ん"},
}

var Katakana_table = map[string][]string{
	"vowels_column":  {"ア", "イ", "ウ", "エ", "オ"},
	"k_column":       {"カ", "キ", "ク", "ケ", "コ"},
	"s_column":       {"サ", "シ", "ス", "セ", "ソ"},
	"t_column":       {"タ", "チ", "ツ", "テ", "ト"},
	"n_column":       {"ナ", "ニ", "ヌ", "ネ", "ノ"},
	"h_column":       {"ハ", "ヒ", "フ", "ヘ", "ホ"},
	"m_column":       {"マ", "ミ", "ム", "メ", "モ"},
	"y_column":       {"ヤ", "ユ", "ヨ"},
	"r_column":       {"ラ", "リ", "ル", "レ", "ロ"},
	"w_column":       {"ワ", "ヲ"},
	"special_column": {"ン"},
}

var Romaji_table = map[string]string{
	Hiragana_table["vowels_column"][0]: "a",
	Hiragana_table["vowels_column"][1]: "i",
	Hiragana_table["vowels_column"][2]: "u",
	Hiragana_table["vowels_column"][3]: "e",
	Hiragana_table["vowels_column"][4]: "o",

	Hiragana_table["k_column"][0]: "ka",
	Hiragana_table["k_column"][1]: "ki",
	Hiragana_table["k_column"][2]: "ku",
	Hiragana_table["k_column"][3]: "ke",
	Hiragana_table["k_column"][4]: "ko",

	Hiragana_table["s_column"][0]: "sa",
	Hiragana_table["s_column"][1]: "shi",
	Hiragana_table["s_column"][2]: "su",
	Hiragana_table["s_column"][3]: "se",
	Hiragana_table["s_column"][4]: "so",

	Hiragana_table["t_column"][0]: "ta",
	Hiragana_table["t_column"][1]: "chi",
	Hiragana_table["t_column"][2]: "tsu",
	Hiragana_table["t_column"][3]: "te",
	Hiragana_table["t_column"][4]: "to",

	Hiragana_table["n_column"][0]: "na",
	Hiragana_table["n_column"][1]: "ni",
	Hiragana_table["n_column"][2]: "nu",
	Hiragana_table["n_column"][3]: "ne",
	Hiragana_table["n_column"][4]: "no",

	Hiragana_table["h_column"][0]: "ha",
	Hiragana_table["h_column"][1]: "hi",
	Hiragana_table["h_column"][2]: "fu",
	Hiragana_table["h_column"][3]: "he",
	Hiragana_table["h_column"][4]: "ho",

	Hiragana_table["m_column"][0]: "ma",
	Hiragana_table["m_column"][1]: "mi",
	Hiragana_table["m_column"][2]: "mu",
	Hiragana_table["m_column"][3]: "me",
	Hiragana_table["m_column"][4]: "mo",

	Hiragana_table["y_column"][0]: "ya",
	Hiragana_table["y_column"][1]: "yu",
	Hiragana_table["y_column"][2]: "yo",

	Hiragana_table["r_column"][0]: "ra",
	Hiragana_table["r_column"][1]: "ri",
	Hiragana_table["r_column"][2]: "ru",
	Hiragana_table["r_column"][3]: "re",
	Hiragana_table["r_column"][4]: "ro",

	Hiragana_table["w_column"][0]: "wa",
	Hiragana_table["w_column"][1]: "wo",

	Hiragana_table["special_column"][0]: "n",

	// Split maybe ?

	Katakana_table["vowels_column"][0]: "a",
	Katakana_table["vowels_column"][1]: "i",
	Katakana_table["vowels_column"][2]: "u",
	Katakana_table["vowels_column"][3]: "e",
	Katakana_table["vowels_column"][4]: "o",

	Katakana_table["k_column"][0]: "ka",
	Katakana_table["k_column"][1]: "ki",
	Katakana_table["k_column"][2]: "ku",
	Katakana_table["k_column"][3]: "ke",
	Katakana_table["k_column"][4]: "ko",

	Katakana_table["s_column"][0]: "sa",
	Katakana_table["s_column"][1]: "shi",
	Katakana_table["s_column"][2]: "su",
	Katakana_table["s_column"][3]: "se",
	Katakana_table["s_column"][4]: "so",

	Katakana_table["t_column"][0]: "ta",
	Katakana_table["t_column"][1]: "chi",
	Katakana_table["t_column"][2]: "tsu",
	Katakana_table["t_column"][3]: "te",
	Katakana_table["t_column"][4]: "to",

	Katakana_table["n_column"][0]: "na",
	Katakana_table["n_column"][1]: "ni",
	Katakana_table["n_column"][2]: "nu",
	Katakana_table["n_column"][3]: "ne",
	Katakana_table["n_column"][4]: "no",

	Katakana_table["h_column"][0]: "ha",
	Katakana_table["h_column"][1]: "hi",
	Katakana_table["h_column"][2]: "fu",
	Katakana_table["h_column"][3]: "he",
	Katakana_table["h_column"][4]: "ho",

	Katakana_table["m_column"][0]: "ma",
	Katakana_table["m_column"][1]: "mi",
	Katakana_table["m_column"][2]: "mu",
	Katakana_table["m_column"][3]: "me",
	Katakana_table["m_column"][4]: "mo",

	Katakana_table["y_column"][0]: "ya",
	Katakana_table["y_column"][1]: "yu",
	Katakana_table["y_column"][2]: "yo",

	Katakana_table["r_column"][0]: "ra",
	Katakana_table["r_column"][1]: "ri",
	Katakana_table["r_column"][2]: "ru",
	Katakana_table["r_column"][3]: "re",
	Katakana_table["r_column"][4]: "ro",

	Katakana_table["w_column"][0]: "wa",
	Katakana_table["w_column"][1]: "wo",

	Katakana_table["special_column"][0]: "n",
}
