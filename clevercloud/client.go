package clevercloud

import "go.clever-cloud.dev/client"

func getClient() *client.Client {
	return client.New(
		client.WithAutoOauthConfig(),
	)
}
