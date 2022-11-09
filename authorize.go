package muttoauth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/compfile"
	"github.com/rwxrob/help"
)

var authCmd = &Z.Cmd{
	Name:     `authorize`,
	Summary:  `Requests refresh token and saves the to the specified path`,
	NumArgs:  1,
	Usage:    `<file>`,
	Comp:     compfile.New(),
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(caller *Z.Cmd, args ...string) error {
		DestFile = args[0]

		authURI := fmt.Sprintf(
			"https://accounts.google.com/o/oauth2/auth?client_id=%s&redirect_uri=%s&scope=%s&response_type=code",
			ClientID, redirectURI, scope,
		)
		fmt.Printf("Open this URL in a web browser:\n%s\n", authURI)

		http.HandleFunc("/", handler)
		return http.ListenAndServe(port, nil)
	},
}

var DestFile string

type Token struct {
	Access  string `json:"access_token"`
	Refresh string `json:"refresh_token"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(
		w, "<h1>Muttoauth | https://github.com/murtaza-u/muttoauth<h1>",
	)

	code := r.URL.Query().Get("code")
	if code == "" {
		return
	}

	params := make(url.Values, 5)
	params.Set("code", code)
	params.Set("client_id", ClientID)
	params.Set("client_secret", ClientSecret)
	params.Set("redirect_uri", redirectURI)
	params.Set("grant_type", "authorization_code")
	resp, err := http.PostForm(tknURI, params)
	if err != nil {
		fmt.Fprintf(w, "An error occured: %s\n", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	tkn := new(Token)
	err = json.Unmarshal(body, tkn)
	if err != nil {
		fmt.Fprintf(w, "An error occured: %s\n", err.Error())
		return
	}

	err = os.WriteFile(DestFile, []byte(tkn.Refresh), 0600)
	if err != nil {
		fmt.Fprintf(w, "An error occured: %s\n", err.Error())
		return
	}

	fmt.Printf(
		"\nRefresh token has been written to `%s`. Press Ctrl-C to exit.\n",
		DestFile,
	)
}
