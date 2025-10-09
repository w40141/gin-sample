package router

import (
	"net/http"
	"time"
)

// Config はアプリケーションの設定を保持します。
type Config struct {
	loc           *time.Location
	timeFormatter string
	addr          string
}

// Loc はタイムゾーンのロケーションを返します。
func (cfg Config) Loc() *time.Location {
	return cfg.loc
}

// TimeFormatter は時間のフォーマット文字列を返します。
func (cfg Config) TimeFormatter() string {
	return cfg.timeFormatter
}

// Addr はサーバーのアドレスを返します。
func (cfg Config) Addr() string {
	return cfg.addr
}

// LoadConfig はアプリケーションの設定を読み込みます。
func LoadConfig() (Config, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return Config{}, err
	}

	timeFormatter := time.RFC3339

	addr := ":8081"

	return Config{
		loc:           loc,
		timeFormatter: timeFormatter,
		addr:          addr,
	}, nil
}

// Start はサーバーを起動します。
func (cfg Config) Start(r http.Handler) error {
	return http.ListenAndServe(cfg.addr, r)
}
