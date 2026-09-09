package helpers

import (
	"reflect"
	"unsafe"

	v4 "github.com/nutanix-cloud-native/prism-go-client/converged/v4"
	v4sdk "github.com/nutanix-cloud-native/prism-go-client/v4"
)

func StripSDKAuthorization(c *v4.Client) {
	clientField := unexportedField(c, "client")
	if !clientField.IsValid() || clientField.IsNil() {
		return
	}
	sdkClient, ok := clientField.Interface().(*v4sdk.Client)
	if !ok || sdkClient == nil || sdkClient.ClustersApiInstance == nil || sdkClient.ClustersApiInstance.ApiClient == nil {
		return
	}
	clearSDKAuthState(sdkClient.ClustersApiInstance.ApiClient)
}

func unexportedField(structPtr any, name string) reflect.Value {
	v := reflect.ValueOf(structPtr)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return reflect.Value{}
	}
	elem := v.Elem()
	if elem.Kind() != reflect.Struct {
		return reflect.Value{}
	}
	f := elem.FieldByName(name)
	if !f.IsValid() {
		return reflect.Value{}
	}
	return reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem()
}

func clearSDKAuthState(apiClient any) {
	v := reflect.ValueOf(apiClient)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return
	}
	elem := v.Elem()
	if elem.Kind() != reflect.Struct {
		return
	}
	if headersField := elem.FieldByName("defaultHeaders"); headersField.IsValid() {
		headers := reflect.NewAt(headersField.Type(), unsafe.Pointer(headersField.UnsafeAddr())).Elem()
		if headers.Kind() == reflect.Map && !headers.IsNil() {
			for _, key := range []string{"Authorization", "X-ntnx-api-key"} {
				headers.SetMapIndex(reflect.ValueOf(key), reflect.Value{})
			}
		}
	}
	for _, name := range []string{"cookie", "previousAuth", "Username", "Password"} {
		if f := elem.FieldByName(name); f.IsValid() && f.Kind() == reflect.String {
			reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem().SetString("")
		}
	}
	if basic := elem.FieldByName("basicAuth"); basic.IsValid() && basic.Kind() == reflect.Ptr && !basic.IsNil() {
		ba := basic.Elem()
		for _, name := range []string{"UserName", "Password"} {
			if f := ba.FieldByName(name); f.IsValid() && f.Kind() == reflect.String {
				reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).Elem().SetString("")
			}
		}
	}
}
