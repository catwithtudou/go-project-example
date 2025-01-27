package setting

import (
	"github.com/fsnotify/fsnotify"
	_ "github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)
/**
 *@Author tudou
 *@Date 2020/7/26
 **/

type Setting struct{
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting,error){
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	for _,config:=range configs{
		if config!=""{
			vp.AddConfigPath(config)
		}
	}
	vp.SetConfigType("yaml")
	err:=vp.ReadInConfig()
	if err!=nil{
		return nil,err
	}
	s:=&Setting{vp}

	return s,nil
}


func (s *Setting)WatchSettingChange(){
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}


