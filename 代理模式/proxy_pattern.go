package main

import (
	"fmt"
	"reflect"
)

type Image interface {
	display()
}
type RealImage struct {
	fileName string
}

func NewRealImage(fileName string) *RealImage {
	return &RealImage{fileName: fileName}
}
func (ri *RealImage) display() {
	fmt.Println("displaying:", ri.fileName)
}

func (ri *RealImage) loadFromDisk() {
	fmt.Println("loading:", ri.fileName)
}

//代理图片
type ImageProxy struct {
	realImage RealImage
	fileName  string
}

func NewImageProxy(fileName string) *ImageProxy {
	return &ImageProxy{
		fileName: fileName,
	}
}
func (p *ImageProxy) display() {
	// 未加载的话需要磁盘中加载
	// 判断为空
	// 或者 p.realImage == (RealImage{})
	if reflect.DeepEqual(p.realImage, RealImage{}) {
		p.realImage = *NewRealImage(p.fileName)
		// 模拟加载
		p.realImage.loadFromDisk()
	}
	p.realImage.display()
}
