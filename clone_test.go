package goclonestruct

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type NestedStruct struct {
	InnerName  string
	InnerValue float64
}

type TestStruct struct {
	Name       string
	Value      int
	Tags       []string
	Properties map[string]int
	Nested     NestedStruct
}

// TestCloneUsingGob tests CloneUsingGob by ensuring that:
// - The cloning function does not return an error when cloning a valid source
//   struct to a valid destination struct.
// - The cloning function returns an error when cloning a nil source struct to a
//   valid destination struct.
// - The cloning function returns an error when cloning a valid source struct to
//   a nil destination struct.
func TestCloneUsingGob(t *testing.T) {
	source := &TestStruct{
		Name:  "Test",
		Value: 42,
		Tags:  []string{"tag1", "tag2"},
		Properties: map[string]int{
			"key1": 1,
			"key2": 2,
		},
		Nested: NestedStruct{
			InnerName:  "InnerTest",
			InnerValue: 3.14,
		},
	}
	var destination TestStruct

	// Test successful cloning
	err := CloneUsingGob(source, &destination)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !cmp.Equal(*source, destination) {
		t.Fatalf("expected destination to be %+v, got %+v", *source, destination)
	}

	// Test nil source
	var nilSource *TestStruct
	err = CloneUsingGob(nilSource, &destination)
	if err == nil {
		t.Fatal("expected error when source is nil, got none")
	} else if err.Error() != "source is nil" {
		t.Fatalf("expected 'source is nil' error, got %v", err)
	}

	// Test nil destination
	var nilDestination *TestStruct
	err = CloneUsingGob(source, nilDestination)
	if err == nil {
		t.Fatal("expected error when destination is nil, got none")
	} else if err.Error() != "destination is nil" {
		t.Fatalf("expected 'destination is nil' error, got %v", err)
	}
}

// TestCloneUsingJson tests CloneUsingJson by ensuring that:
// - The cloning function does not return an error when cloning a valid source
//   struct to a valid destination struct, and that the destination matches the source.
// - The cloning function returns an error when cloning a nil source struct to a
//   valid destination struct.
// - The cloning function returns an error when cloning a valid source struct to
//   a nil destination struct.
func TestCloneUsingJson(t *testing.T) {
	source := &TestStruct{
		Name:  "Test",
		Value: 42,
		Tags:  []string{"tag1", "tag2"},
		Properties: map[string]int{
			"key1": 1,
			"key2": 2,
		},
		Nested: NestedStruct{
			InnerName:  "InnerTest",
			InnerValue: 3.14,
		},
	}
	var destination TestStruct

	// Test successful cloning
	err := CloneUsingJson(source, &destination)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !cmp.Equal(*source, destination) {
		t.Fatalf("expected destination to be %+v, got %+v", *source, destination)
	}

	// Test nil source
	var nilSource *TestStruct
	err = CloneUsingJson(nilSource, &destination)
	if err == nil {
		t.Fatal("expected error when source is nil, got none")
	} else if err.Error() != "source is nil" {
		t.Fatalf("expected 'source is nil' error, got %v", err)
	}

	// Test nil destination
	var nilDestination *TestStruct
	err = CloneUsingJson(source, nilDestination)
	if err == nil {
		t.Fatal("expected error when destination is nil, got none")
	} else if err.Error() != "destination is nil" {
		t.Fatalf("expected 'destination is nil' error, got %v", err)
	}
}

// BenchmarkCloneUsingGob benchmarks the performance of the CloneUsingGob function.
// It sets up a source TestStruct with predefined values and runs the benchmark
// in parallel, cloning the source struct to a destination struct repeatedly.
// The benchmark fails if CloneUsingGob returns an error during the cloning process.
func BenchmarkCloneUsingGob(b *testing.B) {
	source := &TestStruct{
		Name:  "Benchmark Test",
		Value: 123,
		Tags:  []string{"tag1", "tag2", "tag3"},
		Properties: map[string]int{
			"key1": 1,
			"key2": 2,
			"key3": 3,
		},
		Nested: NestedStruct{
			InnerName:  "Inner Benchmark",
			InnerValue: 6.28,
		},
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var destination TestStruct
			err := CloneUsingGob(source, &destination)
			if err != nil {
				b.Fatalf("error cloning using gob: %v", err)
			}
		}
	})
}

// BenchmarkCloneUsingJson benchmarks the performance of the CloneUsingJson function.
// It sets up a source TestStruct with predefined values and runs the benchmark
// in parallel, cloning the source struct to a destination struct repeatedly.
// The benchmark fails if CloneUsingJson returns an error during the cloning process.
func BenchmarkCloneUsingJson(b *testing.B) {
	source := &TestStruct{
		Name:  "Benchmark Test",
		Value: 123,
		Tags:  []string{"tag1", "tag2", "tag3"},
		Properties: map[string]int{
			"key1": 1,
			"key2": 2,
			"key3": 3,
		},
		Nested: NestedStruct{
			InnerName:  "Inner Benchmark",
			InnerValue: 6.28,
		},
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var destination TestStruct
			err := CloneUsingJson(source, &destination)
			if err != nil {
				b.Fatalf("error cloning using json: %v", err)
			}
		}
	})
}
