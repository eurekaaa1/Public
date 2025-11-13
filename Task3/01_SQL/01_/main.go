package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 定义测试模型
type User struct {
	gorm.Model      // 包含 ID, CreatedAt, UpdatedAt, DeletedAt 字段
	Username string `gorm:"size:50;not null"` // 用户名
	Age      int    // 年龄
}

// 连接数据库，返回 gorm.DB 实例和错误
func connectDB() (*gorm.DB, error) {
	// 数据库连接信息
	username := "mark_user"
	password := "123456"
	host := "localhost"
	port := "3306"
	dbName := "mark_db"

	// 构建 DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbName)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 显示 SQL 日志
	})
	if err != nil {
		return nil, fmt.Errorf("GORM 连接数据库失败: %w", err) // 包装错误
	}

	// 自动迁移表（创建 User 表）
	if err := db.AutoMigrate(&User{}); err != nil {
		return nil, fmt.Errorf("自动迁移表失败: %w", err)
	}

	fmt.Println("数据库连接成功")
	return db, nil
}

// 关闭数据库连接
func closeDB(db *gorm.DB) error {
	// 获取底层的 *sql.DB 实例
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("获取底层 SQL 连接失败: %w", err)
	}

	// 关闭连接
	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("关闭数据库连接失败: %w", err)
	}

	fmt.Println("数据库连接已关闭")
	return nil
}

func main() {
	// 1. 连接数据库
	db, err := connectDB()
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	// 注意：这里不直接 defer closeDB，而是在操作完成后显式关闭（也可以用 defer）
	// defer func() {
	// 	if err := closeDB(db); err != nil {
	// 		log.Printf("关闭连接时出错: %v", err)
	// 	}
	// }()

	// 2. 执行数据库操作（插入+查询）
	testUser := User{Username: "test_gorm", Age: 25}
	result := db.Create(&testUser)
	if result.Error != nil {
		log.Fatalf("插入数据失败: %v", result.Error)
	}
	log.Printf("成功插入测试数据, ID: %d", testUser.ID)

	var queryUser User
	db.First(&queryUser, testUser.ID)
	log.Printf("查询到的数据: ID=%d, Username=%s, Age=%d",
		queryUser.ID, queryUser.Username, queryUser.Age)

	// 3. 关闭连接
	if err := closeDB(db); err != nil {
		log.Printf("关闭连接出错: %v", err)
	}

	fmt.Println("test_gorm over")
}