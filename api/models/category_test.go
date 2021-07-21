package models

import "testing"

func TestCategoryModel_Validate(t *testing.T) {
	c := &Category{}
	c.Description = ""

	expected := "category.description can't be empty"

	if err := c.Validate(); err.Error() != expected {
		t.Error(err)
	}
}
