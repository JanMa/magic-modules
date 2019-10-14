package google

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGoogleComputeInstanceList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGoogleComputeInstanceListRead,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"project": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"zone": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"self_link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceGoogleComputeInstanceListRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	zone, err := getZone(d, config)
	if err != nil {
		return err
	}

	instanceList, err := config.clientCompute.Instances.List(project, zone).Filter(d.Get("filter").(string)).Do()
	if err != nil {
		return err
	}

	list := []map[string]string{}
	for _, i := range instanceList.Items {
		list = append(list, map[string]string{
			"name":      i.Name,
			"project":   project,
			"zone":      i.Zone,
			"self_link": i.SelfLink,
		})
	}

	d.Set("self_link", instanceList.SelfLink)
	d.Set("instances", list)
	d.SetId(instanceList.Id)
	return nil
}
