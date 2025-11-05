package mapper

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/casdoor/casdoor-go-sdk/proto"
)

func ProtoUserResponseToUser(user_reponse *proto.UserResponse) *casdoorsdk.User {
	user := &casdoorsdk.User{
		Owner:       user_reponse.User.Owner,
		Name:        user_reponse.User.Name,
		CreatedTime: user_reponse.User.CreatedTime,
		UpdatedTime: user_reponse.User.UpdatedTime,

		Id:              user_reponse.User.Id,
		ExternalId:      user_reponse.User.ExternalId,
		Type:            user_reponse.User.Type,
		Password:        user_reponse.User.Password,
		PasswordSalt:    user_reponse.User.PasswordSalt,
		PasswordType:    user_reponse.User.PasswordType,
		DisplayName:     user_reponse.User.DisplayName,
		FirstName:       user_reponse.User.FirstName,
		LastName:        user_reponse.User.LastName,
		Avatar:          user_reponse.User.Avatar,
		AvatarType:      user_reponse.User.AvatarType,
		PermanentAvatar: user_reponse.User.PermanentAvatar,
		Email:           user_reponse.User.Email,
		EmailVerified:   user_reponse.User.EmailVerified,
		Phone:           user_reponse.User.Phone,
		CountryCode:     user_reponse.User.CountryCode,
		Region:          user_reponse.User.Region,
		Location:        user_reponse.User.Location,
		Address:         user_reponse.User.Address,
		Affiliation:     user_reponse.User.Affiliation,
		Title:           user_reponse.User.Title,
		IdCardType:      user_reponse.User.IdCardType,
		IdCard:          user_reponse.User.IdCard,
		Homepage:        user_reponse.User.Homepage,
		Bio:             user_reponse.User.Bio,
		Tag:             user_reponse.User.Tag,
		Language:        user_reponse.User.Language,
		Gender:          user_reponse.User.Gender,
		Birthday:        user_reponse.User.Birthday,
		Education:       user_reponse.User.Education,

		IsDefaultAvatar:   user_reponse.User.IsDefaultAvatar,
		IsOnline:          user_reponse.User.IsOnline,
		IsAdmin:           user_reponse.User.IsAdmin,
		IsForbidden:       user_reponse.User.IsForbidden,
		IsDeleted:         user_reponse.User.IsDeleted,
		SignupApplication: user_reponse.User.SignupApplication,
		Hash:              user_reponse.User.Hash,
		PreHash:           user_reponse.User.PreHash,
		AccessKey:         user_reponse.User.AccessKey,
		AccessSecret:      user_reponse.User.AccessSecret,

		CreatedIp:      user_reponse.User.CreatedIp,
		LastSigninTime: user_reponse.User.LastSigninTime,
		LastSigninIp:   user_reponse.User.LastSigninIp,
	}
	return user
}
