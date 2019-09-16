package Adaptor

import "testing"

func TestAdaptor(t *testing.T){
	player := PlayerAdaptor{}
	player.play("mp3","告白气球")
	player.play("wma","七里香")
	player.play("mp4","说好不哭")
}