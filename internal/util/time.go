// Package util はユーティリティ関数を提供する
package util

import "time"

var TimeLocation *time.Location

// Initialize はユーティリティ関数を初期化する
func Initialize() error {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	TimeLocation = loc

	return nil
}

func Now() time.Time {
	return time.Now().In(TimeLocation)
}
