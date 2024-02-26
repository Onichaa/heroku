package cmd

import (
  x "mywabot/system"

  "fmt"
"strings"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(opengc|opengroup|closegc|closegroup)",
    Cmd:    []string{"opengc", "closegc"},
    Tags:   "group",
    Prefix: true,
    IsAdmin: true,
    IsBotAdm: true,
    IsGroup: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

    err := sock.WA.SetGroupAnnounce(m.From, strings.Contains(m.Text, "close"))
    if err != nil {
      m.Reply("Error")
      fmt.Println(err.Error())
    }
      m.React("✅")
    },
  })
}
