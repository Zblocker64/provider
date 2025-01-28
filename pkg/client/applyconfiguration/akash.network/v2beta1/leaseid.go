/*
Copyright The Akash Network Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v2beta1

// LeaseIDApplyConfiguration represents a declarative configuration of the LeaseID type for use
// with apply.
type LeaseIDApplyConfiguration struct {
	Owner    *string `json:"owner,omitempty"`
	DSeq     *string `json:"dseq,omitempty"`
	GSeq     *uint32 `json:"gseq,omitempty"`
	OSeq     *uint32 `json:"oseq,omitempty"`
	Provider *string `json:"provider,omitempty"`
}

// LeaseIDApplyConfiguration constructs a declarative configuration of the LeaseID type for use with
// apply.
func LeaseID() *LeaseIDApplyConfiguration {
	return &LeaseIDApplyConfiguration{}
}

// WithOwner sets the Owner field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Owner field is set to the value of the last call.
func (b *LeaseIDApplyConfiguration) WithOwner(value string) *LeaseIDApplyConfiguration {
	b.Owner = &value
	return b
}

// WithDSeq sets the DSeq field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DSeq field is set to the value of the last call.
func (b *LeaseIDApplyConfiguration) WithDSeq(value string) *LeaseIDApplyConfiguration {
	b.DSeq = &value
	return b
}

// WithGSeq sets the GSeq field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GSeq field is set to the value of the last call.
func (b *LeaseIDApplyConfiguration) WithGSeq(value uint32) *LeaseIDApplyConfiguration {
	b.GSeq = &value
	return b
}

// WithOSeq sets the OSeq field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OSeq field is set to the value of the last call.
func (b *LeaseIDApplyConfiguration) WithOSeq(value uint32) *LeaseIDApplyConfiguration {
	b.OSeq = &value
	return b
}

// WithProvider sets the Provider field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Provider field is set to the value of the last call.
func (b *LeaseIDApplyConfiguration) WithProvider(value string) *LeaseIDApplyConfiguration {
	b.Provider = &value
	return b
}
