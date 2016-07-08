package main

import (
    "fmt"
    "os"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/terraform"
)

// Provider is a basic structure that describes a provider: the configuration
// keys it takes, the resources it supports, a callback to configure, etc.
func Provider() terraform.ResourceProvider {
    // The actual provider
    return &schema.Provider{
        Schema: map[string]*schema.Schema{
            "debug": &schema.Schema{
                Type:        schema.TypeBool,
                Optional:    true,
                Default:     false,
           },
            "insecure": &schema.Schema{
                Type:        schema.TypeBool,
                Optional:    true,
                Default:     false,
           },
            "nsxusername": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Default:     os.Getenv("NSXUSERNAME"),
           },
            "nsxpassword": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Default:     os.Getenv("NSXPASSWORD"),
           },
            "nsxserver": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                Default:     os.Getenv("NSXSERVER"),
           },
        },

        ResourcesMap: map[string]*schema.Resource{
            "nsx_logical_switch": resourceLogicalSwitch(),
        },

        ConfigureFunc: providerConfigure,
    }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	debug := d.Get("debug").(bool)
        insecure := d.Get("insecure").(bool)
        nsxusername := d.Get("nsxusername").(string)

	if nsxusername == "" {
		return nil, fmt.Errorf("nsxusername must be provided.")
	}

        nsxpassword := d.Get("nsxpassword").(string)

	if nsxpassword == "" {
		return nil, fmt.Errorf("nsxpassword must be provided.")
	}

        nsxserver := d.Get("nsxserver").(string)

	if nsxserver == "" {
		return nil, fmt.Errorf("nsxserver must be provided.")
	}

	config := Config{
		Debug:       debug,
		Insecure:    insecure,
		NSXUserName: nsxusername,
		NSXPassword: nsxpassword,
		NSXServer:   nsxserver,
	}

	return config.Client()
}
