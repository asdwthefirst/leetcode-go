/*
*

	@author: jiangxi
	@since: 2023/6/18
	@desc: //TODO

*
*/
package designpattern

import "fmt"

type Downloader interface {
	Download(uri string)
}

// template，包含了implement接口，又在行为上继承了Downloader接口。
// 而implement接口内的方法就是要到具体类延迟实现的。Downloader接口的实现就是在父代实现的。
type template struct {
	implement
	uri string
}

type implement interface {
	download()
	save()
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

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

type HTTPDownloader struct {
	*template //继承template实现的Download方法
}

// 重点，这里的操作就是模板方法的灵魂。
func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader) //这里很绕，downloader实现impl，就是延迟实现的方法接口。
	downloader.template = template
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
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}
