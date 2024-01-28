package model

import (
	"testing"
)

func TestCreateUpdateDelete(t *testing.T) {
	myMap := NewAutoIncrementMap()

	// Test Create
	key1 := myMap.Add("Value1", "Value2", "Value3", "Value4", "Value5")
	if key1 != 1 {
		t.Errorf("Expected key for Value1: 1, got: %d", key1)
	}

	// Test Update
	myMap.Update(key1, "UpdatedValue1", "UpdatedValue2", "UpdatedValue3", "UpdatedValue4")
	obj, exists := myMap.GetValueForKey(key1)
	if !exists || obj.Title != "UpdatedValue1" || obj.Content != "UpdatedValue2" || obj.Author != "UpdatedValue3" || obj.Tags != "UpdatedValue4" || obj.PublicationDate != "Value4" {
		t.Errorf("Update failed for key: %d", key1)
	}

	// Test Delete
	myMap.Delete(key1)
	_, exists = myMap.GetValueForKey(key1)
	if exists {
		t.Errorf("Delete failed for key: %d", key1)
	}
}
