package httpMethods

import (
	//"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func Delete(URL string) (*resty.Response, error) {
	client := resty.New()

	resp, err := client.R().
		Delete(URL)

	if err != nil {
		log.Println(err)
	}
	return resp, err

}
