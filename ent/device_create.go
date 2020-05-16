// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/minskylab/collecta/ent/device"
)

// DeviceCreate is the builder for creating a Device entity.
type DeviceCreate struct {
	config
	device *string
}

// SetDevice sets the device field.
func (dc *DeviceCreate) SetDevice(s string) *DeviceCreate {
	dc.device = &s
	return dc
}

// Save creates the Device in the database.
func (dc *DeviceCreate) Save(ctx context.Context) (*Device, error) {
	if dc.device == nil {
		return nil, errors.New("ent: missing required field \"device\"")
	}
	return dc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeviceCreate) SaveX(ctx context.Context) *Device {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DeviceCreate) sqlSave(ctx context.Context) (*Device, error) {
	var (
		d     = &Device{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: device.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: device.FieldID,
			},
		}
	)
	if value := dc.device; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: device.FieldDevice,
		})
		d.Device = *value
	}
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}
