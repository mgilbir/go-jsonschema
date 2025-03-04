// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"
import "net/netip"

type IpMyObject struct {
	// MyIp corresponds to the JSON schema field "myIp".
	MyIp netip.Addr `json:"myIp" yaml:"myIp" mapstructure:"myIp"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IpMyObject) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["myIp"]; !ok || v == nil {
		return fmt.Errorf("field myIp in IpMyObject: required")
	}
	type Plain IpMyObject
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = IpMyObject(plain)
	return nil
}

type Ip struct {
	// MyObject corresponds to the JSON schema field "myObject".
	MyObject *IpMyObject `json:"myObject,omitempty" yaml:"myObject,omitempty" mapstructure:"myObject,omitempty"`
}
