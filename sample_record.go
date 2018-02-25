package sample

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/aedatastore"
	"google.golang.org/appengine/log"
)

type SampleRecord struct {
	KeyName   string `datastore:"-"`
	Timestamp int64
}

type SampleRecordStore struct {
	DatastoreClient datastore.Client
}

func NewSampleRecordStore(ctx context.Context) (*SampleRecordStore, error) {
	ds, err := aedatastore.FromContext(ctx)
	if err != nil {
		log.Errorf(ctx, "failed Datastore New Client: %+v", err)
		return nil, err
	}
	return &SampleRecordStore{ds}, nil
}

func (store *SampleRecordStore) NewKey(uuid string, ctx context.Context, ds datastore.Client) datastore.Key {
	return ds.NameKey("SampleRecord", uuid, nil)
}

func (store *SampleRecordStore) Create(ctx context.Context, e *SampleRecord) (*SampleRecord, error) {
	ds := store.DatastoreClient

	uuid := uuid.New().String()
	key := store.NewKey(uuid, ctx, ds)

	_, err := ds.Put(ctx, key, e)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed put record to Datastore. key=%v", key))
	}
	e.KeyName = uuid
	return e, nil
}

func (store *SampleRecordStore) Get(ctx context.Context, key datastore.Key) (*SampleRecord, error) {
	ds := store.DatastoreClient

	var record SampleRecord
	err := ds.Get(ctx, key, &record)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed get record from Datastore. key=%s", key.Name()))
	}
	record.KeyName = key.Name()

	return &record, nil
}

func (store *SampleRecordStore) Update(ctx context.Context, e *SampleRecord) (*SampleRecord, error) {
	ds := store.DatastoreClient
	key := store.NewKey(e.KeyName, ctx, ds)
	_, err := ds.Put(ctx, key, e)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed put record to Datastore."))
	}
	return e, nil
}

func (store *SampleRecordStore) Delete(ctx context.Context, key datastore.Key) error {
	ds := store.DatastoreClient
	err := ds.Delete(ctx, key)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed delete record from Datastore. key=%s", key.Name))
	}
	return nil
}
