/*
 * Equinix Fabric API v4
 *
 * Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>
 *
 * API version: 4.7
 * Contact: api-support@equinix.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package v4

// ConnectivitySourceType : Type of connectivity. COLO, colocation; BMMR, building meet-me room. The default is COLO. <br> A building meet-me room (BMMR) is a room within the same building where an Equinix IBX customer can connect with a non-Equinix IBX customer.
type ConnectivitySourceType string

// List of ConnectivitySourceType
const (
	COLO_ConnectivitySourceType   ConnectivitySourceType = "COLO"
	BMMR_ConnectivitySourceType   ConnectivitySourceType = "BMMR"
	REMOTE_ConnectivitySourceType ConnectivitySourceType = "REMOTE"
)
