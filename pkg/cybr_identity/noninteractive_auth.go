package cybr_cli

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type tokenSource struct {
	ctx  context.Context
	conf *clientcredentials.Config
}

// OauthCredClient returns a validated OAuth2 Authentication Token based on the following provided information:
//
//	clientID - Username for the OAuth2 Application (e.g. "identity-privilege-integration-user$@example.com")
//	clientSecret - Password for the OAuth2 Application
//	clientAppID - ID for the OAuth2 Application
//	clientURL - URL for the OAuth2 Application (e.g. "example.my.cyberark.cloud")
//	clientScope - Array containing scopes for claim
func OauthCredClient(clientID, clientSecret, clientAppID, clientURL string, clientScope []string) (*oauth2.Token, error) {
	// Establish oauth2/clientcredentials config with user provided data
	var credentialConfig = clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://" + clientURL + "/oauth2/token/" + clientAppID,
		AuthStyle:    0,
		Scopes:       clientScope,
	}

	// Create tokenSource with provided configuration info
	ts := &tokenSource{
		ctx:  context.Background(),
		conf: &credentialConfig,
	}

	// Request new token from OAuth2 server using Client Credentials
	authToken, err := ts.conf.Token(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to obtain OAuth2 Token %w", err)
	}

	return authToken, nil
}

// OauthResourceOwner returns a validated OAuth2 Authentication Token with Refresh Token based on the following provided information:
//
//	clientID - Username for the OAuth2 Application (e.g. "identity-privilege-integration-user$@example.com")
//	clientSecret - Password for the OAuth2 Application
//	clientAppID - ID for the OAuth2 Application
//	clientURL - URL for the OAuth2 Application (e.g. "example.my.idaptive.app")
//	clientScope -
//	resourceUsername - Username for the Resource Owner
//	resourcePassword - Password for the Resource Owner
func OauthResourceOwner(clientID, clientSecret, clientAppID, clientURL, resourceUsername, resourcePassword string, clientScope []string) (*oauth2.Token, error) {
	endpoint := oauth2.Endpoint{
		AuthURL:   "https://" + clientURL + "/oauth2/authorize/" + clientAppID,
		TokenURL:  "https://" + clientURL + "/oauth2/token/" + clientAppID,
		AuthStyle: 0,
	}

	// Establish oauth2/clientcredentials config with user provided data
	var resourceOwnerConfig = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     endpoint,
		Scopes:       clientScope,
	}

	ctx := context.Background()

	authToken, err := resourceOwnerConfig.PasswordCredentialsToken(ctx, resourceUsername, resourcePassword)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain OAuth2 Token %w", err)
	}

	return authToken, nil
}
