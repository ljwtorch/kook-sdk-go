# 贡献指南

感谢你对 KOOK SDK for Go 项目的关注！我们欢迎所有形式的贡献。

## 目录

- [提交 Issue](#提交-issue)
- [提交 Pull Request](#提交-pull-request)
- [代码规范](#代码规范)
- [Commit 规范](#commit-规范)
- [分支策略](#分支策略)

---

## 提交 Issue

### Bug 报告

如果你发现了 Bug，请提交 Issue 并包含以下信息：

**标题格式：** `[Bug] 简短描述`

**内容模板：**

```markdown
## Bug 描述

简要描述 Bug 的表现。

## 复现步骤

1. 执行 '...'
2. 调用 '...'
3. 出现错误 '...'

## 期望行为

描述你期望的行为。

## 实际行为

描述实际发生的行为。

## 环境信息

- Go 版本：1.21.x
- SDK 版本：0.1.0-dev.1
- 操作系统：macOS / Linux / Windows

## 相关代码

```go
// 你的代码
```

## 错误日志

```
// 错误信息
```
```

### 功能建议

如果你有功能建议，请提交 Issue 并包含以下信息：

**标题格式：** `[Feature] 简短描述`

**内容模板：**

```markdown
## 功能描述

简要描述你希望添加的功能。

## 使用场景

描述这个功能的使用场景。

## 建议实现

如果有实现思路，请描述。

## 参考资料

如果有相关文档或链接，请提供。
```

### API 问题

如果你在使用 API 时遇到问题：

**标题格式：** `[API] 简短描述`

**内容模板：**

```markdown
## 问题描述

描述你遇到的问题。

## 调用代码

```go
// 你的代码
```

## 返回结果

```json
// API 返回
```

## 期望结果

描述你期望的返回结果。
```

---

## 提交 Pull Request

### PR 流程

1. **Fork 仓库** - 点击仓库右上角的 Fork 按钮

2. **克隆代码**
   ```bash
   git clone https://github.com/your-username/kook-sdk-go.git
   cd kook-sdk-go
   ```

3. **创建分支**
   ```bash
   git checkout -b feature/your-feature-name
   # 或
   git checkout -b fix/your-fix-name
   ```

4. **修改代码** - 进行你的修改

5. **运行测试**
   ```bash
   go test ./...
   ```

6. **提交代码**
   ```bash
   git add .
   git commit -m "feat: add your feature description"
   ```

7. **推送分支**
   ```bash
   git push origin feature/your-feature-name
   ```

8. **创建 PR** - 在 GitHub 上创建 Pull Request

### PR 标题格式

```
<type>(<scope>): <description>
```

**示例：**
- `feat(api): add batch message sending`
- `fix(gateway): handle connection timeout`
- `docs(readme): update installation guide`

### PR 内容模板

```markdown
## 变更描述

简要描述本次变更的内容。

## 变更类型

- [ ] 新功能 (feat)
- [ ] Bug 修复 (fix)
- [ ] 文档更新 (docs)
- [ ] 代码重构 (refactor)
- [ ] 测试 (test)
- [ ] 其他 (chore)

## 关联 Issue

Closes #123

## 测试

描述你如何测试你的修改。

## 截图（如果适用）

添加相关截图。

## 检查清单

- [ ] 代码遵循项目规范
- [ ] 已添加必要的测试
- [ ] 已更新相关文档
- [ ] 所有测试通过
```

---

## 代码规范

### 命名规范

- **包名：** 全小写，简短，无下划线（如 `model`, `event`, `card`）
- **接口：** 以 `-er` 结尾或描述行为的名词（如 `EventHandler`, `MessageSender`）
- **导出函数/方法：** PascalCase
- **非导出函数/方法：** camelCase
- **常量：** PascalCase（如 `SignalEvent`, `MessageTypeText`）
- **文件名：** snake_case（如 `direct_message.go`, `guild_role.go`）

### API 方法命名

- **List 操作：** `List{Resource}` (如 `ListGuilds`, `ListChannels`)
- **获取详情：** `Get{Resource}` (如 `GetGuild`, `GetChannel`)
- **创建：** `Create{Resource}` (如 `CreateChannel`, `CreateMessage`)
- **更新：** `Update{Resource}` (如 `UpdateChannel`, `UpdateMessage`)
- **删除：** `Delete{Resource}` (如 `DeleteChannel`, `DeleteMessage`)
- **特殊操作：** 用动词+名词 (如 `AddReaction`, `PinMessage`, `KickoutMember`)

### 代码风格

- 每个导出的类型、函数、常量必须有 GoDoc 注释
- 错误处理遵循 Go 惯例：返回 `error`，不使用 panic
- 使用 `context.Context` 作为所有 API 方法的第一个参数
- 使用 Functional Options 模式配置 Client
- 所有 HTTP API 方法返回 `(*Response, error)`

### 注释规范

```go
// GetGuild 获取指定服务器的详细信息。
// guildID 为服务器的唯一标识。
//
// 使用示例:
//
//	guild, err := client.GetGuild(ctx, "guild-id")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(guild.Name)
func (c *Client) GetGuild(ctx context.Context, guildID string) (*model.Guild, error) {
	// ...
}
```

---

## Commit 规范

使用 [Conventional Commits](https://www.conventionalcommits.org/) 规范：

### 格式

```
<type>(<scope>): <description>

[optional body]

[optional footer]
```

### Type 类型

| 类型 | 说明 |
|------|------|
| `feat` | 新功能 |
| `fix` | Bug 修复 |
| `docs` | 文档更新 |
| `style` | 代码格式（不影响功能） |
| `refactor` | 代码重构 |
| `perf` | 性能优化 |
| `test` | 测试相关 |
| `chore` | 构建/工具相关 |

### Scope 范围

| 范围 | 说明 |
|------|------|
| `client` | Client 相关 |
| `api` | API 模块 |
| `ws` | WebSocket 相关 |
| `model` | 数据模型 |
| `event` | 事件定义 |
| `card` | 卡片消息 |
| `docs` | 文档 |

### 示例

```
feat(api): add batch message sending
fix(gateway): handle connection timeout
docs(readme): update installation guide
refactor(client): extract HTTP methods
test(api): add unit tests for guild API
chore(deps): update dependencies
```

### Breaking Changes

如果有破坏性变更，在 footer 中添加：

```
feat(api): change message response format

BREAKING CHANGE: CreateMessage now returns MessageResponse instead of Message
```

---

## 分支策略

- `main` - 稳定发布分支
- `develop` - 开发分支
- `feature/*` - 功能分支
- `fix/*` - 修复分支

### 分支命名

```
feature/add-batch-message
fix/gateway-timeout
docs/update-readme
```

---

## 开发环境

### 环境要求

- Go 1.21+
- Git

### 克隆项目

```bash
git clone https://github.com/ljwtorch/kook-sdk-go.git
cd kook-sdk-go
```

### 运行测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./api/...
go test ./card/...

# 运行并显示覆盖率
go test -cover ./...
```

### 构建

```bash
go build ./...
```

---

## 发布流程

1. 更新 `kook.go` 中的版本号
2. 更新 CHANGELOG.md
3. 创建 Git Tag
4. 推送到 GitHub

```bash
git tag v0.1.0-dev.2
git push origin v0.1.0-dev.2
```

---

## 联系方式

如有任何问题，请通过以下方式联系：

- [GitHub Issues](https://github.com/ljwtorch/kook-sdk-go/issues)
- [Pull Requests](https://github.com/ljwtorch/kook-sdk-go/pulls)

---

感谢你的贡献！
