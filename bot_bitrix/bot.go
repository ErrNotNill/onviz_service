package bot_bitrix

import "net/http"

func BotBitrix(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	w.WriteHeader(http.StatusOK)
}
