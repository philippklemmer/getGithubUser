package githubapi 

import (
	"net/http"
	"time"
	"encoding/json"	
	"log"
)

type UserData struct {
	Login string `json:"login"`
	Name string `json:"name"`
	HtmlUrl string `json:"html_url"`
}

func GetGitubUser(username string) (*UserData, error) {
	baseUrl := "https:github.com/"
	requestUrl := baseUrl + username
	log.Println(requestUrl)

	client := http.Client {
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "getGithubAccount")

	res, err := client.Do(req)
	if err != nil {
		return nil, err	
	}
	defer res.Body.Close()	

	data := UserData{}
	if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil 
}
