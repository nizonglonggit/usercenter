package logic

import (
    "context"
    "github.com/nizonglonggit/usercenter/model"
    "time"

    "github.com/nizonglonggit/usercenter/internal/svc"
    "github.com/nizonglonggit/usercenter/user"

    "github.com/tal-tech/go-zero/core/logx"
)

type RegisterUserLogic struct {
    ctx    context.Context
    svcCtx *svc.ServiceContext
    logx.Logger
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
    return &RegisterUserLogic{
        ctx:    ctx,
        svcCtx: svcCtx,
        Logger: logx.WithContext(ctx),
    }
}

func (l *RegisterUserLogic) RegisterUser(in *user.RegisterUserReq) (*user.BaseResp, error) {
    var headUrl string
    if len(in.HeadPortraitUrl) == 0 {
        headUrl = "default"
    } else {
        headUrl = in.HeadPortraitUrl
    }

    _, err := l.svcCtx.Model.Insert(model.User{
        NickName:        in.NickName,
        PassWord:        in.PassWord,
        Gender:          in.Gender,
        Email:           in.Email,
        HeadPortraitUrl: headUrl,
        RegDate:         time.Now().Format("2006-01-02"),
        Status:          1,
    })
    if err != nil {
        return nil, err
    }

    return &user.BaseResp{
        Code: 0,
        Msg:  "ok",
    }, nil
}
