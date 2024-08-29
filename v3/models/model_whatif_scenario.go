/*
Nutanix Intentful API

# Introduction  **Representational state transfer ( REST )** Application Programming Interface (API) 3.0 is based on an intentful API philosophy. According to the intentful API philosophy the machine should handle the programming instead of the user enabling the datacenter administrator able to focus on the other tasks. You need to specify the end state of an entity and the system will compile and execute a series of steps to achieve the defined end state of the entity. The progress to achieve the desired state is tracked through waits and events. This approach is similar to the Google Kubernetes standard.  All the entities and the list of entities follow a homogenous specification format resulting in ease of programming despite different entities being involved. Also, this enables the user to get familiar with the APIs quickly.  Intentful APIs reduces the number of action verbs required to achieve the desired state of an entity. Most of the changes can be achieved by defining the desired state of an entity rather than through a series of action verbs.  All API users should have valid Prism login credentials to send API calls to the Prism server.  APIs on nutanix.dev portal are publicly available to all valid users without special permissions for viewing purposes.  This document covers APIs that are available in pc2022.11 release.    ## Terminology  | Item                  | Description                                                                                                                                                             | | --------------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | | Entity                | An object that is managed through an API.                                                                                                                               | | Kind                  | Represents the type of entity. For example, VM or alerts.                                                                                                               | | Resource Kind         | An entity that represents a physical or virtual infrastructure resource. For example, VM or an application.                                                             | | Helper Kind           | An entity that is not an infrastructure resource, but represents the entity that is used along with the infrastructure resources. For example, alert, event, or status. | | Basic Resource Kind   | An entity type that is a primary resource. For example, VM.                                                                                                             | | Related Resource Kind | An entity type that is automatically defined by the system and is related to the basic Kind entities. For example, vm_snapshot.                                         | | Type                  | Represents sub-categories or sub-objects of the top-level resource types. For example, reference.                                                                       | | Resource Limit        | An entity type that represents a quota of resources of different types.                                                                                                 | | Snapshot              | Represents the state of resources and policies of an entity at a particular point of time.                                                                              | | Backup                | An automatically created aggregation object that collects all the snapshots for the same object within the same resource pool.                                          | | Profile               | A profile is a set of resources and policy specifications that is applied to an entity.                                                                                 | | Template              | A template is a specification that generates valid entities by providing input parameters.                                                                              | | Spec                  | A section in the request json, which represents an entity, that expresses the desired state of the entity.                                                              | | Status                | A section in the json representation that expresses the current state of the entity.                                                                                    |  ## Rules  - The metadata section has fields that relate to all kinds of entities and can be updated instantly by the system. - The spec section cannot be updated after initialization by the system without a user or automation intervention. - The status section has a copy of the currently enforced version of the spec section and other fields that are managed or updated automatically by the system. Few fields in the status section can be updated by the user as well. - The parent_reference field is mandatory for snapshot and backup entities and is present in the normal entities, if the entities are created with the parent_reference field populated (clone or restore). - Any field that represents a defined type has the type name as a suffix. For example, backup_factory and vm_reference.  ## Known Issues and limitations:  This section describes some of the issues identified in V3 APIs that you might encounter.  - **ENG-321951**- While attempting to get a list of all VMs on Prism Central, we have observed that only AHV VMs are returned, but ESXi VM list is not retrieved.  - **ENG-453487**- Users with insufficient permissions receive an HTTP 200 response instead of an HTTP 403 response for view alert API requests. These users can be like \"Developer\", \"Operator\", or custom roles that do not have sufficient permission to view alerts. This is a generic issue for all V3 API requests with insufficient permissions.   ### [Nutanix API versions – What are they and what does each one do?](https://developer.nutanix.com/2019/01/15/nutanix-api-versions-what-are-they-and-what-does-each-one-do/)   ### [What to expect post pc.2022.1 release?](https://www.nutanix.dev/api_references/prism-central-v3/#/ZG9jOjM2MzQyMjUz-v3-api-behavior-changes-with-pc-2022-1-release)  # Getting Started  ## Components of a REST API  A REST API request and response can be classified into the following components.  The request URI consists of {URI-scheme} :// {URI-host} / {<Port-number>/<Base-path>} ? {query-string}.  | Item                    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | | ----------------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | | URI scheme              | Indicates the protocol that is used to transmit the request. For example, http or https.                                                                                                                                                                                                                                                                                                                                                                                                                                    | | URI host                | Specifies the domain name or IP address of the server where the REST service endpoint is hosted. For example, prism.nutanix.com or 10.11.12.13.                                                                                                                                                                                                                                                                                                                                                                             | | Port number             | Indicates the port number that is used for HTTP communications from the REST API clients to the REST API server.                                                                                                                                                                                                                                                                                                                                                                                                            | | Basepath                | Indicates the URL prefix that is common to all the API server endpoints. For example, for <https://10.11.12.13/PrismGateway/services/rest/v2.0/vms>, the basepath is /PrismGateway/services/rest/v2.0.                                                                                                                                                                                                                                                                                                                      | | Path parameters         | Path parameters are part of the endpoint. For example, <https://10.11.12.13/PrismGateway/services/rest/v2.0/vms/750eb5ba-2de0-4ca4-8d6f-786364817f94/>, the UUID of the VM (750eb5ba-2de0-4ca4-8d6f-786364817f94) is a path parameter.                                                                                                                                                                                                                                                                                      | | Query string parameters | Query string parameters appear after a question mark (?) in the endpoint. The question mark followed by the parameters and their values is referred as the query string. In the query string, each parameter is listed one right after the other with an ampersand (&). The order of the query string parameters does not matter. For example, <https://10.11.12.13/PrismGateway/services/rest/v2.0/vms/750eb5ba-2de0-4ca4-8d6f-786364817f94?include_vm_disk_config=true>, the include_vm_disk_config is a query parameter. |  HTTP defines methods to indicate the desired action to be performed on the identified resource.  | Method | Description                                    | | ------ | :--------------------------------------------- | | Get    | Retrieves a resource                           | | Post   | Creates a resource                             | | Put    | Updates or creates within an existing resource | | Delete | Removes the resource                           |  HTTP request headers provide meta-information about a request. The request body contains the data that client wants to send to the server. For example, Content-Type and Content-Transfer-Encoding.  HTTP response headers provide meta-information about a response. The response consists of status code (three-digit number), content type, and body of the response.  | Item                | Description                                                                      | | ------------------- | :------------------------------------------------------------------------------- | | 1xx - Informational | Communicates transfer protocol-level information.                                | | 2xx - Success       | Indicates that the client request is successfully accepted.                      | | 3xx - Redirection   | Indicates that the client must take some additional steps to complete an action. | | 4xx - Client Error  | Indicates that there is some error from the client side.                         | | 5xx - Server Error  | Indicates that there is some error from the server side.                         |  ## Entity Structure  An entity structure comprises of the following components.  - **Metadata** - Metadata provides high-level information about the related entity including entity uuid, kind, and spec_version. - **Status** - Status provides information about the current state of the related entity including entity name, state, and entity-specific status information. - **Specification** - For a GET request, the specification section of the JSON response displays similar information as status section. However, for of POST request, the specification section provide details of what should be the desired state of the entity. For example, you can use the following JSON payload to create a basic VM. The payload denotes the minimum parameters required to create a VM by using the V3 APIs. All the missing specification items like CPU and RAM are configured with default values.  ``` { \"spec\":{ \"name\":\"NewVM\", \"resources\":{ } }, \"metadata\":{ \"kind\":\"vm\" } } ```  # Authentication  All requests to the Nutanix REST APIs must be authenticated. The following authentication types are supported.  - **Basic Authentication**: The user provides username and password every time a request is sent as the authentication header. - **Session Authentication**: The user credentials are stored in a cookie.  **Cookie Consumption**  To reduce the number of cookies generated in the system, Nutanix recommends that you make the first API call using basic credentials to receive the cookies and then instead of sending the credentials, cache and send the following cookies to authenticate subsequent requests.  **For Prism Central with IAMv2 enabled:**  - NTNX_MERCURY_IAM_SESSION - NTNX_IAM_SESSION - NTNX_MERCURY_IAM_REFRESH_TOKEN  **For Prism Central with IAMv2 disabled or Prism Element:**  - NTNX_IGW_SESSION - NTNX_MERCURY_IGW_SESSION - JSESSIONID(If present)  **Note:** Upon cookie expiration, you must send basic credentials again in the next API call, to obtain the new cookies.  The following are the code snippets to authenticate a user.  **Python**  ``` #!/usr/bin/env python3.7 import urllib3 import requests from base64 import b64encode urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning) request_url = 'https://10.0.0.1:9440/api/nutanix/v3/vms/list' username = 'admin' password = 'password' encoded_credentials = b64encode(bytes(f'{username}:{password}', encoding='ascii')).decode('ascii') auth_header = f'Basic {encoded_credentials}' payload = '{\"kind\":\"vm\"}' headers = {'Accept': 'application/json', 'Content-Type': 'application/json', 'Authorization': f'{auth_header}', 'cache-control': 'no-cache'} response = requests.request('post', request_url, data=payload, headers=headers, verify=False) print(response.status_code) ```  **C#**  ``` using System; using System.Net; using System.Text; namespace ConsoleApp1 { class Program { static void Main(string[] args) { string ClusterUsername = \"admin\"; string ClusterPassword = \"Ntnx2017!\"; string RequestUrl = \"https://192.168.1.131:9440/api/nutanix/v3/vms/list\"; string Payload = \"{\\\"kind\\\":\\\"vm\\\"}\"; string AuthHeader = System.Convert.ToBase64String( System.Text.ASCIIEncoding.ASCII.GetBytes( ClusterUsername + \":\" + ClusterPassword)); ServicePointManager.ServerCertificateValidationCallback = ( (sender, certificate, chain, sslPolicyErrors) => true); ServicePointManager.Expect100Continue = true; ServicePointManager.SecurityProtocol = SecurityProtocolType.Tls12; String authHeader = System.Convert.ToBase64String( System.Text.ASCIIEncoding.ASCII.GetBytes( ClusterUsername + \":\" + ClusterPassword)); var request = (HttpWebRequest)WebRequest.Create(RequestUrl); var requestBody = Encoding.ASCII.GetBytes(Payload); HttpWebResponse HttpResponse = null; request.Headers.Add(\"Authorization\", \"Basic \" + AuthHeader); request.Headers.Add(\"Accept-Encoding\", \"application/json\"); request.ContentType = \"application/json\"; request.Accept = \"application/json\"; request.Method = \"POST\"; var newStream = request.GetRequestStream(); newStream.Write(requestBody, 0, requestBody.Length); newStream.Close(); HttpResponse = (HttpWebResponse)request.GetResponse(); Console.WriteLine(HttpResponse.StatusCode); Console.ReadKey(); } } } ```  **PowerShell**  ``` # set the basic properties for the request $RequestUri = \"https://10.0.0.1:9440/api/nutanix/v3/vms/list\" $Payload = \"{\"\"kind\"\":\"\"vm\"\"}\" $Username = \"admin\" $Password = \"password\" # create the HTTP Basic Authorization header $pair = $Username + \":\" + $Password $bytes = [System.Text.Encoding]::ASCII.GetBytes($pair) $base64 = [System.Convert]::ToBase64String($bytes) $basicAuthValue = \"Basic $base64\" # setup the request headers $Headers = @{ 'Accept' = 'application/json' 'Authorization' = $basicAuthValue 'Content-Type' = 'application/json' } # disable SSL certification verification # you probably shouldn't do this in production ... if (-not ( [System.Management.Automation.PSTypeName] ` 'ServerCertificateValidationCallback').Type) {$certCallback = @\" using System; using System.Net; using System.Net.Security; using System.Security.Cryptography.X509Certificates; public class ServerCertificateValidationCallback { public static void Ignore() { if(ServicePointManager.ServerCertificateValidationCallback ==null) { ServicePointManager.ServerCertificateValidationCallback += delegate ( Object obj, X509Certificate certificate, X509Chain chain, SslPolicyErrors errors ) { return true; }; } } } \"@ Add-Type $certCallback } [ServerCertificateValidationCallback]::Ignore() [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12 # submit the request Invoke-WebRequest ` -Uri $RequestUri ` -Headers $Headers ` -Method \"POST\" ` -Body $Payload ` -TimeoutSec 5 ` -UseBasicParsing ` -DisableKeepAlive ` | ConvertFrom-Json | Select * ```  **PHP**  ``` <?php $username = 'admin'; $password = 'password'; $ch = curl_init(); curl_setopt_array($ch, [ CURLOPT_RETURNTRANSFER => 1, CURLINFO_HEADER_OUT => TRUE, CURLOPT_URL => 'https://10.0.0.1:9440/api/nutanix/v3/vms/list', CURLOPT_USERPWD => $username . ':' . $password, CURLOPT_POST => 1, CURLOPT_SSL_VERIFYHOST => FALSE, CURLOPT_SSL_VERIFYPEER => FALSE, CURLOPT_HTTPHEADER => [ 'Content-Type: application/json', 'Accept-Encoding: application/json', ], CURLOPT_POSTFIELDS => json_encode([ 'kind' => 'vm' ]) ]); $result = curl_exec($ch); curl_close($ch); if(!$result) { die('Error: \"' . curl_error($ch) . '\". Code: ' . curl_errno($ch)); } else { echo($result); } ``` # pc.2022.1 release  ## v3 API behavior changes with pc.2022.1 release   From Prism Central release 2022.1 onwards, the v3 intentful API is enhanced to protect data consistency, especially in concurrent race conditions. As a result of this enhancement, response time for the synchronous API requests is improved and there are some behavior changes from an API consumption perspective.   1.  In the current v3 API workflow, clients will get a synchronous API response that contains task and entity UUID. Post which clients should send GET API on the task periodically until the task succeeds or fails. Clients can also send GET API on the entity.         - **Behavioral changes**: Some of the synchronous errors will be reported asynchronously now.            1. **Invalid idempotency UUID error**:                       **Current API behavior**: If POST API contains invalid idempotency UUID, clients will receive a     synchronous bad request error with an HTTP code 400.           **Expected change**: If POST API contains invalid idempotency UUID, clients will get a synchronous API response with an HTTP code 202(accepted). GET API on the task will return a bad request error.      1. **Authorization errors**:            **Current API behavior**: Clients will get synchronous authorization errors (for example- HTTP code 403 forbidden error).            **Expected change**: Clients will get a synchronous API response with an HTTP code 202 (accepted). GET API on the task will show an authorization error.       1. **Incorrect category mapping error**:            **Current API behavior**: If an intentful API payload has a project reference and uses categories mapping but the category for the project does not exist, the client will get an error response with an HTTP code 404(not found).            **Expected change**: If an intentful API payload has a project reference and uses categories mapping but the category for the project does not exist, the client will receive a synchronous API response with an HTTP code 202 (accepted). GET API on the task will show a “not found” error.      1.  **Invalid user detail error**:            **Current API behavior**: If an intentful API payload has user_uuid or username in owner_reference but the corresponding user does not exist, the client will receive an API response with an HTTP code 404.            **Expected change**: If an intentful API payload has user_uuid or username in owner_reference but the corresponding user does not exist, clients will get a synchronous API response with an HTTP code 202(accepted). GET API on the task will show a bad request error.         1.  . **Missing user detail error**:           **Current API behavior**: If an intentful API payload does not have user_uuid or username in owner_reference and the login user does not exist, the client will get an API response with an HTTP code 404.           **Expected change**: If an intentful API payload does not have user_uuid or username in owner_reference and login user does not exist, clients will get a synchronous API response with an HTTP code 202(accepted). GET API on the task will show a bad request error.      -  **Format changes** for some of the synchronous API responses:                  **Current API behavior**: Response to an intentful POST, PUT, or DELETE API will contain tenant_uuid, owner_uuid, owner_username, service_name, session_log_uuid, trace, user_uuid, username, remote_authorization, session_json_dict, incap_req_id, project_reference or category information.                  **Expected change**: Main processing will occur after a synchronous response to the client. As a result,  synchronous response to an intentful POST, PUT, or DELETE API will not contain tenant_uuid, owner_uuid, owner_username, service_name, session_log_uuid, trace, user_uuid, username, remote_authorization, session_json_dict, incap_req_id, project_reference or category information.                  If API consumers are expecting these properties in the synchronous response before, they have to submit a GET/LIST request to inspect these properties.  2.  The version control workflow for v3 intentful PUT and DELETE APIs:    - **Absolutely no version control check for DELETE APIs**:                  **Current API behavior**: DELETE API will fail if an incorrect spec_version is provided as one of the request arguments.                  **Expected change**:  DELETE API will always delete the entity. No version control is applicable on the DELETE request.       - **Edit conflict error could be reported asynchronously**:        **Current API behavior**: If v3 PUT APIs are conflicting with v2 PUT APIs on the same entities, the client will receive an HTTP code 409 response for v3 PUT APIs.       **Expected change**: If v3 PUT APIs are conflicting with v2 PUT APIs on the same entities, the client will receive an HTTP code 202 response for v3 PUT APIs and will receive the task failed response with an edit conflict error asynchronously.     - **More restrictive version control by using spec_hash**:                 **Current API behavior**: GET/LIST API will return a response that may contain one of the values for version control, it could be spec_version or spec_hash. spec_version is numeric while spec_hash is a string.                  Whichever value is provided in the GET/LIST API, the same value should be included in the next PUT API payload of the same entity.        **Expected change**: GET/LIST API will provide both spec_hash and spec_version in API response. Use of spec_hash is not enforced but could be enforced in future releases.        - **The version control value in the GET API response will not change when a PUT request is in progress**:                  **Current API behavior**: If there are one GET request and one PUT request on the same entity, the version control value from the GET request response might be different depending on the current state of the entity. The behavior is not deterministic.          **Expected change**: If there is one GET request and one PUT request on the same entity, the spec_hash value of the GET request response will always be the same as it was before the PUT request until the PUT is completed.    1. V3 intentful POST API and the following GET API:    -  **Get API request for new entity**:        **Current API behavior**: If a POST API fails asynchronously, the following GET API on the same entity will get an HTTP code 202 response with the entity in an error state.          However, GET API on the same entity after 5+ minutes will show an HTTP code 404 response reporting the entity does not exist.        **Expected change**: If a POST API fails asynchronously, GET API on the same entity will always show an HTTP code 404 response meaning the entity does not exist.        - **The intent spec is an entity we created to record previous and ongoing requests from v3 POST/PUT/DELETE APIs**. There should be only one intent spec for each entity if the intent spec exists:          **Current API behavior:** If the client queries intent spec for an entity immediately after getting an HTTP code 202 response for the v3 POST API on this entity, the intent spec will always be present.          **Expected change**: If the client queries intent spec for an entity immediately after getting an HTTP code 202 response for the v3 POST API on this entity, the intent spec may not be created yet.  4. Prism Central upgrade to pc.2022.1 and following releases:        **Current API behavior**: The pending and running API requests will continue to execute after an upgrade.        **Expected change**: During the pc.2022.1 version upgrade, all pending and running API requests and related tasks will be marked as failed after a restart. This will be shown even though they may be in progress and completed in the middle of the upgrade period.       This behavior only applies to upgrades from releases before pc.2022.1 to releases equal to or after pc.2022.1. All future upgrades from pc.2022.1 release will remain the same as before where pending and running API requests will continue to execute after an upgrade.  5. Multiple nodes upgrade:   -   pc.2022.1 version can take care of the scenarios where Prism Central has the new version and Prism Element is still at the older version.        -   Any task (for existing create, update, delete v3 API operations) that is in running state during the upgrade will be marked failed. This also applies to new requests that come before the upgrade ends. The recommendation is to wait until the upgrade finishes before executing any v3 create, update, delete API.     6. Schema consistency:    - API schema in Prism Central will be used in case there is a      discrepancy between Prism Element and Prism Central. There will be a      change of behavior when Prism Central has a post pc.2022.1 version.        **Current API behavior**: API responses for POST, PUT, and DELETE API on Prism Central will match API schema on Prism Element.            **Expected change**: API response for POST, PUT, and DELETE API on Prism Central will match API schema on Prism Central.  # Use Cases            ## Post /VMs  Perform the following procedure to generate the code snippet for creating a VM based on the input parameters.  1. Under **VMs**, click **POST VMs** and navigate to the end of the page. 2. In the **Code Generation** panel, under the **Settings** tab, enter the value of the following parameters.    1. Enter the username of the Prism Central in the **Username** field.    2. Enter the password in the **Password** field.    3. Enter the IP address of the Prism Central in the **Host** field. 3. Under the **Body** tab, enter the value of all the required parameters as described in the **Request Body** section. Optionally, you can also enter the values of the optional parameters.    1. Under **Specification**, enter the description of VM in a string format against the **description** attribute. The maximum length allowed is 1000 characters only.    2. Under **Resources** do the following.       1. Enter the value of number of logical threads per core in an integer format against the **num_threads_per_core** attribute. The minimum value allowed is 1.       2. Enter the value of current or desired power state of the VM in a string format against the **power_state** attribute.       3. Enter the value of number of vCPUs per socket in an integer format against the **num_vcpus_per_socket** attribute. The minimun value allowed is 1.       4. Enter the value of number of vCPU sockets in an integer format against the **num_sockets** attribute. The minimum value allowed is 1.    3. Under **Name**, enter the name of the VM in a string format. The maximum length allowed is 64 characters only.    4. Under **Metadata** do the following.       1. The **last_update_time** is a read-only attribute indicating UTC date and time in an RFC-339 format, when the VM was last updated in a string format.       2. The **kind** is a read-only attribute indicating the kind name in a string format. The default value is VM.       3. Enter the value of VM uuid in a string format against the **uuid** attribute. The valid pattern is ^\\[a-fA-F0-9]{8}-\\[a-fA-F0-9]{4}-\\[a-fA-F0-9]{4}-\\[a-fA-F0-9]{4}-\\[a-fAF0-9]{12}$. 4. Under the **Code Generation** tab, select the scripting languages to view the code snippet of creating a VM.     You can either select Python or PowerShell or Shell or Go or Java or JavaScript or Node or Obj-C or PHP or Ruby or Swift or C# or C.     The code snippet appears in the selected scripting language. You can use the code snippet to create a VM in your environment.  ## Put /VMs  Perform the following procedure to generate the code snippet for updating an existing VM based on the input parameters.  1. Under **VMs**, click **PUT VMs** and navigate to the end of the page. 2. In the **Code Generation** panel, under the **Settings** tab, enter the value of the following parameters.    1. Enter the UUID of the VM that needs to be updated in the **UUID** field.    2. Enter the username of the Prism Central in the **Username** field.    3. Enter the password in the **Password** field.    4. Enter the IP address of the Prism Central in the **Host** field. 3. Under the **Body** tab, enter the value of all the required parameters as described in the **Request Body** section. Optionally, you can also enter the values of the optional parameters.    1. Under **Specification**, enter the description of VM in a string format against the **description** attribute. The maximum length allowed is 1000 characters only.    2. Under **Resources** do the following.       1. Enter the value of number of logical threads per core in an integer format against the **num_threads_per_core** attribute. The minimum value allowed is 1.       2. Enter the value of current or desired power state of the VM in a string format against the **power_state** attribute.       3. Enter the value of number of vCPUs per socket in an integer format against the **num_vcpus_per_socket** attribute. The minimun value allowed is 1.       4. Enter the value of number of vCPU sockets in an integer format against the **num_sockets** attribute. The minimum value allowed is 1.    3. Under **Name**, enter the name of the VM in a string format. The maximum length allowed is 64 characters only.    4. Under **Metadata** do the following.       1. The **last_update_time** is a read-only attribute indicating UTC date and time in an RFC-339 format, when the VM was last updated in a string format.       2. The **kind** is a read-only attribute indicating the kind name in a string format. The default value is VM.       3. Enter the value of VM uuid in a string format against the **uuid** attribute. The valid pattern is ^\\[a-fA-F0-9]{8}-\\[a-fA-F0-9]{4}-\\[a-fA-F0-9]{4}-\\[a-fA-F0-9]{4}-\\[a-fAF0-9]{12}$. 4. Under the **Code Generation** tab, select the scripting languages to view the code snippet of creating a VM.     You can either select Python or PowerShell or Shell or Go or Java or JavaScript or Node or Obj-C or PHP or Ruby or Swift or C# or C.     The code snippet appears in the selected scripting language. You can use the code snippet to update a VM in your environment. 

API version: 3.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// checks if the WhatifScenario type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WhatifScenario{}

// WhatifScenario Scenario Object.
type WhatifScenario struct {
	// The flag to indicate whether it is a new cluster or not.
	NewCluster *bool `json:"new_cluster,omitempty"`
	// The entity type for the cluster e.g. cluster or nutanix_vcenter__cluster.
	ClusterEntityType *string `json:"cluster_entity_type,omitempty"`
	// The uuid would be automatically generated when created.
	UUID *string `json:"uuid,omitempty" validate:"regexp=^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$"`
	VendorList []string `json:"vendor_list,omitempty"`
	// workload added by user.
	WorkloadList []Workload `json:"workload_list,omitempty"`
	RecommendedRunway *Runway `json:"recommended_runway,omitempty"`
	// Last updated timestamp.
	UpdatedTimeSec *int32 `json:"updated_time_sec,omitempty"`
	// The cluster uuid.
	ClusterUuid *string `json:"cluster_uuid,omitempty" validate:"regexp=^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$"`
	// The target runway.
	TargetRunwayDays *int32 `json:"target_runway_days,omitempty"`
	ClusterSpec *ClusterSpec `json:"cluster_spec,omitempty"`
	Runway *Runway `json:"runway,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WhatifScenario WhatifScenario

// NewWhatifScenario instantiates a new WhatifScenario object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatifScenario() *WhatifScenario {
	this := WhatifScenario{}
	var clusterEntityType string = "cluster"
	this.ClusterEntityType = &clusterEntityType
	var targetRunwayDays int32 = 180
	this.TargetRunwayDays = &targetRunwayDays
	return &this
}

// NewWhatifScenarioWithDefaults instantiates a new WhatifScenario object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatifScenarioWithDefaults() *WhatifScenario {
	this := WhatifScenario{}
	var clusterEntityType string = "cluster"
	this.ClusterEntityType = &clusterEntityType
	var targetRunwayDays int32 = 180
	this.TargetRunwayDays = &targetRunwayDays
	return &this
}

// GetNewCluster returns the NewCluster field value if set, zero value otherwise.
func (o *WhatifScenario) GetNewCluster() bool {
	if o == nil || IsNil(o.NewCluster) {
		var ret bool
		return ret
	}
	return *o.NewCluster
}

// GetNewClusterOk returns a tuple with the NewCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatifScenario) GetNewClusterOk() (*bool, bool) {
	if o == nil || IsNil(o.NewCluster) {
		return nil, false
	}
	return o.NewCluster, true
}

// HasNewCluster returns a boolean if a field has been set.
func (o *WhatifScenario) HasNewCluster() bool {
	if o != nil && !IsNil(o.NewCluster) {
		return true
	}

	return false
}

// SetNewCluster gets a reference to the given bool and assigns it to the NewCluster field.
func (o *WhatifScenario) SetNewCluster(v bool) {
	o.NewCluster = &v
}

// GetClusterEntityType returns the ClusterEntityType field value if set, zero value otherwise.
func (o *WhatifScenario) GetClusterEntityType() string {
	if o == nil || IsNil(o.ClusterEntityType) {
		var ret string
		return ret
	}
	return *o.ClusterEntityType
}

// GetClusterEntityTypeOk returns a tuple with the ClusterEntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatifScenario) GetClusterEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterEntityType) {
		return nil, false
	}
	return o.ClusterEntityType, true
}

// HasClusterEntityType returns a boolean if a field has been set.
func (o *WhatifScenario) HasClusterEntityType() bool {
	if o != nil && !IsNil(o.ClusterEntityType) {
		return true
	}

	return false
}

// SetClusterEntityType gets a reference to the given string and assigns it to the ClusterEntityType field.
func (o *WhatifScenario) SetClusterEntityType(v string) {
	o.ClusterEntityType = &v
}

// GetUUID returns the UUID field value if set, zero value otherwise.
func (o *WhatifScenario) GetUUID() string {
	if o == nil || IsNil(o.UUID) {
		var ret string
		return ret
	}
	return *o.UUID
}

// GetUUIDOk returns a tuple with the UUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatifScenario) GetUUIDOk() (*string, bool) {
	if o == nil || IsNil(o.UUID) {
		return nil, false
	}
	return o.UUID, true
}

// HasUUID returns a boolean if a field has been set.
func (o *WhatifScenario) HasUUID() bool {
	if o != nil && !IsNil(o.UUID) {
		return true
	}

	return false
}

// SetUUID gets a reference to the given string and assigns it to the UUID field.
func (o *WhatifScenario) SetUUID(v string) {
	o.UUID = &v
}

// GetVendorList returns the VendorList field value if set, zero value otherwise.
func (o *WhatifScenario) GetVendorList() []string {
	if o == nil || IsNil(o.VendorList) {
		var ret []string
		return ret
	}
	return o.VendorList
}

// GetVendorListOk returns a tuple with the VendorList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatifScenario) GetVendorListOk() ([]string, bool) {
	if o == nil || IsNil(o.VendorList) {
		return nil, false
	}
	return o.VendorList, true
}

// HasVendorList returns a boolean if a field has been set.
func (o *WhatifScenario) HasVendorList() bool {
	if o != nil && !IsNil(o.VendorList) {
		return true
	}

	return false
}

// SetVendorList gets a reference to the given []string and assigns it to the VendorList field.
func (o *WhatifScenario) SetVendorList(v []string) {
	o.VendorList = v
}

// GetWorkloadList returns the WorkloadList field value if set, zero value otherwise.
func (o *WhatifScenario) GetWorkloadList() []Workload {
	if o == nil || IsNil(o.WorkloadList) {
		var ret []Workload
		return ret
	}
	return o.WorkloadList
}

// GetWorkloadListOk returns a tuple with the WorkloadList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatifScenario) GetWorkloadListOk() ([]Workload, bool) {
	if o == nil || IsNil(o.WorkloadList) {
		return nil, false
	}
	return o.WorkloadList, true
}

// HasWorkloadList returns a boolean if a field has been set.
func (o *WhatifScenario) HasWorkloadList() bool {
	if o != nil && !IsNil(o.WorkloadList) {
		return true
	}

	return false
}

// SetWorkloadList gets a reference to the given []Workload and assigns it to the WorkloadList field.
func (o *WhatifScenario) SetWorkloadList(v []Workload) {
	o.WorkloadList = v
}

// GetRecommendedRunway returns the RecommendedRunway field value if set, zero value otherwise.
func (o *WhatifScenario) GetRecommendedRunway() Runway {
	if o == nil || IsNil(o.RecommendedRunway) {
		var ret Runway
		return ret
	}
	return *o.RecommendedRunway
}

// GetRecommendedRunwayOk returns a tuple with the RecommendedRunway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatifScenario) GetRecommendedRunwayOk() (*Runway, bool) {
	if o == nil || IsNil(o.RecommendedRunway) {
		return nil, false
	}
	return o.RecommendedRunway, true
}

// HasRecommendedRunway returns a boolean if a field has been set.
func (o *WhatifScenario) HasRecommendedRunway() bool {
	if o != nil && !IsNil(o.RecommendedRunway) {
		return true
	}

	return false
}

// SetRecommendedRunway gets a reference to the given Runway and assigns it to the RecommendedRunway field.
func (o *WhatifScenario) SetRecommendedRunway(v Runway) {
	o.RecommendedRunway = &v
}

// GetUpdatedTimeSec returns the UpdatedTimeSec field value if set, zero value otherwise.
func (o *WhatifScenario) GetUpdatedTimeSec() int32 {
	if o == nil || IsNil(o.UpdatedTimeSec) {
		var ret int32
		return ret
	}
	return *o.UpdatedTimeSec
}

// GetUpdatedTimeSecOk returns a tuple with the UpdatedTimeSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatifScenario) GetUpdatedTimeSecOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedTimeSec) {
		return nil, false
	}
	return o.UpdatedTimeSec, true
}

// HasUpdatedTimeSec returns a boolean if a field has been set.
func (o *WhatifScenario) HasUpdatedTimeSec() bool {
	if o != nil && !IsNil(o.UpdatedTimeSec) {
		return true
	}

	return false
}

// SetUpdatedTimeSec gets a reference to the given int32 and assigns it to the UpdatedTimeSec field.
func (o *WhatifScenario) SetUpdatedTimeSec(v int32) {
	o.UpdatedTimeSec = &v
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *WhatifScenario) GetClusterUuid() string {
	if o == nil || IsNil(o.ClusterUuid) {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatifScenario) GetClusterUuidOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterUuid) {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *WhatifScenario) HasClusterUuid() bool {
	if o != nil && !IsNil(o.ClusterUuid) {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *WhatifScenario) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetTargetRunwayDays returns the TargetRunwayDays field value if set, zero value otherwise.
func (o *WhatifScenario) GetTargetRunwayDays() int32 {
	if o == nil || IsNil(o.TargetRunwayDays) {
		var ret int32
		return ret
	}
	return *o.TargetRunwayDays
}

// GetTargetRunwayDaysOk returns a tuple with the TargetRunwayDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatifScenario) GetTargetRunwayDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.TargetRunwayDays) {
		return nil, false
	}
	return o.TargetRunwayDays, true
}

// HasTargetRunwayDays returns a boolean if a field has been set.
func (o *WhatifScenario) HasTargetRunwayDays() bool {
	if o != nil && !IsNil(o.TargetRunwayDays) {
		return true
	}

	return false
}

// SetTargetRunwayDays gets a reference to the given int32 and assigns it to the TargetRunwayDays field.
func (o *WhatifScenario) SetTargetRunwayDays(v int32) {
	o.TargetRunwayDays = &v
}

// GetClusterSpec returns the ClusterSpec field value if set, zero value otherwise.
func (o *WhatifScenario) GetClusterSpec() ClusterSpec {
	if o == nil || IsNil(o.ClusterSpec) {
		var ret ClusterSpec
		return ret
	}
	return *o.ClusterSpec
}

// GetClusterSpecOk returns a tuple with the ClusterSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatifScenario) GetClusterSpecOk() (*ClusterSpec, bool) {
	if o == nil || IsNil(o.ClusterSpec) {
		return nil, false
	}
	return o.ClusterSpec, true
}

// HasClusterSpec returns a boolean if a field has been set.
func (o *WhatifScenario) HasClusterSpec() bool {
	if o != nil && !IsNil(o.ClusterSpec) {
		return true
	}

	return false
}

// SetClusterSpec gets a reference to the given ClusterSpec and assigns it to the ClusterSpec field.
func (o *WhatifScenario) SetClusterSpec(v ClusterSpec) {
	o.ClusterSpec = &v
}

// GetRunway returns the Runway field value if set, zero value otherwise.
func (o *WhatifScenario) GetRunway() Runway {
	if o == nil || IsNil(o.Runway) {
		var ret Runway
		return ret
	}
	return *o.Runway
}

// GetRunwayOk returns a tuple with the Runway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatifScenario) GetRunwayOk() (*Runway, bool) {
	if o == nil || IsNil(o.Runway) {
		return nil, false
	}
	return o.Runway, true
}

// HasRunway returns a boolean if a field has been set.
func (o *WhatifScenario) HasRunway() bool {
	if o != nil && !IsNil(o.Runway) {
		return true
	}

	return false
}

// SetRunway gets a reference to the given Runway and assigns it to the Runway field.
func (o *WhatifScenario) SetRunway(v Runway) {
	o.Runway = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WhatifScenario) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatifScenario) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WhatifScenario) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WhatifScenario) SetName(v string) {
	o.Name = &v
}

func (o WhatifScenario) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WhatifScenario) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewCluster) {
		toSerialize["new_cluster"] = o.NewCluster
	}
	if !IsNil(o.ClusterEntityType) {
		toSerialize["cluster_entity_type"] = o.ClusterEntityType
	}
	if !IsNil(o.UUID) {
		toSerialize["uuid"] = o.UUID
	}
	if !IsNil(o.VendorList) {
		toSerialize["vendor_list"] = o.VendorList
	}
	if !IsNil(o.WorkloadList) {
		toSerialize["workload_list"] = o.WorkloadList
	}
	if !IsNil(o.RecommendedRunway) {
		toSerialize["recommended_runway"] = o.RecommendedRunway
	}
	if !IsNil(o.UpdatedTimeSec) {
		toSerialize["updated_time_sec"] = o.UpdatedTimeSec
	}
	if !IsNil(o.ClusterUuid) {
		toSerialize["cluster_uuid"] = o.ClusterUuid
	}
	if !IsNil(o.TargetRunwayDays) {
		toSerialize["target_runway_days"] = o.TargetRunwayDays
	}
	if !IsNil(o.ClusterSpec) {
		toSerialize["cluster_spec"] = o.ClusterSpec
	}
	if !IsNil(o.Runway) {
		toSerialize["runway"] = o.Runway
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WhatifScenario) UnmarshalJSON(data []byte) (err error) {
	varWhatifScenario := _WhatifScenario{}

	err = json.Unmarshal(data, &varWhatifScenario)

	if err != nil {
		return err
	}

	*o = WhatifScenario(varWhatifScenario)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "new_cluster")
		delete(additionalProperties, "cluster_entity_type")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "vendor_list")
		delete(additionalProperties, "workload_list")
		delete(additionalProperties, "recommended_runway")
		delete(additionalProperties, "updated_time_sec")
		delete(additionalProperties, "cluster_uuid")
		delete(additionalProperties, "target_runway_days")
		delete(additionalProperties, "cluster_spec")
		delete(additionalProperties, "runway")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWhatifScenario struct {
	value *WhatifScenario
	isSet bool
}

func (v NullableWhatifScenario) Get() *WhatifScenario {
	return v.value
}

func (v *NullableWhatifScenario) Set(val *WhatifScenario) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatifScenario) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatifScenario) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatifScenario(val *WhatifScenario) *NullableWhatifScenario {
	return &NullableWhatifScenario{value: val, isSet: true}
}

func (v NullableWhatifScenario) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatifScenario) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


