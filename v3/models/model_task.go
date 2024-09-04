/*
Nutanix Intentful API

# Introduction  **Representational state transfer ( REST )** Application Programming Interface (API) 3.0 is based on an intentful API philosophy. According to the intentful API philosophy the machine should handle the programming instead of the user enabling the datacenter administrator able to focus on the other tasks. You need to specify the end state of an entity and the system will compile and execute a series of steps to achieve the defined end state of the entity. The progress to achieve the desired state is tracked through waits and events. This approach is similar to the Google Kubernetes standard.  All the entities and the list of entities follow a homogenous specification format resulting in ease of programming despite different entities being involved. Also, this enables the user to get familiar with the APIs quickly.  Intentful APIs reduces the number of action verbs required to achieve the desired state of an entity. Most of the changes can be achieved by defining the desired state of an entity rather than through a series of action verbs.  All API users should have valid Prism login credentials to send API calls to the Prism server.  APIs on nutanix.dev portal are publicly available to all valid users without special permissions for viewing purposes.  This document covers APIs that are available in pc2022.11 release.    ## Terminology  | Item                  | Description                                                                                                                                                             | | --------------------- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | | Entity                | An object that is managed through an API.                                                                                                                               | | Kind                  | Represents the type of entity. For example, VM or alerts.                                                                                                               | | Resource Kind         | An entity that represents a physical or virtual infrastructure resource. For example, VM or an application.                                                             | | Helper Kind           | An entity that is not an infrastructure resource, but represents the entity that is used along with the infrastructure resources. For example, alert, event, or status. | | Basic Resource Kind   | An entity type that is a primary resource. For example, VM.                                                                                                             | | Related Resource Kind | An entity type that is automatically defined by the system and is related to the basic Kind entities. For example, vm_snapshot.                                         | | Type                  | Represents sub-categories or sub-objects of the top-level resource types. For example, reference.                                                                       | | Resource Limit        | An entity type that represents a quota of resources of different types.                                                                                                 | | Snapshot              | Represents the state of resources and policies of an entity at a particular point of time.                                                                              | | Backup                | An automatically created aggregation object that collects all the snapshots for the same object within the same resource pool.                                          | | Profile               | A profile is a set of resources and policy specifications that is applied to an entity.                                                                                 | | Template              | A template is a specification that generates valid entities by providing input parameters.                                                                              | | Spec                  | A section in the request json, which represents an entity, that expresses the desired state of the entity.                                                              | | Status                | A section in the json representation that expresses the current state of the entity.                                                                                    |  ## Rules  - The metadata section has fields that relate to all kinds of entities and can be updated instantly by the system. - The spec section cannot be updated after initialization by the system without a user or automation intervention. - The status section has a copy of the currently enforced version of the spec section and other fields that are managed or updated automatically by the system. Few fields in the status section can be updated by the user as well. - The parent_reference field is mandatory for snapshot and backup entities and is present in the normal entities, if the entities are created with the parent_reference field populated (clone or restore). - Any field that represents a defined type has the type name as a suffix. For example, backup_factory and vm_reference.  ## Known Issues and limitations:  This section describes some of the issues identified in V3 APIs that you might encounter.  - **ENG-321951**- While attempting to get a list of all VMs on Prism Central, we have observed that only AHV VMs are returned, but ESXi VM list is not retrieved.  - **ENG-453487**- Users with insufficient permissions receive an HTTP 200 response instead of an HTTP 403 response for view alert API requests. These users can be like \"Developer\", \"Operator\", or custom roles that do not have sufficient permission to view alerts. This is a generic issue for all V3 API requests with insufficient permissions.   ### [Nutanix API versions – What are they and what does each one do?](https://developer.nutanix.com/2019/01/15/nutanix-api-versions-what-are-they-and-what-does-each-one-do/)   ### [What to expect post pc.2022.1 release?](https://www.nutanix.dev/api_references/prism-central-v3/#/ZG9jOjM2MzQyMjUz-v3-api-behavior-changes-with-pc-2022-1-release)  # Getting Started  ## Components of a REST API  A REST API request and response can be classified into the following components.  The request URI consists of {URI-scheme} :// {URI-host} / {<Port-number>/<Base-path>} ? {query-string}.  | Item                    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | | ----------------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | | URI scheme              | Indicates the protocol that is used to transmit the request. For example, http or https.                                                                                                                                                                                                                                                                                                                                                                                                                                    | | URI host                | Specifies the domain name or IP address of the server where the REST service endpoint is hosted. For example, prism.nutanix.com or 10.11.12.13.                                                                                                                                                                                                                                                                                                                                                                             | | Port number             | Indicates the port number that is used for HTTP communications from the REST API clients to the REST API server.                                                                                                                                                                                                                                                                                                                                                                                                            | | Basepath                | Indicates the URL prefix that is common to all the API server endpoints. For example, for <https://10.11.12.13/PrismGateway/services/rest/v2.0/vms>, the basepath is /PrismGateway/services/rest/v2.0.                                                                                                                                                                                                                                                                                                                      | | Path parameters         | Path parameters are part of the endpoint. For example, <https://10.11.12.13/PrismGateway/services/rest/v2.0/vms/750eb5ba-2de0-4ca4-8d6f-786364817f94/>, the UUID of the VM (750eb5ba-2de0-4ca4-8d6f-786364817f94) is a path parameter.                                                                                                                                                                                                                                                                                      | | Query string parameters | Query string parameters appear after a question mark (?) in the endpoint. The question mark followed by the parameters and their values is referred as the query string. In the query string, each parameter is listed one right after the other with an ampersand (&). The order of the query string parameters does not matter. For example, <https://10.11.12.13/PrismGateway/services/rest/v2.0/vms/750eb5ba-2de0-4ca4-8d6f-786364817f94?include_vm_disk_config=true>, the include_vm_disk_config is a query parameter. |  HTTP defines methods to indicate the desired action to be performed on the identified resource.  | Method | Description                                    | | ------ | :--------------------------------------------- | | Get    | Retrieves a resource                           | | Post   | Creates a resource                             | | Put    | Updates or creates within an existing resource | | Delete | Removes the resource                           |  HTTP request headers provide meta-information about a request. The request body contains the data that client wants to send to the server. For example, Content-Type and Content-Transfer-Encoding.  HTTP response headers provide meta-information about a response. The response consists of status code (three-digit number), content type, and body of the response.  | Item                | Description                                                                      | | ------------------- | :------------------------------------------------------------------------------- | | 1xx - Informational | Communicates transfer protocol-level information.                                | | 2xx - Success       | Indicates that the client request is successfully accepted.                      | | 3xx - Redirection   | Indicates that the client must take some additional steps to complete an action. | | 4xx - Client Error  | Indicates that there is some error from the client side.                         | | 5xx - Server Error  | Indicates that there is some error from the server side.                         |  ## Entity Structure  An entity structure comprises of the following components.  - **Metadata** - Metadata provides high-level information about the related entity including entity uuid, kind, and spec_version. - **Status** - Status provides information about the current state of the related entity including entity name, state, and entity-specific status information. - **Specification** - For a GET request, the specification section of the JSON response displays similar information as status section. However, for of POST request, the specification section provide details of what should be the desired state of the entity. For example, you can use the following JSON payload to create a basic VM. The payload denotes the minimum parameters required to create a VM by using the V3 APIs. All the missing specification items like CPU and RAM are configured with default values.  ``` { \"spec\":{ \"name\":\"NewVM\", \"resources\":{ } }, \"metadata\":{ \"kind\":\"vm\" } } ```  # Authentication  All requests to the Nutanix REST APIs must be authenticated. The following authentication types are supported.  - **Basic Authentication**: The user provides username and password every time a request is sent as the authentication header. - **Session Authentication**: The user credentials are stored in a cookie.  **Cookie Consumption**  To reduce the number of cookies generated in the system, Nutanix recommends that you make the first API call using basic credentials to receive the cookies and then instead of sending the credentials, cache and send the following cookies to authenticate subsequent requests.  **For Prism Central with IAMv2 enabled:**  - NTNX_MERCURY_IAM_SESSION - NTNX_IAM_SESSION - NTNX_MERCURY_IAM_REFRESH_TOKEN  **For Prism Central with IAMv2 disabled or Prism Element:**  - NTNX_IGW_SESSION - NTNX_MERCURY_IGW_SESSION - JSESSIONID(If present)  **Note:** Upon cookie expiration, you must send basic credentials again in the next API call, to obtain the new cookies.  The following are the code snippets to authenticate a user.  **Python**  ``` #!/usr/bin/env python3.7 import urllib3 import requests from base64 import b64encode urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning) request_url = 'https://10.0.0.1:9440/api/nutanix/v3/vms/list' username = 'admin' password = 'password' encoded_credentials = b64encode(bytes(f'{username}:{password}', encoding='ascii')).decode('ascii') auth_header = f'Basic {encoded_credentials}' payload = '{\"kind\":\"vm\"}' headers = {'Accept': 'application/json', 'Content-Type': 'application/json', 'Authorization': f'{auth_header}', 'cache-control': 'no-cache'} response = requests.request('post', request_url, data=payload, headers=headers, verify=False) print(response.status_code) ```  **C#**  ``` using System; using System.Net; using System.Text; namespace ConsoleApp1 { class Program { static void Main(string[] args) { string ClusterUsername = \"admin\"; string ClusterPassword = \"Ntnx2017!\"; string RequestUrl = \"https://192.168.1.131:9440/api/nutanix/v3/vms/list\"; string Payload = \"{\\\"kind\\\":\\\"vm\\\"}\"; string AuthHeader = System.Convert.ToBase64String( System.Text.ASCIIEncoding.ASCII.GetBytes( ClusterUsername + \":\" + ClusterPassword)); ServicePointManager.ServerCertificateValidationCallback = ( (sender, certificate, chain, sslPolicyErrors) => true); ServicePointManager.Expect100Continue = true; ServicePointManager.SecurityProtocol = SecurityProtocolType.Tls12; String authHeader = System.Convert.ToBase64String( System.Text.ASCIIEncoding.ASCII.GetBytes( ClusterUsername + \":\" + ClusterPassword)); var request = (HttpWebRequest)WebRequest.Create(RequestUrl); var requestBody = Encoding.ASCII.GetBytes(Payload); HttpWebResponse HttpResponse = null; request.Headers.Add(\"Authorization\", \"Basic \" + AuthHeader); request.Headers.Add(\"Accept-Encoding\", \"application/json\"); request.ContentType = \"application/json\"; request.Accept = \"application/json\"; request.Method = \"POST\"; var newStream = request.GetRequestStream(); newStream.Write(requestBody, 0, requestBody.Length); newStream.Close(); HttpResponse = (HttpWebResponse)request.GetResponse(); Console.WriteLine(HttpResponse.StatusCode); Console.ReadKey(); } } } ```  **PowerShell**  ``` # set the basic properties for the request $RequestUri = \"https://10.0.0.1:9440/api/nutanix/v3/vms/list\" $Payload = \"{\"\"kind\"\":\"\"vm\"\"}\" $Username = \"admin\" $Password = \"password\" # create the HTTP Basic Authorization header $pair = $Username + \":\" + $Password $bytes = [System.Text.Encoding]::ASCII.GetBytes($pair) $base64 = [System.Convert]::ToBase64String($bytes) $basicAuthValue = \"Basic $base64\" # setup the request headers $Headers = @{ 'Accept' = 'application/json' 'Authorization' = $basicAuthValue 'Content-Type' = 'application/json' } # disable SSL certification verification # you probably shouldn't do this in production ... if (-not ( [System.Management.Automation.PSTypeName] ` 'ServerCertificateValidationCallback').Type) {$certCallback = @\" using System; using System.Net; using System.Net.Security; using System.Security.Cryptography.X509Certificates; public class ServerCertificateValidationCallback { public static void Ignore() { if(ServicePointManager.ServerCertificateValidationCallback ==null) { ServicePointManager.ServerCertificateValidationCallback += delegate ( Object obj, X509Certificate certificate, X509Chain chain, SslPolicyErrors errors ) { return true; }; } } } \"@ Add-Type $certCallback } [ServerCertificateValidationCallback]::Ignore() [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12 # submit the request Invoke-WebRequest ` -Uri $RequestUri ` -Headers $Headers ` -Method \"POST\" ` -Body $Payload ` -TimeoutSec 5 ` -UseBasicParsing ` -DisableKeepAlive ` | ConvertFrom-Json | Select * ```  **PHP**  ``` <?php $username = 'admin'; $password = 'password'; $ch = curl_init(); curl_setopt_array($ch, [ CURLOPT_RETURNTRANSFER => 1, CURLINFO_HEADER_OUT => TRUE, CURLOPT_URL => 'https://10.0.0.1:9440/api/nutanix/v3/vms/list', CURLOPT_USERPWD => $username . ':' . $password, CURLOPT_POST => 1, CURLOPT_SSL_VERIFYHOST => FALSE, CURLOPT_SSL_VERIFYPEER => FALSE, CURLOPT_HTTPHEADER => [ 'Content-Type: application/json', 'Accept-Encoding: application/json', ], CURLOPT_POSTFIELDS => json_encode([ 'kind' => 'vm' ]) ]); $result = curl_exec($ch); curl_close($ch); if(!$result) { die('Error: \"' . curl_error($ch) . '\". Code: ' . curl_errno($ch)); } else { echo($result); } ``` # pc.2022.1 release  ## v3 API behavior changes with pc.2022.1 release   From Prism Central release 2022.1 onwards, the v3 intentful API is enhanced to protect data consistency, especially in concurrent race conditions. As a result of this enhancement, response time for the synchronous API requests is improved and there are some behavior changes from an API consumption perspective.   1.  In the current v3 API workflow, clients will get a synchronous API response that contains task and entity UUID. Post which clients should send GET API on the task periodically until the task succeeds or fails. Clients can also send GET API on the entity.         - **Behavioral changes**: Some of the synchronous errors will be reported asynchronously now.            1. **Invalid idempotency UUID error**:                       **Current API behavior**: If POST API contains invalid idempotency UUID, clients will receive a     synchronous bad request error with an HTTP code 400.           **Expected change**: If POST API contains invalid idempotency UUID, clients will get a synchronous API response with an HTTP code 202(accepted). GET API on the task will return a bad request error.      1. **Authorization errors**:            **Current API behavior**: Clients will get synchronous authorization errors (for example- HTTP code 403 forbidden error).            **Expected change**: Clients will get a synchronous API response with an HTTP code 202 (accepted). GET API on the task will show an authorization error.       1. **Incorrect category mapping error**:            **Current API behavior**: If an intentful API payload has a project reference and uses categories mapping but the category for the project does not exist, the client will get an error response with an HTTP code 404(not found).            **Expected change**: If an intentful API payload has a project reference and uses categories mapping but the category for the project does not exist, the client will receive a synchronous API response with an HTTP code 202 (accepted). GET API on the task will show a “not found” error.      1.  **Invalid user detail error**:            **Current API behavior**: If an intentful API payload has user_uuid or username in owner_reference but the corresponding user does not exist, the client will receive an API response with an HTTP code 404.            **Expected change**: If an intentful API payload has user_uuid or username in owner_reference but the corresponding user does not exist, clients will get a synchronous API response with an HTTP code 202(accepted). GET API on the task will show a bad request error.         1.  . **Missing user detail error**:           **Current API behavior**: If an intentful API payload does not have user_uuid or username in owner_reference and the login user does not exist, the client will get an API response with an HTTP code 404.           **Expected change**: If an intentful API payload does not have user_uuid or username in owner_reference and login user does not exist, clients will get a synchronous API response with an HTTP code 202(accepted). GET API on the task will show a bad request error.      -  **Format changes** for some of the synchronous API responses:                  **Current API behavior**: Response to an intentful POST, PUT, or DELETE API will contain tenant_uuid, owner_uuid, owner_username, service_name, session_log_uuid, trace, user_uuid, username, remote_authorization, session_json_dict, incap_req_id, project_reference or category information.                  **Expected change**: Main processing will occur after a synchronous response to the client. As a result,  synchronous response to an intentful POST, PUT, or DELETE API will not contain tenant_uuid, owner_uuid, owner_username, service_name, session_log_uuid, trace, user_uuid, username, remote_authorization, session_json_dict, incap_req_id, project_reference or category information.                  If API consumers are expecting these properties in the synchronous response before, they have to submit a GET/LIST request to inspect these properties.  2.  The version control workflow for v3 intentful PUT and DELETE APIs:    - **Absolutely no version control check for DELETE APIs**:                  **Current API behavior**: DELETE API will fail if an incorrect spec_version is provided as one of the request arguments.                  **Expected change**:  DELETE API will always delete the entity. No version control is applicable on the DELETE request.       - **Edit conflict error could be reported asynchronously**:        **Current API behavior**: If v3 PUT APIs are conflicting with v2 PUT APIs on the same entities, the client will receive an HTTP code 409 response for v3 PUT APIs.       **Expected change**: If v3 PUT APIs are conflicting with v2 PUT APIs on the same entities, the client will receive an HTTP code 202 response for v3 PUT APIs and will receive the task failed response with an edit conflict error asynchronously.     - **More restrictive version control by using spec_hash**:                 **Current API behavior**: GET/LIST API will return a response that may contain one of the values for version control, it could be spec_version or spec_hash. spec_version is numeric while spec_hash is a string.                  Whichever value is provided in the GET/LIST API, the same value should be included in the next PUT API payload of the same entity.        **Expected change**: GET/LIST API will provide both spec_hash and spec_version in API response. Use of spec_hash is not enforced but could be enforced in future releases.        - **The version control value in the GET API response will not change when a PUT request is in progress**:                  **Current API behavior**: If there are one GET request and one PUT request on the same entity, the version control value from the GET request response might be different depending on the current state of the entity. The behavior is not deterministic.          **Expected change**: If there is one GET request and one PUT request on the same entity, the spec_hash value of the GET request response will always be the same as it was before the PUT request until the PUT is completed.    1. V3 intentful POST API and the following GET API:    -  **Get API request for new entity**:        **Current API behavior**: If a POST API fails asynchronously, the following GET API on the same entity will get an HTTP code 202 response with the entity in an error state.          However, GET API on the same entity after 5+ minutes will show an HTTP code 404 response reporting the entity does not exist.        **Expected change**: If a POST API fails asynchronously, GET API on the same entity will always show an HTTP code 404 response meaning the entity does not exist.        - **The intent spec is an entity we created to record previous and ongoing requests from v3 POST/PUT/DELETE APIs**. There should be only one intent spec for each entity if the intent spec exists:          **Current API behavior:** If the client queries intent spec for an entity immediately after getting an HTTP code 202 response for the v3 POST API on this entity, the intent spec will always be present.          **Expected change**: If the client queries intent spec for an entity immediately after getting an HTTP code 202 response for the v3 POST API on this entity, the intent spec may not be created yet.  4. Prism Central upgrade to pc.2022.1 and following releases:        **Current API behavior**: The pending and running API requests will continue to execute after an upgrade.        **Expected change**: During the pc.2022.1 version upgrade, all pending and running API requests and related tasks will be marked as failed after a restart. This will be shown even though they may be in progress and completed in the middle of the upgrade period.       This behavior only applies to upgrades from releases before pc.2022.1 to releases equal to or after pc.2022.1. All future upgrades from pc.2022.1 release will remain the same as before where pending and running API requests will continue to execute after an upgrade.  5. Multiple nodes upgrade:   -   pc.2022.1 version can take care of the scenarios where Prism Central has the new version and Prism Element is still at the older version.        -   Any task (for existing create, update, delete v3 API operations) that is in running state during the upgrade will be marked failed. This also applies to new requests that come before the upgrade ends. The recommendation is to wait until the upgrade finishes before executing any v3 create, update, delete API.     6. Schema consistency:    - API schema in Prism Central will be used in case there is a      discrepancy between Prism Element and Prism Central. There will be a      change of behavior when Prism Central has a post pc.2022.1 version.        **Current API behavior**: API responses for POST, PUT, and DELETE API on Prism Central will match API schema on Prism Element.            **Expected change**: API response for POST, PUT, and DELETE API on Prism Central will match API schema on Prism Central.  # Use Cases            ## Post /VMs  Perform the following procedure to generate the code snippet for creating a VM based on the input parameters.  1. Under **VMs**, click **POST VMs** and navigate to the end of the page. 2. In the **Code Generation** panel, under the **Settings** tab, enter the value of the following parameters.    1. Enter the username of the Prism Central in the **Username** field.    2. Enter the password in the **Password** field.    3. Enter the IP address of the Prism Central in the **Host** field. 3. Under the **Body** tab, enter the value of all the required parameters as described in the **Request Body** section. Optionally, you can also enter the values of the optional parameters.    1. Under **Specification**, enter the description of VM in a string format against the **description** attribute. The maximum length allowed is 1000 characters only.    2. Under **Resources** do the following.       1. Enter the value of number of logical threads per core in an integer format against the **num_threads_per_core** attribute. The minimum value allowed is 1.       2. Enter the value of current or desired power state of the VM in a string format against the **power_state** attribute.       3. Enter the value of number of vCPUs per socket in an integer format against the **num_vcpus_per_socket** attribute. The minimun value allowed is 1.       4. Enter the value of number of vCPU sockets in an integer format against the **num_sockets** attribute. The minimum value allowed is 1.    3. Under **Name**, enter the name of the VM in a string format. The maximum length allowed is 64 characters only.    4. Under **Metadata** do the following.       1. The **last_update_time** is a read-only attribute indicating UTC date and time in an RFC-339 format, when the VM was last updated in a string format.       2. The **kind** is a read-only attribute indicating the kind name in a string format. The default value is VM.       3. Enter the value of VM uuid in a string format against the **uuid** attribute. The valid pattern is ^\\[a-fA-F0-9]{8}-\\[a-fA-F0-9]{4}-\\[a-fA-F0-9]{4}-\\[a-fA-F0-9]{4}-\\[a-fAF0-9]{12}$. 4. Under the **Code Generation** tab, select the scripting languages to view the code snippet of creating a VM.     You can either select Python or PowerShell or Shell or Go or Java or JavaScript or Node or Obj-C or PHP or Ruby or Swift or C# or C.     The code snippet appears in the selected scripting language. You can use the code snippet to create a VM in your environment.  ## Put /VMs  Perform the following procedure to generate the code snippet for updating an existing VM based on the input parameters.  1. Under **VMs**, click **PUT VMs** and navigate to the end of the page. 2. In the **Code Generation** panel, under the **Settings** tab, enter the value of the following parameters.    1. Enter the UUID of the VM that needs to be updated in the **UUID** field.    2. Enter the username of the Prism Central in the **Username** field.    3. Enter the password in the **Password** field.    4. Enter the IP address of the Prism Central in the **Host** field. 3. Under the **Body** tab, enter the value of all the required parameters as described in the **Request Body** section. Optionally, you can also enter the values of the optional parameters.    1. Under **Specification**, enter the description of VM in a string format against the **description** attribute. The maximum length allowed is 1000 characters only.    2. Under **Resources** do the following.       1. Enter the value of number of logical threads per core in an integer format against the **num_threads_per_core** attribute. The minimum value allowed is 1.       2. Enter the value of current or desired power state of the VM in a string format against the **power_state** attribute.       3. Enter the value of number of vCPUs per socket in an integer format against the **num_vcpus_per_socket** attribute. The minimun value allowed is 1.       4. Enter the value of number of vCPU sockets in an integer format against the **num_sockets** attribute. The minimum value allowed is 1.    3. Under **Name**, enter the name of the VM in a string format. The maximum length allowed is 64 characters only.    4. Under **Metadata** do the following.       1. The **last_update_time** is a read-only attribute indicating UTC date and time in an RFC-339 format, when the VM was last updated in a string format.       2. The **kind** is a read-only attribute indicating the kind name in a string format. The default value is VM.       3. Enter the value of VM uuid in a string format against the **uuid** attribute. The valid pattern is ^\\[a-fA-F0-9]{8}-\\[a-fA-F0-9]{4}-\\[a-fA-F0-9]{4}-\\[a-fA-F0-9]{4}-\\[a-fAF0-9]{12}$. 4. Under the **Code Generation** tab, select the scripting languages to view the code snippet of creating a VM.     You can either select Python or PowerShell or Shell or Go or Java or JavaScript or Node or Obj-C or PHP or Ruby or Swift or C# or C.     The code snippet appears in the selected scripting language. You can use the code snippet to update a VM in your environment. 

API version: 3.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"time"
)

// checks if the Task type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Task{}

// Task Task details.
type Task struct {
	// Current state of the task.
	Status *string `json:"status,omitempty"`
	// UTC date and time in RFC-3339 format when task was last updated. 
	LastUpdateTime *time.Time `json:"last_update_time,omitempty"`
	// In case of task failure this field will provide the error description. 
	ErrorDetail *string `json:"error_detail,omitempty"`
	// Number of times the task has been updated. The value increases sequentially with each update of the task and can be used to verify if there have been changes to the task. 
	LogicalTimestamp *int64 `json:"logical_timestamp,omitempty"`
	// Final expected state of the task. It is set when the task is aborted. 
	RequestedStatus *string `json:"requested_status,omitempty"`
	EntityReferenceList []Reference `json:"entity_reference_list,omitempty"`
	// UTC date and time in RFC-3339 format when Task execution started. 
	StartTime *time.Time `json:"start_time,omitempty"`
	// UTC date and time in RFC-3339 format when task was created. 
	CreationTime *time.Time `json:"creation_time,omitempty"`
	// UUID of the task.
	UUID *string `json:"uuid,omitempty" validate:"regexp=^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$"`
	// Time in microseconds from epoch when the task execution started. 
	StartTimeUsecs *int64 `json:"start_time_usecs,omitempty"`
	ClusterReference *ClusterReference `json:"cluster_reference,omitempty"`
	// Reference to the sub-tasks.
	SubtaskReferenceList []TaskReference `json:"subtask_reference_list,omitempty"`
	// UTC date and time in RFC-3339 format when Task execution completed. 
	CompletionTime *time.Time `json:"completion_time,omitempty"`
	// Time in microseconds from epoch when task was created. 
	CreationTimeUsecs *int64 `json:"creation_time_usecs,omitempty"`
	// Description of what currently the task is doing.
	ProgressMessage *string `json:"progress_message,omitempty"`
	// Type of the operation tracked by the task.
	OperationType *string `json:"operation_type,omitempty"`
	// Time in microseconds from epoch when task execution completed. 
	CompletionTimeUsecs *int64 `json:"completion_time_usecs,omitempty"`
	// In case of task failure this field will provide the error code. 
	ErrorCode *string `json:"error_code,omitempty"`
	// The completion percentage for the task.
	PercentageComplete *int32 `json:"percentage_complete,omitempty"`
	// API Version of the Nutanix v3 API framework.
	ApiVersion *string `json:"api_version,omitempty"`
	ParentTaskReference *TaskReference `json:"parent_task_reference,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Task Task

// NewTask instantiates a new Task object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTask() *Task {
	this := Task{}
	return &this
}

// NewTaskWithDefaults instantiates a new Task object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskWithDefaults() *Task {
	this := Task{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Task) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Task) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Task) SetStatus(v string) {
	o.Status = &v
}

// GetLastUpdateTime returns the LastUpdateTime field value if set, zero value otherwise.
func (o *Task) GetLastUpdateTime() time.Time {
	if o == nil || IsNil(o.LastUpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdateTime
}

// GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetLastUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdateTime) {
		return nil, false
	}
	return o.LastUpdateTime, true
}

// HasLastUpdateTime returns a boolean if a field has been set.
func (o *Task) HasLastUpdateTime() bool {
	if o != nil && !IsNil(o.LastUpdateTime) {
		return true
	}

	return false
}

// SetLastUpdateTime gets a reference to the given time.Time and assigns it to the LastUpdateTime field.
func (o *Task) SetLastUpdateTime(v time.Time) {
	o.LastUpdateTime = &v
}

// GetErrorDetail returns the ErrorDetail field value if set, zero value otherwise.
func (o *Task) GetErrorDetail() string {
	if o == nil || IsNil(o.ErrorDetail) {
		var ret string
		return ret
	}
	return *o.ErrorDetail
}

// GetErrorDetailOk returns a tuple with the ErrorDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetErrorDetailOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDetail) {
		return nil, false
	}
	return o.ErrorDetail, true
}

// HasErrorDetail returns a boolean if a field has been set.
func (o *Task) HasErrorDetail() bool {
	if o != nil && !IsNil(o.ErrorDetail) {
		return true
	}

	return false
}

// SetErrorDetail gets a reference to the given string and assigns it to the ErrorDetail field.
func (o *Task) SetErrorDetail(v string) {
	o.ErrorDetail = &v
}

// GetLogicalTimestamp returns the LogicalTimestamp field value if set, zero value otherwise.
func (o *Task) GetLogicalTimestamp() int64 {
	if o == nil || IsNil(o.LogicalTimestamp) {
		var ret int64
		return ret
	}
	return *o.LogicalTimestamp
}

// GetLogicalTimestampOk returns a tuple with the LogicalTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetLogicalTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.LogicalTimestamp) {
		return nil, false
	}
	return o.LogicalTimestamp, true
}

// HasLogicalTimestamp returns a boolean if a field has been set.
func (o *Task) HasLogicalTimestamp() bool {
	if o != nil && !IsNil(o.LogicalTimestamp) {
		return true
	}

	return false
}

// SetLogicalTimestamp gets a reference to the given int64 and assigns it to the LogicalTimestamp field.
func (o *Task) SetLogicalTimestamp(v int64) {
	o.LogicalTimestamp = &v
}

// GetRequestedStatus returns the RequestedStatus field value if set, zero value otherwise.
func (o *Task) GetRequestedStatus() string {
	if o == nil || IsNil(o.RequestedStatus) {
		var ret string
		return ret
	}
	return *o.RequestedStatus
}

// GetRequestedStatusOk returns a tuple with the RequestedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetRequestedStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RequestedStatus) {
		return nil, false
	}
	return o.RequestedStatus, true
}

// HasRequestedStatus returns a boolean if a field has been set.
func (o *Task) HasRequestedStatus() bool {
	if o != nil && !IsNil(o.RequestedStatus) {
		return true
	}

	return false
}

// SetRequestedStatus gets a reference to the given string and assigns it to the RequestedStatus field.
func (o *Task) SetRequestedStatus(v string) {
	o.RequestedStatus = &v
}

// GetEntityReferenceList returns the EntityReferenceList field value if set, zero value otherwise.
func (o *Task) GetEntityReferenceList() []Reference {
	if o == nil || IsNil(o.EntityReferenceList) {
		var ret []Reference
		return ret
	}
	return o.EntityReferenceList
}

// GetEntityReferenceListOk returns a tuple with the EntityReferenceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetEntityReferenceListOk() ([]Reference, bool) {
	if o == nil || IsNil(o.EntityReferenceList) {
		return nil, false
	}
	return o.EntityReferenceList, true
}

// HasEntityReferenceList returns a boolean if a field has been set.
func (o *Task) HasEntityReferenceList() bool {
	if o != nil && !IsNil(o.EntityReferenceList) {
		return true
	}

	return false
}

// SetEntityReferenceList gets a reference to the given []Reference and assigns it to the EntityReferenceList field.
func (o *Task) SetEntityReferenceList(v []Reference) {
	o.EntityReferenceList = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *Task) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *Task) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *Task) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *Task) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *Task) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *Task) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetUUID returns the UUID field value if set, zero value otherwise.
func (o *Task) GetUUID() string {
	if o == nil || IsNil(o.UUID) {
		var ret string
		return ret
	}
	return *o.UUID
}

// GetUUIDOk returns a tuple with the UUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetUUIDOk() (*string, bool) {
	if o == nil || IsNil(o.UUID) {
		return nil, false
	}
	return o.UUID, true
}

// HasUUID returns a boolean if a field has been set.
func (o *Task) HasUUID() bool {
	if o != nil && !IsNil(o.UUID) {
		return true
	}

	return false
}

// SetUUID gets a reference to the given string and assigns it to the UUID field.
func (o *Task) SetUUID(v string) {
	o.UUID = &v
}

// GetStartTimeUsecs returns the StartTimeUsecs field value if set, zero value otherwise.
func (o *Task) GetStartTimeUsecs() int64 {
	if o == nil || IsNil(o.StartTimeUsecs) {
		var ret int64
		return ret
	}
	return *o.StartTimeUsecs
}

// GetStartTimeUsecsOk returns a tuple with the StartTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetStartTimeUsecsOk() (*int64, bool) {
	if o == nil || IsNil(o.StartTimeUsecs) {
		return nil, false
	}
	return o.StartTimeUsecs, true
}

// HasStartTimeUsecs returns a boolean if a field has been set.
func (o *Task) HasStartTimeUsecs() bool {
	if o != nil && !IsNil(o.StartTimeUsecs) {
		return true
	}

	return false
}

// SetStartTimeUsecs gets a reference to the given int64 and assigns it to the StartTimeUsecs field.
func (o *Task) SetStartTimeUsecs(v int64) {
	o.StartTimeUsecs = &v
}

// GetClusterReference returns the ClusterReference field value if set, zero value otherwise.
func (o *Task) GetClusterReference() ClusterReference {
	if o == nil || IsNil(o.ClusterReference) {
		var ret ClusterReference
		return ret
	}
	return *o.ClusterReference
}

// GetClusterReferenceOk returns a tuple with the ClusterReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetClusterReferenceOk() (*ClusterReference, bool) {
	if o == nil || IsNil(o.ClusterReference) {
		return nil, false
	}
	return o.ClusterReference, true
}

// HasClusterReference returns a boolean if a field has been set.
func (o *Task) HasClusterReference() bool {
	if o != nil && !IsNil(o.ClusterReference) {
		return true
	}

	return false
}

// SetClusterReference gets a reference to the given ClusterReference and assigns it to the ClusterReference field.
func (o *Task) SetClusterReference(v ClusterReference) {
	o.ClusterReference = &v
}

// GetSubtaskReferenceList returns the SubtaskReferenceList field value if set, zero value otherwise.
func (o *Task) GetSubtaskReferenceList() []TaskReference {
	if o == nil || IsNil(o.SubtaskReferenceList) {
		var ret []TaskReference
		return ret
	}
	return o.SubtaskReferenceList
}

// GetSubtaskReferenceListOk returns a tuple with the SubtaskReferenceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetSubtaskReferenceListOk() ([]TaskReference, bool) {
	if o == nil || IsNil(o.SubtaskReferenceList) {
		return nil, false
	}
	return o.SubtaskReferenceList, true
}

// HasSubtaskReferenceList returns a boolean if a field has been set.
func (o *Task) HasSubtaskReferenceList() bool {
	if o != nil && !IsNil(o.SubtaskReferenceList) {
		return true
	}

	return false
}

// SetSubtaskReferenceList gets a reference to the given []TaskReference and assigns it to the SubtaskReferenceList field.
func (o *Task) SetSubtaskReferenceList(v []TaskReference) {
	o.SubtaskReferenceList = v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *Task) GetCompletionTime() time.Time {
	if o == nil || IsNil(o.CompletionTime) {
		var ret time.Time
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetCompletionTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CompletionTime) {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *Task) HasCompletionTime() bool {
	if o != nil && !IsNil(o.CompletionTime) {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given time.Time and assigns it to the CompletionTime field.
func (o *Task) SetCompletionTime(v time.Time) {
	o.CompletionTime = &v
}

// GetCreationTimeUsecs returns the CreationTimeUsecs field value if set, zero value otherwise.
func (o *Task) GetCreationTimeUsecs() int64 {
	if o == nil || IsNil(o.CreationTimeUsecs) {
		var ret int64
		return ret
	}
	return *o.CreationTimeUsecs
}

// GetCreationTimeUsecsOk returns a tuple with the CreationTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetCreationTimeUsecsOk() (*int64, bool) {
	if o == nil || IsNil(o.CreationTimeUsecs) {
		return nil, false
	}
	return o.CreationTimeUsecs, true
}

// HasCreationTimeUsecs returns a boolean if a field has been set.
func (o *Task) HasCreationTimeUsecs() bool {
	if o != nil && !IsNil(o.CreationTimeUsecs) {
		return true
	}

	return false
}

// SetCreationTimeUsecs gets a reference to the given int64 and assigns it to the CreationTimeUsecs field.
func (o *Task) SetCreationTimeUsecs(v int64) {
	o.CreationTimeUsecs = &v
}

// GetProgressMessage returns the ProgressMessage field value if set, zero value otherwise.
func (o *Task) GetProgressMessage() string {
	if o == nil || IsNil(o.ProgressMessage) {
		var ret string
		return ret
	}
	return *o.ProgressMessage
}

// GetProgressMessageOk returns a tuple with the ProgressMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetProgressMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ProgressMessage) {
		return nil, false
	}
	return o.ProgressMessage, true
}

// HasProgressMessage returns a boolean if a field has been set.
func (o *Task) HasProgressMessage() bool {
	if o != nil && !IsNil(o.ProgressMessage) {
		return true
	}

	return false
}

// SetProgressMessage gets a reference to the given string and assigns it to the ProgressMessage field.
func (o *Task) SetProgressMessage(v string) {
	o.ProgressMessage = &v
}

// GetOperationType returns the OperationType field value if set, zero value otherwise.
func (o *Task) GetOperationType() string {
	if o == nil || IsNil(o.OperationType) {
		var ret string
		return ret
	}
	return *o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetOperationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OperationType) {
		return nil, false
	}
	return o.OperationType, true
}

// HasOperationType returns a boolean if a field has been set.
func (o *Task) HasOperationType() bool {
	if o != nil && !IsNil(o.OperationType) {
		return true
	}

	return false
}

// SetOperationType gets a reference to the given string and assigns it to the OperationType field.
func (o *Task) SetOperationType(v string) {
	o.OperationType = &v
}

// GetCompletionTimeUsecs returns the CompletionTimeUsecs field value if set, zero value otherwise.
func (o *Task) GetCompletionTimeUsecs() int64 {
	if o == nil || IsNil(o.CompletionTimeUsecs) {
		var ret int64
		return ret
	}
	return *o.CompletionTimeUsecs
}

// GetCompletionTimeUsecsOk returns a tuple with the CompletionTimeUsecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetCompletionTimeUsecsOk() (*int64, bool) {
	if o == nil || IsNil(o.CompletionTimeUsecs) {
		return nil, false
	}
	return o.CompletionTimeUsecs, true
}

// HasCompletionTimeUsecs returns a boolean if a field has been set.
func (o *Task) HasCompletionTimeUsecs() bool {
	if o != nil && !IsNil(o.CompletionTimeUsecs) {
		return true
	}

	return false
}

// SetCompletionTimeUsecs gets a reference to the given int64 and assigns it to the CompletionTimeUsecs field.
func (o *Task) SetCompletionTimeUsecs(v int64) {
	o.CompletionTimeUsecs = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *Task) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *Task) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *Task) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetPercentageComplete returns the PercentageComplete field value if set, zero value otherwise.
func (o *Task) GetPercentageComplete() int32 {
	if o == nil || IsNil(o.PercentageComplete) {
		var ret int32
		return ret
	}
	return *o.PercentageComplete
}

// GetPercentageCompleteOk returns a tuple with the PercentageComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetPercentageCompleteOk() (*int32, bool) {
	if o == nil || IsNil(o.PercentageComplete) {
		return nil, false
	}
	return o.PercentageComplete, true
}

// HasPercentageComplete returns a boolean if a field has been set.
func (o *Task) HasPercentageComplete() bool {
	if o != nil && !IsNil(o.PercentageComplete) {
		return true
	}

	return false
}

// SetPercentageComplete gets a reference to the given int32 and assigns it to the PercentageComplete field.
func (o *Task) SetPercentageComplete(v int32) {
	o.PercentageComplete = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *Task) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *Task) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *Task) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetParentTaskReference returns the ParentTaskReference field value if set, zero value otherwise.
func (o *Task) GetParentTaskReference() TaskReference {
	if o == nil || IsNil(o.ParentTaskReference) {
		var ret TaskReference
		return ret
	}
	return *o.ParentTaskReference
}

// GetParentTaskReferenceOk returns a tuple with the ParentTaskReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetParentTaskReferenceOk() (*TaskReference, bool) {
	if o == nil || IsNil(o.ParentTaskReference) {
		return nil, false
	}
	return o.ParentTaskReference, true
}

// HasParentTaskReference returns a boolean if a field has been set.
func (o *Task) HasParentTaskReference() bool {
	if o != nil && !IsNil(o.ParentTaskReference) {
		return true
	}

	return false
}

// SetParentTaskReference gets a reference to the given TaskReference and assigns it to the ParentTaskReference field.
func (o *Task) SetParentTaskReference(v TaskReference) {
	o.ParentTaskReference = &v
}

func (o Task) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Task) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.LastUpdateTime) {
		toSerialize["last_update_time"] = o.LastUpdateTime
	}
	if !IsNil(o.ErrorDetail) {
		toSerialize["error_detail"] = o.ErrorDetail
	}
	if !IsNil(o.LogicalTimestamp) {
		toSerialize["logical_timestamp"] = o.LogicalTimestamp
	}
	if !IsNil(o.RequestedStatus) {
		toSerialize["requested_status"] = o.RequestedStatus
	}
	if !IsNil(o.EntityReferenceList) {
		toSerialize["entity_reference_list"] = o.EntityReferenceList
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.UUID) {
		toSerialize["uuid"] = o.UUID
	}
	if !IsNil(o.StartTimeUsecs) {
		toSerialize["start_time_usecs"] = o.StartTimeUsecs
	}
	if !IsNil(o.ClusterReference) {
		toSerialize["cluster_reference"] = o.ClusterReference
	}
	if !IsNil(o.SubtaskReferenceList) {
		toSerialize["subtask_reference_list"] = o.SubtaskReferenceList
	}
	if !IsNil(o.CompletionTime) {
		toSerialize["completion_time"] = o.CompletionTime
	}
	if !IsNil(o.CreationTimeUsecs) {
		toSerialize["creation_time_usecs"] = o.CreationTimeUsecs
	}
	if !IsNil(o.ProgressMessage) {
		toSerialize["progress_message"] = o.ProgressMessage
	}
	if !IsNil(o.OperationType) {
		toSerialize["operation_type"] = o.OperationType
	}
	if !IsNil(o.CompletionTimeUsecs) {
		toSerialize["completion_time_usecs"] = o.CompletionTimeUsecs
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !IsNil(o.PercentageComplete) {
		toSerialize["percentage_complete"] = o.PercentageComplete
	}
	if !IsNil(o.ApiVersion) {
		toSerialize["api_version"] = o.ApiVersion
	}
	if !IsNil(o.ParentTaskReference) {
		toSerialize["parent_task_reference"] = o.ParentTaskReference
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Task) UnmarshalJSON(data []byte) (err error) {
	varTask := _Task{}

	err = json.Unmarshal(data, &varTask)

	if err != nil {
		return err
	}

	*o = Task(varTask)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "last_update_time")
		delete(additionalProperties, "error_detail")
		delete(additionalProperties, "logical_timestamp")
		delete(additionalProperties, "requested_status")
		delete(additionalProperties, "entity_reference_list")
		delete(additionalProperties, "start_time")
		delete(additionalProperties, "creation_time")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "start_time_usecs")
		delete(additionalProperties, "cluster_reference")
		delete(additionalProperties, "subtask_reference_list")
		delete(additionalProperties, "completion_time")
		delete(additionalProperties, "creation_time_usecs")
		delete(additionalProperties, "progress_message")
		delete(additionalProperties, "operation_type")
		delete(additionalProperties, "completion_time_usecs")
		delete(additionalProperties, "error_code")
		delete(additionalProperties, "percentage_complete")
		delete(additionalProperties, "api_version")
		delete(additionalProperties, "parent_task_reference")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTask struct {
	value *Task
	isSet bool
}

func (v NullableTask) Get() *Task {
	return v.value
}

func (v *NullableTask) Set(val *Task) {
	v.value = val
	v.isSet = true
}

func (v NullableTask) IsSet() bool {
	return v.isSet
}

func (v *NullableTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTask(val *Task) *NullableTask {
	return &NullableTask{value: val, isSet: true}
}

func (v NullableTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


