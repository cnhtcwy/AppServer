package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)
type SysUser struct {
	Id     int64 `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email string `db:"email"`
}
type User struct {
	Id int `json:"id" db:"id"`
	CreatedAt string `json:"created_at" db:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" time_format:"2006-01-02 15:04:05"`
	UUID        uuid.UUID    `json:"uuid" db:"uuid" gorm:"comment:用户UUID"`
	Username    string       `json:"username" db:"username" gorm:"comment:用户登录名"`
	Password    string       `json:"password" db:"password"  gorm:"comment:用户登录密码"`
	NickName    string       `json:"nick_name" db:"nick_name" gorm:"default:系统用户;comment:用户昵称" `
	HeaderImg   string       `json:"header_img" db:"header_img" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	AuthorityId string       `json:"authority_id" db:"authority_id" gorm:"default:888;comment:用户角色ID"`
}

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/qmplus?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return
}

// 查询单条数据示例
func queryRowDemo() {
	//sqlStr := "select id, username, password,email from sys_user where id=?"
	var u SysUser
	err := db.Get(&u, `select id, username, password,email from sys_user where id=?`, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.Username, u.Password, u.Email)
}
// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, username, password,email from sys_user"
	var users []SysUser
	err := db.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}
func Regist() error{
	res,err := db.Exec("insert into sys_users(created_at,updated_at,uuid,username,password,nick_name,header_img,authority_id) values (?,?,?,?,?,?,?,?)",time.Now(),time.Now(),uuid.NewV4(),"root","123","sonic","http://www.baidu.com","123")
	fmt.Println(res)
	return err
}
func Delete(id int) error {
	_,err:=db.Exec("delete from sys_users where id = ?",id)
	return err
}
// 格式化为:2006-01-02 15:04:05
func GetNormalTimeString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
// 字符串转时间
func GetTimeByString(timestring string) (time.Time,error){
	if timestring == "" {
		return time.Time{},nil
	}
	return time.ParseInLocation("20060102150405", timestring, time.Local)
}

// 标准字符串转时间
func GetTimeByNormalString(timestring string) (time.Time){
	if timestring == "" {
		return time.Time{}
	}
	time1,_ := time.ParseInLocation("2006-01-02 15:04:05", timestring, time.Local)
	return time1
}
var user1 = User{CreatedAt: GetNormalTimeString(time.Now()),UpdatedAt: time.Now(),UUID: uuid.NewV4(),Username: "test1",Password: "111111",NickName: "littleNick",HeaderImg: "https://tupiam.com",AuthorityId: "888"}
var user2 = User{CreatedAt: GetNormalTimeString(time.Now()),UpdatedAt: time.Now(),UUID: uuid.NewV4(),Username: "test2",Password: "111111",NickName: "littleNick",HeaderImg: "https://tupiam.com",AuthorityId: "888"}
var user3 = User{CreatedAt: GetNormalTimeString(time.Now()),UpdatedAt: time.Now(),UUID: uuid.NewV4(),Username: "test3",Password: "111111",NickName: "littleNick",HeaderImg: "https://tupiam.com",AuthorityId: "888"}
//插入单条struct
func InsertOnlyOneByStruct(user User) error {
	fmt.Println(user)
	sqlStr := "insert into sys_users(created_at,updated_at,uuid,username,password,nick_name,header_img,authority_id) values (:created_at,now(),:uuid,:username,:password,:nick_name,:header_img,:authority_id)"
	m := make(map[string]interface{})
	j,_ := json.Marshal(user)
	json.Unmarshal(j,&m)
	fmt.Println(m)
	_,err := db.NamedExec(sqlStr, m)
	return err
}
func BatchInsertUser1(users []*User) error {
	// 存放 (?, ?) 的slice
	valueStrings := make([]string, 0, len(users))
	// 存放values的slice
	valueArgs := make([]interface{}, 0, len(users) * 2)
	// 遍历users准备相关数据
	for _, u := range users {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?,?,?,?,?,?,?,?)")
		valueArgs = append(valueArgs, u.CreatedAt)
		valueArgs = append(valueArgs, u.UpdatedAt)
		valueArgs = append(valueArgs, u.UUID)
		valueArgs = append(valueArgs, u.Username)
		valueArgs = append(valueArgs, u.Password)
		valueArgs = append(valueArgs, u.NickName)
		valueArgs = append(valueArgs, u.HeaderImg)
		valueArgs = append(valueArgs, u.AuthorityId)
	}
	// 自行拼接要执行的具体语句
	stmt := fmt.Sprintf("INSERT INTO sys_users(created_at,updated_at,uuid,username,password,nick_name,header_img,authority_id) VALUES %s",
		strings.Join(valueStrings, ","))
	fmt.Println(stmt)
	fmt.Println(valueArgs)
	_, err := db.Exec(stmt, valueArgs...)
	return err
}
// BatchInsertUsers3 使用NamedExec实现批量插入
func BatchInsertUsers3(users []*User) error {
	_, err := db.NamedExec("INSERT INTO sys_users(created_at,updated_at,uuid,username,password,nick_name,header_img,authority_id) VALUES (:created_at,:updated_at,:uuid,:username,:password,:nick_name,:header_img,:authority_id)", users)
	return err
}
// QueryByIDs 根据给定ID查询
func QueryByIDs(ids []int)(users []User, err error){

	// 动态填充id
	query, args, err := sqlx.In("SELECT created_at,updated_at,uuid,username,password,nick_name,header_img,authority_id FROM sys_users WHERE id IN (?)", ids)
	if err != nil {
		return
	}
	fmt.Println(query)
	fmt.Println(args)
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)
	fmt.Println(query)
	err = db.Select(&users, query, args...)
	return
}
func DeleteByIDs(ids []int) error  {
	// 动态填充id
	query, args, err := sqlx.In("DELETE FROM sys_users WHERE id IN (?)", ids)
	if err != nil {
		return err
	}
	fmt.Println(query)
	fmt.Println(args)
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	//query = db.Rebind(query)
	//fmt.Println(query)
	_,err = db.Exec(query, args...)
	return err
}
func main() {
	if err :=initDB();err!= nil{
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}
	fmt.Println("init DB success...")
	//queryRowDemo()
	//queryMultiRowDemo()
	//Delete(4)
	//Regist()
	//users := []*User{&user1,&user2,&user3}
	//BatchInsertUser1(users)

	//InsertOnlyOneByStruct(user1)
	//users,_:=QueryByIDs([]int{2, 5, 6, 1})
	//fmt.Println(users)
	//for u,_ := range users {
	//	fmt.Println(users[u].UUID)
	//}
	DeleteByIDs([]int{9,10,11})
}

