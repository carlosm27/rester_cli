package httpMethods

import (
	//"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func Get(URL string) (*resty.Response, error) {
	client := resty.New()

	resp, err := client.R().EnableTrace().Get(URL)

	if err != nil {
		log.Println(err)
	}
	return resp, err

}
