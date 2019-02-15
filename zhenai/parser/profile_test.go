package parser

import (
	"imooc.com/u2pppw/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "安静的雪")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+"element; but was %v", result.Items)
	}

	profileStringList := result.Items[0].(model.ProfileStringList)

	expected := model.ProfileStringList{
		Info: []string{"安静的雪", "未婚", "21岁", " 魔羯座(12.22-01.19)", "170cm", "工作地:阿坝阿坝县", "月收入:3千以下", "高中及以下"},
	}

	if len(profileStringList.Info) != len(expected.Info) {
		t.Errorf("expected %v: but was %v", expected, profileStringList)
	}
}
