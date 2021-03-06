/*
 * TechTrain MISSION Game API
 *
 * TechTrain MISSION ゲームAPI入門仕様  まずはこのAPI仕様に沿って機能を実装してみましょう。    この画面の各APIの「Try it out」->「Execute」を利用することで  ローカル環境で起動中のAPIにAPIリクエストをすることができます。
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"techtrain-mission/go/db"
	"techtrain-mission/go/helper"
	"techtrain-mission/go/model"
)

// UserCreatePost - ユーザ情報作成API
func UserCreatePost(c *gin.Context) {
	var req UserCreateRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}

	trans, err := db.GetDB().Begin()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
	}

	user := model.User{Name: req.Name}

	err = trans.Insert(&user)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
		return
	}

	token, err := helper.GenerateToken(user.ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
		return
	}

	if err := trans.Commit(); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
		return
	}

	c.JSON(http.StatusOK, UserCreateResponse{Token: token})
}

// UserGetGet - ユーザ情報取得API
func UserGetGet(c *gin.Context) {
	token := c.GetHeader("X-Token")
	userID, err := helper.GetUserIDFromToken(token)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err).SetType(gin.ErrorTypePrivate)
		return
	}

	obj, err := db.GetDB().Get(model.User{}, userID)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
		return
	}
	user := obj.(*model.User)

	c.JSON(http.StatusOK, UserGetResponse{Name: user.Name})
}

// UserUpdatePut - ユーザ情報更新API
func UserUpdatePut(c *gin.Context) {
	token := c.GetHeader("X-Token")
	userID, err := helper.GetUserIDFromToken(token)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err).SetType(gin.ErrorTypePrivate)
		return
	}

	var req UserUpdateRequest
	if err := c.BindJSON(&req); err != nil {
		return
	}

	// TODO: Userカラムが増えた場合どうするのか
	user := model.User{
		ID:   userID,
		Name: req.Name,
	}

	_, err = db.GetDB().Update(&user)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
		return
	}
}
