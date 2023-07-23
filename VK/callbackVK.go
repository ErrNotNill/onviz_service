package VK

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/callback"
	"github.com/SevereCloud/vksdk/v2/events"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var SecretKey = "AA1379657aa"
var Token = "vk1.a.uuGYW0HLLXGopIeCS0mwuHG6GnuBJwDMa8GaYgIwy07nKgjNRrE-gTUUg868oQvep7pjozOiTixAD9j_CpZXgxbCY37NFWoIm392Mxp41-4XfX6U_fhOXH2fg_o50dbGIJtRCNsH7J8YdHNxwcVYOhHC24X2lCqf8YghKs8FWJa4r7sW6u7qK4W_3CR2uSnT3hiXR9azxqjp8ZW89VyMSQ"
var ConfirmationToken = "17d6f039"

var newUserId int
var newPostId int

type Fields struct {
	Fields struct {
		TITLE              string `json:"TITLE"`
		NAME               string `json:"NAME"`
		COMMENTS           string `json:"COMMENTS"`
		SOURCE_DESCRIPTION string `json:"SOURCE_DESCRIPTION"`
		SOURCE_ID          string `json:"SOURCE_ID"`
	} `json:"fields"`
}

func CallBack(w http.ResponseWriter, r *http.Request) {
	cb := callback.NewCallback()
	fmt.Println("Callback service started")
	cb.ConfirmationKey = ConfirmationToken
	cb.SecretKey = SecretKey

	fmt.Println("Confirmation accepted")

	cb.WallReplyNew(func(ctx context.Context, obj events.WallReplyNewObject) {

		newUserId = obj.FromID
		newPostId = obj.PostID

		convPostToStr := strconv.Itoa(newPostId)
		convUserIdToStr := strconv.Itoa(newUserId)

		UrlOnUser := fmt.Sprintf("https://vk.com/id%v", convUserIdToStr)
		UrlOnPost := fmt.Sprintf("https://vk.com/onviz?w=wall-165775952_%v", convPostToStr)

		tn := &Fields{struct {
			TITLE              string `json:"TITLE"`
			NAME               string `json:"NAME"`
			COMMENTS           string `json:"COMMENTS"`
			SOURCE_DESCRIPTION string `json:"SOURCE_DESCRIPTION"`
			SOURCE_ID          string `json:"SOURCE_ID"`
		}{TITLE: "Комментарий из ВК", SOURCE_ID: "ВКонтакте - Вконтакте", NAME: UrlOnUser, COMMENTS: obj.Text, SOURCE_DESCRIPTION: UrlOnPost}}

		jsnm, err := json.Marshal(tn)
		if err != nil {
			log.Println("Error to convert json fields from struct")
		}
		r := bytes.NewReader(jsnm)

		is_words := regexp.MustCompile(`[a-zA-Zа-яА-Я]`).MatchString(obj.Text)
		//is_numbers := regexp.MustCompile(`[0-9]`).MatchString(txt)

		bytesSlice := []byte(obj.Text)
		bytesRune := bytes.Runes(bytesSlice)

		if is_words == true && len(bytesRune) > 1 && obj.Text != "" && obj.FromID != 628998745 && obj.FromID != 629352947 && obj.FromID != 642491603 {
			fmt.Println(obj.FromID)
			_, err = http.Post("https://onviz.bitrix24.ru/rest/13938/pqq6j4ohvutvzfmi/crm.lead.add", "application/json", r)
			if err != nil {
				log.Println("Error http:post request to Bitrix24")
			}
			log.Println("Lead was send to Bitrix24")
			log.Printf("User Url: %v / Post Url: %v / Comment Text: %v", UrlOnUser, UrlOnPost, obj.Text)
		} else {
			fmt.Println("Not correctly commentary. Maybe contains only nums or nil commentary or this commentary was wrote by group administrator")
		}
	})
	cb.HandleFunc(w, r)
}
