package moe

import (
	"encoding/json"
	"fmt"
	"time"
)

var spinnersJSON = []byte(`{
	"dots": {
		"speed": 80,
		"frames": [
			"⠋",
			"⠙",
			"⠹",
			"⠸",
			"⠼",
			"⠴",
			"⠦",
			"⠧",
			"⠇",
			"⠏"
		]
	},
	"dots2": {
		"speed": 80,
		"frames": [
			"⣾",
			"⣽",
			"⣻",
			"⢿",
			"⡿",
			"⣟",
			"⣯",
			"⣷"
		]
	},
	"dots3": {
		"speed": 80,
		"frames": [
			"⠋",
			"⠙",
			"⠚",
			"⠞",
			"⠖",
			"⠦",
			"⠴",
			"⠲",
			"⠳",
			"⠓"
		]
	},
	"dots4": {
		"speed": 80,
		"frames": [
			"⠄",
			"⠆",
			"⠇",
			"⠋",
			"⠙",
			"⠸",
			"⠰",
			"⠠",
			"⠰",
			"⠸",
			"⠙",
			"⠋",
			"⠇",
			"⠆"
		]
	},
	"dots5": {
		"speed": 80,
		"frames": [
			"⠋",
			"⠙",
			"⠚",
			"⠒",
			"⠂",
			"⠂",
			"⠒",
			"⠲",
			"⠴",
			"⠦",
			"⠖",
			"⠒",
			"⠐",
			"⠐",
			"⠒",
			"⠓",
			"⠋"
		]
	},
	"dots6": {
		"speed": 80,
		"frames": [
			"⠁",
			"⠉",
			"⠙",
			"⠚",
			"⠒",
			"⠂",
			"⠂",
			"⠒",
			"⠲",
			"⠴",
			"⠤",
			"⠄",
			"⠄",
			"⠤",
			"⠴",
			"⠲",
			"⠒",
			"⠂",
			"⠂",
			"⠒",
			"⠚",
			"⠙",
			"⠉",
			"⠁"
		]
	},
	"dots7": {
		"speed": 80,
		"frames": [
			"⠈",
			"⠉",
			"⠋",
			"⠓",
			"⠒",
			"⠐",
			"⠐",
			"⠒",
			"⠖",
			"⠦",
			"⠤",
			"⠠",
			"⠠",
			"⠤",
			"⠦",
			"⠖",
			"⠒",
			"⠐",
			"⠐",
			"⠒",
			"⠓",
			"⠋",
			"⠉",
			"⠈"
		]
	},
	"dots8": {
		"speed": 80,
		"frames": [
			"⠁",
			"⠁",
			"⠉",
			"⠙",
			"⠚",
			"⠒",
			"⠂",
			"⠂",
			"⠒",
			"⠲",
			"⠴",
			"⠤",
			"⠄",
			"⠄",
			"⠤",
			"⠠",
			"⠠",
			"⠤",
			"⠦",
			"⠖",
			"⠒",
			"⠐",
			"⠐",
			"⠒",
			"⠓",
			"⠋",
			"⠉",
			"⠈",
			"⠈"
		]
	},
	"dots9": {
		"speed": 80,
		"frames": [
			"⢹",
			"⢺",
			"⢼",
			"⣸",
			"⣇",
			"⡧",
			"⡗",
			"⡏"
		]
	},
	"dots10": {
		"speed": 80,
		"frames": [
			"⢄",
			"⢂",
			"⢁",
			"⡁",
			"⡈",
			"⡐",
			"⡠"
		]
	},
	"dots11": {
		"speed": 100,
		"frames": [
			"⠁",
			"⠂",
			"⠄",
			"⡀",
			"⢀",
			"⠠",
			"⠐",
			"⠈"
		]
	},
	"dots12": {
		"speed": 80,
		"frames": [
			"⢀⠀",
			"⡀⠀",
			"⠄⠀",
			"⢂⠀",
			"⡂⠀",
			"⠅⠀",
			"⢃⠀",
			"⡃⠀",
			"⠍⠀",
			"⢋⠀",
			"⡋⠀",
			"⠍⠁",
			"⢋⠁",
			"⡋⠁",
			"⠍⠉",
			"⠋⠉",
			"⠋⠉",
			"⠉⠙",
			"⠉⠙",
			"⠉⠩",
			"⠈⢙",
			"⠈⡙",
			"⢈⠩",
			"⡀⢙",
			"⠄⡙",
			"⢂⠩",
			"⡂⢘",
			"⠅⡘",
			"⢃⠨",
			"⡃⢐",
			"⠍⡐",
			"⢋⠠",
			"⡋⢀",
			"⠍⡁",
			"⢋⠁",
			"⡋⠁",
			"⠍⠉",
			"⠋⠉",
			"⠋⠉",
			"⠉⠙",
			"⠉⠙",
			"⠉⠩",
			"⠈⢙",
			"⠈⡙",
			"⠈⠩",
			"⠀⢙",
			"⠀⡙",
			"⠀⠩",
			"⠀⢘",
			"⠀⡘",
			"⠀⠨",
			"⠀⢐",
			"⠀⡐",
			"⠀⠠",
			"⠀⢀",
			"⠀⡀"
		]
	},
	"line": {
		"speed": 130,
		"frames": [
			"-",
			"\\",
			"|",
			"/"
		]
	},
	"line2": {
		"speed": 100,
		"frames": [
			"⠂",
			"-",
			"–",
			"—",
			"–",
			"-"
		]
	},
	"pipe": {
		"speed": 100,
		"frames": [
			"┤",
			"┘",
			"┴",
			"└",
			"├",
			"┌",
			"┬",
			"┐"
		]
	},
	"simpleDots": {
		"speed": 400,
		"frames": [
			".  ",
			".. ",
			"...",
			"   "
		]
	},
	"simpleDotsScrolling": {
		"speed": 200,
		"frames": [
			".  ",
			".. ",
			"...",
			" ..",
			"  .",
			"   "
		]
	},
	"star": {
		"speed": 70,
		"frames": [
			"✶",
			"✸",
			"✹",
			"✺",
			"✹",
			"✷"
		]
	},
	"star2": {
		"speed": 80,
		"frames": [
			"+",
			"x",
			"*"
		]
	},
	"hamburger": {
		"speed": 100,
		"frames": [
			"☱",
			"☲",
			"☴"
		]
	},
	"growVertical": {
		"speed": 120,
		"frames": [
			"▁",
			"▃",
			"▄",
			"▅",
			"▆",
			"▇",
			"▆",
			"▅",
			"▄",
			"▃"
		]
	},
	"growHorizontal": {
		"speed": 120,
		"frames": [
			"▏",
			"▎",
			"▍",
			"▌",
			"▋",
			"▊",
			"▉",
			"▊",
			"▋",
			"▌",
			"▍",
			"▎"
		]
	},
	"balloon": {
		"speed": 140,
		"frames": [
			" ",
			".",
			"o",
			"O",
			"@",
			"*",
			" "
		]
	},
	"balloon2": {
		"speed": 120,
		"frames": [
			".",
			"o",
			"O",
			"°",
			"O",
			"o",
			"."
		]
	},
	"noise": {
		"speed": 100,
		"frames": [
			"▓",
			"▒",
			"░"
		]
	},
	"bounce": {
		"speed": 120,
		"frames": [
			"⠁",
			"⠂",
			"⠄",
			"⠂"
		]
	},
	"boxBounce": {
		"speed": 120,
		"frames": [
			"▖",
			"▘",
			"▝",
			"▗"
		]
	},
	"boxBounce2": {
		"speed": 100,
		"frames": [
			"▌",
			"▀",
			"▐",
			"▄"
		]
	},
	"triangle": {
		"speed": 50,
		"frames": [
			"◢",
			"◣",
			"◤",
			"◥"
		]
	},
	"arc": {
		"speed": 100,
		"frames": [
			"◜",
			"◠",
			"◝",
			"◞",
			"◡",
			"◟"
		]
	},
	"circle": {
		"speed": 120,
		"frames": [
			"◡",
			"⊙",
			"◠"
		]
	},
	"squareCorners": {
		"speed": 180,
		"frames": [
			"◰",
			"◳",
			"◲",
			"◱"
		]
	},
	"circleQuarters": {
		"speed": 120,
		"frames": [
			"◴",
			"◷",
			"◶",
			"◵"
		]
	},
	"circleHalves": {
		"speed": 50,
		"frames": [
			"◐",
			"◓",
			"◑",
			"◒"
		]
	},
	"squish": {
		"speed": 100,
		"frames": [
			"╫",
			"╪"
		]
	},
	"toggle": {
		"speed": 250,
		"frames": [
			"⊶",
			"⊷"
		]
	},
	"toggle2": {
		"speed": 80,
		"frames": [
			"▫",
			"▪"
		]
	},
	"toggle3": {
		"speed": 120,
		"frames": [
			"□",
			"■"
		]
	},
	"toggle4": {
		"speed": 100,
		"frames": [
			"■",
			"□",
			"▪",
			"▫"
		]
	},
	"toggle5": {
		"speed": 100,
		"frames": [
			"▮",
			"▯"
		]
	},
	"toggle6": {
		"speed": 300,
		"frames": [
			"ဝ",
			"၀"
		]
	},
	"toggle7": {
		"speed": 80,
		"frames": [
			"⦾",
			"⦿"
		]
	},
	"toggle8": {
		"speed": 100,
		"frames": [
			"◍",
			"◌"
		]
	},
	"toggle9": {
		"speed": 100,
		"frames": [
			"◉",
			"◎"
		]
	},
	"toggle10": {
		"speed": 100,
		"frames": [
			"㊂",
			"㊀",
			"㊁"
		]
	},
	"toggle11": {
		"speed": 50,
		"frames": [
			"⧇",
			"⧆"
		]
	},
	"toggle12": {
		"speed": 120,
		"frames": [
			"☗",
			"☖"
		]
	},
	"toggle13": {
		"speed": 80,
		"frames": [
			"=",
			"*",
			"-"
		]
	},
	"arrow": {
		"speed": 100,
		"frames": [
			"←",
			"↖",
			"↑",
			"↗",
			"→",
			"↘",
			"↓",
			"↙"
		]
	},
	"arrow2": {
		"speed": 80,
		"frames": [
			"⬆️ ",
			"↗️ ",
			"➡️ ",
			"↘️ ",
			"⬇️ ",
			"↙️ ",
			"⬅️ ",
			"↖️ "
		]
	},
	"arrow3": {
		"speed": 120,
		"frames": [
			"▹▹▹▹▹",
			"▸▹▹▹▹",
			"▹▸▹▹▹",
			"▹▹▸▹▹",
			"▹▹▹▸▹",
			"▹▹▹▹▸"
		]
	},
	"bouncingBar": {
		"speed": 80,
		"frames": [
			"[    ]",
			"[=   ]",
			"[==  ]",
			"[=== ]",
			"[ ===]",
			"[  ==]",
			"[   =]",
			"[    ]",
			"[   =]",
			"[  ==]",
			"[ ===]",
			"[====]",
			"[=== ]",
			"[==  ]",
			"[=   ]"
		]
	},
	"bouncingBall": {
		"speed": 80,
		"frames": [
			"( ●    )",
			"(  ●   )",
			"(   ●  )",
			"(    ● )",
			"(     ●)",
			"(    ● )",
			"(   ●  )",
			"(  ●   )",
			"( ●    )",
			"(●     )"
		]
	},
	"smiley": {
		"speed": 200,
		"frames": [
			"😄 ",
			"😝 "
		]
	},
	"monkey": {
		"speed": 300,
		"frames": [
			"🙈 ",
			"🙈 ",
			"🙉 ",
			"🙊 "
		]
	},
	"hearts": {
		"speed": 100,
		"frames": [
			"💛 ",
			"💙 ",
			"💜 ",
			"💚 ",
			"❤️ "
		]
	},
	"clock": {
		"speed": 100,
		"frames": [
			"🕐 ",
			"🕑 ",
			"🕒 ",
			"🕓 ",
			"🕔 ",
			"🕕 ",
			"🕖 ",
			"🕗 ",
			"🕘 ",
			"🕙 ",
			"🕚 "
		]
	},
	"earth": {
		"speed": 180,
		"frames": [
			"🌍 ",
			"🌎 ",
			"🌏 "
		]
	},
	"moon": {
		"speed": 80,
		"frames": [
			"🌑 ",
			"🌒 ",
			"🌓 ",
			"🌔 ",
			"🌕 ",
			"🌖 ",
			"🌗 ",
			"🌘 "
		]
	},
	"runner": {
		"speed": 140,
		"frames": [
			"🚶 ",
			"🏃 "
		]
	},
	"pong": {
		"speed": 80,
		"frames": [
			"▐⠂       ▌",
			"▐⠈       ▌",
			"▐ ⠂      ▌",
			"▐ ⠠      ▌",
			"▐  ⡀     ▌",
			"▐  ⠠     ▌",
			"▐   ⠂    ▌",
			"▐   ⠈    ▌",
			"▐    ⠂   ▌",
			"▐    ⠠   ▌",
			"▐     ⡀  ▌",
			"▐     ⠠  ▌",
			"▐      ⠂ ▌",
			"▐      ⠈ ▌",
			"▐       ⠂▌",
			"▐       ⠠▌",
			"▐       ⡀▌",
			"▐      ⠠ ▌",
			"▐      ⠂ ▌",
			"▐     ⠈  ▌",
			"▐     ⠂  ▌",
			"▐    ⠠   ▌",
			"▐    ⡀   ▌",
			"▐   ⠠    ▌",
			"▐   ⠂    ▌",
			"▐  ⠈     ▌",
			"▐  ⠂     ▌",
			"▐ ⠠      ▌",
			"▐ ⡀      ▌",
			"▐⠠       ▌"
		]
	},
	"shark": {
		"speed": 120,
		"frames": [
			"▐|\\____________▌",
			"▐_|\\___________▌",
			"▐__|\\__________▌",
			"▐___|\\_________▌",
			"▐____|\\________▌",
			"▐_____|\\_______▌",
			"▐______|\\______▌",
			"▐_______|\\_____▌",
			"▐________|\\____▌",
			"▐_________|\\___▌",
			"▐__________|\\__▌",
			"▐___________|\\_▌",
			"▐____________|\\▌",
			"▐____________/|▌",
			"▐___________/|_▌",
			"▐__________/|__▌",
			"▐_________/|___▌",
			"▐________/|____▌",
			"▐_______/|_____▌",
			"▐______/|______▌",
			"▐_____/|_______▌",
			"▐____/|________▌",
			"▐___/|_________▌",
			"▐__/|__________▌",
			"▐_/|___________▌",
			"▐/|____________▌"
		]
	},
	"dqpb": {
		"speed": 100,
		"frames": [
			"d",
			"q",
			"p",
			"b"
		]
	},
	"weather": {
		"speed": 100,
		"frames": [
			"☀️ ",
			"☀️ ",
			"☀️ ",
			"🌤 ",
			"⛅️ ",
			"🌥 ",
			"☁️ ",
			"🌧 ",
			"🌨 ",
			"🌧 ",
			"🌨 ",
			"🌧 ",
			"🌨 ",
			"⛈ ",
			"🌨 ",
			"🌧 ",
			"🌨 ",
			"☁️ ",
			"🌥 ",
			"⛅️ ",
			"🌤 ",
			"☀️ ",
			"☀️ "
		]
	},
	"christmas": {
		"speed": 400,
		"frames": [
			"🌲", 
			"🎄"
		]
	}
}`)

// SpinnerMap spinners
var SpinnerMap map[string]Moe

func init() {
	var tempMap map[string]map[string]interface{}
	SpinnerMap = make(map[string]Moe)

	err := json.Unmarshal(spinnersJSON, &tempMap)
	if err != nil {
		fmt.Println("Can't decode json message", err)
		return
	}

	for key, value := range tempMap {
		speed := time.Duration(value["speed"].(float64)) * time.Millisecond
		types := value["frames"].([]interface{})
		var paramSlice []string
		for _, param := range types {
			paramSlice = append(paramSlice, param.(string))
		}
		SpinnerMap[key] = Moe{speed: speed, frames: paramSlice}
	}

}
