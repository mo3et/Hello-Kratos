// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-cqrs/app/logger/service/internal/data/ent/predicate"
	"kratos-cqrs/app/logger/service/internal/data/ent/sensor"
	"kratos-cqrs/app/logger/service/internal/data/ent/sensordata"
	"sync"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeSensor     = "Sensor"
	TypeSensorData = "SensorData"
)

// SensorMutation represents an operation that mutates the Sensor nodes in the graph.
type SensorMutation struct {
	config
	op            Op
	typ           string
	id            *int
	_type         *string
	location      *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Sensor, error)
	predicates    []predicate.Sensor
}

var _ ent.Mutation = (*SensorMutation)(nil)

// sensorOption allows management of the mutation configuration using functional options.
type sensorOption func(*SensorMutation)

// newSensorMutation creates new mutation for the Sensor entity.
func newSensorMutation(c config, op Op, opts ...sensorOption) *SensorMutation {
	m := &SensorMutation{
		config:        c,
		op:            op,
		typ:           TypeSensor,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSensorID sets the ID field of the mutation.
func withSensorID(id int) sensorOption {
	return func(m *SensorMutation) {
		var (
			err   error
			once  sync.Once
			value *Sensor
		)
		m.oldValue = func(ctx context.Context) (*Sensor, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Sensor.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSensor sets the old Sensor of the mutation.
func withSensor(node *Sensor) sensorOption {
	return func(m *SensorMutation) {
		m.oldValue = func(context.Context) (*Sensor, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SensorMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SensorMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Sensor entities.
func (m *SensorMutation) SetID(id int) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SensorMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *SensorMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Sensor.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetType sets the "type" field.
func (m *SensorMutation) SetType(s string) {
	m._type = &s
}

// GetType returns the value of the "type" field in the mutation.
func (m *SensorMutation) GetType() (r string, exists bool) {
	v := m._type
	if v == nil {
		return
	}
	return *v, true
}

// OldType returns the old "type" field's value of the Sensor entity.
// If the Sensor object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SensorMutation) OldType(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldType: %w", err)
	}
	return oldValue.Type, nil
}

// ResetType resets all changes to the "type" field.
func (m *SensorMutation) ResetType() {
	m._type = nil
}

// SetLocation sets the "location" field.
func (m *SensorMutation) SetLocation(s string) {
	m.location = &s
}

// Location returns the value of the "location" field in the mutation.
func (m *SensorMutation) Location() (r string, exists bool) {
	v := m.location
	if v == nil {
		return
	}
	return *v, true
}

// OldLocation returns the old "location" field's value of the Sensor entity.
// If the Sensor object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SensorMutation) OldLocation(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLocation is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLocation requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLocation: %w", err)
	}
	return oldValue.Location, nil
}

// ResetLocation resets all changes to the "location" field.
func (m *SensorMutation) ResetLocation() {
	m.location = nil
}

// Where appends a list predicates to the SensorMutation builder.
func (m *SensorMutation) Where(ps ...predicate.Sensor) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *SensorMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Sensor).
func (m *SensorMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SensorMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m._type != nil {
		fields = append(fields, sensor.FieldType)
	}
	if m.location != nil {
		fields = append(fields, sensor.FieldLocation)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SensorMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case sensor.FieldType:
		return m.GetType()
	case sensor.FieldLocation:
		return m.Location()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SensorMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case sensor.FieldType:
		return m.OldType(ctx)
	case sensor.FieldLocation:
		return m.OldLocation(ctx)
	}
	return nil, fmt.Errorf("unknown Sensor field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SensorMutation) SetField(name string, value ent.Value) error {
	switch name {
	case sensor.FieldType:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetType(v)
		return nil
	case sensor.FieldLocation:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLocation(v)
		return nil
	}
	return fmt.Errorf("unknown Sensor field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SensorMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SensorMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SensorMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Sensor numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SensorMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SensorMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SensorMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Sensor nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SensorMutation) ResetField(name string) error {
	switch name {
	case sensor.FieldType:
		m.ResetType()
		return nil
	case sensor.FieldLocation:
		m.ResetLocation()
		return nil
	}
	return fmt.Errorf("unknown Sensor field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SensorMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SensorMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SensorMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SensorMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SensorMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SensorMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SensorMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Sensor unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SensorMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Sensor edge %s", name)
}

// SensorDataMutation represents an operation that mutates the SensorData nodes in the graph.
type SensorDataMutation struct {
	config
	op             Op
	typ            string
	id             *int64
	time           *int64
	addtime        *int64
	sensor_id      *int
	addsensor_id   *int
	temperature    *float64
	addtemperature *float64
	cpu            *float64
	addcpu         *float64
	clearedFields  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*SensorData, error)
	predicates     []predicate.SensorData
}

var _ ent.Mutation = (*SensorDataMutation)(nil)

// sensordataOption allows management of the mutation configuration using functional options.
type sensordataOption func(*SensorDataMutation)

// newSensorDataMutation creates new mutation for the SensorData entity.
func newSensorDataMutation(c config, op Op, opts ...sensordataOption) *SensorDataMutation {
	m := &SensorDataMutation{
		config:        c,
		op:            op,
		typ:           TypeSensorData,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSensorDataID sets the ID field of the mutation.
func withSensorDataID(id int64) sensordataOption {
	return func(m *SensorDataMutation) {
		var (
			err   error
			once  sync.Once
			value *SensorData
		)
		m.oldValue = func(ctx context.Context) (*SensorData, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().SensorData.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSensorData sets the old SensorData of the mutation.
func withSensorData(node *SensorData) sensordataOption {
	return func(m *SensorDataMutation) {
		m.oldValue = func(context.Context) (*SensorData, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SensorDataMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SensorDataMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of SensorData entities.
func (m *SensorDataMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SensorDataMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *SensorDataMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().SensorData.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTime sets the "time" field.
func (m *SensorDataMutation) SetTime(i int64) {
	m.time = &i
	m.addtime = nil
}

// Time returns the value of the "time" field in the mutation.
func (m *SensorDataMutation) Time() (r int64, exists bool) {
	v := m.time
	if v == nil {
		return
	}
	return *v, true
}

// OldTime returns the old "time" field's value of the SensorData entity.
// If the SensorData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SensorDataMutation) OldTime(ctx context.Context) (v *int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTime: %w", err)
	}
	return oldValue.Time, nil
}

// AddTime adds i to the "time" field.
func (m *SensorDataMutation) AddTime(i int64) {
	if m.addtime != nil {
		*m.addtime += i
	} else {
		m.addtime = &i
	}
}

// AddedTime returns the value that was added to the "time" field in this mutation.
func (m *SensorDataMutation) AddedTime() (r int64, exists bool) {
	v := m.addtime
	if v == nil {
		return
	}
	return *v, true
}

// ClearTime clears the value of the "time" field.
func (m *SensorDataMutation) ClearTime() {
	m.time = nil
	m.addtime = nil
	m.clearedFields[sensordata.FieldTime] = struct{}{}
}

// TimeCleared returns if the "time" field was cleared in this mutation.
func (m *SensorDataMutation) TimeCleared() bool {
	_, ok := m.clearedFields[sensordata.FieldTime]
	return ok
}

// ResetTime resets all changes to the "time" field.
func (m *SensorDataMutation) ResetTime() {
	m.time = nil
	m.addtime = nil
	delete(m.clearedFields, sensordata.FieldTime)
}

// SetSensorID sets the "sensor_id" field.
func (m *SensorDataMutation) SetSensorID(i int) {
	m.sensor_id = &i
	m.addsensor_id = nil
}

// SensorID returns the value of the "sensor_id" field in the mutation.
func (m *SensorDataMutation) SensorID() (r int, exists bool) {
	v := m.sensor_id
	if v == nil {
		return
	}
	return *v, true
}

// OldSensorID returns the old "sensor_id" field's value of the SensorData entity.
// If the SensorData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SensorDataMutation) OldSensorID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSensorID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSensorID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSensorID: %w", err)
	}
	return oldValue.SensorID, nil
}

// AddSensorID adds i to the "sensor_id" field.
func (m *SensorDataMutation) AddSensorID(i int) {
	if m.addsensor_id != nil {
		*m.addsensor_id += i
	} else {
		m.addsensor_id = &i
	}
}

// AddedSensorID returns the value that was added to the "sensor_id" field in this mutation.
func (m *SensorDataMutation) AddedSensorID() (r int, exists bool) {
	v := m.addsensor_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetSensorID resets all changes to the "sensor_id" field.
func (m *SensorDataMutation) ResetSensorID() {
	m.sensor_id = nil
	m.addsensor_id = nil
}

// SetTemperature sets the "temperature" field.
func (m *SensorDataMutation) SetTemperature(f float64) {
	m.temperature = &f
	m.addtemperature = nil
}

// Temperature returns the value of the "temperature" field in the mutation.
func (m *SensorDataMutation) Temperature() (r float64, exists bool) {
	v := m.temperature
	if v == nil {
		return
	}
	return *v, true
}

// OldTemperature returns the old "temperature" field's value of the SensorData entity.
// If the SensorData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SensorDataMutation) OldTemperature(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTemperature is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTemperature requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTemperature: %w", err)
	}
	return oldValue.Temperature, nil
}

// AddTemperature adds f to the "temperature" field.
func (m *SensorDataMutation) AddTemperature(f float64) {
	if m.addtemperature != nil {
		*m.addtemperature += f
	} else {
		m.addtemperature = &f
	}
}

// AddedTemperature returns the value that was added to the "temperature" field in this mutation.
func (m *SensorDataMutation) AddedTemperature() (r float64, exists bool) {
	v := m.addtemperature
	if v == nil {
		return
	}
	return *v, true
}

// ResetTemperature resets all changes to the "temperature" field.
func (m *SensorDataMutation) ResetTemperature() {
	m.temperature = nil
	m.addtemperature = nil
}

// SetCPU sets the "cpu" field.
func (m *SensorDataMutation) SetCPU(f float64) {
	m.cpu = &f
	m.addcpu = nil
}

// CPU returns the value of the "cpu" field in the mutation.
func (m *SensorDataMutation) CPU() (r float64, exists bool) {
	v := m.cpu
	if v == nil {
		return
	}
	return *v, true
}

// OldCPU returns the old "cpu" field's value of the SensorData entity.
// If the SensorData object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *SensorDataMutation) OldCPU(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCPU is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCPU requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCPU: %w", err)
	}
	return oldValue.CPU, nil
}

// AddCPU adds f to the "cpu" field.
func (m *SensorDataMutation) AddCPU(f float64) {
	if m.addcpu != nil {
		*m.addcpu += f
	} else {
		m.addcpu = &f
	}
}

// AddedCPU returns the value that was added to the "cpu" field in this mutation.
func (m *SensorDataMutation) AddedCPU() (r float64, exists bool) {
	v := m.addcpu
	if v == nil {
		return
	}
	return *v, true
}

// ResetCPU resets all changes to the "cpu" field.
func (m *SensorDataMutation) ResetCPU() {
	m.cpu = nil
	m.addcpu = nil
}

// Where appends a list predicates to the SensorDataMutation builder.
func (m *SensorDataMutation) Where(ps ...predicate.SensorData) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *SensorDataMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (SensorData).
func (m *SensorDataMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SensorDataMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.time != nil {
		fields = append(fields, sensordata.FieldTime)
	}
	if m.sensor_id != nil {
		fields = append(fields, sensordata.FieldSensorID)
	}
	if m.temperature != nil {
		fields = append(fields, sensordata.FieldTemperature)
	}
	if m.cpu != nil {
		fields = append(fields, sensordata.FieldCPU)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SensorDataMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case sensordata.FieldTime:
		return m.Time()
	case sensordata.FieldSensorID:
		return m.SensorID()
	case sensordata.FieldTemperature:
		return m.Temperature()
	case sensordata.FieldCPU:
		return m.CPU()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SensorDataMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case sensordata.FieldTime:
		return m.OldTime(ctx)
	case sensordata.FieldSensorID:
		return m.OldSensorID(ctx)
	case sensordata.FieldTemperature:
		return m.OldTemperature(ctx)
	case sensordata.FieldCPU:
		return m.OldCPU(ctx)
	}
	return nil, fmt.Errorf("unknown SensorData field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SensorDataMutation) SetField(name string, value ent.Value) error {
	switch name {
	case sensordata.FieldTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTime(v)
		return nil
	case sensordata.FieldSensorID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSensorID(v)
		return nil
	case sensordata.FieldTemperature:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTemperature(v)
		return nil
	case sensordata.FieldCPU:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCPU(v)
		return nil
	}
	return fmt.Errorf("unknown SensorData field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SensorDataMutation) AddedFields() []string {
	var fields []string
	if m.addtime != nil {
		fields = append(fields, sensordata.FieldTime)
	}
	if m.addsensor_id != nil {
		fields = append(fields, sensordata.FieldSensorID)
	}
	if m.addtemperature != nil {
		fields = append(fields, sensordata.FieldTemperature)
	}
	if m.addcpu != nil {
		fields = append(fields, sensordata.FieldCPU)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SensorDataMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case sensordata.FieldTime:
		return m.AddedTime()
	case sensordata.FieldSensorID:
		return m.AddedSensorID()
	case sensordata.FieldTemperature:
		return m.AddedTemperature()
	case sensordata.FieldCPU:
		return m.AddedCPU()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SensorDataMutation) AddField(name string, value ent.Value) error {
	switch name {
	case sensordata.FieldTime:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddTime(v)
		return nil
	case sensordata.FieldSensorID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddSensorID(v)
		return nil
	case sensordata.FieldTemperature:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddTemperature(v)
		return nil
	case sensordata.FieldCPU:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCPU(v)
		return nil
	}
	return fmt.Errorf("unknown SensorData numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SensorDataMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(sensordata.FieldTime) {
		fields = append(fields, sensordata.FieldTime)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SensorDataMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SensorDataMutation) ClearField(name string) error {
	switch name {
	case sensordata.FieldTime:
		m.ClearTime()
		return nil
	}
	return fmt.Errorf("unknown SensorData nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SensorDataMutation) ResetField(name string) error {
	switch name {
	case sensordata.FieldTime:
		m.ResetTime()
		return nil
	case sensordata.FieldSensorID:
		m.ResetSensorID()
		return nil
	case sensordata.FieldTemperature:
		m.ResetTemperature()
		return nil
	case sensordata.FieldCPU:
		m.ResetCPU()
		return nil
	}
	return fmt.Errorf("unknown SensorData field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SensorDataMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SensorDataMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SensorDataMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SensorDataMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SensorDataMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SensorDataMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SensorDataMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown SensorData unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SensorDataMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown SensorData edge %s", name)
}
