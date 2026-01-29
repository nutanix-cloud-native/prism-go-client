// Package v4 provides the Prism Central V4 API implementation of the converged client.
// It implements the interfaces defined in the converged package (AntiAffinityPolicies,
// Clusters, Categories, Images, StorageContainers, Subnets, VMs, Tasks, VolumeGroups,
// DomainManager, Users) using the Nutanix V4 REST API and SDK.
//
// # Creating a converged client with cache
//
// Create and use the converged V4 client via a ClientCache. The cache reuses clients
// per key and invalidates them when the management endpoint or credentials change,
// so long-lived processes (e.g. controllers) should always get a client through the cache.
//
// 1. Create a cache (typically once per process):
//
//	cache := v4.NewClientCache()
//
// 2. Implement types.CachedClientParams for your cluster/endpoint. It must provide Key()
//    (unique per Prism Central instance, e.g. cluster name) and ManagementEndpoint()
//    (address, credentials, TLS options from github.com/nutanix-cloud-native/prism-go-client/environment/types):
//
//	type myCachedClientParams struct {
//	    name         string
//	    mgmtEndpoint types.ManagementEndpoint
//	}
//	func (p *myCachedClientParams) Key() string { return p.name }
//	func (p *myCachedClientParams) ManagementEndpoint() types.ManagementEndpoint { return p.mgmtEndpoint }
//
// 3. Get or create the converged client for a given endpoint:
//
//	params := &myCachedClientParams{name: "cluster1", mgmtEndpoint: managementEndpoint}
//	client, err := cache.GetOrCreate(params)
//	if err != nil {
//	    return err
//	}
//	// Use client.Clusters, client.VMs, client.Images, etc.
//
// The cache hashes the management endpoint. If credentials or endpoint URL change,
// GetOrCreate detects the change and creates a new client. To force recreation (e.g.
// after credential rotation), remove the entry and call GetOrCreate again:
//
//	cache.Delete(params)
//	client, err = cache.GetOrCreate(params)
//
// GetOrCreate accepts optional types.ClientOption[v4prismGoClient.Client] values
// (e.g. for TLS or proxy) when creating a new client.
//
// # Usage examples (client from cache.GetOrCreate)
//
// List clusters and get Prism Central version:
//
//	clusters, err := client.Clusters.List(ctx)
//	version, err := client.DomainManager.GetPrismCentralVersion(ctx)
//
// List VMs with pagination and filter:
//
//	vms, err := client.VMs.List(ctx,
//	    converged.WithPage(0),
//	    converged.WithLimit(50),
//	    converged.WithFilter("name eq 'my-vm'"),
//	)
//
// Create a VM asynchronously and wait for the result:
//
//	op, err := client.VMs.CreateAsync(ctx, vmSpec)
//	if err != nil { return err }
//	createdVMs, err := op.Wait(ctx)
//
// Stream all images with an iterator (no full list in memory):
//
//	it := client.Images.NewIterator(ctx)
//	for img, err := range it {
//	    if err != nil { return err }
//	    use(img)
//	}
//
// Get a cluster, list its hosts, and power on a VM:
//
//	cluster, err := client.Clusters.Get(ctx, clusterUUID)
//	hosts, err := client.Clusters.ListClusterHosts(ctx, clusterUUID, converged.WithLimit(100))
//	powerOp, err := client.VMs.PowerOnVM(vmUUID)
//	_, err = powerOp.Wait(ctx)
//
// Reserve IPs on a subnet (returns a task reference for the async operation):
//
//	taskRef, err := client.Subnets.ReserveIpsBySubnetId(ctx, subnetExtId, ipReserveSpec)
//
// All list/get/create/update/delete methods accept a context.Context for cancellation and timeouts.
package v4
