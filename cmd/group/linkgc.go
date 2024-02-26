package cmd

import (
  x "mywabot/system"

  "fmt"

)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(linkgc|linkgroup)",
    Cmd:    []string{"linkgc"},
    Tags:   "group",
    Prefix: true,
    IsBotAdm: true,
    IsGroup: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      resp, err := sock.WA.GetGroupInviteLink(m.From, false)
      if err != nil {
        m.Reply("Failed to get the group link.")
      } else {
        m.Reply(fmt.Sprintf("Link group: %s", resp))
      }
      m.React("✅")
    },
  })
}
