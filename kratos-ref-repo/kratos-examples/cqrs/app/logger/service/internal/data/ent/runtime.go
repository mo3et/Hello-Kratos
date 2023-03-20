// Code generated by ent, DO NOT EDIT.

package ent

import (
	"kratos-cqrs/app/logger/service/internal/data/ent/schema"
	"kratos-cqrs/app/logger/service/internal/data/ent/sensor"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	sensorFields := schema.Sensor{}.Fields()
	_ = sensorFields
	// sensorDescType is the schema descriptor for type field.
	sensorDescType := sensorFields[1].Descriptor()
	// sensor.DefaultType holds the default value on creation for the type field.
	sensor.DefaultType = sensorDescType.Default.(string)
	// sensor.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	sensor.TypeValidator = sensorDescType.Validators[0].(func(string) error)
	// sensorDescLocation is the schema descriptor for location field.
	sensorDescLocation := sensorFields[2].Descriptor()
	// sensor.DefaultLocation holds the default value on creation for the location field.
	sensor.DefaultLocation = sensorDescLocation.Default.(string)
	// sensor.LocationValidator is a validator for the "location" field. It is called by the builders before save.
	sensor.LocationValidator = sensorDescLocation.Validators[0].(func(string) error)
}