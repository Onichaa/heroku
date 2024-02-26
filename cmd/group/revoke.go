package cmd

import (
  x "mywabot/system"

  "fmt"

)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(revoke|resetlink|resetlinkgc)",
    Cmd:    []string{"resetlinkgc"},
    Tags:   "group",
    Prefix: true,
    IsAdmin: true,
    IsBotAdm: true,
    IsGroup: true,
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

    resp, err := sock.WA.GetGroupInviteLink(m.From, true)
    if err != nil {
      m.Reply("Failed to reset the link group.")
    } else {
      m.Reply(fmt.Sprintf("New group link: %s", resp))
    }
      m.React("✅")
    },
  })
}
