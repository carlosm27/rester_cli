package httpMethods

import (
	//"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func Patch(URL string, body string) (*resty.Response, error) {
	client := resty.New()

	resp, err := client.R().
		SetBody(body).
		Patch(URL)

	if err != nil {
		log.Println(err)
	}
	return resp, err

}
