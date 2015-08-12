package audio

import (
    "log"
    "os"
    "golang.org/x/mobile/exp/audio"
)

// Singleton
var instantiated *Player = nil

func PlayerPtr() *Player {
    if instantiated == nil {
        instantiated = new(Player);
    }
    return instantiated;
}

type Player struct {
    player *audio.Player
    file *os.File
}

func (p *Player) File(filename string) (*Player, error) {

    if p.player != nil {
        p.player.Close()
    }

    file, err := os.Open(filename)
    if err != nil {
        return p, err
    }

    p.file = file

    player, err := audio.NewPlayer(file, 0, 0)
    if err != nil {
        log.Fatal(err)
    }
    p.player = player

    return p, nil
}

func (p *Player) Play() {

    if p.player == nil {
        return
    }

    p.player.Play()
}

func (p *Player) Stop(){
    if p.player == nil {
        return
    }
    p.player.Stop()
}

func (p *Player) Close() {
    if p.player == nil {
        return
    }

    p.player.Close()
}