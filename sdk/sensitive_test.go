package sdk

import (
	"os"
	"testing"
)

func TestRongCloud_SensitiveAdd(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		Region{
			primaryDomain: os.Getenv("PRIMARY_DOMAIN"),
			backupDomain:  os.Getenv("BACKUP_DOMAIN"),
		},
	)
	err := rc.SensitiveAdd(
		"7Szq13MKRVortoknTAk7W8",
		"7Szq13MKRVortoknTAk7W8",
		1,
	)
	t.Log(err)
}

func TestRongCloud_SensitiveGetList(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	rep, err := rc.SensitiveGetList()
	t.Log(err)
	t.Log(rep)
}

func TestRongCloud_SensitiveRemove(t *testing.T) {
	rc := NewRongCloud(
		os.Getenv("APP_KEY"),
		os.Getenv("APP_SECRET"),
		REGION_BJ,
	)
	err := rc.SensitiveRemove(
		[]string{"7Szq13MKRVortoknTAk7W8"},
	)
	t.Log(err)
}
