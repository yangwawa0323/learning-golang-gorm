// suffix  XXXX_test.go

package testing

import (
	"fmt"
	"testing"
	"time"

	"github.com/jaswdr/faker"
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

	p := faker.New().Person()
	for i := 0; i < 100; i++ {
		var user = model.User{
			Name: p.Name(),
			Age:  p.Faker.UInt8(),
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

func testDbFirst(t *testing.T) {
	var user model.User
	db.First(&user)
	t.Logf("%#v", user)
}

func testDbTake(t *testing.T) {
	var user model.User
	db.Take(&user)
	t.Logf("%#v", user)
}

func testDbLast(t *testing.T) {
	var user model.User
	db.Last(&user)
	t.Logf("%v", user)
}

func testDbModelFirst(t *testing.T) {
	result := map[string]interface{}{}
	db.Model(&model.User{}).First(&result)

	t.Logf("%#v", result)
}

func testDbModelTake(t *testing.T) {
	result := map[string]interface{}{}
	db.Model(&model.User{}).Take(&result)

	t.Logf("%#v", result)
}

func testDbTableTake(t *testing.T) {
	result := map[string]interface{}{}
	query_res := db.Table("country").Take(&result)

	if query_res.Error != nil {
		t.Fail()
	}
	t.Logf("%#v", result)
}

func testDbTableFirst(t *testing.T) {
	result := map[string]interface{}{}
	query_res := db.Table("country").First(&result)
	if query_res != nil {
		t.Fail()
	}
	t.Logf("%#v", result)
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
	// t.Run("Use map batch create two users", testBatchCreateUsers)

	// db.First   order by id ASC limit 1
	// t.Run("db.First", testDbFirst)

	// db.Take   limit 1
	// t.Run("db.Take", testDbTake)

	// db.Last  order by id DESC limit 1
	//t.Run("db.Last", testDbLast)

	// db.Model().First()
	// t.Run("db.Model.First", testDbModelFirst)

	// db.Model().Take()
	// t.Run("db.Model.Take", testDbModelTake)

	// db.Table().Take()
	t.Run("db.Table.Take", testDbTableTake)

	// db.Table().First()  👎
	t.Run("db.Table.First", testDbTableFirst)
}

// Author 1-->N  Books
