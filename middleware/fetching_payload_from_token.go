package middleware

import (
	"firebase.google.com/go/auth"
	"github.com/Kien-Lai/senme-common/constant"
	"github.com/Kien-Lai/senme-common/firebase"
	"github.com/Kien-Lai/senme-common/utils"
	"github.com/gin-gonic/gin"
	"regexp"
)

type UserContext struct {
	UserId      string
	FirebaseId  string
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	ProviderID  string `json:"providerId,omitempty"`
	Roles       []string
}

func FetchingPayloadFromToken(context *gin.Context) {
	if regexp.MustCompile("/actuator/*").MatchString(context.FullPath()) {
		context.Next()
		return
	}
	bearerToken := context.GetHeader(constant.HEADER_AUTHORIZATION)
	if utils.IsBlank(bearerToken) {
		context.AbortWithStatus(401)
		return
	}
	accessToken := bearerToken[len(constant.HEADER_BEARER_PREFIX):len(bearerToken)]
	firebaseToken, err := firebase.INSTANCE.VerifyIDToken(context, accessToken)
	if err != nil {
		context.AbortWithStatus(401)
		return
	}
	userContext := UserContext{}
	if firebaseToken != nil {
		roles, uid := fetchRolesAndUserId(firebaseToken)
		userContext.UserId = uid
		userContext.FirebaseId = firebaseToken.UID
		userContext.Roles = roles
	}

	context.Set(constant.CONTEXT_USER_KEY, &userContext)
	context.Next()
}

func fetchRolesAndUserId(firebaseToken *auth.Token) ([]string, string) {
	var roles []string
	var uid string
	for k, v := range firebaseToken.Claims {
		if k == constant.FIREBASE_ROLES {
			for _, role := range v.([]interface{}) {
				roles = append(roles, role.(string))
			}
		}
		if k == constant.USER_ID {
			uid = v.(string)
		}
	}
	return roles, uid
}

func GetUserContext(ctx *gin.Context) *UserContext {
	value, _ := ctx.Get(constant.CONTEXT_USER_KEY)
	return value.(*UserContext)
}
