package common

const (
	// DefaultSecret used for some verification
	DefaultSecret = "inertia"

	// MsgDaemonOK is the OK response upon successfully reaching daemon
	MsgDaemonOK = "I'm a little Webhook, short and stout!"

	// DefaultPort defines the standard daemon port
	DefaultPort = "8081"
)

// DaemonRequest is the configurable body of a request to the daemon.
type DaemonRequest struct {
	Stream     bool        `json:"stream"`
	Container  string      `json:"container,omitempty"`
	Project    string      `json:"project"`
	BuildType  string      `json:"build_type"`
	GitOptions *GitOptions `json:"git_options"`
	Secret     string      `json:"secret"`
}

// GitOptions represents GitHub-related deployment options
type GitOptions struct {
	RemoteURL string `json:"remote"`
	Branch    string `json:"branch"`
}