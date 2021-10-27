package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	// logger = log.New(os.Stdout, "", log.LstdFlags)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	viper.SetConfigName("revproxy")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.revproxy")
	viper.ReadInConfig()
	log.Printf("config file used: %s\n", viper.ConfigFileUsed())
}

func main() {
	baseUrl := MustGetPropAsStringFromYaml("target.base-url")
	port := MustGetPropAsStringFromYaml("server.port")
	prefixToTrim := viper.GetString("target.prefix-to-trim")
	remote, err := url.Parse(baseUrl)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy, prefixToTrim))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}

}

func handler(p *httputil.ReverseProxy, prefixToTrim string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("X-Ben", "Rad")
		if prefixToTrim != "" {
			r.URL.Path = strings.TrimPrefix(r.URL.Path, prefixToTrim)
		}
		log.Print(r.URL.Path)
		p.ServeHTTP(w, r)
	}
}
