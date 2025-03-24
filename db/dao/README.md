#model生成示例
#官方文档
https://gorm.io/zh_CN/gen/

#请根据自身需求生成，不要每次全部替换
gentool -dsn "root:root@tcp(localhost:3306)/store?charset=utf8mb4&parseTime=True&loc=Local" -outPath="/project/db/mysql/dao"  -tables=""

