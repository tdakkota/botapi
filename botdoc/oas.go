package botdoc

import (
	"github.com/ogen-go/ogen"
)

// OAS generates OpenAPI v3 Specification from API definition.
func (a API) OAS() *ogen.Spec {
	c := &ogen.Components{
		Schemas: map[string]ogen.Schema{},
	}
	p := ogen.Paths{
		"/getMe": ogen.PathItem{
			Description: "A simple method for testing your bot's auth token. " +
				"Requires no parameters. " +
				"Returns basic information about the bot in form of a User object.",
			Get: &ogen.Operation{
				OperationID: "getMe",
				Responses: ogen.Responses{
					"200": ogen.Response{
						Description: "Basic information about the bot",
						Content: map[string]ogen.Media{
							"application/json": {
								Schema: ogen.Schema{
									Ref: "#/components/schemas/User",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, d := range a.Types {
		s := ogen.Schema{
			Description: d.Description,
			Type:        "object",
		}
		c.Schemas[d.Name] = s
	}

	return &ogen.Spec{
		OpenAPI: "3.0.0",
		Info: ogen.Info{
			Title:          "Telegram Bot API",
			TermsOfService: "https://telegram.org/tos",
			Description:    "API for Telegram bots",
			Version:        "5.3",
		},
		Servers: []ogen.Server{
			{
				Description: "production",
				URL:         "https://api.telegram.org/",
			},
		},
		Paths:      p,
		Components: c,
	}
}
