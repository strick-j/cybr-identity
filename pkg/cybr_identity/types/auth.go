package types

type StartAuth struct {
	Username string `json:"User"`
	Version  string `json:"Version"`
}

type AuthResponse struct {
	Success         bool        `json:"success"`
	Result          Result      `json:"Result"`
	Message         interface{} `json:"Message"`
	MessageID       interface{} `json:"MessageID"`
	Exception       interface{} `json:"Exception"`
	ErrorID         interface{} `json:"ErrorID"`
	ErrorCode       interface{} `json:"ErrorCode"`
	IsSoftError     bool        `json:"IsSoftError"`
	InnerExceptions interface{} `json:"InnerExceptions"`
}

type Result struct {
	AuthLevel          string       `json:"AuthLevel,omitempty"`
	DisplayName        string       `json:"DisplayName,omitempty"`
	Token              string       `json:"Token,omitempty"`
	Auth               string       `json:"Auth,omitempty"`
	UserID             string       `json:"UserId,omitempty"`
	EmailAddress       string       `json:"EmailAddress,omitempty"`
	UserDirectory      string       `json:"UserDirectory,omitempty"`
	PodFqdn            string       `json:"PodFqdn,omitempty"`
	User               string       `json:"User,omitempty"`
	CustomerID         string       `json:"Customer,omitempty"`
	SystemID           string       `json:"SystemID,omitempty"`
	SourceDsType       string       `json:"SourceDsType,omitempty"`
	Summary            string       `json:"Summary"`
	ClientHints        ClientHints  `json:"ClientHints,omitempty"`
	Version            string       `json:"Version,omitempty"`
	SessionID          string       `json:"SessionId,omitempty"`
	EventDescription   interface{}  `json:"EventDescription,omitempty"`
	RetryWaitingTime   int          `json:"RetryWaitingTime,omitempty"`
	SecurityImageName  interface{}  `json:"SecurityImageName,omitempty"`
	AllowLoginMfaCache bool         `json:"AllowLoginMfaCache,omitempty"`
	Challenges         []Challenges `json:"Challenges,omitempty"`
	TenantID           string       `json:"TenantId,omitempty"`
}

type ClientHints struct {
	PersistDefault                bool `json:"PersistDefault"`
	AllowPersist                  bool `json:"AllowPersist"`
	AllowForgotPassword           bool `json:"AllowForgotPassword"`
	EndpointAuthenticationEnabled bool `json:"EndpointAuthenticationEnabled"`
}

type Mechanisms struct {
	MaskedEmailAddress string `json:"MaskedEmailAddress,omitempty"`
	AnswerType         string `json:"AnswerType"`
	Name               string `json:"Name"`
	PromptMechChosen   string `json:"PromptMechChosen,omitempty"`
	PromptSelectMech   string `json:"PromptSelectMech"`
	PartialAddress     string `json:"PartialAddress,omitempty"`
	MechanismID        string `json:"MechanismId"`
}

type Challenges struct {
	Mechanisms []Mechanisms `json:"Mechanisms"`
}
