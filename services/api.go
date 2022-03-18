package services

import (
	"encoding/json"
	"fmt"
	"gopls-workspace/models"

	"github.com/go-resty/resty/v2"
)



type ApiServices interface {
	UpdateUser(request models.GetUser) (models.Returnblock, error)
}

/*WhatsAppService ... */
type apiServices struct {
	updateuser     string
}

// New Instance
func ApiService() ApiServices {
	return &apiServices{
		updateuser:    "https://6e12-41-169-160-57.ngrok.io" + "api/user/ReturnUser",
	}
}

func (f *apiServices) UpdateUser(request models.GetUser) (models.Returnblock, error) {
	//  Create a Resty Client
	r := models.Returnblock{}

	data, _ := json.Marshal(request)

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(string(data)).
		Post(f.updateuser)

	if err != nil {
		fmt.Println(err)
	} else {
		// Explore response object
		fmt.Println("Response Info:")
		fmt.Println("  Error      :", err)
		fmt.Println("  Status Code:", resp.StatusCode())
		fmt.Println("  Status     :", resp.Status())
		fmt.Println("  Proto      :", resp.Proto())
		fmt.Println("  Time       :", resp.Time())
		fmt.Println("  Received At:", resp.ReceivedAt())
		fmt.Println("  Body       :\n", resp)
		fmt.Println()
		fmt.Println(string(resp.Body()))

		json.Unmarshal([]byte(string(resp.Body())), &r)

	}

	return r, err
}