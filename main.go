package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func init() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("error loading .env file")
	}

}

func main() {

os.Setenv("SLACK_BOT_TOKEN", os.Getenv("SLACK_BOT_TOKEN"))
os.Setenv("CHANNEL_ID", os.Getenv("CHANNEL_ID"))
api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
channelArr := []string{os.Getenv("CHANNEL_ID")}
fileArr := []string{"busanharbor1.jpeg"}

for i := 0; i<len(fileArr); i++ {
	params := slack.FileUploadParameters {
		Channels: channelArr,
		File: fileArr[i],
	}
	file, err := api.UploadFile(params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
}
}