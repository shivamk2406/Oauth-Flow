package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	config = &oauth2.Config{
		ClientID:     "70082559410-n9vesipeijugsq9lu6u4auj3i5j2dv60.apps.googleusercontent.com",
		ClientSecret: "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",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	}
	config2 = &oauth2.Config{
		ClientID:     "70082559410-f14svtrpqskh9evtuku2seilipo0g0cf.apps.googleusercontent.com",
		ClientSecret: "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",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	}
	randomState = "random"
)

var (
	AccessTypeOnline oauth2.AuthCodeOption = oauth2.SetAuthURLParam("access_type", "online")
)

func Start() {

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	http.HandleFunc("/onLogin", handleonLogin)
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	var html = `<html><body><a href="\login">Google Login</a></body></html>`
	fmt.Fprint(w, html)

}

func handleLogin(w http.ResponseWriter, r *http.Request) {

	url := config.AuthCodeURL(randomState)
	url1 := config2.AuthCodeURL(randomState)
	fmt.Println(url1)
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
	if r.FormValue("state") != randomState {
		fmt.Println("State is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Println(r.FormValue("code"))
	fmt.Println(time.Now())
	token, err := config.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		fmt.Println("could not create token")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Println(token.Expiry)

	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Println("could not request token")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	//defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not request token")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	//http.Redirect(w, r, "/onLogin", http.StatusTemporaryRedirect)
	fmt.Fprintf(w, "Response %s", content)
}

func handleonLogin(w http.ResponseWriter, r *http.Request) {
	var html = `<html><body>Login Successful</body></html>`
	fmt.Fprint(w, html)
}
