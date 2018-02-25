package sample

import (
	"fmt"
	"net/http"
	"time"

	"google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/create", sampleCreateHandler)
	http.HandleFunc("/read", sampleReadHandler)
	http.HandleFunc("/update", sampleUpdateHandler)
	http.HandleFunc("/delete", sampleDeleteHandler)
}

func sampleCreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	es, err := NewSampleRecordStore(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to init store: %s", err.Error())
	}

	record := &SampleRecord{
		Timestamp: time.Now().Unix(),
	}
	record, err = es.Create(ctx, record)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to put record: %s", err.Error())
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success to create record. key: %s", record.KeyName)
}

func sampleReadHandler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("uuid")

	ctx := appengine.NewContext(r)
	es, err := NewSampleRecordStore(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to init store: %s", err.Error())
	}

	record, err := es.Get(ctx, es.DatastoreClient.NameKey("SampleRecord", param, nil))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to put record: %s", err.Error())
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success to read record: %d", record.Timestamp)
}

func sampleUpdateHandler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("uuid")

	ctx := appengine.NewContext(r)
	es, err := NewSampleRecordStore(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to init store: %s", err.Error())
	}

	key := es.DatastoreClient.NameKey("SampleRecord", param, nil)
	record, err := es.Get(ctx, key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to put record: %s", err.Error())
	}

	record.Timestamp = time.Now().Unix()
	_, err = es.Update(ctx, record)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to put record: %s", err.Error())
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success to update record")
}

func sampleDeleteHandler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("uuid")

	ctx := appengine.NewContext(r)
	es, err := NewSampleRecordStore(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to init store: %s", err.Error())
	}

	err = es.Delete(ctx, es.DatastoreClient.NameKey("SampleRecord", param, nil))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to delete record: %s", err.Error())
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success to delete record: %s", param)
}
