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

var refreshCmd = &Z.Cmd{
	Name:     `refresh`,
	Summary:  `Refreshes access token and outputs it to stdout`,
	Usage:    `<file>`,
	NumArgs:  1,
	Comp:     compfile.New(),
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(caller *Z.Cmd, args ...string) error {
		refTkn, err := os.ReadFile(args[0])
		if err != nil {
			return err
		}

		params := make(url.Values, 5)
		params.Set("client_id", ClientID)
		params.Set("client_secret", ClientSecret)
		params.Set("refresh_token", string(refTkn))
		params.Set("grant_type", "refresh_token")

		resp, err := http.PostForm(tknURI, params)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		tkn := new(Token)
		err = json.Unmarshal(body, tkn)
		if err != nil {
			return err
		}

		fmt.Println(tkn.Access)
		return nil
	},
}
