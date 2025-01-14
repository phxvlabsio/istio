// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import "istio.io/pkg/monitoring"

func init() {
	monitoring.MustRegister(
		providerLookupClusterFailures,
	)
}

var providerLookupClusterFailures = monitoring.NewSum(
	"provider_lookup_cluster_failures",
	"Number of times a cluster lookup failed",
	monitoring.WithLabels(typeTag),
)

func IncLookupClusterFailures(provider string) {
	providerLookupClusterFailures.With(typeTag.Value(provider)).Increment()
}
