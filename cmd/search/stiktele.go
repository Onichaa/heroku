package cmd

import (
  x "mywabot/system"

  "fmt"
  "time"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(stiktele|stik)",
    Cmd:    []string{"stiktele"},
    Tags:   "search",
    Prefix: true,
    IsQuery: true,
    ValueQ: ".stiktele vernalta",
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      resp, err := x.Stiktele(m.Query)
        if err != nil {
          fmt.Println(err)
        }
      fmt.Println(resp)

      for _, sticker := range resp["sticker"].([]string) {
        time.Sleep(2 * time.Second)

        bytes, err := x.ToByte(sticker)
        if err != nil {
          m.Reply(err.Error())
          return
        }

        s := x.StickerApi(&x.Sticker{
          File: bytes,
          Tipe: func() x.MediaType {
          if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
            return x.IMAGE
          } else if m.IsVideo || m.IsQuotedVideo {
            return x.VIDEO
          } else {
            return x.TEKS
          }
          }(),
        }, &x.MetadataSticker{
          Author:    m.PushName,
          Pack:      "https://s.id/ryuubot",
          KeepScale: true,
          Removebg:  true,
          Circle: func() bool {
            if m.Query == "-c" {
              return true
            } else {
              return false
            }
          }(),
        })

        sock.SendSticker(m.From, s.Build(), *m)
        
        }

      m.React("✅")
    },
  })
}
