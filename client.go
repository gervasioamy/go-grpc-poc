package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gervasioamy/go-grpc-poc/proto"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewNotificationServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		input := strings.SplitN(text, " ", 2)
		switch input[0] {
		case "":
			// nothing to do
		case "list":
			getNotifications(c)
		case "remove":
			if len(input) < 2 {
				fmt.Println("ID missed, try again")
			} else {
				removeNotification(c, input[1])
			}
		case "send":
			if len(input) < 2 {
				fmt.Println("Notification missed, try again")
			} else {
				sendNotification(c, input[1])
			}

		default:
			// nothing
		}
	}
}

func getNotifications(c proto.NotificationServiceClient) {
	response, err := c.GetNotifications(context.Background(), &proto.GetNotificationsRequest{})
	if err != nil {
		log.Fatalf("Error when calling GetNotifications: %s", err)
	}
	log.Printf("Response from Server: ")
	for i, n := range response.Notifications {
		log.Printf("  - [%v] (%s) - %s", i, n.Timestamp, n.Body)
	}
}

func removeNotification(c proto.NotificationServiceClient, idFromConsole string) {
	i, err := strconv.Atoi(idFromConsole)
	if err != nil {
		log.Printf("Id typed is NaN %v", err)
		return
	}
	response, err := c.RemoveNotification(context.Background(), &proto.RemoveNotificationRequest{Id: int32(i)})
	if err != nil {
		log.Fatalf("Error when calling RemoveNotification: %s", err)
		return
	}
	log.Printf("Response from Server: Was removed? %t", response.Removed)
}

func sendNotification(c proto.NotificationServiceClient, text string) {
	n := proto.Notification{
		Body:      text,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
	}
	response, err := c.SendNotification(context.Background(), &n)
	if err != nil {
		log.Fatalf("Error when calling SendNotification: %s", err)
	}
	log.Printf("Response from Server: ID = %v", response.Id)
}
