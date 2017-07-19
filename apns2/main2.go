package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sideshow/apns2"
	"gopkg.in/alecthomas/kingpin.v2"
	apnskey "github.com/sideshow/apns2/token"
)

var (
	keyPath = kingpin.Flag("keypath", "Path to private key file.").Required().Short('k').String()
	topic           = kingpin.Flag("topic", "The topic of the remote notification, which is typically the bundle ID for your app").Required().Short('t').String()
	mode            = kingpin.Flag("mode", "APNS server to send notifications to. `production` or `development`. Defaults to `production`").Default("production").Short('m').String()
)

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("0.1").Author("Alisson Sales")
	kingpin.CommandLine.Help = `Listens to STDIN to send notifications and writes APNS response code and reason to STDOUT.
	The expected format is: <DeviceToken> <APNS Payload>
	Example: aff0c63d9eaa63ad161bafee732d5bc2c31f66d552054718ff19ce314371e5d0 {"aps": {"alert": "hi"}}`
	kingpin.Parse()

	key, err := apnskey.ECKeyFromFile(*keyPath)
	if err != nil {
		log.Fatalf("Error retrieving private key `%v`: %v", keyPath, err)
	}

	apns := apnskey.NewToken(key,"L9KK6GBQT9", "BNS3ZXYS9X")


	client := apns2.NewClientWithAPNSToken(apns)

	if *mode == "development" {
		client.Development()
	} else {
		client.Production()
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		in := scanner.Text()
		notificationArgs := strings.SplitN(in, " ", 2)
		token := notificationArgs[0]
		payload := notificationArgs[1]

		notification := &apns2.Notification{
			DeviceToken: token,
			Topic:       *topic,
			Payload:     payload,
		}

		res, err := client.Push(notification)

		if err != nil {
			log.Fatal("Error: ", err)
		} else {
			fmt.Printf("%v: '%v'\n", res.StatusCode, res.Reason)
		}
	}
}
