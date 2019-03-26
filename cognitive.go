package neocortex

// CognitiveServiceType represents the type of your cognitive service (ibm watson, dialog flow, aws chat, etc...)
type CognitiveServiceType string

// Watson is a ibm watson communicate service
const Watson CognitiveServiceType = "watson"

// DialogFlow is a dialogflow service
const DialogFlow CognitiveServiceType = "dflow"

// CognitiveService is any service can process the natural language conversation
type CognitiveService struct {
	Type      CognitiveServiceType
	URL       string
	APIKey    string
	SecretKey string
}
