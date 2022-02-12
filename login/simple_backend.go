package login

const SimpleProviderName = "simple"

func init() {
	RegisterProvider(
		&ProviderDescription{
			Name: SimpleProviderName,
		},
		SimpleBackendFactory)
}

func SimpleBackendFactory(config map[string]string) (Backend, error) {
	userPassword := map[string]string{}
	for k, v := range config {
		if k != "provider" && k != "name" {
			userPassword[k] = v
		}
	}
	return NewSimpleBackend(userPassword), nil
}

// Simple backend, working on a map of username password pairs
type SimpleBackend struct {
	userPassword map[string]string
}

// NewSimpleBackend creates a new SIMPLE Backend and verifies the parameters.
func NewSimpleBackend(userPassword map[string]string) *SimpleBackend {
	return &SimpleBackend{
		userPassword: userPassword,
	}
}

func (sb *SimpleBackend) Authenticate(username, password string) (bool, UserInfo, error) {
	if p, exist := sb.userPassword[username]; exist && p == password {
		return true, UserInfo{Username: username}, nil
	}
	return false, UserInfo{}, nil
}