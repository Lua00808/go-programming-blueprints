package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ErrNoAvatarURL はAvatarインスタンスがアバターのURLを返すことが出来ない
// 場合に発生するエラーです。
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

// Avatar はユーザーのプロフィール画像を表す型です。
type Avatar interface {
	// GetAvatarURLは指定されたクライアントのアバターのURLを返します。
	//問題が発生した場合にはエラーを返します。特に、URLを取得できなかった
	//場合にはErrNoAvatarURLを返します。
	GetAvatarURL(c *client) (string, error)
}
type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "//www.gravatar.com/avatar/" + useridStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			if files, err := ioutil.ReadDir("avatars"); err == nil {
				for _, file := range files {
					if file.IsDir() {
						continue // for 文の先頭に戻る(後続の処理をしない)
					}
					if match, _ := filepath.Match(useridStr+"*", file.Name()); match {
						return "/avatars/" + file.Name(), nil
					}
				}
			}
		}
	}
	return "", ErrNoAvatarURL
}

// リファクタリングと最適化から
