package utils

import (
	"net/http"
	"os"
)

func SendTelegramMessage(chat_id, msg string) {
	url := "https://api.telegram.org/bot"
	token := os.Getenv("TELEGRAM_APITOKEN")
	query := "/sendMessage?chat_id=" + chat_id + "&text=" + msg
	urlwithtoken := url + token + query
	http.Get(urlwithtoken)
}

// https://api.telegram.org/bot5633579826:AAHHsgj7HxHihsnbHfiFvEqIgld5IBZi3SY/sendMessage?chat_id=518057868&text=hello
// https://api.telegram.org/bot5633579826:AAHHsgj7HxHihsnbHfiFvEqIgld5IBZi3SY/getUpdates
