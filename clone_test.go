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
