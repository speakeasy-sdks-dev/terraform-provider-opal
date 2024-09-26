// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/internal/hooks"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/internal/utils"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/retry"
	"net/http"
	"time"
)

const (
	// Production
	ServerProd string = "prod"
)

// ServerList contains the list of servers available to the SDK
var ServerList = map[string]string{
	ServerProd: "https://api.opal.dev/v1",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	Server            string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	if c.Server == "" {
		c.Server = "prod"
	}

	return ServerList[c.Server], nil
}

// OpalAPI - Opal API: Your Home For Developer Resources.
type OpalAPI struct {
	// Operations related to apps
	Apps *Apps
	// Operations related to configuration templates
	ConfigurationTemplates *ConfigurationTemplates
	// Operations related to events
	Events        *Events
	GroupBindings *GroupBindings
	// Operations related to groups
	Groups *Groups
	// Operations related to message channels
	MessageChannels    *MessageChannels
	NonHumanIdentities *NonHumanIdentities
	// Operations related to on-call schedules
	OnCallSchedules *OnCallSchedules
	// Operations related to owners
	Owners *Owners
	// Operations related to requests
	Requests *Requests
	// Operations related to resources
	Resources *Resources
	// Operations related to sessions
	Sessions *Sessions
	// Operations related to tags
	Tags *Tags
	// Operations related to uars
	Uars *Uars
	// Operations related to users
	Users *Users

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*OpalAPI)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *OpalAPI) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *OpalAPI) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServer allows the overriding of the default server by name
func WithServer(server string) SDKOption {
	return func(sdk *OpalAPI) {
		_, ok := ServerList[server]
		if !ok {
			panic(fmt.Errorf("server %s not found", server))
		}

		sdk.sdkConfiguration.Server = server
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *OpalAPI) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *OpalAPI) {
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *OpalAPI) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *OpalAPI) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *OpalAPI) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *OpalAPI {
	sdk := &OpalAPI{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0",
			SDKVersion:        "0.0.1",
			GenVersion:        "2.425.1",
			UserAgent:         "speakeasy-sdk/go 0.0.1 2.425.1 1.0 github.com/opalsecurity/terraform-provider-opal/internal/sdk",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Apps = newApps(sdk.sdkConfiguration)

	sdk.ConfigurationTemplates = newConfigurationTemplates(sdk.sdkConfiguration)

	sdk.Events = newEvents(sdk.sdkConfiguration)

	sdk.GroupBindings = newGroupBindings(sdk.sdkConfiguration)

	sdk.Groups = newGroups(sdk.sdkConfiguration)

	sdk.MessageChannels = newMessageChannels(sdk.sdkConfiguration)

	sdk.NonHumanIdentities = newNonHumanIdentities(sdk.sdkConfiguration)

	sdk.OnCallSchedules = newOnCallSchedules(sdk.sdkConfiguration)

	sdk.Owners = newOwners(sdk.sdkConfiguration)

	sdk.Requests = newRequests(sdk.sdkConfiguration)

	sdk.Resources = newResources(sdk.sdkConfiguration)

	sdk.Sessions = newSessions(sdk.sdkConfiguration)

	sdk.Tags = newTags(sdk.sdkConfiguration)

	sdk.Uars = newUars(sdk.sdkConfiguration)

	sdk.Users = newUsers(sdk.sdkConfiguration)

	return sdk
}
