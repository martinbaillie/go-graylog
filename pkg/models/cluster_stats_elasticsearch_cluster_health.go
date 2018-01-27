// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClusterStatsElasticsearchClusterHealth cluster stats elasticsearch cluster health
// swagger:model clusterStatsElasticsearchClusterHealth
type ClusterStatsElasticsearchClusterHealth struct {

	// active primary shards
	ActivePrimaryShards int64 `json:"active_primary_shards,omitempty"`

	// active shards
	ActiveShards int64 `json:"active_shards,omitempty"`

	// initializing shards
	InitializingShards int64 `json:"initializing_shards,omitempty"`

	// number of data nodes
	NumberOfDataNodes int64 `json:"number_of_data_nodes,omitempty"`

	// number of nodes
	NumberOfNodes int64 `json:"number_of_nodes,omitempty"`

	// pending tasks
	PendingTasks int64 `json:"pending_tasks,omitempty"`

	// pending tasks time in queue
	PendingTasksTimeInQueue []int64 `json:"pending_tasks_time_in_queue"`

	// relocating shards
	RelocatingShards int64 `json:"relocating_shards,omitempty"`

	// timed out
	TimedOut bool `json:"timed_out,omitempty"`

	// unassigned shards
	UnassignedShards int64 `json:"unassigned_shards,omitempty"`
}

// Validate validates this cluster stats elasticsearch cluster health
func (m *ClusterStatsElasticsearchClusterHealth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePendingTasksTimeInQueue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterStatsElasticsearchClusterHealth) validatePendingTasksTimeInQueue(formats strfmt.Registry) error {

	if swag.IsZero(m.PendingTasksTimeInQueue) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterStatsElasticsearchClusterHealth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterStatsElasticsearchClusterHealth) UnmarshalBinary(b []byte) error {
	var res ClusterStatsElasticsearchClusterHealth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}