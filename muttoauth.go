package muttoauth

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

const (
	port        = ":5000"
	redirectURI = "http://localhost:5000"
	scope       = "https://mail.google.com/"
	tknURI      = "https://accounts.google.com/o/oauth2/token"
)

var Cmd = &Z.Cmd{
	Name:     `muttoauth`,
	Summary:  `Google OAuth2 authorization script for Mutt E-mail client`,
	Commands: []*Z.Cmd{help.Cmd, authCmd, refreshCmd},
}
