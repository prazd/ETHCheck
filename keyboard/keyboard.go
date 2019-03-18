package keyboard

import tb "gopkg.in/tucnak/telebot.v2"

var (
	EthButton = tb.ReplyButton{Text: "ETH"}

	EtcButton = tb.ReplyButton{Text: "ETC"}

	BtcButton = tb.ReplyButton{Text: "BTC"}

	BchButton = tb.ReplyButton{Text: "BCH"}

	LtcButton = tb.ReplyButton{Text: "LTC"}

	MainMenu = [][]tb.ReplyButton{
		[]tb.ReplyButton{EthButton, EtcButton},
		[]tb.ReplyButton{BtcButton, BchButton, LtcButton},
	}

	// TODO: Balance check
)