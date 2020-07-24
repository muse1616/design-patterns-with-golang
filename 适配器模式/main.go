package main

func main() {
	var mediaPlayer MediaPlayer
	mediaPlayer = new(Mp3Player)
	mediaPlayer.play("mp3")
	mediaPlayer.play("mp4")
	mediaPlayer.play("vlc")
}
