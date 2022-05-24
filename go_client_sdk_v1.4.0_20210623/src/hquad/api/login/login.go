package login

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"net/http"
)

/* https://mholt.github.io/json-to-go/
const LOGIN_JSON string = `{
  "User": {
    "login_id": "string",
    "password": "string"
  }
}`
*/

type LoginJson struct {
	User struct {
		LoginID  string `json:"login_id"`
		Password string `json:"password"`
	} `json:"User"`
}

type LoginRequest struct {
	Url        string
	Username   string
	Password   string
	ForceLogin bool
}

type LoginResponse struct {
	BS_SESSION_ID string
}

const BS_SESSION_ID = "bs-session-id"

var LoginMap = make(map[string]string)
var CurrentUser string
var CurrentPassword string
var HttpClient *http.Client

const (
	LOGIN_URL    = "http://192.168.1.220/api/login" // "https://192.168.1.220/api/login" // will be set in web app as BIOSTAR2 SERVER URL
	CONTEXT_TYPE = "application/json"
)

func GetBSSessionId() (*string, error) {
	bs_session_id, ok := LoginMap[CurrentUser]
	if !ok {
		return nil, errors.New("must supply url/username/password")
	}
	if bs_session_id == "" {
		if CurrentUser == "" {
			return nil, errors.New("should supply url/username/password")
		}
		loginResp, err := Login(&LoginRequest{Url: LOGIN_URL, Username: CurrentUser, Password: CurrentPassword})
		if err != nil {
			return nil, errors.New("unable to connect with existing url/username/password")
		}
		return &loginResp.BS_SESSION_ID, nil
	}
	return &bs_session_id, nil
}

// Login returns bs-session-id given login url (i.e, http://192.168.1.220/api/login), contextType (application/json) and userCredentials (login/pwd)
func Login(loginRequest *LoginRequest) (*LoginResponse, error) {

	var loginJson LoginJson
	loginJson.User.LoginID = loginRequest.Username
	loginJson.User.Password = loginRequest.Password

	if !loginRequest.ForceLogin && LoginMap[loginJson.User.LoginID] != "" {
		return &LoginResponse{BS_SESSION_ID: LoginMap[loginJson.User.LoginID]}, nil
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	HttpClient = &http.Client{Transport: tr}

	json, err := json.Marshal(loginJson)
	if err != nil {
		return nil, err
	}
	resp, err := HttpClient.Post(loginRequest.Url, CONTEXT_TYPE, bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}
	bs_session_id := resp.Header.Get(BS_SESSION_ID)
	LoginMap[loginRequest.Username] = bs_session_id
	CurrentUser = loginRequest.Username
	CurrentPassword = loginRequest.Password

	return &LoginResponse{BS_SESSION_ID: bs_session_id}, nil
}
