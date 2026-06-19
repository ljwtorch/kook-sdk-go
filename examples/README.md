# KOOK SDK for Go - 示例代码

本目录包含 KOOK SDK for Go 的可运行示例代码，展示各模块的使用方法。

## 示例列表

| 示例 | 说明 | 目录 |
|------|------|------|
| [user](user/) | 用户 API（获取用户、上下线） | `examples/user/` |
| [guild](guild/) | 服务器 API（列表、详情、成员管理） | `examples/guild/` |
| [channel](channel/) | 频道 API（创建、编辑、删除） | `examples/channel/` |
| [message](message/) | 消息 API（发送、编辑、删除、回应） | `examples/message/` |
| [card](card/) | 卡片消息（构建和发送） | `examples/card/` |
| [direct_message](direct_message/) | 私聊 API（会话、消息） | `examples/direct_message/` |
| [websocket](websocket/) | WebSocket 事件监听 | `examples/websocket/` |

## 环境变量

所有示例都需要设置以下环境变量：

| 环境变量 | 说明 | 必填 |
|----------|------|------|
| `KOOK_BOT_TOKEN` | Bot Token | 是 |
| `KOOK_GUILD_ID` | 服务器 ID | 部分示例需要 |
| `KOOK_CHANNEL_ID` | 频道 ID | 部分示例需要 |
| `KOOK_TARGET_USER_ID` | 目标用户 ID | 部分示例需要 |

### 获取 Token

1. 访问 [KOOK 开放平台](https://developer.kookapp.cn/)
2. 创建或选择你的 Bot
3. 在 Bot 设置页面获取 Token

### 获取 ID

- **服务器 ID**：在 KOOK 客户端中，右键服务器名称，复制服务器 ID
- **频道 ID**：在 KOOK 客户端中，右键频道名称，复制频道 ID
- **用户 ID**：在 KOOK 客户端中，右键用户头像，复制用户 ID

## 运行示例

```bash
# 设置 Token
export KOOK_BOT_TOKEN="your-bot-token"

# 运行用户示例
go run examples/user/main.go

# 查看帮助
go run examples/user/main.go --help
```

## 测试模式

每个示例都包含三种测试模式（通过环境变量 `KOOK_TEST_MODE` 控制）：

| 模式 | 说明 | 副作用 |
|------|------|--------|
| `read` | 读操作测试 | 无副作用，可反复执行 |
| `write` | 写操作测试 | 有副作用，会创建/修改资源 |
| `dangerous` | 危险操作测试 | 不可逆，会删除资源 |

### 完整测试流程

```bash
# 1. 执行 read 模式确认基础连接正常
export KOOK_TEST_MODE=read
go run examples/message/main.go

# 2. 执行 write 模式测试写操作
export KOOK_TEST_MODE=write
go run examples/message/main.go

# 3. 将 write 模式输出的临时 ID 设为环境变量，执行 dangerous 模式清理资源
export KOOK_TEST_MSG_ID="xxx"
export KOOK_TEST_MODE=dangerous
go run examples/message/main.go
```

## 注意事项

1. **测试环境**：建议在测试服务器中运行示例，避免影响生产环境
2. **权限要求**：Bot 需要有足够的权限执行相关操作
3. **清理资源**：使用 `dangerous` 模式清理 `write` 模式创建的资源
4. **频率限制**：示例已内置速率限制处理，无需额外处理

## 更多文档

- [快速开始](../docs/getting-started.md)
- [API 参考](../docs/api-reference.md)
- [WebSocket 事件](../docs/websocket.md)
- [卡片消息](../docs/card-message.md)
