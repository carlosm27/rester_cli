package httpMethods

import (
	//"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func Put(URL string, body string) (*resty.Response, error) {
	client := resty.New()

	resp, err := client.R().
		SetBody(body).
		Put(URL)

	if err != nil {
		log.Println(err)
	}
	return resp, err

}
