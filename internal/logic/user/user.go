package user

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"online-study-room/internal/dao"
	"online-study-room/internal/svc"
	"online-study-room/internal/types"
)

type User struct {
	logx.Logger
	ctx context.Context

	//room dao.StudyRoomModel
	//seat dao.StudySeatModel
	user dao.StudyUserModel
}

func NewLogic(ctx context.Context, svcCtx *svc.ServiceContext) User {
	return User{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		//room:   svcCtx.Room,
		//seat:   svcCtx.Seat,
		user: svcCtx.User,
	}
}

const (
	SIGNINCODE_OK = iota
	SIGNINCODE_NOUSER
	SIGNINCODE_PASSWORDERROR
	SIGNINCODE_HASUSER
)

func (c *User) Signin(userName string, password string) (code int, err error) {
	u, err := c.user.FindByUserName(userName)
	if err != nil {
		return 0, err
	}

	if u.ID == 0 {
		return SIGNINCODE_NOUSER, errors.New("用户名不存在")
	}

	if u.Password != password {
		return SIGNINCODE_PASSWORDERROR, errors.New("用户密码错误")
	}

	return SIGNINCODE_OK, nil
}

func (c *User) Register(userName string, password string, nikeName string) (code int, err error) {
	u, err := c.user.FindByUserName(userName)
	if err != nil {
		return 0, err
	}

	if u.ID != 0 {
		return SIGNINCODE_HASUSER, errors.New("用户名已经注册")
	}

	_, err = c.user.Insert(userName, password, nikeName)
	if err != nil {
		return 0, err
	}
	return SIGNINCODE_OK, nil
}

func (c *User) User(id int) (types.User, error) {
	u, err := c.user.Find(id)
	if err != nil {
		return types.User{}, err
	}

	return types.User{
		UserName: u.UserName,
		Password: u.Password,
		NikeName: u.NikeName,
	}, nil
}
