package main

import "fmt"

// 下载器接口
type Downloader interface {
	Download(uri string)
}

// 下载保存接口 公共接口
type implement interface {
	download()
	save()
}

// 魔板 父类
type template struct {
	uri string
	// 持有接口
	// 这个的作用是让父类可以使用子类的方法
	implement
}

func NewTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

// 公共方法
func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

func (t *template) save() {
	fmt.Print("default save\n")
}

// 两个子类 持有父类引用
type HTTPDownloader struct {
	// 继承父类
	*template
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	// 让父类持有自己
	downloader.template = NewTemplate(downloader)
	return downloader
}
func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}
func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	downloader.template = NewTemplate(downloader)
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}
