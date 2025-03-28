//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Some standard accessors for the shard struct.
// It is important to NEVER access the shard struct directly, because we lazy load shards, so the information might not be there.
package db

import (
	"github.com/weaviate/weaviate/adapters/repos/db/indexcounter"
	"github.com/weaviate/weaviate/adapters/repos/db/inverted"
	"github.com/weaviate/weaviate/adapters/repos/db/lsmkv"
	"github.com/weaviate/weaviate/entities/schema"
)

func (s *Shard) Queue() *VectorIndexQueue {
	return s.queue
}

func (s *Shard) Queues() map[string]*VectorIndexQueue {
	return s.queues
}

// GetVectorIndexQueue retrieves a vector index queue associated with the targetVector.
// Empty targetVector is treated as a request to access a queue for the legacy vector index.
func (s *Shard) GetVectorIndexQueue(targetVector string) (*VectorIndexQueue, bool) {
	if targetVector == "" {
		return s.queue, s.queue != nil
	}

	queue, ok := s.queues[targetVector]
	return queue, ok
}

// GetVectorIndex retrieves a vector index queue associated with the targetVector.
// Empty targetVector is treated as a request to access a queue for the legacy vector index.
func (s *Shard) GetVectorIndex(targetVector string) (VectorIndex, bool) {
	if targetVector == "" {
		return s.vectorIndex, s.vectorIndex != nil
	}

	index, ok := s.vectorIndexes[targetVector]
	return index, ok
}

func (s *Shard) QueueForName(targetVector string) *VectorIndexQueue {
	return s.queues[targetVector]
}

func (s *Shard) VectorIndex() VectorIndex {
	return s.vectorIndex
}

func (s *Shard) VectorIndexes() map[string]VectorIndex {
	return s.vectorIndexes
}

func (s *Shard) VectorIndexForName(targetVector string) VectorIndex {
	return s.vectorIndexes[targetVector]
}

func (s *Shard) Versioner() *shardVersioner {
	return s.versioner
}

func (s *Shard) Index() *Index {
	return s.index
}

// Shard name(identifier?)
func (s *Shard) Name() string {
	return s.name
}

// The physical data store
func (s *Shard) Store() *lsmkv.Store {
	return s.store
}

func (s *Shard) Counter() *indexcounter.Counter {
	return s.counter
}

// Tracks the lengths of all properties.  Must be updated on inserts/deletes.
func (s *Shard) GetPropertyLengthTracker() *inverted.JsonShardMetaData {
	return s.propLenTracker
}

// Tracks the lengths of all properties.  Must be updated on inserts/deletes.
func (s *Shard) SetPropertyLengthTracker(tracker *inverted.JsonShardMetaData) {
	s.propLenTracker = tracker
}

// Grafana metrics
func (s *Shard) Metrics() *Metrics {
	return s.metrics
}

func (s *Shard) setFallbackToSearchable(fallback bool) {
	s.fallbackToSearchable = fallback
}

func (s *Shard) addJobToQueue(job job) {
	s.centralJobQueue <- job
}

func (s *Shard) hasGeoIndex() bool {
	s.propertyIndicesLock.RLock()
	defer s.propertyIndicesLock.RUnlock()

	for _, idx := range s.propertyIndices {
		if idx.Type == schema.DataTypeGeoCoordinates {
			return true
		}
	}
	return false
}
