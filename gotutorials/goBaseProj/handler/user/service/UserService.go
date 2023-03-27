package service

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"../../../model"
)

type UserService struct {
	DB *sql.DB
}

// 등록
func (self UserService) PostUser(param model.User) model.UserResult {
	result := model.UserResult{}

	// 1. 정산 정보 등록
	sql := fmt.Sprintf("INSERT INTO user( id, f_name, l_name, email, created_at) VALUES(%s, '%s', '%s', '%s', now());",
		strconv.Itoa(param.ID),
		param.FirstName,
		param.LastName,
		param.Email)

	log.Print("▶ UserService : Insert : sql : ")
	log.Println(sql)

	res, postErr := self.DB.Exec(sql)

	if postErr != nil {
		result.Status = http.StatusInternalServerError
		result.Message = "정산 정보 등록에 실패하였습니다."
		result.Cmd = "INSERT"
		return result
	}

	// 등록된 정산정보의 srl값 받아오기
	lastInsertId, getSrlerr := res.LastInsertId()
	if getSrlerr != nil {
		result.Status = http.StatusInternalServerError
		result.Message = "정산 정보 등록에 실패하였습니다."
		result.Cmd = "INSERT"
		return result
	}

	result.Status = http.StatusOK
	//result.Message = "정산 정보 등록에 성공하였습니다."
	result.Message = strconv.FormatInt(lastInsertId, 10)
	result.Cmd = "INSERT"

	log.Println("정보 등록에 성공하였습니다.", lastInsertId)

	return result
}
