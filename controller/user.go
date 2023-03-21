package controller

import (
	"gin-project/model"
	"gin-project/util"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {

}



// GetUserById 查询用户
func (receiver User) GetUserById(ctx *gin.Context)  {
	id := ctx.Param("id")
	if id == "" {
		util.FailJsonResp(ctx, 400, "用户id不能为空")
		return
	}
	user, err := model.GetUserById(id)
	if err != nil {
		util.FailJsonResp(ctx, 400, err.Error())
	} else {
		util.SuccessJsonResp(ctx, user, 1)
	}
}

// RegisterUser 注册用户
func (receiver User) RegisterUser(ctx *gin.Context)  {
	var postData model.User
	if ctx.ShouldBind(&postData) != nil {

		util.FailJsonResp(ctx, 400, "用户名和密码不能为空")
		return
	}
	if user := model.GetUserByUserName(postData.UserName); user != nil {
		util.FailJsonResp(ctx, 400, "用户名已被注册")
		return
	}
	user, err := model.RegisterUser(postData.UserName, postData.Password)
	if err != nil {
		util.FailJsonResp(ctx, 400, err.Error())
	} else {
		util.SuccessJsonResp(ctx, user, 1)
	}
}

// LoginUser 用户登陆
func (receiver User) LoginUser(ctx *gin.Context)  {
	var postData model.User
	if ctx.ShouldBind(&postData) != nil {
		util.FailJsonResp(ctx, 400, "用户名和密码不能为空")
		return
	}

	user, err := model.LoginUser(postData.UserName, postData.Password)
	if err != nil {
		util.FailJsonResp(ctx, 400, err.Error())
	} else {

		tokenStr, err := util.GenerateTokenStr(postData.UserName, time.Minute * 5)
		if err != nil {
			util.FailJsonResp(ctx, 200, err.Error())
		} else {
			ctx.Request.SetBasicAuth(strconv.FormatUint(uint64(user.ID), 10), tokenStr)
			util.SuccessJsonResp(ctx, gin.H{
				"user": user,
				"token": ctx.GetHeader("Authorization"),
			}, 1)
		}

	}
}

// LogoutUser 退出登陆
func (receiver User) LogoutUser(ctx *gin.Context)  {
	ctx.Request.SetBasicAuth("", "")
	util.SuccessJsonResp(ctx, nil, 0)
}

// UploadMedia  用户上传视频
func (receiver User) UploadMedia(ctx *gin.Context)  {
	file, err := ctx.FormFile("media")
	if err != nil {
		util.FailJsonResp(ctx, 400, "请选择视频")
		return
	}
	// 创建文件夹

	// 生成唯一name
	name := util.RandStr() + file.Filename

	// 文件存储本地路径
	localUrl := filepath.Join(util.GetMediaStoreDir(), name )

	err = ctx.SaveUploadedFile(file, localUrl)
	if err != nil {
		util.FailJsonResp(ctx, 400, "视频存储错误：" + err.Error())
		return
	}
	
	media := &model.Media{
		OriginName: file.Filename,
		Size: file.Size,
		Name: name,
		UserID: 1,
		LocalUrl: localUrl,
		Status: 0,

	}
	err = model.CreateMedia(media)
	if err != nil {
		util.FailJsonResp(ctx,400, err.Error())
	} else {
		util.SuccessJsonResp(ctx, media, 1)
	}

}

func (receiver User) ChangeMediaStatus(ctx *gin.Context, status int)  {
	mediaId := ctx.Param("id")
	userId, _, _  := ctx.Request.BasicAuth()

	if mediaId == "" {
		util.FailJsonResp(ctx, 400, "视频id不能为空")
		return
	}
	if err := model.ProbMedia(mediaId, userId);  err != nil {
		util.FailJsonResp(ctx, 400, err.Error())
		return
	}
	if media, err := model.UpdateMedia(mediaId, status); err != nil {
		util.FailJsonResp(ctx, 400, err.Error())
	} else {
		util.SuccessJsonResp(ctx, media, 1)
	}
}

// PublicMedia 用户发布视频
func (receiver User) PublicMedia(ctx *gin.Context)  {

	receiver.ChangeMediaStatus(ctx, 1)
}

// BanMedia  用户下架视频
func (receiver User) BanMedia(ctx *gin.Context)  {

	receiver.ChangeMediaStatus(ctx, -1)
}