package main

/**
适配器持有旧接口，实现新接口
 */
type newPlayer interface {
	play(filename string, typ string)
}

type oldPlayer struct {
}

func (*oldPlayer) playMp3(filename string){
	println("play mp3:" + filename)
}

func (*oldPlayer) playMp4(filename string){
	println("play mp4:" + filename)
}

type playerAdaptor struct {
	old oldPlayer
}

func (p *playerAdaptor) play(filename string, typ string) {
	switch typ {
	case "mp3":
		p.old.playMp3(filename)
	case "mp4":
		p.old.playMp3(filename)
	default:
		println("err")

	}
}
