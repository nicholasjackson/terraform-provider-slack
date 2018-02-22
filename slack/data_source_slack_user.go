package slack

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nlopes/slack"
)

func dataSourceSlackUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSlackUserRead,

		Schema: map[string]*schema.Schema{
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"team_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"real_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_owner": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_admin": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_bot": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceSlackUserRead(d *schema.ResourceData, meta interface{}) error {
	api := slack.New(meta.(*Config).APIKey)

	email := d.Get("email").(string)
	d.SetId(email)

	user, err := api.GetUserByEmail(email)
	if err != nil {
		return fmt.Errorf("Could not fetch user: %s", err)
	}

	d.Set("id", user.ID)
	d.Set("team_id", user.TeamID)
	d.Set("real_name", user.Profile.RealName)
	d.Set("is_owner", user.IsOwner)
	d.Set("is_admin", user.IsAdmin)
	d.Set("is_bot", user.IsBot)

	return nil
}
