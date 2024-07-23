/*
This is an experimental package that provides a SDK for interacting with the Nutanix Prism API over both
v3 and v4 versions. The SDK is designed to be used by services that need to interact with the Prism Central APIs.
The aim is to abstract the v3-v4 away from the caller and provide a consistent interface to the caller regardless
of the version of the Prism Central API present in the environment. Note that the SDK is not a 1:1 mapping of the
Prism Central API, but rather a subset of the API that is commonly used by services. We will continue to add more
APIs to the SDK as needed.

Note: This package is experimental and is subject to change without notice.
*/
package adapter
