package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SLACK_API_KEY", nil),
				Description: "The api key to use for interacting with your Slack team.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"slack_bot": resourceBot(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		APIKey: d.Get("api_key").(string),
	}

	return config, nil
}
