package urlcollection

import (
	"crypto/sha256"
	"fmt"
	"github.com/meedoed/url-shorterer/internal/config"
)

type URL struct {
	SourceURL string `json:"source_url"`
	ShortURL  string `json:"short_url"`
}

func (url *URL) GenURL() {
	conf := config.GetConfig()
	sha := sha256.Sum256([]byte(url.SourceURL))
	url.ShortURL = fmt.Sprintf("%s:%s/%x", conf.Storage.Host, conf.Storage.Port, sha[:5])
}
