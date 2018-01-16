package spiralMemory

import "testing"

func TestStorageSet(t *testing.T) {
	storage := NewStorage()
	storage.Set(0, 0, 24)
	if storage.Data[0][0] != 24 {
		t.Fatalf("SET has failed. Wanted %d, Got %d", 24, storage.Data[0][0])
	}
	if storage.Get(0, 0) != 24 {
		t.Fatalf("GET has failed")
	}
	if storage.Get(91247, 123) != 0 {
		t.Fatalf("GET has failed for unset elements")
	}
}
