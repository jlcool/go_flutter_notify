package notifyFlutter

import (
	"errors"

	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/martinlindhe/notify"
)

//  Make sure to use the same channel name as was used on the Flutter client side.
const channelName = "com.jlcool.notifyFlutter"
const (
	PARAM_APP_NAME  = "appName"
	PARAM_TITLE     = "title"
	PARAM_TEXT      = "text"
	PARAM_ICON_PATH = "iconPath"
)

type NotifyFlutterPlugin struct {
	channel *plugin.MethodChannel
}

var _ flutter.Plugin = &NotifyFlutterPlugin{} // compile-time type check

func (p *NotifyFlutterPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	p.channel = plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	p.channel.HandleFunc("notify", handleNotify)
	p.channel.HandleFunc("alert", handleAlert)
	return nil // no error
}

func handleNotify(arguments interface{}) (reply interface{}, err error) {
	var ok bool
	var args map[interface{}]interface{}
	if args, ok = arguments.(map[interface{}]interface{}); !ok {
		return nil, errors.New("invalid arguments")
	}
	var param_app_name string
	var param_title string
	var param_text string
	var param_icon_path string

	if appName, ok := args[PARAM_APP_NAME]; ok {
		param_app_name = appName.(string)
	}
	if title, ok := args[PARAM_TITLE]; ok {
		param_title = title.(string)
	}
	if text, ok := args[PARAM_TEXT]; ok {
		param_text = text.(string)
	}
	if iconPath, ok := args[PARAM_ICON_PATH]; ok {
		param_icon_path = iconPath.(string)
	}
	notify.Notify(param_app_name, param_title, param_text, param_icon_path)
	return nil, nil
}
func handleAlert(arguments interface{}) (reply interface{}, err error) {
	var ok bool
	var args map[interface{}]interface{}
	if args, ok = arguments.(map[interface{}]interface{}); !ok {
		return nil, errors.New("invalid arguments")
	}
	var param_app_name string
	var param_title string
	var param_text string
	var param_icon_path string

	if appName, ok := args[PARAM_APP_NAME]; ok {
		param_app_name = appName.(string)
	}
	if title, ok := args[PARAM_TITLE]; ok {
		param_title = title.(string)
	}
	if text, ok := args[PARAM_TEXT]; ok {
		param_text = text.(string)
	}
	if iconPath, ok := args[PARAM_ICON_PATH]; ok {
		param_icon_path = iconPath.(string)
	}
	notify.Alert(param_app_name, param_title, param_text, param_icon_path)
	return nil, nil
}
