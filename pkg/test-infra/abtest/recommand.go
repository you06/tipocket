// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package abtest

import "github.com/pingcap/tipocket/pkg/test-infra/tidb"

// Recommendation defines 2 clusters for abtest
type Recommendation struct {
	Cluster1 *tidb.TiDBClusterRecommendation
	Cluster2 *tidb.TiDBClusterRecommendation
	NS       string
	Name     string
}

// RecommendedCluster gives recommand cluster
func RecommendedCluster(ns, name string) *Recommendation {
	return &Recommendation{
		Cluster1: tidb.RecommendedTiDBCluster(ns, name),
		Cluster2: tidb.RecommendedTiDBCluster(ns, name),
		NS:       ns,
		Name:     name,
	}
}