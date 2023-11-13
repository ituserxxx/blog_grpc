package user

import (
	"brl/rpc/user/userclient"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"brl/http/internal/svc"
	"brl/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(ctx context.Context, req *types.LoginReq) (resp *types.LoginResp, err error) {
	res, err := l.svcCtx.UserRpc.UserLogin(ctx, &userclient.UserLoginReq{Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, err
	}

	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	now := time.Now().Unix()
	accessToken, err := l.GenToken(now, l.svcCtx.Config.JwtAuth.AccessSecret, map[string]interface{}{"uid": res.Id}, accessExpire)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Id:           res.Id,
		Username:     res.Username,
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
func (l *UserLoginLogic) GenToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
