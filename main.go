package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/toriato/katia"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const Context string = "db"

var ErrUnsupportedSource = errors.New("지원하지 않는 데이터베이스 종류입니다")

var Plugin = katia.Plugin{
	Name:        "katia-plugin-database",
	Description: "GORM 라이브러리를 통해 데이터베이스를 제공합니다",
	Author:      "Sangha Lee <totoriato@gmail.com>",
	Version:     [3]int{0, 1, 0},

	OnEnable: func(bot *katia.Bot, plugin *katia.Plugin) error {
		// 설정 파일 불러오기
		var config Config
		{
			raw, err := ioutil.ReadFile(
				filepath.Join(plugin.Base(), "config.yaml"))
			if err != nil {
				return err
			}

			if err := yaml.Unmarshal(raw, &config); err != nil {
				return err
			}
		}

		var dialector interface{}

		switch config.Type {
		case PostgreSQL:
			dialector = postgres.Open(config.Source)
		default:
			return ErrUnsupportedSource
		}

		db, err := gorm.Open(dialector.(gorm.Dialector), &gorm.Config{})
		if err != nil {
			return err
		}

		bot.Set(Context, db)
		return nil
	},
}
