// suffix  XXXX_test.go

package testing

import (
	"fmt"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/yangwawa0323/learning-golang-gorm/model"
	"gorm.io/gorm"
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

func testDbFirstWithKey(t *testing.T) {
	var user model.User
	// db.First(&user, 10) // 👍 select * from users where id = 10
	db.First(&user, "10") //  👍 select * from users where id = 10

	t.Log(user)
}

func testDbFindThreeUser(t *testing.T) {
	var users []model.User
	db.Find(&users, []int{1, 2000, 3000})
	t.Log(users)
}

func testDbFirstWithId(t *testing.T) {
	var user = model.User{Model: gorm.Model{ID: 10}} // 👍
	db.First(&user)
	// var user model.User   // 👎
	// db.Model(&model.User{Model: gorm.Model{ID: 10}}).First(&user)
	t.Log(user)
}

func printUsers(t *testing.T, users []model.User) {
	var total = len(users)
	for idx, user := range users {
		t.Log(fmt.Sprintf("%d/%d : ", idx, total), user)
	}

}

func testDbWhereCondition(t *testing.T) {
	var users []model.User

	// SELECT * FROM users WHERE name = "xxx"
	// db.Where("name = ?", "Mr. Gaetano Bartell").Find(&users)

	// SELECT * FROM users WHERE name <> "xxx"
	// db.Where("name <> ?", "Mr. Gaetano Bartell").Find(&users)

	// SELECT * FROM users WHERE name IN ("xxx", "YYY","ZZZ")
	// db.Where("name IN ?", []string{"Mr. Gaetano Bartell", "Yang kun", "Claudie Effertz"}).Find(&users)

	// SELECT * FROM users name LIKE "Gaet%"
	// db.Where("name LIKE ?", "%Gaet%").Find(&users)
	// db.Where("name LIKE ?", "Mr.%").Find(&users)

	// SELECT * FROM users WHERE name LIKE "Mr.%" AND age < ?
	// db.Where("name LIKE ? AND age < ?", "Mr.%", "105").Find(&users)

	//
	db.Where(map[string]interface{}{
		"name": "Mr. Wade Ernser DDS",
		"age":  25,
	}).Find(&users)
	// t.Log(users)
	printUsers(t, users)
}

func testDbNotAndOrCondition(t *testing.T) {
	var users []model.User

	// SELECT * FROM users WHERE NOT name = "xxxx"
	// db.Not("name = ?", "Mr. Gaetano Bartell").Find(&users)

	// SELECT * FROM users WHERE id NOT IN (1,30,40)
	// db.Not([]int64{1, 30, 40, 50}).Find(&users)

	// db.Not(model.User{Name: "Geovanni Schneide", Age: 154}).Find(&users)

	db.Where("name = ?", "Geovanni Schneider").Or("email is not null").Find(&users)

	printUsers(t, users)
}

func testDbSelect(t *testing.T) {
	var users []model.User
	db.Select("email", "age").Find(&users)

	printUsers(t, users)
}

func testDbOrder(t *testing.T) {
	var users []model.User
	// SELECT * FROM users ORDER BY age desc, name;
	// db.Order("age desc , name").Find(&users)
	db.Where("age > ? and age < ? ", 4, 14).Order("age desc").Order("name").Find(&users)
	printUsers(t, users)
}

func testDbLimitOffset(t *testing.T) {
	var users []model.User

	// SELECT * FROM users LIMIT 3;
	// db.Order("age desc").Limit(3).Find(&users)  👍

	// page_num - 1 * 10   3 - 1 * 10 = 20
	db.Limit(10).Offset(20).Order("ID").Find(&users) // 2-11, 12-21, 22-31
	printUsers(t, users)
}

func testDbGroupHaving(t *testing.T) {
	type result struct {
		CountryCode string
		Total       int
	}

	// var results []result

	// ORM
	// SELECT country_code, sum(population) as total from
	// cities GROUP BY country_code HAVING total < 500;
	// db.Model(&model.City{}).Select("country_code as CountryCode, sum(population) as Total").Group("country_code").Having("Total < ? ", 500).Find(&results)

	// SELECT country_code, sum(population) as Total from
	// cities GROUP BY country_code ORDER BY Total DESC limit 1;
	// t.Logf("%#v\n", results)
	var r result
	db.Model(&model.City{}).Select("country_code as CountryCode, sum(population) as Total").Group("country_code").Order("Total desc").Limit(1).First(&r)
	t.Logf("%#v", r)
}

func testObjectUpdate(t *testing.T) {
	// var user model.User
	// db.Order("ID desc").First(&user)

	// user.Name = "DDDDD Harvey"

	// gorm.Save()  if has ID use UPDATE

	db.Save(&model.User{Model: gorm.Model{ID: 100}, Name: "Jinzhu", Age: 99})
	// gorm.Sace()  if has not ID use INSERT INTO
	// db.Save(&model.User{Name: "Jinzhu", Age: 99})

	// t.Log("user id:", user.ID)

}

func testDbUpdate(t *testing.T) {
	result := db.Model(&model.User{}).Where("id = ?", 100).Update("name", "Jinzhu").Update("age", 99)

	if result.Error != nil {
		t.Fail()
	}

	t.Logf("Update %d row(s)", result.RowsAffected)

}

func testDbUpdates(t *testing.T) {
	// result := db.Model(&model.User{}).Where("id = ? ", 100).Updates(&model.User{Age: 56, Name: "Hello"})

	result := db.Model(&model.User{}).Where("id = ? ", 100).Updates(map[string]interface{}{
		"name": "world", "age": 77,
	})
	if result.Error != nil {
		t.Fail()
	}

	t.Logf("Update %d row(s)", result.RowsAffected)
}

func testDbSelectUpdate(t *testing.T) {
	// db.Model(&model.User{}).Where("id = ? ", 100).Select("age").Updates(map[string]interface{}{
	// 	"name": "hello_world_55", "age": 55, "email": "unknown66@qq.com",
	// })

	var email = "xxxx@qq.com"
	db.Model(&model.User{}).Where("id = ? ", 99).Select("age").Updates(
		&model.User{Name: "xxxx", Age: 44, Email: &email},
	)
}

// Update does work without Where() use the Primary update if have the primary
func testDbWhereUpdates(t *testing.T) {
	// result := db.Model(&model.User{}).Where("id IN ?", []int{80, 99}).Updates(model.User{Name: "Hello", Age: 18})

	result := db.Model(&model.User{}).Where("id <> ? ", "-1").Updates(model.User{Name: "Hello", Age: 18})
	if result.Error != nil {
		t.Fail()
	}

	t.Logf("Updated %d row(s)\n", result.RowsAffected)
}

// Update age = age + 1
func testUpdateGormExpr(t *testing.T) {
	result := db.Model(&model.User{}).Where("id <> ?", -1).Update("age", gorm.Expr("age + ?", 1))
	if result.Error != nil {
		t.Fail()
	}

	t.Logf("Updated %d row(s)\n", result.RowsAffected)

}

func testDbUpdateColumnGormExpr(t *testing.T) {
	result := db.Model(&model.City{}).Where("population < 500").UpdateColumn("population", gorm.Expr("population * ?", 1.1))
	if result.Error != nil {
		t.Fail()
	}

	t.Logf("Updated %d row(s)\n", result.RowsAffected)
}

type Result struct {
	ID   int
	Name string
	Age  int
}

func testDbRawScan(t *testing.T) {
	// var results []Result  // 👍
	var results []map[string]interface{}
	db.Raw("SELECT name,age FROM friends LIMIT 10").Scan(&results)
	// t.Log(results)
	for _, result := range results {
		for key, value := range result {
			t.Log(key, " : ", value)
		}
		t.Log("===============")
	}
}

func testDbExec(t *testing.T) {
	db.Exec("TRUNCATE TABLE friends")
}

func testDbSession(t *testing.T) {
	var city model.City
	stmt := dryDb.Where("name LIKE ? AND population < ?", "Til%", 500).First(&city).Statement
	t.Log(stmt.SQL.String())
	t.Log(stmt.Vars)
}

func testDbToSQL(t *testing.T) {
	var sql string = db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		var users []model.User
		return tx.Model(&model.User{}).Where("id IN ?", []int{10, 20, 30}).Limit(10).Order("age desc").Find(&users)

	})

	t.Log(sql)
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
	// t.Run("db.Table.Take", testDbTableTake)

	// db.Table().First()  👎
	// t.Run("db.Table.First", testDbTableFirst)

	// db.First() by primary key
	// t.Run("db.First() by primary key", testDbFirstWithKey)

	// db.Find() query three users
	// t.Run("db.Find() query 3 users", testDbFindThreeUser)

	// db.First() with id
	// t.Run("db.First() with id ", testDbFirstWithId)

	// db.Where()
	// t.Run("db.Where() ", testDbWhereCondition)

	// db.Not()
	// t.Run("db.Not()", testDbNotAndOrCondition)

	// db.Select()
	// t.Run("db.Select()", testDbSelect)

	// db.Order()
	// t.Run("db.Order()", testDbOrder)

	// db.Limit()
	// t.Run("db.Limit()", testDbLimitOffset)

	// db.Select().Group()
	// t.Run("db.Select().Group()", testDbGroupHaving)

	// object update
	// t.Run("object update", testObjectUpdate)

	// db.Update()  update single field
	// t.Run("db.Update()", testDbUpdate)

	// db.Updates() update multiple field
	// t.Run("db.Updates()", testDbUpdates)

	// db.Select().Updates()
	// t.Run("db.Select().Updates()", testDbSelectUpdate)

	// t.Run("db.Where().Update()", testDbWhereUpdates)

	//  gorm.Expr
	// t.Run("db.Update( gorm.Expr() )", testUpdateGormExpr)

	// t.Run("db.UpdateColumn()", testDbUpdateColumnGormExpr)

	// Raw query
	// t.Run("db.Raw().Scan()", testDbRawScan)

	// Execution
	// t.Run("db.Exec()", testDbExec)

	// Dry Run
	t.Run("db.Session() Dry Run", testDbSession)

	// db.ToSQL
	t.Run("db.ToSQL()", testDbToSQL)
}

// Author 1-->N  Books
