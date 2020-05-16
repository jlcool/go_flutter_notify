# GO_FLUTTER_NOTIFY

go-flutter 的通知插件，依照我的另一个flutter版插件写的 [flutter_plugin_notify](https://github.com/jlcool/flutter_plugin_notify)，接口可以在这里看

根据 https://github.com/martinlindhe/notify 编写

windows 中文乱码问题参考 https://github.com/go-toast/toast/issues/18 解决
## Usage

#### 1. 安装 https://github.com/jlcool/flutter_plugin_notify

#### 2. 在cmd/options.go中添加

Import as:

```go
import "github.com/jlcool/go_flutter_notify"
```

```go
flutter.AddPlugin(&notifyFlutter.NotifyFlutterPlugin{}),
```
