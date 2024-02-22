package cmd

import (
  x "mywabot/system"

)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(donasi|donate)",
    Cmd:    []string{"donasi"},
    Tags:   "main",
    Prefix: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {

      teks := "Jangan lupa donasi ya kak se berapa persen pun kami terima 🐣"
      res := "https://telegra.ph//file/024a6a1fe28189278a5f7.jpg" 

     
      sock.SendImage(m.From, res, teks, *m)
      
    },
  })
}
