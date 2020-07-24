package main

import "fmt"

// MediaPlayer只能播放MP3 现在有一个AdvancedMediaPlayer可以播放mp4、vlc 通过适配器使得MediaPlayer播放mp4

type MediaPlayer interface {
	play(mediaType string)
}
type Mp3Player struct {
	mediaAdapter *MediaAdapter
}

func (m3p Mp3Player) play(mediaType string) {
	//内置原始功能
	if mediaType == "mp3" {
		fmt.Println("播放mp3")
		return
	}
	//	 适配扩展
	if mediaType == "mp4" {
		m3p.mediaAdapter = NewAdapter(Mp4Player{})
		m3p.mediaAdapter.advanceMediaPlayer.playMp4()
	} else if mediaType == "vlc" {
		m3p.mediaAdapter = NewAdapter(VlcPlayer{})
		m3p.mediaAdapter.advanceMediaPlayer.playVlc()
	}
}

// advancePlayer接口及实现类
type AdvancedMediaPlayer interface {
	playMp4()
	playVlc()
}
type Mp4Player struct {
}

func (m4p Mp4Player) playMp4() {
	fmt.Println("播放mp4")
}
func (m4p Mp4Player) playVlc() {
	fmt.Println("do nothing")
}

type VlcPlayer struct {
}

func (v VlcPlayer) playMp4() {
	fmt.Println("do nothing")
}
func (v VlcPlayer) playVlc() {
	fmt.Println("播放Vlc")
}

// 适配器
type MediaAdapter struct {
	advanceMediaPlayer AdvancedMediaPlayer
}

//Constructor
func NewAdapter(ad AdvancedMediaPlayer) *MediaAdapter {
	return &MediaAdapter{advanceMediaPlayer: ad}
}

