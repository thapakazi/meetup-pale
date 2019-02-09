package pale

import (
	"fmt"
	"io/ioutil"
	"os"
)

func (m *MeetupM) get_url(urlPath string) ([]byte, error) {
	var body []byte
	apiUrl := os.Getenv("MEETUP_API_URL") + urlPath
	fmt.Println(apiUrl)
	response, err := m.client.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
		return body, err
	}
	defer response.Body.Close()

	// err = json.NewDecoder(response.Body).Decode(&body)
	// if err != nil {
	// 	fmt.Println("failed to read queue body contents", err)
	// 	return body
	// }
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return body, err
	}
	return body, nil
}
