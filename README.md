# 介绍
这是一个 Go 的 Repository contract 库，旨在实现多库之间（例如 MySQL、MongoDB）无缝切换。以最小的代价来降低切库带来的风险。<br>
同时，对于返回结果集进行封装，提供便捷的、安全的 Collection、Dict 操作。

> 切库常见于从 MySQL 切到 MongoDB。 
> 
> 但本仓库目标不仅仅是这两个数据库的切换

# 特性
1. 多数据库之间无缝切换
2. 封装结果集，并提供便捷操作
3. 支持软删除

# 约定
1. ### ID 命名
    建议 结构体中字段使用`ID`，SQL 数据库使用 `id` 作为主键， MongoDB 使用 `_id` 作为主键

2. ### 使用分布式 ID 生成
    建议使用 [go-id](https://github.com/ace-zhaoy/go-id)生成 int64 的 ID， 而不是使用 MySQL 的主键自增或 Mongo 的 ObjectId。
    虽然 Repository 支持不同类型的 ID，但切库时往往变得很复杂。

3. ### 使用 `IDField()` 获取 ID 的字段名
    对于需要指定 ID 字段的操作，使用 `IDField()` 获取主键字段名，而不是固定写`_id` OR `id`。
    除非保证无切库需求。

4. ### 软删字段名使用 `deleted_at`（DeletedAt），值为 int64 类型
    软删并非必须项，但若使用软删，则使用 int64 的 deleted_at 字段。默认值为 0，删除写入当前秒级时间戳。

5. ### “未找到文档/记录”的错误处理
    对于单记录查询，建议使用 `errors.Is(err, repository.ErrNotFound)` 统一判断。<br>
    因为不同库返回的错误并不一致，比如 gorm 返回 `gorm.ErrRecordNotFound`， mongo 返回 `mongo.ErrNoDocuments`。
    如果使用各自的 err 判断，切库就存在漏判风险。<br>
    > 业务错误建议使用 [errors](https://github.com/ace-zhaoy/errors)，支持错误码、栈信息、便捷的错误检查与处理

> 默认使用的 Collection、Dict 是无锁的（非并发安全），若需要并发安全，可以调用 Safe(true)，将返回带“读写锁”的对象