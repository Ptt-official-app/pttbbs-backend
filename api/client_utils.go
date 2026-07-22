package api

import (
	"encoding/json"

	"github.com/Ptt-official-app/pttbbs-backend/schema"
)

func checkClient(clientID string, clientSecret string) (isValid bool, client *schema.Client) {
	client, err := schema.GetClient(clientID)
	if err != nil {
		return false, nil
	}
	if client == nil {
		return false, nil
	}

	return true, client
}

func getClientInfo(client *schema.Client) string {
	clientInfo := &ClientInfo{
		ClientID:   client.ClientID,
		ClientType: client.ClientType,
	}
	str, _ := json.Marshal(clientInfo)
	return string(str)
}
