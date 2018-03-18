package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

type User struct {
	Id    int    `orm:"pk;auto"`
	Email string `orm:"size(50)"`
	Name  string `orm:"size(50)"`
	Pass  string `orm:"size(100)"`
}

type Profile struct {
	Id        int    `orm:pk`
	Sex       int8
	Age       int8
	TelePhone int64
	Address   string `orm:"size(256)"`
	UserId    int
}

func init() {
	orm.RegisterModel(new(User), new(Profile))

}

func (this *User) AddUser(data map[string]string) string {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	user := new(User)
	profile := new(Profile)
	user.Name = data["name"]
	user.Email = data["email"]
	user.Pass = data["pass"]

	id, err := o.Insert(user)
	if err != nil {
		return "0"
	}

	sex, _ := strconv.ParseInt(data["sex"], 10, 64)
	profile.Sex = int8(sex)
	age, _ := strconv.ParseInt(data["sex"], 10, 64)
	profile.Age = int8(age)
	profile.TelePhone, _ = strconv.ParseInt(data["telephone"], 10, 64)
	profile.Address = data["address"]

	profile.UserId = int(id)

	_, err = o.Insert(profile)
	return strconv.FormatInt(id, 10)
}

func (this *User) GetUserInfoByEmail(email string) (user User, err error) {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	err = o.QueryTable(user).Filter("email", email).One(&user)

	if err != nil {
		return User{}, err
	}
	return user, err
}
