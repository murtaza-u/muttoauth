# `muttoauth2`

> Google `OAuth2` authorization script for Mutt E-mail client

## Prerequisite

* Create a [Google Cloud project](https://cloud.google.com/resource-manager/docs/creating-managing-projects)
* Configure [OAuth consent screen](https://developers.google.com/workspace/guides/configure-oauth-consent)
* Enable [Gmail API](https://console.cloud.google.com/apis/library/gmail.googleapis.com)
* [Create](https://developers.google.com/workspace/guides/create-credentials#oauth-client-id)
  OAuth `client id` and `client secret`

## Edit `muttoauth`

```bash
# client id
id="*********************************************.apps.googleusercontent.com"

# client secret
secret="***********************************"

# gpg key id(gpg --list-secret-keys --keyid-format LONG)
keyid="****************"
```

* Set `id` equal to client id
* Set `secret` equal to client secret
* Get `gpg` key id

```bash
gpg --list-secret-keys --keyid-format LONG

sec   rsa4096/{keyid} 2022-06-06 [SC]
```

* Set `keyid` equal to `gpg` key id

## Check for missing dependencies

```bash
muttoauth -d

[x] curl
[ ] jq
[x] gpg
[x] shred
```

## Authorize

```bash
muttoauth -a /save/to/tkn_file

Open this URL in a web browser
Copy the authorization code and paste it below
https://accounts.google.com/o/oauth2/auth?client_id=*********************************************.apps.googleusercontent.com&redirect_uri=urn:ietf:wg:oauth:2.0:oob&scope=https://mail.google.com/&response_type=code
Authorization Code: *********************************************
```

## Configure Mutt to use `OAuth2`

```diff
 set imap_user = "test@gmail.com"
 set smtp_url = "smtps://test@gmail.com@smtp.gmail.com:465"

-set my_pass = "`pass show email`"
-set imap_pass = $my_pass
-set smtp_pass = $my_pass
+set imap_authenticators="oauthbearer"
+set imap_oauth_refresh_command = "muttoauth -r /path/to/tkn_file"
+set smtp_authenticators=${imap_authenticators}
+set smtp_oauth_refresh_command=${imap_oauth_refresh_command}
```
