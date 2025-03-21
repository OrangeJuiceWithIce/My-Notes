# Gorm
## 定义模型
```Go
type User struct{
    ID uint
    Name string
    Email *String  //指针类型，可为空
    MemberNumber sql.NullString  //来自database/sql包的可空字段
}
```
1.Gorm会将结构体的名称转换为snake_case并加上复数形式，例如User->users.
>如何避免复数形式?
```
db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{
    NamingStrategy: schema.NamingStrategy{
        SingularTable: true
    },
})
```
2.允许结构体嵌套
## 连接数据库
```go
dsn :="user:pass@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
```
```go
db, err := gorm.Open(mysql.New(mysql.Config{
  DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
  DefaultStringSize: 256, // string 类型字段的默认长度
  DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
  DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
  DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
  SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
}), &gorm.Config{})
```