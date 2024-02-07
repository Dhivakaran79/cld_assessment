package worker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"assessment_cld/model"
)

// ProcessMessages concurrently processes incoming messages from the channel and sends them to the webhook URL
func ProcessMessages(channel <-chan model.ContactFormSubmission) {
	for msg := range channel {
		go func(msg model.ContactFormSubmission) {
			// Convert the message into the desired format
			payload := map[string]interface{}{
				"event":            msg.Ev,
				"event_type":       msg.Et,
				"app_id":           msg.Id,
				"user_id":          msg.Uid,
				"message_id":       msg.Mid,
				"page_title":       msg.T,
				"page_url":         msg.P,
				"browser_language": msg.L,
				"screen_size":      msg.Sc,
				"attributes": map[string]interface{}{
					msg.Atrk1: map[string]interface{}{
						"value": msg.Atrv1,
						"type":  msg.Atrt1,
					},
					msg.Atrk2: map[string]interface{}{
						"value": msg.Atrv2,
						"type":  msg.Atrt2,
					},
				},
				"traits": map[string]interface{}{
					msg.Uatrk1: map[string]interface{}{
						"value": msg.Uatrv1,
						"type":  msg.Uatrt1,
					},
					msg.Uatrk2: map[string]interface{}{
						"value": msg.Uatrv2,
						"type":  msg.Uatrt2,
					},
					msg.Uatrk3: map[string]interface{}{
						"value": msg.Uatrv3,
						"type":  msg.Uatrt3,
					},
				},
			}

			// Convert the payload to JSON
			payloadBytes, err := json.Marshal(payload)
			if err != nil {
				fmt.Println("Error marshalling JSON:", err)
				return
			}

			// Send the payload to the webhook
			webhookURL := "https://webhook.site/f3413769-6abc-4d38-aba4-f066eac7318b"
			resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payloadBytes))
			if err != nil {
				fmt.Println("Error sending payload to webhook:", err)
				return
			}
			defer resp.Body.Close()

			// Log the response from the webhook
			fmt.Println("Webhook response:", resp.Status)
		}(msg)
	}
}
