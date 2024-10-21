package goclonestruct

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"sync"
)

// Declare a sync.Pool for bytes.Buffer
var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// CloneUsingGob clones the given source struct to the given destination struct using encoding/gob.
// It returns an error if either the source or destination is nil, or if there is an error during encoding or decoding.
func CloneUsingGob[S any, D any](source *S, destination *D) error {
	if source == nil {
		return fmt.Errorf("source is nil")
	}
	if destination == nil {
		return fmt.Errorf("destination is nil")
	}

	// Get a buffer from the pool
	buf := bufferPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()         // Reset the buffer for reuse
		bufferPool.Put(buf) // Return the buffer to the pool after use
	}()

	// Use a single encoder and decoder during the process.
	encoder := gob.NewEncoder(buf)
	if err := encoder.Encode(source); err != nil {
		return fmt.Errorf("error encoding source: %w", err)
	}

	decoder := gob.NewDecoder(buf)
	if err := decoder.Decode(destination); err != nil {
		return fmt.Errorf("error decoding destination: %w", err)
	}

	return nil
}

// CloneUsingJson clones the given source struct to the given destination struct using encoding/json.
// It returns an error if either the source or destination is nil, or if there is an error during encoding or decoding.
func CloneUsingJson[S any, D any](source *S, destination *D) error {
	if source == nil {
		return fmt.Errorf("source is nil")
	}
	if destination == nil {
		return fmt.Errorf("destination is nil")
	}

	// Get a buffer from the pool
	buf := bufferPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()         // Reset the buffer for reuse
		bufferPool.Put(buf) // Return the buffer to the pool after use
	}()

	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(source); err != nil {
		return fmt.Errorf("error encoding source: %w", err)
	}

	decoder := json.NewDecoder(buf)
	if err := decoder.Decode(destination); err != nil {
		return fmt.Errorf("error decoding destination: %w", err)
	}

	return nil
}
