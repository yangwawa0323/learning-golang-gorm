// suffix  XXXX_test.go

package testing

import (
	"fmt"
	"testing"
	"time"

	"github.com/yangwawa0323/learning-golang-gorm/model"
)

type DataMap map[string]interface{} // reflect 👎

func testCreateUser(t *testing.T) {
	birth := time.Now()
	var user = model.User{Name: "Jinzhu", Age: 18, Birthday: &birth}

	result := db.Create(&user)
	if result.Error != nil {
		t.Fail()
	}

	t.Log("create a user succeed.")
	// user.ID
	// result.Error
	// result.RowsAffected
}

func testBatchCreateUsers(t *testing.T) {
	var users []model.User = make([]model.User, 0)
	for i := 0; i < 100; i++ {
		var user = model.User{
			Name: fmt.Sprintf("JinZhu%03d", i),
			Age:  18,
		}
		users = append(users, user)
	}

	result := db.Create(&users)
	t.Log(fmt.Sprintf("create %d users", result.RowsAffected))

}

func testCreateMapUser(t *testing.T) {
	result := db.Model(&model.User{}).Create(
		map[string]interface{}{
			"Name":     "JinZhu",
			"Age":      18,
			"Birthday": time.Now(),
		})

	if result.Error != nil {
		t.Fail()
	}
	t.Log("use map create user succeed")
}

func testBatchCreateMapUsers(t *testing.T) {
	result := db.Model(&model.User{}).Create(
		// []DataMap{   👎
		[]map[string]interface{}{
			{
				"Name":     "JinZhu",
				"Age":      18,
				"Birthday": time.Now(),
			},
			{
				"Name": "Yangwawa",
				"Age":  49,
			},
		})

	if result.Error != nil {
		t.Fail()
	}
	t.Log("use map create user succeed")
}

// TestXXXX(t *testing.T)
func TestModel(t *testing.T) {

	// t.Run("Create user model", testCreateUserModel)
	// t.Run("Update user", testUpdateUser)

	// create a user
	// t.Run("Create a user", testCreateUser)

	// create 100 users
	// t.Run("Batch create 100 users", testBatchCreateUsers)

	// use map to create user
	// t.Run("Use map to create user", testCreateMapUser)

	// use map batch create two users
	t.Run("Use map batch create two users", testBatchCreateMapUsers)
}

// Author 1-->N  Books
