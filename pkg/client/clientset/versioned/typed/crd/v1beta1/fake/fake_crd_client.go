// Copyright 2023 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "antrea.io/antrea/pkg/client/clientset/versioned/typed/crd/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCrdV1beta1 struct {
	*testing.Fake
}

func (c *FakeCrdV1beta1) AntreaAgentInfos() v1beta1.AntreaAgentInfoInterface {
	return &FakeAntreaAgentInfos{c}
}

func (c *FakeCrdV1beta1) AntreaControllerInfos() v1beta1.AntreaControllerInfoInterface {
	return &FakeAntreaControllerInfos{c}
}

func (c *FakeCrdV1beta1) ClusterGroups() v1beta1.ClusterGroupInterface {
	return &FakeClusterGroups{c}
}

func (c *FakeCrdV1beta1) ExternalIPPools() v1beta1.ExternalIPPoolInterface {
	return &FakeExternalIPPools{c}
}

func (c *FakeCrdV1beta1) Groups(namespace string) v1beta1.GroupInterface {
	return &FakeGroups{c, namespace}
}

func (c *FakeCrdV1beta1) Tiers() v1beta1.TierInterface {
	return &FakeTiers{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCrdV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
