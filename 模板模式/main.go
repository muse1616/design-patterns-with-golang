package main

func main() {
	downloader1 := NewHTTPDownloader()
	downloader1.Download("http://example.com/abc.zip")

	downloader1 = NewFTPDownloader()
	downloader1.Download("ftp://example.com/abc.zip")

}
