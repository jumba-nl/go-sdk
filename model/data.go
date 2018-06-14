package model

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	. "github.com/jumba-nl/jumba/ext/config"
	"github.com/jumba-nl/jumba/util"
)

type Data struct {
	Profile              *Profile
	ClientSocialSettings *ClientSocialSettings
	Statistics           map[string]interface{}
	Payload              interface{}
	Request              interface{}
	Total                int64 // deprecated
	Config               *Configuration
	NotFound             bool     // deprecated
	TotalForSale         int64    // deprecated
	CDN                  string   // deprecated
	CSS                  []string // deprecated
	JS                   []string // deprecated
	SavedPath            string
	Canonical            string
	Cookies              Cookies
	Meta                 DataMeta
	News                 []NewsHeader
}

type DataMeta struct {
	// defines the page title
	Headline string
	// defines the meta title / title (title is only set when headline is empty)
	Title string
	// defines the meta description / description
	Description string
	// defines the Open graph type
	Type string
	// defines the Open graph Image
	Image string
}

type Cookies struct {
	ShowNotice    bool
	AllowTracking bool
}

type ClientSocialSettings struct {
	Social struct {
		Facebook struct {
			Enabled bool
		}
		Twitter struct {
			Enabled bool
		}
		Linkedin struct {
			Enabled bool
		}
		Google struct {
			Enabled bool
		}
	}
}

func NewData(w http.ResponseWriter, r *http.Request) (d Data) {
	d = Data{}
	LoadDefaultData(w, r, &d)
	return
}

func DefaultTemplateData(w http.ResponseWriter, r *http.Request) (d *Data) {
	d = new(Data)
	LoadDefaultData(w, r, d)
	return
}

func Onboarding(w http.ResponseWriter, r *http.Request, data *Data) (redirected bool) {
	if data.Profile != nil && data.Profile.JumbaFtu && !strings.Contains(r.URL.Path, "/mijn/onboarding") {
		http.Redirect(w, r, "/mijn/onboarding", http.StatusTemporaryRedirect)
		return true
	}

	return false
}

func LoadDefaultData(w http.ResponseWriter, r *http.Request, data *Data) {
	data.Config = Config
	data.CDN = Config.CDN
	data.News = GetLatestNewsHeaders()

	exists := false
	qs := r.URL.Query().Encode()
	path := r.URL.Path

	for _, cookie := range r.Cookies() {
		switch cookie.Name {
		case "jmb_ac":
			exists = true
			data.Cookies.AllowTracking, _ = strconv.ParseBool(cookie.Value)
		case "jmb_cpth":
			path = cookie.Value
		}
	}

	// show cookiebanner when:
	// 1. the banner hasn't been accepted/denied before.
	// 2. it current URL isn't /privacy.
	// 3. the current user-agent isn't a crawler.
	if !exists && r.URL.Path != "/cookies" && !util.IsAgentCrawler(r.Header.Get("User-Agent")) {
		data.Cookies.ShowNotice = true
	}

	saveURL := r.URL.Query().Get("jmb_surl")
	if saveURL != "" {
		parsedURL, errURL := url.Parse(saveURL)
		if errURL == nil {
			path = parsedURL.String()

			savePath(w, parsedURL.String())
		}
	} else if !isURLBlacklisted(r.URL.Path) {
		if (len(r.URL.Path) > 6 && r.URL.Path[:6] == "/order") || (len(r.URL.Path) > 7 && r.URL.Path[:7] == "/jumba/") {
			// do nothing as it seems that len(r.URL.path) breaks every one like even /
		} else {
			qs = r.URL.Query().Encode()
			path = r.URL.Path
			if len(qs) > 0 {
				path += "?" + qs
			}

			savePath(w, path)
		}
	}

	// means we don't have a redirect cookie or if a user is redirect to a logged in page
	// if we don't do this user will end up in an endless loop
	if path == r.URL.Path || (len(path) >= 6 && path[:6] == "/mijn/" && r.URL.Path == "/logout") {
		data.SavedPath = "/"
	} else {
		data.SavedPath = path
	}

	saveReferrer(w, r)
}

var matchingURLBlacklist = []string{"/login", "/404", "/logout", "/registratie", "/registratie/verstuurd", "/payment/data", "/favicon.ico", "/img/favicon.ico"}
var containingURLBlacklist = []string{"/v1/", "/v2/", "/vx/"}

func savePath(w http.ResponseWriter, path string) {
	path = strings.Replace(path, "?lc=true", "?", -1)
	path = strings.Replace(path, "&lc=true", "", -1)
	if path[len(path)-1:] == "?" {
		path = path[:len(path)-1]
	}

	cookie := &http.Cookie{Name: "jmb_cpth", Value: path, Path: "/", Expires: time.Now().Add(time.Hour * 1)}
	http.SetCookie(w, cookie)
}

func saveReferrer(w http.ResponseWriter, r *http.Request) {
	referrer := r.Referer()
	if referrer != "" {
		if uri, err := url.Parse(r.Referer()); err == nil {
			http.SetCookie(w, &http.Cookie{Name: "jmb_referrer", Value: uri.Path, Path: "/"})
		}
	}
}

func isURLBlacklisted(url string) bool {
	contains := false

	for _, s := range matchingURLBlacklist {
		if url == s {
			contains = true
		}
	}

	for _, s := range containingURLBlacklist {
		if strings.Contains(url, s) {
			contains = true
		}
	}

	return contains
}
