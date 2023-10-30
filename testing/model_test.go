// suffix  XXXX_test.go

package testing

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/yangwawa0323/learning-golang-gorm/model"
)

func testCreatUserCompany(t *testing.T) {
	fake := faker.New()
	var company = model.Company{
		Name: fake.Company().Name(),
	}
	db.Create(&company)
	for i := 0; i < 3; i++ {
		var user = model.User{
			Company: company,
			Name:    fake.Person().Name(),
		}
		t.Log("create a user")
		db.Create(&user)
	}
}

func testPreload(t *testing.T) {
	var user model.User

	// SELECT * FROM companies;   ---> Memory
	// SELECT * FROM users WHERE id = 6;
	// db.First(&user, 6)   // 👎
	db.Preload("Company").First(&user, 6) // 💯

	stmt := dryDb.Preload("Company").First(&user, 6).Statement

	t.Log("user company : ", user.Company.Name)
	t.Log("user : ", user)
	t.Log("==================")
	t.Log(stmt.SQL.String())
	t.Logf("\n\n")
}

func testJoin(t *testing.T) {
	var user model.User

	// SELECT * FROM users LEFT JOIN companies
	// ON users.company_id = companies.id WHERE user.id = 6
	db.Joins("Company").First(&user, 6)

	stmt := dryDb.Joins("Company").First(&user, 6).Statement

	t.Log("user company : ", user.Company.Name)
	t.Log("user : ", user)
	t.Log("==================")
	t.Log(stmt.SQL.String())
	t.Logf("\n\n")

}

func testPreloadUserForCompany(t *testing.T) {
	var company model.Company
	db.Preload("User").First(&company, 13)

	t.Log(company)
}

func testCreateUserCreditCard(t *testing.T) {
	fake := faker.New()

	var user model.User
	db.First(&user, 7)

	var card model.CreditCard = model.CreditCard{
		Number: fake.Payment().CreditCardNumber(),
	}
	db.Create(&card)

	user.CreditCard = card
	db.Save(user)
}

// TestXXXX(t *testing.T)
func TestModel(t *testing.T) {
	// t.Run("Create users and companies", testCreatUserCompany)

	// Preload data
	// t.Run("Preload data", testPreload)

	// Join
	// t.Run("Join tables", testJoin)

	// t.Run("Preload users for company", testPreloadUserForCompany)  // :-1

	t.Run("create user credit card", testCreateUserCreditCard)
}

// Author 1-->N  Books
