package notification

import "fmt"

type (
	EmailNotification struct{ Email string }
	SmsNotification   struct{ Phone string }
	PushNotification  struct{ DeviceToken string }
)

func (e *EmailNotification) SendNotification(message string) error {
	fmt.Printf("Sending email to %s: %s\n", e.Email, message)
	return nil
}

func (s *SmsNotification) SendNotification(message string) error {
	fmt.Printf("Sending SMS to %s: %s\n", s.Phone, message)
	return nil
}

func (p *PushNotification) SendNotification(message string) error {
	fmt.Printf("Sending push notification to device %s: %s\n", p.DeviceToken, message)
	return nil
}
