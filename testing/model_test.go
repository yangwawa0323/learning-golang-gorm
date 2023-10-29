// suffix  XXXX_test.go

package testing

import (
	"fmt"
	"testing"

	"github.com/yangwawa0323/learning-golang-gorm/model"
)

func testBatchCreateUsers(t *testing.T) {
	var users []model.User = make([]model.User, 0)
	for i := 0; i < 100; i++ {
		var user = model.User{
			Name: fmt.Sprintf("JinZhu%03d", i),
			Age:  18,
		}
		users = append(users, user)
	}
}

func testUpdateUser(t *testing.T) {

}

// TestXXXX(t *testing.T)
func TestModel(t *testing.T) {

	// t.Run("Create user model", testCreateUserModel)
	t.Run("Update user", testUpdateUser)
}
