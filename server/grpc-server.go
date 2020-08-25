package server

import (
	"errors"
	"log"

	"github.com/gervasioamy/go-grpc-poc/proto"
	"golang.org/x/net/context"
)

// Notification is the notification to be persisted in the server
type Notification struct {
	body      string
	timestamp string
}

// NotificationsServer is the grpc server
type NotificationsServer struct {
	Notifications []Notification
}

// SendNotification is the impl of remote method defined in protobuf
func (s *NotificationsServer) SendNotification(c context.Context, n *proto.Notification) (*proto.SendNotificationResponse, error) {
	s.Notifications = append(s.Notifications, Notification{n.Body, n.Timestamp})
	newID := int32(len(s.Notifications)) - 1
	// TODO error handling
	log.Printf("Received notification: '%s' -- Number of notifications: %v", n.Body, len(s.Notifications))
	return &proto.SendNotificationResponse{Id: newID}, nil
}

// GetNotifications is the impl of remote method defined in protobuf
func (s *NotificationsServer) GetNotifications(context.Context, *proto.GetNotificationsRequest) (*proto.GetNotificationsResponse, error) {
	log.Printf("GetNotifications called. Returning %v notifications", len(s.Notifications))
	response := &proto.GetNotificationsResponse{}
	for i := 0; i < len(s.Notifications); i++ {
		n := proto.Notification{
			Body:      s.Notifications[i].body,
			Timestamp: s.Notifications[i].timestamp,
		}
		response.Notifications = append(response.Notifications, &n)
	}
	return response, nil
}

// RemoveNotification is the impl of remote method defined in protobuf
func (s *NotificationsServer) RemoveNotification(c context.Context, req *proto.RemoveNotificationRequest) (*proto.RemoveNotificationResponse, error) {
	log.Printf("Removing Notification: %v", req.Id)
	if req.Id < 0 || req.Id > int32(len(s.Notifications)) {
		return nil, errors.New("ID to be remove doesn't exist")
	}
	s.Notifications = append(s.Notifications[:req.Id], s.Notifications[req.Id+1:]...)
	return &proto.RemoveNotificationResponse{Removed: true}, nil
}

/*
type NotificationServiceServer interface {
	SendNotification(context.Context, *Notification) (*SendNotificationResponse, error)
	GetNotifications(context.Context, *GetNotificationsRequest) (*GetNotificationsResponse, error)
	RemoveNotification(context.Context, *RemoveNotificationRequest) (*RemoveNotificationResponse, error)
}
*/
