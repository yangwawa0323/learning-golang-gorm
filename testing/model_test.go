// suffix  XXXX_test.go

package testing

import (
	"math/rand"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/yangwawa0323/learning-golang-gorm/model"
)

func testCreateUsers(t *testing.T) {
	fake := faker.New()
	var users []model.User
	max := 70
	min := 56
	for i := 0; i < 30; i++ {
		users = append(users, model.User{
			Name: fake.Person().Name(),
			Age:  uint8(rand.Intn(max-min) + min),
		})
	}
	db.Save(&users)
}

// TestXXXX(t *testing.T)
func TestModel(t *testing.T) {
	t.Run("create users ", testCreateUsers)
}

// Author 1-->N  Books
