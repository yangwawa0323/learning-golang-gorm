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

	db.Create(&users)
}

func testDbDelete(t *testing.T) {
	db.Where("id = ?", 99).Delete(&model.User{})
}

func testDbFind(t *testing.T) {
	var user model.User

	// SELECT * FROM users WHERE deleted_at IS NULL AND id = 99 LIMIT 1
	db.Where("id = ?", 99).First(&user)
	t.Log(user)
}

func testDbUnscopedFind(t *testing.T) {

	var user model.User
	// SELECT * FROM users WHERE id = 99
	db.Unscoped().Where("id = ?", 99).First(&user)
	t.Log(user)

}

// TestXXXX(t *testing.T)
func TestModel(t *testing.T) {

	// t.Run("Create user model", testCreateUserModel)

	// t.Run("Create 100 jinzhu users", testBatchCreateUsers)

	t.Run("Hard delete", testDbDelete)

	t.Run("Find id = 99 user", testDbFind)

	t.Run("db.Unscoped().Find()", testDbUnscopedFind)
}
