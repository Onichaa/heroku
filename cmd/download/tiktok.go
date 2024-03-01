package cmd

import (
  x "mywabot/system"
    "fmt"
    "net/http"
    "net/url"
   // "time"
     "encoding/json"
    "io/ioutil"
    //"strconv"
    "strings"
    //"os"
    "regexp"
)

func init() {
  x.NewCmd(&x.ICmd{
    Name:   "(tt|ttnowm|tiktoknowm|tiktok)",
    Cmd:    []string{"tiktok"},
    Tags:   "download",
    Prefix: true,
    IsQuery: true,
    ValueQ: ".tt https://vt.tiktok.com/ZSFk78cDv/",
    Exec: func(sock *x.Nc, m *x.IMsg) {
       m.React("⏱️")

      if !strings.Contains(m.Query, "tiktok") {
        m.Reply("Itu bukan link tiktok")
      return
      }  

      regex := regexp.MustCompile(`(https?:\/\/[^\s]+)`)
       newLink := regex.FindStringSubmatch(m.Query) 

      

type Result struct {
	Status    bool   `json:"status"`
	Code      int    `json:"code"`
	Creator   string `json:"creator"`
	Result    struct {
		Author      string `json:"author"`
		Status      string `json:"status"`
		InfoVideo   struct {
			Title       string `json:"title"`
			Region      string `json:"region"`
			Thumbnail   string `json:"thumbnail"`
			Duration    int    `json:"duration"`
			TotalDownload int `json:"total_download"`
			TotalPlay   int `json:"total_play"`
			TotalShare  int `json:"total_share"`
			TotalComment int `json:"total_comment"`
		} `json:"info_video"`
		AuthorInfo struct {
			Nickname string `json:"nickname"`
			ID       string `json:"id"`
			Profile  string `json:"profile"`
		} `json:"author_info"`
		URL struct {
			Nowm  string `json:"nowm"`
			Wm    string `json:"wm"`
			Audio string `json:"audio"`
		} `json:"url"`
	} `json:"result"`
}


	url := "https://aemt.me/download/tikdl?url="+url.QueryEscape(newLink[0])
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var result Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

      sock.SendVideo(m.From, result.Result.URL.Nowm, "", *m)
	fmt.Println("Result:", result.Result.URL.Nowm)




      /*
      type TikTokData struct {
      Creator       string `json:"creator"`
      Code          int    `json:"code"`
      Msg           string `json:"msg"`
      ProcessedTime float64 `json:"processed_time"`
      Data          struct {
        ID              string `json:"id"`
        Region          string `json:"region"`
        Title           string `json:"title"`
        Cover           string `json:"cover"`
        OriginCover     string `json:"origin_cover"`
        Duration        int    `json:"duration"`
        Play            string `json:"play"`
        WmPlay          string `json:"wmplay"`
        HdPlay          string `json:"hdplay"`
        Size            int    `json:"size"`
        WmSize          int    `json:"wm_size"`
        HdSize          int    `json:"hd_size"`
        Music           string `json:"music"`
        MusicInfo       struct {
          ID       string `json:"id"`
          Title    string `json:"title"`
          Play     string `json:"play"`
          Cover    string `json:"cover"`
          Author   string `json:"author"`
          Original bool   `json:"original"`
          Duration int    `json:"duration"`
          Album    string `json:"album"`
        } `json:"music_info"`
        PlayCount     int `json:"play_count"`
        DiggCount     int `json:"digg_count"`
        CommentCount  int `json:"comment_count"`
        ShareCount    int `json:"share_count"`
        DownloadCount int `json:"download_count"`
        CollectCount  int `json:"collect_count"`
        CreateTime    int `json:"create_time"`

        Author              struct {
          ID        string `json:"id"`
          UniqueID  string `json:"unique_id"`
          Nickname  string `json:"nickname"`
          Avatar    string `json:"avatar"`
        } `json:"author"`
        Images    []string `json:"images"`
      } `json:"data"`
      }

     

      url := "https://skizo.tech/api/tiktok?url="+url.QueryEscape(newLink[0])+"&apikey=batu"//+os.Getenv("KEY")

      response, err := http.Get(url)
      if err != nil {
      fmt.Println("Error:", err)
      return
      }
      defer response.Body.Close()


      body, err := ioutil.ReadAll(response.Body)
      if err != nil {
      fmt.Println("Error:", err)
      return
      }


      var tiktokData TikTokData
      err = json.Unmarshal(body, &tiktokData)
      if err != nil {
      fmt.Println("Error:", err)
      return
      }


      if tiktokData.Data.Duration == 0 {
      for _, i := range tiktokData.Data.Images {
        x.Sleep(2 * time.Second)

        bytes, err := x.ToByte(i)
        if err != nil {
          m.Reply(err.Error())
          return
        }
        
        sock.SendImage(m.From, bytes, "", *m) 
      }

      } else { 

        teks := `*TIKTOK NO WATERMARK*

𖦹 *ID:* ` + tiktokData.Data.ID + `
𖦹 *Author:* ` + tiktokData.Data.Author.UniqueID + `
𖦹 *Region:* ` + tiktokData.Data.Region + `
𖦹 *Judul:* ` + tiktokData.Data.Title + `
𖦹 *Durasi:* ` + strconv.Itoa(tiktokData.Data.Duration) + `
𖦹 *Music:* ` + tiktokData.Data.Music + `
𖦹 *Info Musik:*
  - *Judul:* ` + tiktokData.Data.MusicInfo.Title + `
  - *Author:* ` + tiktokData.Data.MusicInfo.Author + `
𖦹 *Jumlah Komentar:* ` + strconv.Itoa(tiktokData.Data.CommentCount) + `
𖦹 *Jumlah Share:* ` + strconv.Itoa(tiktokData.Data.ShareCount) + `
𖦹 *Didownload:* ` + strconv.Itoa(tiktokData.Data.DownloadCount) + ` kali`

      bytes, err := x.ToByte(tiktokData.Data.Play)
      if err != nil {
      m.Reply(err.Error())
      return
      }
      sock.SendVideo(m.From, bytes, teks, *m)
      }
      */
      
      
      m.React("✅")
    },
  })
}
