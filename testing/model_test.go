// suffix  XXXX_test.go

package testing

import (
	"math/rand"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/yangwawa0323/learning-golang-gorm/model"
)

func testCreateUsers(t *testing.T) {
	fake := faker.New()
	for i := 0; i < 9; i++ {
		var user = model.User{Name: fake.Person().Name()}
		db.Create(&user)
	}
}

var languages = [3]model.Language{
	{Name: "Chinese"},
	{Name: "English"},
	{Name: "Japanese"},
}

func testCreateLanguages(t *testing.T) {
	db.Create(&languages)
}

func generateLanguages() []model.Language {
	rand.Seed(time.Now().UnixMilli())
	var ret []model.Language
	var langs []model.Language
	db.Find(&langs) // with ID
	for i := 0; i < rand.Intn(len(langs)); i++ {
		ret = append(ret, langs[i])
	}
	return ret
}

func testCreateUserLanguages(t *testing.T) {
	var users []model.User
	db.Find(&users)

	for _, user := range users {
		var langs = generateLanguages() // --> []model.Langauge
		user.Languages = langs
		db.Save(&user)
	}

}

func testPrintOutUserLanguages(t *testing.T) {
	var user model.User
	db.Preload("Languages").Find(&user, 2)

	t.Log(user)
	for idx, lang := range user.Languages {
		t.Log(idx, ", user language: ", lang.Name)
	}
}

// TestXXXX(t *testing.T)
func TestModel(t *testing.T) {
	// t.Run("create users ", testCreateUsers)

	// t.Run("create languages", testCreateLanguages)

	// t.Run("create user languages", testCreateUserLanguages)

	t.Run("print out user languages", testPrintOutUserLanguages)
}

// Author 1-->N  Books
