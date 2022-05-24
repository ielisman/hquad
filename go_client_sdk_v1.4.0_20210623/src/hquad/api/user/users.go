package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"hquad/api/login"
	"io/ioutil"
	"net/http"
	"time"
)

// https://mholt.github.io/json-to-go/
type UserResponse struct {
	User struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		Photo       string `json:"photo"`
		Phone       string `json:"phone"`
		UserID      string `json:"user_id"`
		UserGroupID struct {
			ID string `json:"id"`
		} `json:"user_group_id"`
		Disabled       string    `json:"disabled"`
		StartDatetime  time.Time `json:"start_datetime"`
		ExpiryDatetime time.Time `json:"expiry_datetime"`
		Permission     struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Operators   []struct {
				Name    string `json:"name"`
				OwnerID string `json:"owner_id"`
				UserID  string `json:"user_id"`
			} `json:"operators"`
		} `json:"permission"`
		AccessGroups struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"access_groups"`
		Password             string `json:"password"`
		Pin                  string `json:"pin"`
		LoginID              string `json:"login_id"`
		UserIP               string `json:"user_ip"`
		IdxUserID            string `json:"idx_user_id"`
		IdxUserIDNum         string `json:"idx_user_id_num"`
		IdxPhone             string `json:"idx_phone"`
		IdxEmail             string `json:"idx_email"`
		IdxLastModified      string `json:"idx_last_modified"`
		FingerprintTemplates []struct {
			Template0  string `json:"template0"`
			Template1  string `json:"template1"`
			FingerMask bool   `json:"finger_mask"`
			IsNew      bool   `json:"isNew"`
		} `json:"fingerprint_templates"`
		Credentials struct {
			Faces []struct {
				RawImage  string `json:"raw_image"`
				Templates []struct {
					Template string `json:"template"`
				} `json:"templates"`
				Flag       string `json:"flag"`
				UseProfile string `json:"useProfile"`
				Index      string `json:"index"`
			} `json:"faces"`
			VisualFaces []struct { // TODO: not in specification, added up, also template_ex_picture for photos
				RawImage string `json:"template_ex_normalized_image"`
				PicImage string `json:"template_ex_picture"`
				/*
					Templates []struct {
						Template string `json:"template"`
					} `json:"templates"` */
				// Flag       string `json:"flag"`
				// UseProfile string `json:"useProfile"`
				Index string `json:"index"`
			} `json:"visualFaces"`
		} `json:"credentials"`
		Cards []struct {
			CardType struct {
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				Mode string `json:"mode"`
			} `json:"card_type"`
			CardID          string `json:"card_id"`
			DisplayCardID   string `json:"display_card_id"`
			ID              string `json:"id"`
			CardId          string `json:"cardId"`
			WiegandFormatID struct {
				ID string `json:"id"`
			} `json:"wiegand_format_id"`
			WiegandFormatName string `json:"wiegand_format_name"`
		} `json:"cards"`
	} `json:"User"`
	Response struct {
		Code    string `json:"code"`
		Link    string `json:"link"`
		Message string `json:"message"`
	} `json:"Response"`
	HTTPResponseStatus int `json:"httpResponseStatus"`
}

var BS_SESSION_ID string

const USERS_URL = "http://192.168.1.220/api/users"

func SessionActive() error {
	sess_id, err := login.GetBSSessionId()
	if err != nil || sess_id == nil {
		resp, err := login.Login(&login.LoginRequest{Url: login.LOGIN_URL, Username: "admin", Password: "bc-ielisman-123"})
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to login %v", err))
		}
		sess_id = &resp.BS_SESSION_ID
	}
	if *sess_id == "" {
		return errors.New("cant retrieve session id")
	}
	BS_SESSION_ID = *sess_id
	return nil
}

func GetUser(id string) (*UserResponse, error) {
	if err := SessionActive(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", USERS_URL, id), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", login.CONTEXT_TYPE)
	req.Header.Set("bs-session-id", BS_SESSION_ID)
	resp, err := login.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	jsonUser := string(jsonData)
	fmt.Println(jsonUser)

	var userResponse UserResponse
	err = json.Unmarshal(jsonData, &userResponse)
	if err != nil {
		fmt.Println("Cant unmarshal", err)
		return nil, err
	}
	face := userResponse.User.Credentials.VisualFaces[0]
	var image string
	if face.RawImage != "" {
		image = face.RawImage
	} else {
		image = face.PicImage
	}

	fmt.Println("Image\n", image)
	fmt.Println("Index: ", face.Index)
	fmt.Println("Length: ", len(userResponse.User.Credentials.VisualFaces))
	if len(userResponse.User.Cards) > 0 {
		card := userResponse.User.Cards[0]
		fmt.Println("Card CardID", card.CardID)
		fmt.Println("Card DisplayCardID", card.DisplayCardID)
		fmt.Println("Card CardId", card.CardId)
		fmt.Println("Card WiegandFormatName", card.WiegandFormatName)
		fmt.Println("Card WiegandFormatID", card.WiegandFormatID.ID)
		fmt.Println("Card ID", card.ID)
		fmt.Printf("Card Type=%v ID=%v Name=%v Mode=%v\n", card.CardType.Type, card.CardType.ID, card.CardType.Name, card.CardType.Mode)
	}

	return &userResponse, nil
}

func GetUsers() {

}

func main() {
	GetUser("3")
}
