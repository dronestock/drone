# drone

`Drone`插件基础框架，提供如下功能

- 基础配置
    - 重试
    - 背压
    - 配置解析
    - 命名
- 接口抽象
    - 步骤
    - 配置
    - 插件

## 编写插件

使用`drone`库编写插件非常简单，只需要两步

### 插件主文件`plugin.go`

用于描述插件的主文件，主要是实现`drone.Plugin`接口，大致代码如下

```go
package main

import (
    `github.com/dronestock/drone`
)

type plugin struct {
    drone.PluginBase

    // 远程仓库地址
    Remote string `default:"${PLUGIN_REMOTE=${REMOTE=${DRONE_GIT_HTTP_URL}}}" validate:"required"`
    // 模式
    Mode string `default:"${PLUGIN_MODE=${MODE=push}}"`
    // SSH密钥
    SSHKey string `default:"${PLUGIN_SSH_KEY=${SSH_KEY}}"`
    // 目录
    Folder string `default:"${PLUGIN_FOLDER=${FOLDER=.}}" validate:"required"`
    // 目录列表
    Folders []string `default:"${PLUGIN_FOLDERS=${FOLDERS}}"`
    // 分支
    Branch string `default:"${PLUGIN_BRANCH=${BRANCH=master}}" validate:"required_without=Commit"`
    // 标签
    Tag string `default:"${PLUGIN_TAG=${TAG}}"`
    // 作者
    Author string `default:"${PLUGIN_AUTHOR=${AUTHOR=${DRONE_COMMIT_AUTHOR}}}"`
    // 邮箱
    Email string `default:"${PLUGIN_EMAIL=${EMAIL=${DRONE_COMMIT_AUTHOR_EMAIL}}}"`
    // 提交消息
    Message string `default:"${PLUGIN_MESSAGE=${MESSAGE=${PLUGIN_COMMIT_MESSAGE=drone}}}"`
    // 是否强制提交
    Force bool `default:"${PLUGIN_FORCE=${FORCE=true}}"`

    // 子模块
    Submodules bool `default:"${PLUGIN_SUBMODULES=${SUBMODULES=true}}"`
    // 深度
    Depth int `default:"${PLUGIN_DEPTH=${DEPTH=50}}"`
    // 提交
    Commit string `default:"${PLUGIN_COMMIT=${COMMIT=${DRONE_COMMIT}}}" validate:"required_without=Branch"`

    // 是否清理
    Clear bool `default:"${PLUGIN_CLEAR=${CLEAR=true}}"`
}

func newPlugin() drone.Plugin {
    return new(plugin)
}

func (p *plugin) Config() drone.Config {
    return p
}

func (p *plugin) Steps() []*drone.Step {
    return []*drone.Step{
        drone.NewStep(p.github, drone.Name(`Github加速`)),
        drone.NewStep(p.clear, drone.Name(`清理Git目录`)),
        drone.NewStep(p.ssh, drone.Name(`写入SSH配置`)),
        drone.NewStep(p.pull, drone.Name(`拉代码`)),
        drone.NewStep(p.push, drone.Name(`推代码`)),
    }
}

// 业务逻辑代码
func (p *plugin) github() (undo bool, err error) {
    return
}

// 业务逻辑代码
func (p *plugin) clear() (undo bool, err error) {
    return
}

// 业务逻辑代码
func (p *plugin) ssh() (undo bool, err error) {
    return
}

// 业务逻辑代码
func (p *plugin) pull() (undo bool, err error) {
    return
}

// 业务逻辑代码
func (p *plugin) push() (undo bool, err error) {
    return
}

```

其中，需要实现的方法步骤有

- `github`
- `clear`
- `ssh`
- `pull`
- `push`

### 启动文件`main.go`

启动文件负责启动整个程序，是一个非常轻量的壳

```go
package main

import (
	`github.com/dronestock/drone`
)

func main() {
	panic(drone.Bootstrap(newPlugin, drone.Configs(`FOLDERS`)))
}
```
