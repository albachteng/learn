package store

import (
	"fmt"
	"os"
	"testing"
)

var (
	setupCalled    bool
	store          *Store
	teardownCalled bool
	entries        = []struct {
		key   string
		value string
	}{
		{
			key:   "a",
			value: "123",
		}, {
			key:   "b",
			value: "789",
		}, {
			key:   "a",
			value: "567",
		}, {
			key:   "c",
			value: "456",
		}, {
			key:   "d",
			value: "789",
		}, {
			key:   "a",
			value: "101",
		}, {
			key:   "b",
			value: "112",
		}, {
			key:   "e",
			value: "131",
		}, {
			key:   "f",
			value: "415",
		}, {
			key:   "d",
			value: "161",
		},
	}
)

func setup() {
	fmt.Println("running setup...")
	setupCalled = true
	var err error
	store, err = NewStore("test", 17)
	if err != nil {
		panic(err)
	}
}

func teardown() {
	fmt.Println("running teardown...")
	for i := 0; i < int(store.len); i++ {
		if err := os.Remove(fmt.Sprintf("test-%d.txt", i)); err != nil {
			continue
		}
	}
	teardownCalled = true
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	if !teardownCalled {
		panic("teardown was not called properly")
	}
	os.Exit(code)
}

func TestSet(t *testing.T) {
	if !setupCalled {
		t.Error("setup function not called")
	}
	for _, entry := range entries {
		err := store.Set(entry.key, entry.value)
		assertNoError(t, err)
	}
	v, err := store.Get("a")
	assertNoError(t, err)
	assertEqualStrings(t, v, entries[5].value) // 5 is index of last "a" key
	v, err = store.Get("b")
	assertNoError(t, err)
	assertEqualStrings(t, v, entries[6].value) // index of last "b" key
	v, err = store.Get("f")
	assertNoError(t, err)
	assertEqualStrings(t, v, entries[8].value) // index of last "b" key
}

func TestCompaction(t *testing.T) {
	err := store.Compaction()
	assertNoError(t, err)
	if len(store.segments) != int(store.len) {
		t.Errorf("expected store.len: %q to be equal to the length of the segments slice: %q", store.len, len(store.segments))
	}
	v, err := store.Get("a")
	assertNoError(t, err)
	assertEqualStrings(t, v, entries[5].value) // 5 is index of last "a" key
	v, err = store.Get("b")
	assertNoError(t, err)
	assertEqualStrings(t, v, entries[6].value) // index of last "b" key
	v, err = store.Get("f")
	assertNoError(t, err)
	assertEqualStrings(t, v, entries[8].value) // index of last "f" key
}

func TestMerge(t *testing.T) {
	err := store.Merge()
	if len(store.segments) != int(store.len) {
		t.Errorf("expected store.len: %d to be equal to the length of the segments slice: %d", store.len, len(store.segments))
	}
	assertNoError(t, err)
	v, err := store.Get("a")
	assertNoError(t, err)
	assertEqualStrings(t, v, entries[5].value) // 5 is index of last "a" key
	v, err = store.Get("b")
	assertNoError(t, err)
	assertEqualStrings(t, v, entries[6].value) // index of last "b" key
	v, err = store.Get("f")
	assertNoError(t, err)
	assertEqualStrings(t, v, entries[8].value) // index of last "f" key
}

func assertEqualStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected no error got %q", err)
	}
}
