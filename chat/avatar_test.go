package main

import (
	"io/ioutil" // 入出力を扱う
	"os"
	"path/filepath" // 物理ファイルを扱うためのもの?
	"testing"

	gomniauthtest "github.com/stretchr/gomniauth/test"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("", ErrNoAvatarURL)
	testChatUser := &chatUser{User: testUser}
	url, err := authAvatar.GetAvatarURL(testChatUser)
	if err != ErrNoAvatarURL {
		t.Error("値が存在しない場合、AuthAvatar.GetAvatarURLは" +
			"ErrNoAvatarURLを返すべきです。")
	}
	// 値をセットしていく
	testURL := "http://url-to-avatar/"
	testUser = &gomniauthtest.TestUser{}
	testChatUser.User = testUser
	testUser.On("AvatarURL").Return(testURL, nil)
	url, err = authAvatar.GetAvatarURL(testChatUser)
	if err != nil {
		t.Error("値が存在する場合、AuthAvatar.GetAvatarURLは" +
			"エラーを返すべきではありません。")
	} else {
		if url != testURL {
			t.Error("AuthAvatar.GetAvatarURLは正しいURLを返すべきです。")
		}
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := gravatarAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("GravatarAvatar.GetAvatarURLはエラーを返すべきではありません")
	}
	if url !=
		"//www.gravatar.com/avatar/abc" {
		t.Error("GravatarAvatar.GetAvatarURLが%sという誤った値を返しました。", url)
	}

}
func TestFileSystemAvatar(t *testing.T) {

	// テスト用のアバターのファイルを生成
	filename := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0777)
	defer func() { os.Remove(filename) }()

	var fileSystemAvatar FileSystemAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := fileSystemAvatar.GetAvatarURL(user)
	if err != nil {
		t.Errorf("FileSysyemAvatar.GetAvatarURLはエラーを返すべきではありません")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("FileSysyemAvatar.GetAvatarURLが%sという誤った値を返しました", url)
	}

}
