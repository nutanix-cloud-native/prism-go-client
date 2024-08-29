/*
Nutanix Intentful API

# Introduction  **Representational state transfer ( REST )** Application Programming Interface (API) 3.0 is based on an intentful API philosophy. According to the intentful API philosophy the machine should handle the programming instead of the user enabling the datacenter administrator able to focus on the other tasks. You need to specify the end state of an entity and the system will compile and execute a series of steps to achieve the defined end state of the entity. The progress to achieve the desired state is tracked through waits and events. This approach is similar to the Google Kubernetes standard.  All the entities and the list of entities follow a homogenous specification format resulting in ease of programming despite different entities being involved. Also, this enables the user to get familiar with the APIs quickly.  Intentful APIs reduces the number of action verbs required to achieve the desired state of an entity. Most of the changes can be achieved by defining the desired state of an entity rather than through a series of action verbs.  All API users should have valid Prism login credentials to send API calls to the Prism server.  APIs on nutanix.dev portal are publicly available to all valid users without special permissions for viewing purposes.  This document covers APIs that are available in pc2022.11 release.    ## Terminology  | Item                  | Description                                                                                                                                                             | | --------------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | | Entity                | An object that is managed through an API.                                                                                                                               | | Kind                  | Represents the type of entity. For example, VM or alerts.                                                                                                               | | Resource Kind         | An entity that represents a physical or virtual infrastructure resource. For example, VM or an application.                                                             | | Helper Kind           | An entity that is not an infrastructure resource, but represents the entity that is used along with the infrastructure resources. For example, alert, event, or status. | | Basic Resource Kind   | An entity type that is a primary resource. For example, VM.                                                                                                             | | Related Resource Kind | An entity type that is automatically defined by the system and is related to the basic Kind entities. For example, vm_snapshot.                                         | | Type                  | Represents sub-categories or sub-objects of the top-level resource types. For example, reference.                                                                       | | Resource Limit        | An entity type that represents a quota of resources of different types.                                                                                                 | | Snapshot              | Represents the state of resources and policies of an entity at a particular point of time.                                                                              | | Backup                | An automatically created aggregation object that collects all the snapshots for the same object within the same resource pool.                                          | | Profile               | A profile is a set of resources and policy specifications that is applied to an entity.                                                                                 | | Template              | A template is a specification that generates valid entities by providing input parameters.                                                                              | | Spec                  | A section in the request json, which represents an entity, that expresses the desired state of the entity.                                                              | | Status                | A section in the json representation that expresses the current state of the entity.                                                                                    |  ## Rules  - The metadata section has fields that relate to all kinds of entities and can be updated instantly by the system. - The spec section cannot be updated after initialization by the system without a user or automation intervention. - The status section has a copy of the currently enforced version of the spec section and other fields that are managed or updated automatically by the system. Few fields in the status section can be updated by the user as well. - The parent_reference field is mandatory for snapshot and backup entities and is present in the normal entities, if the entities are created with the parent_reference field populated (clone or restore). - Any field that represents a defined type has the type name as a suffix. For example, backup_factory and vm_reference.  ## Known Issues and limitations:  This section describes some of the issues identified in V3 APIs that you might encounter.  - **ENG-321951**- While attempting to get a list of all VMs on Prism Central, we have observed that only AHV VMs are returned, but ESXi VM list is not retrieved.  - **ENG-453487**- Users with insufficient permissions receive an HTTP 200 response instead of an HTTP 403 response for view alert API requests. These users can be like \"Developer\", \"Operator\", or custom roles that do not have sufficient permission to view alerts. This is a generic issue for all V3 API requests with insufficient permissions.   ### [Nutanix API versions – What are they and what does each one do?](https://developer.nutanix.com/2019/01/15/nutanix-api-versions-what-are-they-and-what-does-each-one-do/)   ### [What to expect post pc.2022.1 release?](https://www.nutanix.dev/api_references/prism-central-v3/#/ZG9jOjM2MzQyMjUz-v3-api-behavior-changes-with-pc-2022-1-release)  # Getting Started  ## Components of a REST API  A REST API request and response can be classified into the following components.  The request URI consists of {URI-scheme} :// {URI-host} / {<Port-number>/<Base-path>} ? {query-string}.  | Item                    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | | ----------------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | | URI scheme              | Indicates the protocol that is used to transmit the request. For example, http or https.                                                                                                                                                                                                                                                                                                                                                                                                                                    | | URI host                | Specifies the domain name or IP address of the server where the REST service endpoint is hosted. For example, prism.nutanix.com or 10.11.12.13.                                                                                                                                                                                                                                                                                                                                                                             | | Port number             | Indicates the port number that is used for HTTP communications from the REST API clients to the REST API server.                                                                                                                                                                                                                                                                                                                                                                                                            | | Basepath                | Indicates the URL prefix that is common to all the API server endpoints. For example, for <https://10.11.12.13/PrismGateway/services/rest/v2.0/vms>, the basepath is /PrismGateway/services/rest/v2.0.                                                                                                                                                                                                                                                                                                                      | | Path parameters         | Path parameters are part of the endpoint. For example, <https://10.11.12.13/PrismGateway/services/rest/v2.0/vms/750eb5ba-2de0-4ca4-8d6f-786364817f94/>, the UUID of the VM (750eb5ba-2de0-4ca4-8d6f-786364817f94) is a path parameter.                                                                                                                                                                                                                                                                                      | | Query string parameters | Query string parameters appear after a question mark (?) in the endpoint. The question mark followed by the parameters and their values is referred as the query string. In the query string, each parameter is listed one right after the other with an ampersand (&). The order of the query string parameters does not matter. For example, <https://10.11.12.13/PrismGateway/services/rest/v2.0/vms/750eb5ba-2de0-4ca4-8d6f-786364817f94?include_vm_disk_config=true>, the include_vm_disk_config is a query parameter. |  HTTP defines methods to indicate the desired action to be performed on the identified resource.  | Method | Description                                    | | ------ | :--------------------------------------------- | | Get    | Retrieves a resource                           | | Post   | Creates a resource                             | | Put    | Updates or creates within an existing resource | | Delete | Removes the resource                           |  HTTP request headers provide meta-information about a request. The request body contains the data that client wants to send to the server. For example, Content-Type and Content-Transfer-Encoding.  HTTP response headers provide meta-information about a response. The response consists of status code (three-digit number), content type, and body of the response.  | Item                | Description                                                                      | | ------------------- | :------------------------------------------------------------------------------- | | 1xx - Informational | Communicates transfer protocol-level information.                                | | 2xx - Success       | Indicates that the client request is successfully accepted.                      | | 3xx - Redirection   | Indicates that the client must take some additional steps to complete an action. | | 4xx - Client Error  | Indicates that there is some error from the client side.                         | | 5xx - Server Error  | Indicates that there is some error from the server side.                         |  ## Entity Structure  An entity structure comprises of the following components.  - **Metadata** - Metadata provides high-level information about the related entity including entity uuid, kind, and spec_version. - **Status** - Status provides information about the current state of the related entity including entity name, state, and entity-specific status information. - **Specification** - For a GET request, the specification section of the JSON response displays similar information as status section. However, for of POST request, the specification section provide details of what should be the desired state of the entity. For example, you can use the following JSON payload to create a basic VM. The payload denotes the minimum parameters required to create a VM by using the V3 APIs. All the missing specification items like CPU and RAM are configured with default values.  ``` { \"spec\":{ \"name\":\"NewVM\", \"resources\":{ } }, \"metadata\":{ \"kind\":\"vm\" } } ```  # Authentication  All requests to the Nutanix REST APIs must be authenticated. The following authentication types are supported.  - **Basic Authentication**: The user provides username and password every time a request is sent as the authentication header. - **Session Authentication**: The user credentials are stored in a cookie.  **Cookie Consumption**  To reduce the number of cookies generated in the system, Nutanix recommends that you make the first API call using basic credentials to receive the cookies and then instead of sending the credentials, cache and send the following cookies to authenticate subsequent requests.  **For Prism Central with IAMv2 enabled:**  - NTNX_MERCURY_IAM_SESSION - NTNX_IAM_SESSION - NTNX_MERCURY_IAM_REFRESH_TOKEN  **For Prism Central with IAMv2 disabled or Prism Element:**  - NTNX_IGW_SESSION - NTNX_MERCURY_IGW_SESSION - JSESSIONID(If present)  **Note:** Upon cookie expiration, you must send basic credentials again in the next API call, to obtain the new cookies.  The following are the code snippets to authenticate a user.  **Python**  ``` #!/usr/bin/env python3.7 import urllib3 import requests from base64 import b64encode urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning) request_url = 'https://10.0.0.1:9440/api/nutanix/v3/vms/list' username = 'admin' password = 'password' encoded_credentials = b64encode(bytes(f'{username}:{password}', encoding='ascii')).decode('ascii') auth_header = f'Basic {encoded_credentials}' payload = '{\"kind\":\"vm\"}' headers = {'Accept': 'application/json', 'Content-Type': 'application/json', 'Authorization': f'{auth_header}', 'cache-control': 'no-cache'} response = requests.request('post', request_url, data=payload, headers=headers, verify=False) print(response.status_code) ```  **C#**  ``` using System; using System.Net; using System.Text; namespace ConsoleApp1 { class Program { static void Main(string[] args) { string ClusterUsername = \"admin\"; string ClusterPassword = \"Ntnx2017!\"; string RequestUrl = \"https://192.168.1.131:9440/api/nutanix/v3/vms/list\"; string Payload = \"{\\\"kind\\\":\\\"vm\\\"}\"; string AuthHeader = System.Convert.ToBase64String( System.Text.ASCIIEncoding.ASCII.GetBytes( ClusterUsername + \":\" + ClusterPassword)); ServicePointManager.ServerCertificateValidationCallback = ( (sender, certificate, chain, sslPolicyErrors) => true); ServicePointManager.Expect100Continue = true; ServicePointManager.SecurityProtocol = SecurityProtocolType.Tls12; String authHeader = System.Convert.ToBase64String( System.Text.ASCIIEncoding.ASCII.GetBytes( ClusterUsername + \":\" + ClusterPassword)); var request = (HttpWebRequest)WebRequest.Create(RequestUrl); var requestBody = Encoding.ASCII.GetBytes(Payload); HttpWebResponse HttpResponse = null; request.Headers.Add(\"Authorization\", \"Basic \" + AuthHeader); request.Headers.Add(\"Accept-Encoding\", \"application/json\"); request.ContentType = \"application/json\"; request.Accept = \"application/json\"; request.Method = \"POST\"; var newStream = request.GetRequestStream(); newStream.Write(requestBody, 0, requestBody.Length); newStream.Close(); HttpResponse = (HttpWebResponse)request.GetResponse(); Console.WriteLine(HttpResponse.StatusCode); Console.ReadKey(); } } } ```  **PowerShell**  ``` # set the basic properties for the request $RequestUri = \"https://10.0.0.1:9440/api/nutanix/v3/vms/list\" $Payload = \"{\"\"kind\"\":\"\"vm\"\"}\" $Username = \"admin\" $Password = \"password\" # create the HTTP Basic Authorization header $pair = $Username + \":\" + $Password $bytes = [System.Text.Encoding]::ASCII.GetBytes($pair) $base64 = [System.Convert]::ToBase64String($bytes) $basicAuthValue = \"Basic $base64\" # setup the request headers $Headers = @{ 'Accept' = 'application/json' 'Authorization' = $basicAuthValue 'Content-Type' = 'application/json' } # disable SSL certification verification # you probably shouldn't do this in production ... if (-not ( [System.Management.Automation.PSTypeName] ` 'ServerCertificateValidationCallback').Type) {$certCallback = @\" using System; using System.Net; using System.Net.Security; using System.Security.Cryptography.X509Certificates; public class ServerCertificateValidationCallback { public static void Ignore() { if(ServicePointManager.ServerCertificateValidationCallback ==null) { ServicePointManager.ServerCertificateValidationCallback += delegate ( Object obj, X509Certificate certificate, X509Chain chain, SslPolicyErrors errors ) { return true; }; } } } \"@ Add-Type $certCallback } [ServerCertificateValidationCallback]::Ignore() [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12 # submit the request Invoke-WebRequest ` -Uri $RequestUri ` -Headers $Headers ` -Method \"POST\" ` -Body $Payload ` -TimeoutSec 5 ` -UseBasicParsing ` -DisableKeepAlive ` | ConvertFrom-Json | Select * ```  **PHP**  ``` <?php $username = 'admin'; $password = 'password'; $ch = curl_init(); curl_setopt_array($ch, [ CURLOPT_RETURNTRANSFER => 1, CURLINFO_HEADER_OUT => TRUE, CURLOPT_URL => 'https://10.0.0.1:9440/api/nutanix/v3/vms/list', CURLOPT_USERPWD => $username . ':' . $password, CURLOPT_POST => 1, CURLOPT_SSL_VERIFYHOST => FALSE, CURLOPT_SSL_VERIFYPEER => FALSE, CURLOPT_HTTPHEADER => [ 'Content-Type: application/json', 'Accept-Encoding: application/json', ], CURLOPT_POSTFIELDS => json_encode([ 'kind' => 'vm' ]) ]); $result = curl_exec($ch); curl_close($ch); if(!$result) { die('Error: \"' . curl_error($ch) . '\". Code: ' . curl_errno($ch)); } else { echo($result); } ``` # pc.2022.1 release  ## v3 API behavior changes with pc.2022.1 release   From Prism Central release 2022.1 onwards, the v3 intentful API is enhanced to protect data consistency, especially in concurrent race conditions. As a result of this enhancement, response time for the synchronous API requests is improved and there are some behavior changes from an API consumption perspective.   1.  In the current v3 API workflow, clients will get a synchronous API response that contains task and entity UUID. Post which clients should send GET API on the task periodically until the task succeeds or fails. Clients can also send GET API on the entity.         - **Behavioral changes**: Some of the synchronous errors will be reported asynchronously now.            1. **Invalid idempotency UUID error**:                       **Current API behavior**: If POST API contains invalid idempotency UUID, clients will receive a     synchronous bad request error with an HTTP code 400.           **Expected change**: If POST API contains invalid idempotency UUID, clients will get a synchronous API response with an HTTP code 202(accepted). GET API on the task will return a bad request error.      1. **Authorization errors**:            **Current API behavior**: Clients will get synchronous authorization errors (for example- HTTP code 403 forbidden error).            **Expected change**: Clients will get a synchronous API response with an HTTP code 202 (accepted). GET API on the task will show an authorization error.       1. **Incorrect category mapping error**:            **Current API behavior**: If an intentful API payload has a project reference and uses categories mapping but the category for the project does not exist, the client will get an error response with an HTTP code 404(not found).            **Expected change**: If an intentful API payload has a project reference and uses categories mapping but the category for the project does not exist, the client will receive a synchronous API response with an HTTP code 202 (accepted). GET API on the task will show a “not found” error.      1.  **Invalid user detail error**:            **Current API behavior**: If an intentful API payload has user_uuid or username in owner_reference but the corresponding user does not exist, the client will receive an API response with an HTTP code 404.            **Expected change**: If an intentful API payload has user_uuid or username in owner_reference but the corresponding user does not exist, clients will get a synchronous API response with an HTTP code 202(accepted). GET API on the task will show a bad request error.         1.  . **Missing user detail error**:           **Current API behavior**: If an intentful API payload does not have user_uuid or username in owner_reference and the login user does not exist, the client will get an API response with an HTTP code 404.           **Expected change**: If an intentful API payload does not have user_uuid or username in owner_reference and login user does not exist, clients will get a synchronous API response with an HTTP code 202(accepted). GET API on the task will show a bad request error.      -  **Format changes** for some of the synchronous API responses:                  **Current API behavior**: Response to an intentful POST, PUT, or DELETE API will contain tenant_uuid, owner_uuid, owner_username, service_name, session_log_uuid, trace, user_uuid, username, remote_authorization, session_json_dict, incap_req_id, project_reference or category information.                  **Expected change**: Main processing will occur after a synchronous response to the client. As a result,  synchronous response to an intentful POST, PUT, or DELETE API will not contain tenant_uuid, owner_uuid, owner_username, service_name, session_log_uuid, trace, user_uuid, username, remote_authorization, session_json_dict, incap_req_id, project_reference or category information.                  If API consumers are expecting these properties in the synchronous response before, they have to submit a GET/LIST request to inspect these properties.  2.  The version control workflow for v3 intentful PUT and DELETE APIs:    - **Absolutely no version control check for DELETE APIs**:                  **Current API behavior**: DELETE API will fail if an incorrect spec_version is provided as one of the request arguments.                  **Expected change**:  DELETE API will always delete the entity. No version control is applicable on the DELETE request.       - **Edit conflict error could be reported asynchronously**:        **Current API behavior**: If v3 PUT APIs are conflicting with v2 PUT APIs on the same entities, the client will receive an HTTP code 409 response for v3 PUT APIs.       **Expected change**: If v3 PUT APIs are conflicting with v2 PUT APIs on the same entities, the client will receive an HTTP code 202 response for v3 PUT APIs and will receive the task failed response with an edit conflict error asynchronously.     - **More restrictive version control by using spec_hash**:                 **Current API behavior**: GET/LIST API will return a response that may contain one of the values for version control, it could be spec_version or spec_hash. spec_version is numeric while spec_hash is a string.                  Whichever value is provided in the GET/LIST API, the same value should be included in the next PUT API payload of the same entity.        **Expected change**: GET/LIST API will provide both spec_hash and spec_version in API response. Use of spec_hash is not enforced but could be enforced in future releases.        - **The version control value in the GET API response will not change when a PUT request is in progress**:                  **Current API behavior**: If there are one GET request and one PUT request on the same entity, the version control value from the GET request response might be different depending on the current state of the entity. The behavior is not deterministic.          **Expected change**: If there is one GET request and one PUT request on the same entity, the spec_hash value of the GET request response will always be the same as it was before the PUT request until the PUT is completed.    1. V3 intentful POST API and the following GET API:    -  **Get API request for new entity**:        **Current API behavior**: If a POST API fails asynchronously, the following GET API on the same entity will get an HTTP code 202 response with the entity in an error state.          However, GET API on the same entity after 5+ minutes will show an HTTP code 404 response reporting the entity does not exist.        **Expected change**: If a POST API fails asynchronously, GET API on the same entity will always show an HTTP code 404 response meaning the entity does not exist.        - **The intent spec is an entity we created to record previous and ongoing requests from v3 POST/PUT/DELETE APIs**. There should be only one intent spec for each entity if the intent spec exists:          **Current API behavior:** If the client queries intent spec for an entity immediately after getting an HTTP code 202 response for the v3 POST API on this entity, the intent spec will always be present.          **Expected change**: If the client queries intent spec for an entity immediately after getting an HTTP code 202 response for the v3 POST API on this entity, the intent spec may not be created yet.  4. Prism Central upgrade to pc.2022.1 and following releases:        **Current API behavior**: The pending and running API requests will continue to execute after an upgrade.        **Expected change**: During the pc.2022.1 version upgrade, all pending and running API requests and related tasks will be marked as failed after a restart. This will be shown even though they may be in progress and completed in the middle of the upgrade period.       This behavior only applies to upgrades from releases before pc.2022.1 to releases equal to or after pc.2022.1. All future upgrades from pc.2022.1 release will remain the same as before where pending and running API requests will continue to execute after an upgrade.  5. Multiple nodes upgrade:   -   pc.2022.1 version can take care of the scenarios where Prism Central has the new version and Prism Element is still at the older version.        -   Any task (for existing create, update, delete v3 API operations) that is in running state during the upgrade will be marked failed. This also applies to new requests that come before the upgrade ends. The recommendation is to wait until the upgrade finishes before executing any v3 create, update, delete API.     6. Schema consistency:    - API schema in Prism Central will be used in case there is a      discrepancy between Prism Element and Prism Central. There will be a      change of behavior when Prism Central has a post pc.2022.1 version.        **Current API behavior**: API responses for POST, PUT, and DELETE API on Prism Central will match API schema on Prism Element.            **Expected change**: API response for POST, PUT, and DELETE API on Prism Central will match API schema on Prism Central.  # Use Cases            ## Post /VMs  Perform the following procedure to generate the code snippet for creating a VM based on the input parameters.  1. Under **VMs**, click **POST VMs** and navigate to the end of the page. 2. In the **Code Generation** panel, under the **Settings** tab, enter the value of the following parameters.    1. Enter the username of the Prism Central in the **Username** field.    2. Enter the password in the **Password** field.    3. Enter the IP address of the Prism Central in the **Host** field. 3. Under the **Body** tab, enter the value of all the required parameters as described in the **Request Body** section. Optionally, you can also enter the values of the optional parameters.    1. Under **Specification**, enter the description of VM in a string format against the **description** attribute. The maximum length allowed is 1000 characters only.    2. Under **Resources** do the following.       1. Enter the value of number of logical threads per core in an integer format against the **num_threads_per_core** attribute. The minimum value allowed is 1.       2. Enter the value of current or desired power state of the VM in a string format against the **power_state** attribute.       3. Enter the value of number of vCPUs per socket in an integer format against the **num_vcpus_per_socket** attribute. The minimun value allowed is 1.       4. Enter the value of number of vCPU sockets in an integer format against the **num_sockets** attribute. The minimum value allowed is 1.    3. Under **Name**, enter the name of the VM in a string format. The maximum length allowed is 64 characters only.    4. Under **Metadata** do the following.       1. The **last_update_time** is a read-only attribute indicating UTC date and time in an RFC-339 format, when the VM was last updated in a string format.       2. The **kind** is a read-only attribute indicating the kind name in a string format. The default value is VM.       3. Enter the value of VM uuid in a string format against the **uuid** attribute. The valid pattern is ^\\[a-fA-F0-9]{8}-\\[a-fA-F0-9]{4}-\\[a-fA-F0-9]{4}-\\[a-fA-F0-9]{4}-\\[a-fAF0-9]{12}$. 4. Under the **Code Generation** tab, select the scripting languages to view the code snippet of creating a VM.     You can either select Python or PowerShell or Shell or Go or Java or JavaScript or Node or Obj-C or PHP or Ruby or Swift or C# or C.     The code snippet appears in the selected scripting language. You can use the code snippet to create a VM in your environment.  ## Put /VMs  Perform the following procedure to generate the code snippet for updating an existing VM based on the input parameters.  1. Under **VMs**, click **PUT VMs** and navigate to the end of the page. 2. In the **Code Generation** panel, under the **Settings** tab, enter the value of the following parameters.    1. Enter the UUID of the VM that needs to be updated in the **UUID** field.    2. Enter the username of the Prism Central in the **Username** field.    3. Enter the password in the **Password** field.    4. Enter the IP address of the Prism Central in the **Host** field. 3. Under the **Body** tab, enter the value of all the required parameters as described in the **Request Body** section. Optionally, you can also enter the values of the optional parameters.    1. Under **Specification**, enter the description of VM in a string format against the **description** attribute. The maximum length allowed is 1000 characters only.    2. Under **Resources** do the following.       1. Enter the value of number of logical threads per core in an integer format against the **num_threads_per_core** attribute. The minimum value allowed is 1.       2. Enter the value of current or desired power state of the VM in a string format against the **power_state** attribute.       3. Enter the value of number of vCPUs per socket in an integer format against the **num_vcpus_per_socket** attribute. The minimun value allowed is 1.       4. Enter the value of number of vCPU sockets in an integer format against the **num_sockets** attribute. The minimum value allowed is 1.    3. Under **Name**, enter the name of the VM in a string format. The maximum length allowed is 64 characters only.    4. Under **Metadata** do the following.       1. The **last_update_time** is a read-only attribute indicating UTC date and time in an RFC-339 format, when the VM was last updated in a string format.       2. The **kind** is a read-only attribute indicating the kind name in a string format. The default value is VM.       3. Enter the value of VM uuid in a string format against the **uuid** attribute. The valid pattern is ^\\[a-fA-F0-9]{8}-\\[a-fA-F0-9]{4}-\\[a-fA-F0-9]{4}-\\[a-fA-F0-9]{4}-\\[a-fAF0-9]{12}$. 4. Under the **Code Generation** tab, select the scripting languages to view the code snippet of creating a VM.     You can either select Python or PowerShell or Shell or Go or Java or JavaScript or Node or Obj-C or PHP or Ruby or Swift or C# or C.     The code snippet appears in the selected scripting language. You can use the code snippet to update a VM in your environment. 

API version: 3.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
)

// checks if the GroupsGetEntitiesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsGetEntitiesRequest{}

// GroupsGetEntitiesRequest Get Entities Request.
type GroupsGetEntitiesRequest struct {
	// When retrieving latest values, how far back to look as a multiple of the downsampling interval for the metric. 
	NumberOfIntervalsForLatestData *int32 `json:"number_of_intervals_for_latest_data,omitempty"`
	// The name of the attribute that will be used to sort groups. 
	GroupSortAttribute *string `json:"group_sort_attribute,omitempty"`
	// For grouping, how many groups to return.
	NumberOfBuckets *int32 `json:"number_of_buckets,omitempty"`
	// A set of entities that the request will be scoped to.
	EntityIds []string `json:"entity_ids,omitempty"`
	GroupMemberAttributes []GroupsRequestedAttribute `json:"group_member_attributes,omitempty"`
	// The name of the attribute that will be used to sort group members. 
	GroupMemberSortAttribute *string `json:"group_member_sort_attribute,omitempty"`
	// For grouping, the boundary to snap to when grouping.
	BucketBoundary *int32 `json:"bucket_boundary,omitempty"`
	// The offset into the total set of groups to return.
	GroupOffset *int64 `json:"group_offset,omitempty"`
	// Downsampling interval to apply to query if override is desired. 
	DownsamplingInterval *int32 `json:"downsampling_interval,omitempty"`
	// For a time-series query, the start of the interval since the epoch in ms. Default is latest value only. 
	IntervalStartMs *int64 `json:"interval_start_ms,omitempty"`
	// The entity type that will be requested.
	EntityType string `json:"entity_type"`
	// The offset into the total member set to return per group.
	GroupMemberOffset *int64 `json:"group_member_offset,omitempty"`
	// Attribute that will be used to perform a group-by if needed. 
	GroupingAttribute *string `json:"grouping_attribute,omitempty"`
	// Downsampling function to take time series data and resolve to one value for sorting purposes. 
	GroupMemberSortDownsamplingFunction *string `json:"group_member_sort_downsampling_function,omitempty"`
	// Sort order for entities and entity groups.
	GroupSortOrder *string `json:"group_sort_order,omitempty"`
	// Downsampling function to take time series data and resolve to one value for sorting purposes. 
	GroupSortDownsampleFunction *string `json:"group_sort_downsample_function,omitempty"`
	// FIQL filter criteria that will be used to filter the returned data. 
	FilterCriteria *string `json:"filter_criteria,omitempty"`
	// Same as bucket_boundary but supports larger range of values.
	LargeBucketBoundary *int64 `json:"large_bucket_boundary,omitempty"`
	// The scope of availability zones from which to fetch the  data. 
	AvailabilityZoneScope *string `json:"availability_zone_scope,omitempty"`
	// The maximum number of groups to return in the result.
	GroupCount *int64 `json:"group_count,omitempty"`
	GroupAttributes []GroupsRequestedAttribute `json:"group_attributes,omitempty"`
	// For a time-series query, the end of the interval since the epoch in ms. Default is latest value only. 
	IntervalEndMs *int64 `json:"interval_end_ms,omitempty"`
	// The type of an attribute being used for grouping - may be continuous or discrete. 
	GroupingAttributeType *string `json:"grouping_attribute_type,omitempty"`
	// The maximum number of members to return per group.
	GroupMemberCount *int64 `json:"group_member_count,omitempty"`
	// Sort order for entities and entity groups.
	GroupMemberSortOrder *string `json:"group_member_sort_order,omitempty"`
	// A custom name to use for tagging the query when debugging.
	QueryName *string `json:"query_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupsGetEntitiesRequest GroupsGetEntitiesRequest

// NewGroupsGetEntitiesRequest instantiates a new GroupsGetEntitiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsGetEntitiesRequest(entityType string) *GroupsGetEntitiesRequest {
	this := GroupsGetEntitiesRequest{}
	var intervalStartMs int64 = 0
	this.IntervalStartMs = &intervalStartMs
	this.EntityType = entityType
	var availabilityZoneScope string = "LOCAL"
	this.AvailabilityZoneScope = &availabilityZoneScope
	var intervalEndMs int64 = 0
	this.IntervalEndMs = &intervalEndMs
	return &this
}

// NewGroupsGetEntitiesRequestWithDefaults instantiates a new GroupsGetEntitiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsGetEntitiesRequestWithDefaults() *GroupsGetEntitiesRequest {
	this := GroupsGetEntitiesRequest{}
	var intervalStartMs int64 = 0
	this.IntervalStartMs = &intervalStartMs
	var availabilityZoneScope string = "LOCAL"
	this.AvailabilityZoneScope = &availabilityZoneScope
	var intervalEndMs int64 = 0
	this.IntervalEndMs = &intervalEndMs
	return &this
}

// GetNumberOfIntervalsForLatestData returns the NumberOfIntervalsForLatestData field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetNumberOfIntervalsForLatestData() int32 {
	if o == nil || IsNil(o.NumberOfIntervalsForLatestData) {
		var ret int32
		return ret
	}
	return *o.NumberOfIntervalsForLatestData
}

// GetNumberOfIntervalsForLatestDataOk returns a tuple with the NumberOfIntervalsForLatestData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetNumberOfIntervalsForLatestDataOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfIntervalsForLatestData) {
		return nil, false
	}
	return o.NumberOfIntervalsForLatestData, true
}

// HasNumberOfIntervalsForLatestData returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasNumberOfIntervalsForLatestData() bool {
	if o != nil && !IsNil(o.NumberOfIntervalsForLatestData) {
		return true
	}

	return false
}

// SetNumberOfIntervalsForLatestData gets a reference to the given int32 and assigns it to the NumberOfIntervalsForLatestData field.
func (o *GroupsGetEntitiesRequest) SetNumberOfIntervalsForLatestData(v int32) {
	o.NumberOfIntervalsForLatestData = &v
}

// GetGroupSortAttribute returns the GroupSortAttribute field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetGroupSortAttribute() string {
	if o == nil || IsNil(o.GroupSortAttribute) {
		var ret string
		return ret
	}
	return *o.GroupSortAttribute
}

// GetGroupSortAttributeOk returns a tuple with the GroupSortAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetGroupSortAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.GroupSortAttribute) {
		return nil, false
	}
	return o.GroupSortAttribute, true
}

// HasGroupSortAttribute returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasGroupSortAttribute() bool {
	if o != nil && !IsNil(o.GroupSortAttribute) {
		return true
	}

	return false
}

// SetGroupSortAttribute gets a reference to the given string and assigns it to the GroupSortAttribute field.
func (o *GroupsGetEntitiesRequest) SetGroupSortAttribute(v string) {
	o.GroupSortAttribute = &v
}

// GetNumberOfBuckets returns the NumberOfBuckets field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetNumberOfBuckets() int32 {
	if o == nil || IsNil(o.NumberOfBuckets) {
		var ret int32
		return ret
	}
	return *o.NumberOfBuckets
}

// GetNumberOfBucketsOk returns a tuple with the NumberOfBuckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetNumberOfBucketsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfBuckets) {
		return nil, false
	}
	return o.NumberOfBuckets, true
}

// HasNumberOfBuckets returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasNumberOfBuckets() bool {
	if o != nil && !IsNil(o.NumberOfBuckets) {
		return true
	}

	return false
}

// SetNumberOfBuckets gets a reference to the given int32 and assigns it to the NumberOfBuckets field.
func (o *GroupsGetEntitiesRequest) SetNumberOfBuckets(v int32) {
	o.NumberOfBuckets = &v
}

// GetEntityIds returns the EntityIds field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetEntityIds() []string {
	if o == nil || IsNil(o.EntityIds) {
		var ret []string
		return ret
	}
	return o.EntityIds
}

// GetEntityIdsOk returns a tuple with the EntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetEntityIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.EntityIds) {
		return nil, false
	}
	return o.EntityIds, true
}

// HasEntityIds returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasEntityIds() bool {
	if o != nil && !IsNil(o.EntityIds) {
		return true
	}

	return false
}

// SetEntityIds gets a reference to the given []string and assigns it to the EntityIds field.
func (o *GroupsGetEntitiesRequest) SetEntityIds(v []string) {
	o.EntityIds = v
}

// GetGroupMemberAttributes returns the GroupMemberAttributes field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetGroupMemberAttributes() []GroupsRequestedAttribute {
	if o == nil || IsNil(o.GroupMemberAttributes) {
		var ret []GroupsRequestedAttribute
		return ret
	}
	return o.GroupMemberAttributes
}

// GetGroupMemberAttributesOk returns a tuple with the GroupMemberAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetGroupMemberAttributesOk() ([]GroupsRequestedAttribute, bool) {
	if o == nil || IsNil(o.GroupMemberAttributes) {
		return nil, false
	}
	return o.GroupMemberAttributes, true
}

// HasGroupMemberAttributes returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasGroupMemberAttributes() bool {
	if o != nil && !IsNil(o.GroupMemberAttributes) {
		return true
	}

	return false
}

// SetGroupMemberAttributes gets a reference to the given []GroupsRequestedAttribute and assigns it to the GroupMemberAttributes field.
func (o *GroupsGetEntitiesRequest) SetGroupMemberAttributes(v []GroupsRequestedAttribute) {
	o.GroupMemberAttributes = v
}

// GetGroupMemberSortAttribute returns the GroupMemberSortAttribute field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetGroupMemberSortAttribute() string {
	if o == nil || IsNil(o.GroupMemberSortAttribute) {
		var ret string
		return ret
	}
	return *o.GroupMemberSortAttribute
}

// GetGroupMemberSortAttributeOk returns a tuple with the GroupMemberSortAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetGroupMemberSortAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.GroupMemberSortAttribute) {
		return nil, false
	}
	return o.GroupMemberSortAttribute, true
}

// HasGroupMemberSortAttribute returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasGroupMemberSortAttribute() bool {
	if o != nil && !IsNil(o.GroupMemberSortAttribute) {
		return true
	}

	return false
}

// SetGroupMemberSortAttribute gets a reference to the given string and assigns it to the GroupMemberSortAttribute field.
func (o *GroupsGetEntitiesRequest) SetGroupMemberSortAttribute(v string) {
	o.GroupMemberSortAttribute = &v
}

// GetBucketBoundary returns the BucketBoundary field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetBucketBoundary() int32 {
	if o == nil || IsNil(o.BucketBoundary) {
		var ret int32
		return ret
	}
	return *o.BucketBoundary
}

// GetBucketBoundaryOk returns a tuple with the BucketBoundary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetBucketBoundaryOk() (*int32, bool) {
	if o == nil || IsNil(o.BucketBoundary) {
		return nil, false
	}
	return o.BucketBoundary, true
}

// HasBucketBoundary returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasBucketBoundary() bool {
	if o != nil && !IsNil(o.BucketBoundary) {
		return true
	}

	return false
}

// SetBucketBoundary gets a reference to the given int32 and assigns it to the BucketBoundary field.
func (o *GroupsGetEntitiesRequest) SetBucketBoundary(v int32) {
	o.BucketBoundary = &v
}

// GetGroupOffset returns the GroupOffset field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetGroupOffset() int64 {
	if o == nil || IsNil(o.GroupOffset) {
		var ret int64
		return ret
	}
	return *o.GroupOffset
}

// GetGroupOffsetOk returns a tuple with the GroupOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetGroupOffsetOk() (*int64, bool) {
	if o == nil || IsNil(o.GroupOffset) {
		return nil, false
	}
	return o.GroupOffset, true
}

// HasGroupOffset returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasGroupOffset() bool {
	if o != nil && !IsNil(o.GroupOffset) {
		return true
	}

	return false
}

// SetGroupOffset gets a reference to the given int64 and assigns it to the GroupOffset field.
func (o *GroupsGetEntitiesRequest) SetGroupOffset(v int64) {
	o.GroupOffset = &v
}

// GetDownsamplingInterval returns the DownsamplingInterval field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetDownsamplingInterval() int32 {
	if o == nil || IsNil(o.DownsamplingInterval) {
		var ret int32
		return ret
	}
	return *o.DownsamplingInterval
}

// GetDownsamplingIntervalOk returns a tuple with the DownsamplingInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetDownsamplingIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.DownsamplingInterval) {
		return nil, false
	}
	return o.DownsamplingInterval, true
}

// HasDownsamplingInterval returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasDownsamplingInterval() bool {
	if o != nil && !IsNil(o.DownsamplingInterval) {
		return true
	}

	return false
}

// SetDownsamplingInterval gets a reference to the given int32 and assigns it to the DownsamplingInterval field.
func (o *GroupsGetEntitiesRequest) SetDownsamplingInterval(v int32) {
	o.DownsamplingInterval = &v
}

// GetIntervalStartMs returns the IntervalStartMs field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetIntervalStartMs() int64 {
	if o == nil || IsNil(o.IntervalStartMs) {
		var ret int64
		return ret
	}
	return *o.IntervalStartMs
}

// GetIntervalStartMsOk returns a tuple with the IntervalStartMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetIntervalStartMsOk() (*int64, bool) {
	if o == nil || IsNil(o.IntervalStartMs) {
		return nil, false
	}
	return o.IntervalStartMs, true
}

// HasIntervalStartMs returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasIntervalStartMs() bool {
	if o != nil && !IsNil(o.IntervalStartMs) {
		return true
	}

	return false
}

// SetIntervalStartMs gets a reference to the given int64 and assigns it to the IntervalStartMs field.
func (o *GroupsGetEntitiesRequest) SetIntervalStartMs(v int64) {
	o.IntervalStartMs = &v
}

// GetEntityType returns the EntityType field value
func (o *GroupsGetEntitiesRequest) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *GroupsGetEntitiesRequest) SetEntityType(v string) {
	o.EntityType = v
}

// GetGroupMemberOffset returns the GroupMemberOffset field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetGroupMemberOffset() int64 {
	if o == nil || IsNil(o.GroupMemberOffset) {
		var ret int64
		return ret
	}
	return *o.GroupMemberOffset
}

// GetGroupMemberOffsetOk returns a tuple with the GroupMemberOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetGroupMemberOffsetOk() (*int64, bool) {
	if o == nil || IsNil(o.GroupMemberOffset) {
		return nil, false
	}
	return o.GroupMemberOffset, true
}

// HasGroupMemberOffset returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasGroupMemberOffset() bool {
	if o != nil && !IsNil(o.GroupMemberOffset) {
		return true
	}

	return false
}

// SetGroupMemberOffset gets a reference to the given int64 and assigns it to the GroupMemberOffset field.
func (o *GroupsGetEntitiesRequest) SetGroupMemberOffset(v int64) {
	o.GroupMemberOffset = &v
}

// GetGroupingAttribute returns the GroupingAttribute field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetGroupingAttribute() string {
	if o == nil || IsNil(o.GroupingAttribute) {
		var ret string
		return ret
	}
	return *o.GroupingAttribute
}

// GetGroupingAttributeOk returns a tuple with the GroupingAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetGroupingAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.GroupingAttribute) {
		return nil, false
	}
	return o.GroupingAttribute, true
}

// HasGroupingAttribute returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasGroupingAttribute() bool {
	if o != nil && !IsNil(o.GroupingAttribute) {
		return true
	}

	return false
}

// SetGroupingAttribute gets a reference to the given string and assigns it to the GroupingAttribute field.
func (o *GroupsGetEntitiesRequest) SetGroupingAttribute(v string) {
	o.GroupingAttribute = &v
}

// GetGroupMemberSortDownsamplingFunction returns the GroupMemberSortDownsamplingFunction field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetGroupMemberSortDownsamplingFunction() string {
	if o == nil || IsNil(o.GroupMemberSortDownsamplingFunction) {
		var ret string
		return ret
	}
	return *o.GroupMemberSortDownsamplingFunction
}

// GetGroupMemberSortDownsamplingFunctionOk returns a tuple with the GroupMemberSortDownsamplingFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetGroupMemberSortDownsamplingFunctionOk() (*string, bool) {
	if o == nil || IsNil(o.GroupMemberSortDownsamplingFunction) {
		return nil, false
	}
	return o.GroupMemberSortDownsamplingFunction, true
}

// HasGroupMemberSortDownsamplingFunction returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasGroupMemberSortDownsamplingFunction() bool {
	if o != nil && !IsNil(o.GroupMemberSortDownsamplingFunction) {
		return true
	}

	return false
}

// SetGroupMemberSortDownsamplingFunction gets a reference to the given string and assigns it to the GroupMemberSortDownsamplingFunction field.
func (o *GroupsGetEntitiesRequest) SetGroupMemberSortDownsamplingFunction(v string) {
	o.GroupMemberSortDownsamplingFunction = &v
}

// GetGroupSortOrder returns the GroupSortOrder field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetGroupSortOrder() string {
	if o == nil || IsNil(o.GroupSortOrder) {
		var ret string
		return ret
	}
	return *o.GroupSortOrder
}

// GetGroupSortOrderOk returns a tuple with the GroupSortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetGroupSortOrderOk() (*string, bool) {
	if o == nil || IsNil(o.GroupSortOrder) {
		return nil, false
	}
	return o.GroupSortOrder, true
}

// HasGroupSortOrder returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasGroupSortOrder() bool {
	if o != nil && !IsNil(o.GroupSortOrder) {
		return true
	}

	return false
}

// SetGroupSortOrder gets a reference to the given string and assigns it to the GroupSortOrder field.
func (o *GroupsGetEntitiesRequest) SetGroupSortOrder(v string) {
	o.GroupSortOrder = &v
}

// GetGroupSortDownsampleFunction returns the GroupSortDownsampleFunction field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetGroupSortDownsampleFunction() string {
	if o == nil || IsNil(o.GroupSortDownsampleFunction) {
		var ret string
		return ret
	}
	return *o.GroupSortDownsampleFunction
}

// GetGroupSortDownsampleFunctionOk returns a tuple with the GroupSortDownsampleFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetGroupSortDownsampleFunctionOk() (*string, bool) {
	if o == nil || IsNil(o.GroupSortDownsampleFunction) {
		return nil, false
	}
	return o.GroupSortDownsampleFunction, true
}

// HasGroupSortDownsampleFunction returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasGroupSortDownsampleFunction() bool {
	if o != nil && !IsNil(o.GroupSortDownsampleFunction) {
		return true
	}

	return false
}

// SetGroupSortDownsampleFunction gets a reference to the given string and assigns it to the GroupSortDownsampleFunction field.
func (o *GroupsGetEntitiesRequest) SetGroupSortDownsampleFunction(v string) {
	o.GroupSortDownsampleFunction = &v
}

// GetFilterCriteria returns the FilterCriteria field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetFilterCriteria() string {
	if o == nil || IsNil(o.FilterCriteria) {
		var ret string
		return ret
	}
	return *o.FilterCriteria
}

// GetFilterCriteriaOk returns a tuple with the FilterCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetFilterCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.FilterCriteria) {
		return nil, false
	}
	return o.FilterCriteria, true
}

// HasFilterCriteria returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasFilterCriteria() bool {
	if o != nil && !IsNil(o.FilterCriteria) {
		return true
	}

	return false
}

// SetFilterCriteria gets a reference to the given string and assigns it to the FilterCriteria field.
func (o *GroupsGetEntitiesRequest) SetFilterCriteria(v string) {
	o.FilterCriteria = &v
}

// GetLargeBucketBoundary returns the LargeBucketBoundary field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetLargeBucketBoundary() int64 {
	if o == nil || IsNil(o.LargeBucketBoundary) {
		var ret int64
		return ret
	}
	return *o.LargeBucketBoundary
}

// GetLargeBucketBoundaryOk returns a tuple with the LargeBucketBoundary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetLargeBucketBoundaryOk() (*int64, bool) {
	if o == nil || IsNil(o.LargeBucketBoundary) {
		return nil, false
	}
	return o.LargeBucketBoundary, true
}

// HasLargeBucketBoundary returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasLargeBucketBoundary() bool {
	if o != nil && !IsNil(o.LargeBucketBoundary) {
		return true
	}

	return false
}

// SetLargeBucketBoundary gets a reference to the given int64 and assigns it to the LargeBucketBoundary field.
func (o *GroupsGetEntitiesRequest) SetLargeBucketBoundary(v int64) {
	o.LargeBucketBoundary = &v
}

// GetAvailabilityZoneScope returns the AvailabilityZoneScope field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetAvailabilityZoneScope() string {
	if o == nil || IsNil(o.AvailabilityZoneScope) {
		var ret string
		return ret
	}
	return *o.AvailabilityZoneScope
}

// GetAvailabilityZoneScopeOk returns a tuple with the AvailabilityZoneScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetAvailabilityZoneScopeOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityZoneScope) {
		return nil, false
	}
	return o.AvailabilityZoneScope, true
}

// HasAvailabilityZoneScope returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasAvailabilityZoneScope() bool {
	if o != nil && !IsNil(o.AvailabilityZoneScope) {
		return true
	}

	return false
}

// SetAvailabilityZoneScope gets a reference to the given string and assigns it to the AvailabilityZoneScope field.
func (o *GroupsGetEntitiesRequest) SetAvailabilityZoneScope(v string) {
	o.AvailabilityZoneScope = &v
}

// GetGroupCount returns the GroupCount field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetGroupCount() int64 {
	if o == nil || IsNil(o.GroupCount) {
		var ret int64
		return ret
	}
	return *o.GroupCount
}

// GetGroupCountOk returns a tuple with the GroupCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetGroupCountOk() (*int64, bool) {
	if o == nil || IsNil(o.GroupCount) {
		return nil, false
	}
	return o.GroupCount, true
}

// HasGroupCount returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasGroupCount() bool {
	if o != nil && !IsNil(o.GroupCount) {
		return true
	}

	return false
}

// SetGroupCount gets a reference to the given int64 and assigns it to the GroupCount field.
func (o *GroupsGetEntitiesRequest) SetGroupCount(v int64) {
	o.GroupCount = &v
}

// GetGroupAttributes returns the GroupAttributes field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetGroupAttributes() []GroupsRequestedAttribute {
	if o == nil || IsNil(o.GroupAttributes) {
		var ret []GroupsRequestedAttribute
		return ret
	}
	return o.GroupAttributes
}

// GetGroupAttributesOk returns a tuple with the GroupAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetGroupAttributesOk() ([]GroupsRequestedAttribute, bool) {
	if o == nil || IsNil(o.GroupAttributes) {
		return nil, false
	}
	return o.GroupAttributes, true
}

// HasGroupAttributes returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasGroupAttributes() bool {
	if o != nil && !IsNil(o.GroupAttributes) {
		return true
	}

	return false
}

// SetGroupAttributes gets a reference to the given []GroupsRequestedAttribute and assigns it to the GroupAttributes field.
func (o *GroupsGetEntitiesRequest) SetGroupAttributes(v []GroupsRequestedAttribute) {
	o.GroupAttributes = v
}

// GetIntervalEndMs returns the IntervalEndMs field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetIntervalEndMs() int64 {
	if o == nil || IsNil(o.IntervalEndMs) {
		var ret int64
		return ret
	}
	return *o.IntervalEndMs
}

// GetIntervalEndMsOk returns a tuple with the IntervalEndMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetIntervalEndMsOk() (*int64, bool) {
	if o == nil || IsNil(o.IntervalEndMs) {
		return nil, false
	}
	return o.IntervalEndMs, true
}

// HasIntervalEndMs returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasIntervalEndMs() bool {
	if o != nil && !IsNil(o.IntervalEndMs) {
		return true
	}

	return false
}

// SetIntervalEndMs gets a reference to the given int64 and assigns it to the IntervalEndMs field.
func (o *GroupsGetEntitiesRequest) SetIntervalEndMs(v int64) {
	o.IntervalEndMs = &v
}

// GetGroupingAttributeType returns the GroupingAttributeType field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetGroupingAttributeType() string {
	if o == nil || IsNil(o.GroupingAttributeType) {
		var ret string
		return ret
	}
	return *o.GroupingAttributeType
}

// GetGroupingAttributeTypeOk returns a tuple with the GroupingAttributeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetGroupingAttributeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GroupingAttributeType) {
		return nil, false
	}
	return o.GroupingAttributeType, true
}

// HasGroupingAttributeType returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasGroupingAttributeType() bool {
	if o != nil && !IsNil(o.GroupingAttributeType) {
		return true
	}

	return false
}

// SetGroupingAttributeType gets a reference to the given string and assigns it to the GroupingAttributeType field.
func (o *GroupsGetEntitiesRequest) SetGroupingAttributeType(v string) {
	o.GroupingAttributeType = &v
}

// GetGroupMemberCount returns the GroupMemberCount field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetGroupMemberCount() int64 {
	if o == nil || IsNil(o.GroupMemberCount) {
		var ret int64
		return ret
	}
	return *o.GroupMemberCount
}

// GetGroupMemberCountOk returns a tuple with the GroupMemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetGroupMemberCountOk() (*int64, bool) {
	if o == nil || IsNil(o.GroupMemberCount) {
		return nil, false
	}
	return o.GroupMemberCount, true
}

// HasGroupMemberCount returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasGroupMemberCount() bool {
	if o != nil && !IsNil(o.GroupMemberCount) {
		return true
	}

	return false
}

// SetGroupMemberCount gets a reference to the given int64 and assigns it to the GroupMemberCount field.
func (o *GroupsGetEntitiesRequest) SetGroupMemberCount(v int64) {
	o.GroupMemberCount = &v
}

// GetGroupMemberSortOrder returns the GroupMemberSortOrder field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetGroupMemberSortOrder() string {
	if o == nil || IsNil(o.GroupMemberSortOrder) {
		var ret string
		return ret
	}
	return *o.GroupMemberSortOrder
}

// GetGroupMemberSortOrderOk returns a tuple with the GroupMemberSortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetGroupMemberSortOrderOk() (*string, bool) {
	if o == nil || IsNil(o.GroupMemberSortOrder) {
		return nil, false
	}
	return o.GroupMemberSortOrder, true
}

// HasGroupMemberSortOrder returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasGroupMemberSortOrder() bool {
	if o != nil && !IsNil(o.GroupMemberSortOrder) {
		return true
	}

	return false
}

// SetGroupMemberSortOrder gets a reference to the given string and assigns it to the GroupMemberSortOrder field.
func (o *GroupsGetEntitiesRequest) SetGroupMemberSortOrder(v string) {
	o.GroupMemberSortOrder = &v
}

// GetQueryName returns the QueryName field value if set, zero value otherwise.
func (o *GroupsGetEntitiesRequest) GetQueryName() string {
	if o == nil || IsNil(o.QueryName) {
		var ret string
		return ret
	}
	return *o.QueryName
}

// GetQueryNameOk returns a tuple with the QueryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsGetEntitiesRequest) GetQueryNameOk() (*string, bool) {
	if o == nil || IsNil(o.QueryName) {
		return nil, false
	}
	return o.QueryName, true
}

// HasQueryName returns a boolean if a field has been set.
func (o *GroupsGetEntitiesRequest) HasQueryName() bool {
	if o != nil && !IsNil(o.QueryName) {
		return true
	}

	return false
}

// SetQueryName gets a reference to the given string and assigns it to the QueryName field.
func (o *GroupsGetEntitiesRequest) SetQueryName(v string) {
	o.QueryName = &v
}

func (o GroupsGetEntitiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsGetEntitiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumberOfIntervalsForLatestData) {
		toSerialize["number_of_intervals_for_latest_data"] = o.NumberOfIntervalsForLatestData
	}
	if !IsNil(o.GroupSortAttribute) {
		toSerialize["group_sort_attribute"] = o.GroupSortAttribute
	}
	if !IsNil(o.NumberOfBuckets) {
		toSerialize["number_of_buckets"] = o.NumberOfBuckets
	}
	if !IsNil(o.EntityIds) {
		toSerialize["entity_ids"] = o.EntityIds
	}
	if !IsNil(o.GroupMemberAttributes) {
		toSerialize["group_member_attributes"] = o.GroupMemberAttributes
	}
	if !IsNil(o.GroupMemberSortAttribute) {
		toSerialize["group_member_sort_attribute"] = o.GroupMemberSortAttribute
	}
	if !IsNil(o.BucketBoundary) {
		toSerialize["bucket_boundary"] = o.BucketBoundary
	}
	if !IsNil(o.GroupOffset) {
		toSerialize["group_offset"] = o.GroupOffset
	}
	if !IsNil(o.DownsamplingInterval) {
		toSerialize["downsampling_interval"] = o.DownsamplingInterval
	}
	if !IsNil(o.IntervalStartMs) {
		toSerialize["interval_start_ms"] = o.IntervalStartMs
	}
	toSerialize["entity_type"] = o.EntityType
	if !IsNil(o.GroupMemberOffset) {
		toSerialize["group_member_offset"] = o.GroupMemberOffset
	}
	if !IsNil(o.GroupingAttribute) {
		toSerialize["grouping_attribute"] = o.GroupingAttribute
	}
	if !IsNil(o.GroupMemberSortDownsamplingFunction) {
		toSerialize["group_member_sort_downsampling_function"] = o.GroupMemberSortDownsamplingFunction
	}
	if !IsNil(o.GroupSortOrder) {
		toSerialize["group_sort_order"] = o.GroupSortOrder
	}
	if !IsNil(o.GroupSortDownsampleFunction) {
		toSerialize["group_sort_downsample_function"] = o.GroupSortDownsampleFunction
	}
	if !IsNil(o.FilterCriteria) {
		toSerialize["filter_criteria"] = o.FilterCriteria
	}
	if !IsNil(o.LargeBucketBoundary) {
		toSerialize["large_bucket_boundary"] = o.LargeBucketBoundary
	}
	if !IsNil(o.AvailabilityZoneScope) {
		toSerialize["availability_zone_scope"] = o.AvailabilityZoneScope
	}
	if !IsNil(o.GroupCount) {
		toSerialize["group_count"] = o.GroupCount
	}
	if !IsNil(o.GroupAttributes) {
		toSerialize["group_attributes"] = o.GroupAttributes
	}
	if !IsNil(o.IntervalEndMs) {
		toSerialize["interval_end_ms"] = o.IntervalEndMs
	}
	if !IsNil(o.GroupingAttributeType) {
		toSerialize["grouping_attribute_type"] = o.GroupingAttributeType
	}
	if !IsNil(o.GroupMemberCount) {
		toSerialize["group_member_count"] = o.GroupMemberCount
	}
	if !IsNil(o.GroupMemberSortOrder) {
		toSerialize["group_member_sort_order"] = o.GroupMemberSortOrder
	}
	if !IsNil(o.QueryName) {
		toSerialize["query_name"] = o.QueryName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupsGetEntitiesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entity_type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGroupsGetEntitiesRequest := _GroupsGetEntitiesRequest{}

	err = json.Unmarshal(data, &varGroupsGetEntitiesRequest)

	if err != nil {
		return err
	}

	*o = GroupsGetEntitiesRequest(varGroupsGetEntitiesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "number_of_intervals_for_latest_data")
		delete(additionalProperties, "group_sort_attribute")
		delete(additionalProperties, "number_of_buckets")
		delete(additionalProperties, "entity_ids")
		delete(additionalProperties, "group_member_attributes")
		delete(additionalProperties, "group_member_sort_attribute")
		delete(additionalProperties, "bucket_boundary")
		delete(additionalProperties, "group_offset")
		delete(additionalProperties, "downsampling_interval")
		delete(additionalProperties, "interval_start_ms")
		delete(additionalProperties, "entity_type")
		delete(additionalProperties, "group_member_offset")
		delete(additionalProperties, "grouping_attribute")
		delete(additionalProperties, "group_member_sort_downsampling_function")
		delete(additionalProperties, "group_sort_order")
		delete(additionalProperties, "group_sort_downsample_function")
		delete(additionalProperties, "filter_criteria")
		delete(additionalProperties, "large_bucket_boundary")
		delete(additionalProperties, "availability_zone_scope")
		delete(additionalProperties, "group_count")
		delete(additionalProperties, "group_attributes")
		delete(additionalProperties, "interval_end_ms")
		delete(additionalProperties, "grouping_attribute_type")
		delete(additionalProperties, "group_member_count")
		delete(additionalProperties, "group_member_sort_order")
		delete(additionalProperties, "query_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupsGetEntitiesRequest struct {
	value *GroupsGetEntitiesRequest
	isSet bool
}

func (v NullableGroupsGetEntitiesRequest) Get() *GroupsGetEntitiesRequest {
	return v.value
}

func (v *NullableGroupsGetEntitiesRequest) Set(val *GroupsGetEntitiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsGetEntitiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsGetEntitiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsGetEntitiesRequest(val *GroupsGetEntitiesRequest) *NullableGroupsGetEntitiesRequest {
	return &NullableGroupsGetEntitiesRequest{value: val, isSet: true}
}

func (v NullableGroupsGetEntitiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsGetEntitiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


