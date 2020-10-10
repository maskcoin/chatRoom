package model

import (
	"encoding/json"
	"fmt"
	"mygithub_code/chatRoom/common/message"

	"github.com/garyburd/redigo/redis"
)

// 我们在服务器启动后，就初始化一个UserDao实例
// 把它做成全局的变量，在需要和redis操作时，就直接使用即可
var (
	MyUserDao *UserDao
)

// UserDao 定义一个结构体，完成对User结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

// NewUserDao 使用工厂模式，创建一个UserDao的实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}

	return
}

// GetUserByID 思考一下在UserDao 应该提供哪些方法给我们
// 1.根据用户id 返回一个 User实例或者error
func (dao *UserDao) getUserByID(conn redis.Conn, id int) (user *message.User, err error) {
	// 通过给定的id 去redis查询用户
	res, err := redis.String(conn.Do("HGET", "users", id))

	if err != nil {
		// 错误
		if err == redis.ErrNil { // 表示在 users 哈希中没有找到对应的hash
			err = ErrorUserNotExists
		}
		return
	}

	// 这里我们需要把res 反序列化成User实例
	user = &message.User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(res), &user) err =", err)
		return
	}

	return
}

// Login 完成对用户的校验
// 如果用户的ID和pwd都正确，则返回一个user实例
// 如果用户的ID或者pwd有错误，则返回对应的错误信息
func (dao *UserDao) Login(userID int, userPwd string) (user *message.User, err error) {
	// 先从UserDao的连接池中取出一个连接
	conn := dao.pool.Get()
	defer conn.Close()

	user, err = dao.getUserByID(conn, userID)
	if err != nil {
		return
	}

	// 这时证明这个用户是获取到了
	if user.UserPwd != userPwd {
		err = ErrorUserPwd
		return
	}

	return
}

// Register comment
func (dao *UserDao) Register(user *message.User) (err error) {
	// 先从UserDao的连接池中取出一个连接
	conn := dao.pool.Get()
	defer conn.Close()

	_, err = dao.getUserByID(conn, user.UserID)
	if err == nil {
		err = ErrorUserExists
		return
	} else if err == ErrorUserNotExists {
		err = nil
		data, err2 := json.Marshal(user)
		if err2 != nil {
			fmt.Println("json.Marshal(user) err=", err2)
			err = err2
			return
		}

		_, err = conn.Do("HSet", "users", fmt.Sprintf("%d", user.UserID), string(data))
		if err != nil {
			return
		}
	} else {
		return
	}

	return
}
