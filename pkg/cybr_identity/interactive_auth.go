package cybr_identity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/strick-j/cybr_identity/pkg/cybr_identity/types"
)

// StartAuthentication performs the initial step in an interactive authentication process:
//
//	 clientURL - URL for the CyberArk Identity tenant (e.g. "example.my.cyberark.cloud")
//	 clientVersion - API Version - Default 1.0
//		clientUsername - Username for interactive authentication (e.g. "identity-privilege-integration-user$@example.com")
func StartAuthentication(clientUsername, clientVersion, clientURL string) (*types.AuthResponse, error) {
	// Establish body for StartAuthentication API Call
	startAuth := types.StartAuth{
		Username: clientUsername,
		Version:  clientVersion,
	}

	payload, err := json.Marshal(startAuth)
	if err != nil {
		return nil, err
	}

	authUrl := makeRouterURL(clientURL, "Security", "StartAuthentication").String()

	var httpClient *http.Client
	req, err := http.NewRequest("GET", authUrl, bytes.NewBuffer(payload))
	if err != nil {
		err = fmt.Errorf("error generating the authentication client request : %s", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("Accept", "*/*")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Generate Conjur Authentication Token from Conjur Byte Response
	authResponse, err := types.AuthResponse(body)
	if err != nil {
		err = fmt.Errorf("unable to generate Authentication Response from Start Authentication Byte Response. Error: %s", err)
		return nil, err
	}

	return authResponse, nil
}
