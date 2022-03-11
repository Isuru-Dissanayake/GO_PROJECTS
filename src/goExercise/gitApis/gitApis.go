package gitApis

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type UserData struct {
	Id       int    `json:"id"`
	Username string `json:"login"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Location string `json:"location"`
}

type UserFollowers struct {
	Followers []struct {
		Name string `json:"login"`
	} `json:"data"`
}

type UserRepos struct {
	Repos []struct {
		Name string `json:"name"`
	} `json:"data"`
}

func GetUserData(userName string, privateKey string) UserData {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.SetBasicAuth(userName, privateKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

	var userInfo UserData

	if resp.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &userInfo)
	}
	return userInfo
}

func GetUserFollowers() UserFollowers {
	req, err := http.NewRequest("GET", "https://api.github.com/users/Isuru-Dissanayake/followers", nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

	var followerInfo UserFollowers

	if resp.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(body)
		bodyString = `{"data":` + bodyString + "}"
		bodyByte := []byte(bodyString)
		json.Unmarshal(bodyByte, &followerInfo)
	}
	return followerInfo
}

func GetRepos() UserRepos {
	req, err := http.NewRequest("GET", "https://api.github.com/users/Isuru-Dissanayake/repos", nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

	var repoInfo UserRepos

	if resp.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(body)
		bodyString = `{"data":` + bodyString + "}"
		bodyByte := []byte(bodyString)
		json.Unmarshal(bodyByte, &repoInfo)
	}
	return repoInfo
}
