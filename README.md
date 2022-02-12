# login-service

login-service is a standalone minimalistic login server providing a (JWT)[https://jwt.io/] login for multiple login backends.

## Abstract

login-service provides a minimal endpoint for authentication. The login is
then performed against the providers and returned as Json Web Token.

## Supported Provider
The following providers (login backends) are supported.

- (OSIAM)[http://osiam.org/]
OSIAM is a secure identity management solution providing REST based services for authentication and authorization.
It implements the multplie OAuth2 flows, as well as SCIM for managing the user data.
- Simple (user/password pairs by configuration)

## Future Planed Features
- Support for 3-leged-Oauth2 flow (OSIAM, Google, Facebook login)
- Backend for checking agains .htaccess file
- Caddyserver middleware

## API

### GET /login

Returns a simple bootstrap styled login form.

The returned html follows the ui composition conventions from (lib-compose)[https://github.com/tarent/lib-compose],
so it can be embedded into an existing layout.

### POST /login

Does the login and returns the JWT. Depending on the content-type, and parameters a classical JSON-Rest or a redirect can be performed.

#### Parameters

| Parameter-Type    | Parameter                                        | Description                                               |          | 
| ------------------|--------------------------------------------------|-----------------------------------------------------------|----------|
| Http-Header       | Accept: text/html                                | Set the JWT-Token as Cookie 'jwt_token'.                  | default  |
| Http-Header       | Accept: application/jwt                          | Returns the JWT-Token within the body. No Cookie is set.  |          |
| Http-Header       | Content-Type: application/x-www-form-urlencoded  | Expect the credentials as form encoded parameters.        | default  |
| Http-Header       | Content-Type: application/json                   | Take the credentials from the provided json object.       |          |
| Post-Parameter    | username                                         | The username                                              |          |
| Post-Parameter    | password                                         | The password                                              |          |
| Config-Parameter  | success-url                                      | The url to redirect on success                            | (default /) |

#### Possible Return Codes

| Code | Meaning               | Description                |
|------| ----------------------|----------------------------|
| 200  | OK                    | Successfully authenticated |
| 403  | Forbidden             | The Credentials are wrong  |
| 400  | Bad Request           | Missing parameters         |
| 500  | Internal Server Error | Internal error, e.g. the login provider is not available or failed    |
| 303  | See Other             | Sets the JWT as a cookie, if the login succeeds and redirect to the urls provided in `redirectSuccess` or `redirectError` |

Hint: The status `401 Unauthorized` is not used as a return code to not conflict with an Http BasicAuth Authentication.

#### Example:
Default is to return the token as Content-Type application/jwt within the body.
```
curl -i --data "username=foo&password=bar" http://127.0.0.1:6789/login
HTTP/1.1 200 OK
Content-Type: application/jwt
Date: Fri, 11 Feb 2022 21:32:27 EST
Content-Length: 100
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJmb28ifQ.wlTUsPIX9A5guOq0TrsoFXWDdeX-gI1auvmIuiWq7VE
```

#### Example: Credentials as JSON
The Credentials also could be send as JSON encoded.
```
curl -i -H 'Content-Type: application/json'  --data '{"username": "foo", "password": "bar"}' http://127.0.0.1:6789/login
HTTP/1.1 200 OK
Content-Type: application/jwt
Date: Fri, 11 Feb 2022 21:34:22 EST
Content-Length: 100
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJmb28ifQ.wlTUsPIX9A5guOq0TrsoFXWDdeX-gI1auvmIuiWq7VE
```

#### Example: web based flow with 'Accept: text/html'
Sets the jwt token as cookie and redirects to a web page.
```
curl -i -H 'Accept: text/html' --data "username=foo&password=bar" http://127.0.0.1:6789/login
HTTP/1.1 303 See Other
Location: /
Set-Cookie: jwt_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJmb28ifQ.wlTUsPIX9A5guOq0TrsoFXWDdeX-gI1auvmIuiWq7VE; HttpOnly
```


## Provider

### Osiam
To start loginsrv against the default osiam configuration on the same machine, use the following example.
```
loginsrv --jwt-secret=jwtsecret --text-logging -backend 'provider=osiam,endpoint=http://localhost:8080,clientId=example-client,clientSecret=secret'
```

Then go to http://127.0.0.1:6789/login and login with `admin/password`.
