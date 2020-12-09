package model

import (
    "database/sql"
    "fmt"
    "github.com/tal-tech/go-zero/core/stores/cache"
    "github.com/tal-tech/go-zero/core/stores/sqlc"
    "github.com/tal-tech/go-zero/core/stores/sqlx"
    "github.com/tal-tech/go-zero/core/stringx"
    "github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
    "strings"
    "time"
)

var (
    userFieldNames          = builderx.FieldNames(&User{})
    userRows                = strings.Join(userFieldNames, ",")
    userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "create_time", "update_time"), ",")
    userRowsExpectAutoSetID   = strings.Join(stringx.Remove(userFieldNames, "id", "create_time", "update_time"), ",")
    userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "user", "create_time", "update_time"), "=?,") + "=?"

    userPrefix = "cache#user#user#"
)

type (
    UserModel struct {
        sqlc.CachedConn
        table string
    }

    User struct {
        ID              int32     `db:"id"`
        NickName        string    `db:"nick_name"`
        PassWord        string    `db:"password"`
        Gender          int32     `db:"gender"`
        Phone           string    `db:"phone"`
        Email           string    `db:"email"`
        HeadPortraitUrl string    `db:"head_portrait_url"`
        RegDate         string    `db:"reg_date"`
        Description     string    `db:"description"`
        Status          int32     `db:"status"`
        CreateTime      time.Time `db:"create_time"`
        UpdateTime      time.Time `db:"update_time"`
        DeleteTime      time.Time `db:"delete_time"`
    }
)

func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, table string) *UserModel {
    return &UserModel{
        CachedConn: sqlc.NewConn(conn, c),
        table:      table,
    }
}

func (m *UserModel) Insert(data User) (sql.Result, error) {
    query := `insert into ` + m.table + ` (` + userRowsExpectAutoSetID + `) values (?, ?)`
    return m.ExecNoCache(query, data.NickName, data.PassWord, data.Gender, data.Phone, data.Email, data.HeadPortraitUrl,
        data.RegDate, data.Description, data.Status)
}

func (m *UserModel) FindOne(book string) (*User, error) {
    bookKey := fmt.Sprintf("%s%v", userPrefix, book)
    var resp User
    err := m.QueryRow(&resp, bookKey, func(conn sqlx.SqlConn, v interface{}) error {
        query := `select ` + userRows + ` from ` + m.table + ` where book = ? limit 1`
        return conn.QueryRow(v, query, book)
    })
    switch err {
    case nil:
        return &resp, nil
    case sqlc.ErrNotFound:
        return nil, ErrNotFound
    default:
        return nil, err
    }
}

//func (m *UserModel) Update(data User) error {
//    bookKey := fmt.Sprintf("%s%v", userPrefix, data.ID)
//    _, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
//        query := `update ` + m.table + ` set ` + bookRowsWithPlaceHolder + ` where book = ?`
//        return conn.Exec(query, data.Price, data.Book)
//    }, bookKey)
//    return err
//}

func (m *UserModel) Delete(book string) error {
    bookKey := fmt.Sprintf("%s%v", userPrefix, book)
    _, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
        query := `delete from ` + m.table + ` where book = ?`
        return conn.Exec(query, book)
    }, bookKey)
    return err
}