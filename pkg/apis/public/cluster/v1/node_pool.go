// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import "fmt"

// Checks whether a node pool with the same name already exists.
func (c *Cluster) HasNodePool(npName string) bool {
	for _, existingNP := range c.Spec.NodePools {
		if existingNP.Name == npName {
			return true
		}
	}
	return false
}

// Returns the node pool with the specified name.
func (c *Cluster) NodePool(npName string) (NodePool, error) {
	for _, existingNP := range c.Spec.NodePools {
		if existingNP.Name == npName {
			return *existingNP, nil
		}
	}
	return NodePool{}, fmt.Errorf("NodePool %s doesn't exist", npName)
}

// Adds or updates a node pool in the GDCH cluster.
func (c *Cluster) ApplyNodePool(np NodePool) {
	newNodePools := []*NodePool{}
	for _, existingNP := range c.Spec.NodePools {
		if existingNP.Name != np.Name {
			newNodePools = append(newNodePools, existingNP)
		}
	}
	newNodePools = append(newNodePools, &np)
	c.Spec.NodePools = newNodePools
}

// Deletes a node pool in the GDCH cluster.
func (c *Cluster) DeleteNodePool(npName string) {
	newNodePools := []*NodePool{}
	for _, existingNP := range c.Spec.NodePools {
		if existingNP.Name != npName {
			newNodePools = append(newNodePools, existingNP)
		}
	}
	c.Spec.NodePools = newNodePools
}
