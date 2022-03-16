package main

import (
	"app.go/gitApis"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type meta struct {
	Type      string
	EventId   string
	CreatedAt int64
	TracerId  string
	ServiceId string
}

type payload struct {
	gitApis.UserData
	gitApis.UserFollowers
	gitApis.UserRepos
}

func createIds() (string, string) {
	eventId := uuid.New().String()
	tracerId := uuid.New().String()
	return eventId, tracerId
}

func generateMeta() meta {
	var metaData meta
	eventId, tracerId := createIds()

	metaData.EventId = eventId
	metaData.TracerId = tracerId
	metaData.Type = "UserInfoChanged"
	metaData.ServiceId = "user-service"
	metaData.CreatedAt = time.Now().UnixNano()

	return metaData
}

func generatePayload() payload {
	var payload payload
	userInfo := gitApis.GetUserData("Isuru-Dissanayake", "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")
	followerInfo := gitApis.GetUserFollowers()
	repoInfo := gitApis.GetRepos()

	payload.UserData = userInfo
	payload.UserFollowers = followerInfo
	payload.UserRepos = repoInfo

	return payload
}

func main() {

	metaData := generateMeta()
	fmt.Println(metaData)
	payLoad := generatePayload()
	fmt.Println(payLoad)
}
