package handler

import (
	"api18/member"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerMember struct {
	handlerServe member.Serve
}

func NewHandlerMember(handlerServe member.Serve) *handlerMember {
	return &handlerMember{handlerServe}
}

func (h *handlerMember) SaveHandlerMember(g *gin.Context) {
	var keyHandlerMemberServe member.InputMember
	err := g.ShouldBindJSON(&keyHandlerMemberServe)

	if err != nil {
		// var errors []string
		// for _, e := range err.(validator.ValidationErrors) {
		// 	errors = append(errors, e.Error())
		// }

		// errorsMessage := gin.H{"errors": errors}

		keyErrorsMember := member.NewFormatterError(err)
		errorsMessage := gin.H{"errors": keyErrorsMember}
		keyResponsJson := member.NewRespons("Gagal Input", http.StatusUnprocessableEntity, "Gagal", errorsMessage)
		g.JSON(http.StatusUnprocessableEntity, keyResponsJson)
		return
	} else {
		newHandler, err := h.handlerServe.SaveServeMember(keyHandlerMemberServe)

		if err != nil {
			keyResponsJson := member.NewRespons("Gagal Input", http.StatusOK, "Gagal", nil)
			g.JSON(http.StatusBadRequest, keyResponsJson)
			return
		} else {
			keyFormatterMember := member.NewFormatterMember(newHandler, "token token token token ")
			keyResponsJson := member.NewRespons("Berhasil Input Data Member", http.StatusOK, "Success", keyFormatterMember)
			g.JSON(http.StatusOK, keyResponsJson)
		}

	}

}
