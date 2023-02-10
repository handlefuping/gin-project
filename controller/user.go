package controller

import (
	"gin-project/model"
	"gin-project/util"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strconv"
	"time"
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

// Upload 用户上传视频
func (receiver User) Upload(ctx *gin.Context)  {
	file, err := ctx.FormFile("video")
	if err != nil {
		util.FailJsonResp(ctx, 400, "请选择视频")
		return
	}
	// 创建文件夹

	// 生成唯一name
	name := util.RandStr() + file.Filename

	// 文件存储本地路径
	localUrl := filepath.Join(util.GetVideoStoreDir(), name )

	err = ctx.SaveUploadedFile(file, localUrl)
	if err != nil {
		util.FailJsonResp(ctx, 400, "视频存储错误：" + err.Error())
		return
	}
	//userId, _, _ := ctx.Request.BasicAuth()
	//userIDn, _ := strconv.Atoi(userId)
	video := &model.Video{
		OriginName: file.Filename,
		Size: file.Size,
		Name: name,
		UserID: 1,
		LocalUrl: localUrl,

	}
	err = model.CreateVideo(video)
	if err != nil {
		util.FailJsonResp(ctx,400, err.Error())
	} else {
		util.SuccessJsonResp(ctx, video, 1)
	}

}