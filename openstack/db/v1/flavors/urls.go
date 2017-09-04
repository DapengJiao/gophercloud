package flavors

import "github.com/gophercloud/gophercloud"

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id)
}

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("flavors")
}

func listDsURL(client *gophercloud.ServiceClient, ds string, version string) string {
	return client.ServiceURL("datastores", ds, "versions", version, "flavors")
}
