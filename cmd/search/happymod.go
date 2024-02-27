package cmd

import (
  x "mywabot/system"

  "fmt"
  "encoding/json"
  "net/http"
  "net/url"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(happymod|mod)",
    Cmd:    []string{"happymod"},
    Tags:   "search",
    Prefix: true,
    IsQuery: true,
    ValueQ: ".happymod free fire",
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")


    type Result struct {
      Creator string `json:"creator"`
      Result  []struct {
        Title   string  `json:"title"`
        Icon    string  `json:"icon"`
        Link    string  `json:"link"`
       
      } `json:"result"`
    }


    resp, err := http.Get("https://aemt.me/happymod?query="+url.QueryEscape(m.Query))
    if err != nil {
      fmt.Println("Error:", err)
      return
    }
    defer resp.Body.Close()

    var result Result
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
      fmt.Println("Error:", err)
      return
    }

      teks := `*HAPPY MOD APK*`
    for _, mod := range result.Result {  
      teks += fmt.Sprintf(`
      
𖦹 *Title:* %s
𖦹 *Icon:* %s
𖦹 *Link:* %s
                          
`, mod.Title, mod.Icon, mod.Link)
    }
      
      m.Reply(teks)
      m.React("✅")
    },
  })
}
