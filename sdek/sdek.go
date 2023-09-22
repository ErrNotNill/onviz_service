package sdek

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	CdekCLientId     = os.Getenv("CDEK_CLIENT_ID")
	CdekClientSecret = os.Getenv("CDEK_CLIENT_SECRET")
)

func SdekStart() {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.cdek.ru/v2/oauth/token?grant_type=client_credentials&client_id=%s&client_secret=%s", CdekCLientId, CdekClientSecret)

	post, err := client.Post(url, "application/x-www-form-urlencoded", nil)
	body, err := io.ReadAll(post.Body)

	resp := &Response{}
	json.Unmarshal(body, &resp)

	fmt.Println(post.StatusCode)
	fmt.Println("access_token", resp.AccessToken)
	fmt.Println("token_type", resp.TokenType)
	fmt.Println("expires_in", resp.ExpiresIn)
	fmt.Println("scope", resp.Scope)
	fmt.Println("jti", resp.Jti)

	//order := &Order{}
	var bearer = "Bearer " + resp.AccessToken
	req, err := http.NewRequest("GET", `https://api.cdek.ru/v2/orders?cdek_number=1460493586`, nil)
	req.Header.Add("Authorization", bearer)
	newclient := &http.Client{}
	rez, err := newclient.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer rez.Body.Close()

	newbody, err := io.ReadAll(rez.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println("newBody", string([]byte(newbody)))

	newReq, err := http.NewRequest("GET", `https://api.cdek.ru/v2/registries?date=2023-09-12`, nil)
	newReq.Header.Add("Authorization", bearer)
	newPay := &http.Client{}
	newRez, err := newPay.Do(newReq)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer newRez.Body.Close()

	newBody, err := io.ReadAll(newRez.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println("newBody", string([]byte(newBody)))

	/*get, err := client.Get(`https://api.cdek.ru/v2/orders?cdek_number=1463958253`)
	get.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	get.Header.Add("Authorization", bearer)
	//fmt.Println("bearer::: ", bearer)

	body, err = io.ReadAll(get.Body)
	fmt.Println("StatusCode: ", get.StatusCode)
	json.Unmarshal(body, &order)
	fmt.Println("order", order)
	fmt.Println("get.Body", get.Body)*/

	//client.Get(`https://api.cdek.ru/v2/payment`)

	//req.Header.Set("Content-Type", "x-www-form-urlencoded")
	//fmt.Println("req.Body", req.Body)
}

type Response struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	Jti         string `json:"jti"`
}

type Order struct {
	Entity struct {
		Uuid                  string `json:"uuid"`
		Type                  int    `json:"type"`
		IsReturn              bool   `json:"is_return"`
		IsReverse             bool   `json:"is_reverse"`
		CdekNumber            string `json:"cdek_number"`
		Number                string `json:"number"`
		DeliveryMode          string `json:"delivery_mode"`
		TariffCode            int    `json:"tariff_code"`
		Comment               string `json:"comment"`
		DeliveryRecipientCost struct {
			Value  int     `json:"value"`
			VatSum float64 `json:"vat_sum"`
		} `json:"delivery_recipient_cost"`
		DeliveryRecipientCostAdv []struct {
			Threshold int     `json:"threshold"`
			Sum       float64 `json:"sum"`
			VatSum    float64 `json:"vat_sum"`
		} `json:"delivery_recipient_cost_adv"`
		Sender struct {
			Company string `json:"company"`
			Name    string `json:"name"`
			Phones  []struct {
				Number string `json:"number"`
			} `json:"phones"`
			PassportRequirementsSatisfied bool `json:"passport_requirements_satisfied"`
		} `json:"sender"`
		Seller struct {
			Name string `json:"name"`
		} `json:"seller"`
		Recipient struct {
			Company string `json:"company"`
			Name    string `json:"name"`
			Email   string `json:"email"`
			Phones  []struct {
				Number string `json:"number"`
			} `json:"phones"`
			PassportRequirementsSatisfied bool `json:"passport_requirements_satisfied"`
		} `json:"recipient"`
		FromLocation struct {
			Code        int    `json:"code"`
			PostalCode  string `json:"postal_code"`
			CountryCode string `json:"country_code"`
			Region      string `json:"region"`
			RegionCode  int    `json:"region_code"`
			City        string `json:"city"`
			Address     string `json:"address"`
			Country     string `json:"country"`
		} `json:"from_location"`
		ToLocation struct {
			Code        int     `json:"code"`
			FiasGuid    string  `json:"fias_guid"`
			PostalCode  string  `json:"postal_code"`
			Longitude   float64 `json:"longitude"`
			Latitude    float64 `json:"latitude"`
			CountryCode string  `json:"country_code"`
			Region      string  `json:"region"`
			RegionCode  int     `json:"region_code"`
			City        string  `json:"city"`
			Address     string  `json:"address"`
			Country     string  `json:"country"`
		} `json:"to_location"`
		Services []struct {
			Code      string  `json:"code"`
			Parameter string  `json:"parameter,omitempty"`
			Sum       float64 `json:"sum"`
		} `json:"services"`
		Packages []struct {
			PackageId    string `json:"package_id"`
			Number       string `json:"number"`
			Weight       int    `json:"weight"`
			Length       int    `json:"length"`
			Width        int    `json:"width"`
			WeightVolume int    `json:"weight_volume"`
			WeightCalc   int    `json:"weight_calc"`
			Height       int    `json:"height"`
			Comment      string `json:"comment"`
			Items        []struct {
				Name    string `json:"name"`
				WareKey string `json:"ware_key"`
				Payment struct {
					Value  float64 `json:"value"`
					VatSum float64 `json:"vat_sum"`
				} `json:"payment"`
				Weight         int     `json:"weight"`
				WeightGross    int     `json:"weight_gross"`
				Amount         int     `json:"amount"`
				DeliveryAmount int     `json:"delivery_amount"`
				NameI18N       string  `json:"name_i18n"`
				Url            string  `json:"url"`
				Cost           float64 `json:"cost"`
				Excise         bool    `json:"excise"`
			} `json:"items"`
		} `json:"packages"`
		DeliveryProblem []interface{} `json:"delivery_problem"`
		Statuses        []struct {
			Code     string `json:"code"`
			Name     string `json:"name"`
			DateTime string `json:"date_time"`
			City     string `json:"city"`
		} `json:"statuses"`
		DeliveryDetail struct {
			Date          string        `json:"date"`
			RecipientName string        `json:"recipient_name"`
			PaymentSum    float64       `json:"payment_sum"`
			DeliverySum   float64       `json:"delivery_sum"`
			TotalSum      float64       `json:"total_sum"`
			PaymentInfo   []interface{} `json:"payment_info"`
		} `json:"delivery_detail"`
	} `json:"entity"`
	Requests []struct {
		RequestUuid string `json:"request_uuid"`
		Type        string `json:"type"`
		DateTime    string `json:"date_time"`
		State       string `json:"state"`
	} `json:"requests"`
}
