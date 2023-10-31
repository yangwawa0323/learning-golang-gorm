// suffix  XXXX_test.go

package testing

import (
	"testing"

	"github.com/yangwawa0323/learning-golang-gorm/model"
)

func testBeforeCreateHook(t *testing.T) {
	var user model.User = model.User{
		Name: "jinzhu64",
		Age:  64,
	}
	result := db.Create(&user)
	if result.Error != nil {
		t.Error(result.Error)
		t.Fail()
	}

}

// TestXXXX(t *testing.T)
func TestModel(t *testing.T) {
	t.Run("before create hook", testBeforeCreateHook)
}

// Author 1-->N  Books
