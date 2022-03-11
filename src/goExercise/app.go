package main

import (
	"app.go/gitApis"
	"fmt"
)

func main() {
	userInfo := gitApis.GetUserData("Isuru-Dissanayake", "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")
	followerInfo := gitApis.GetUserFollowers()
	repoInfo := gitApis.GetRepos()
	fmt.Println(userInfo)
	fmt.Println(followerInfo.Followers)
	fmt.Println(repoInfo.Repos)
}
