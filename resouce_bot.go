package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nlopes/slack"
)

func resourceBot() *schema.Resource {
	return &schema.Resource{
		Create: resourceBotCreate,
		Read:   resourceBotRead,
		Update: resourceBotUpdate,
		Delete: resourceBotDelete,

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceBotCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[INFO] Creating Bot")

	api := slack.New(m.(*Config).APIKey)
	api.SetDebug(true)

	log.Println("[INFO] Getting users")
	groups, err := api.GetUsers()
	if err != nil {
		log.Println("[DEBUG] Error getting users: ", err)
		return nil
	}

	for _, group := range groups {
		log.Printf("ID: %s, Name: %s\n", group.Id, group.Name)
	}

	d.SetId("nicj")

	return nil
}

func resourceBotRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[INFO] Creating Bot")

	return nil
}

func resourceBotUpdate(d *schema.ResourceData, m interface{}) error {
	fmt.Println("Updating Bot")

	return nil
}

func resourceBotDelete(d *schema.ResourceData, m interface{}) error {
	fmt.Println("Deleting Bot")

	return nil
}
