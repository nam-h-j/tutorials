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

// 목록가져오기
func  (self UserService) GetUserList() model.UserListResult {
	result := model.UserListResult{}
	
	// prepared statements
	getUserListQuery := "SELECT id, f_name, l_name, email, created_at FROM user" 
	getUserListStmt, err := self.DB.Prepare(getUserListQuery)
	if err != nil {
		fmt.Println(err)
		return result
	}
	defer getUserListStmt.Close()


	// run query
	rows, err := getUserListStmt.Query()
	if err != nil {
		fmt.Println(err)
		return result
	}
	defer rows.Close()
	

	for rows.Next() {
		tmp_item := model.User{}
		err := rows.Scan(
			&tmp_item.ID,
			&tmp_item.FirstName,
			&tmp_item.LastName,
			&tmp_item.Email,
			&tmp_item.CreatedAt,
		)
		if err != nil {
			log.Fatal(err)
		}
		result.UserList = append(result.UserList, tmp_item)
	}
	// 일치하는 정보 없음
	if err != nil {
		log.Println(err)

		result.Status = http.StatusNoContent
		result.Message = "등록된 회원 정보가 없습니다."
		result.Cmd = "SELECT"

		return result
	}

	result.Status = http.StatusOK
	result.Message = "success"
	result.Cmd = "SELECT"

	return result
}

// 단일 유저 취득
func  (self UserService) GetUser(userId string) model.UserResult {
	result := model.UserResult{}
	sql := fmt.Sprintf("SELECT id, f_name, l_name, email, created_at FROM user WHERE id = %s", userId)

	rows, err := self.DB.Query(sql)

	for rows.Next() {
		tmp_item := model.User{}
		err := rows.Scan(
			&tmp_item.ID,
			&tmp_item.FirstName,
			&tmp_item.LastName,
			&tmp_item.Email,
			&tmp_item.CreatedAt,
		)
		if err != nil {
			log.Fatal(err)
		}
		result.UserData = tmp_item
	}
	
	// 일치하는 정보 없음
	if err != nil {
		log.Println(err)

		result.Status = http.StatusNoContent
		result.Message = "등록된 회원 정보가 없습니다."
		result.Cmd = "SELECT"

		return result
	}

	result.Status = http.StatusOK
	result.Message = "success"
	result.Cmd = "SELECT"

	return result
}

// 등록
func (self UserService) PostUser(body model.User) model.UserResult {
	result := model.UserResult{}

	postUserQuery := "INSERT INTO user( id, f_name, l_name, email, created_at) VALUES(?, ?, ?, ?, now());"
	postUserStmt, err := self.DB.Prepare(postUserQuery)
	if err != nil {
		fmt.Println(err)
		return result
	}
	defer postUserStmt.Close()

	// post 요청하기
	res, err := postUserStmt.Exec(strconv.Itoa(body.ID), body.FirstName, body.LastName, body.Email)
	if err != nil {
		result.Status = http.StatusInternalServerError
		result.Message = "정보 등록 실패"
		result.Cmd = "INSERT"
		return result
	}

	// 등록된 정보의 srl값 및 아이디 받아오기
	lastInsertId, getSrlerr := res.LastInsertId()
	if getSrlerr != nil {
		result.Status = http.StatusInternalServerError
		result.Message = "정보 등록 실패"
		result.Cmd = "INSERT"
		return result
	}

	result.Status = http.StatusOK
	result.Message = strconv.FormatInt(lastInsertId, 10)
	result.Cmd = "INSERT"

	log.Println("정보 등록 성공", lastInsertId)

	return result
}

// 업데이트
func (self UserService) PutUser(body model.User) model.UserResult {
	result := model.UserResult{}

	putUserQuery := "UPDATE user SET f_name = '?', l_name = '?' WHERE id = ?"
		putUserStmt, err := self.DB.Prepare(putUserQuery)
	if err != nil {
		fmt.Println(err)
		return result
	}
	defer putUserStmt.Close()

	// post 요청하기
	res, err := putUserStmt.Exec(body.FirstName, body.LastName, strconv.Itoa(body.ID))
	if err != nil {
		result.Status = http.StatusInternalServerError
		result.Message = "정보 갱신 실패"
		result.Cmd = "UPDATE"
		return result
	}

	// 등록된 정산정보의 srl값 받아오기
	lastInsertId, getSrlerr := res.LastInsertId()
	if getSrlerr != nil {
		result.Status = http.StatusInternalServerError
		result.Message = "정보 갱신 실패"
		result.Cmd = "UPDATE"
		return result
	}

	result.Status = http.StatusOK
	//result.Message = "정산 정보 등록에 성공하였습니다."
	result.Message = strconv.FormatInt(lastInsertId, 10)
	result.Cmd = "INSERT"

	log.Println("정보 등록에 성공하였습니다.", lastInsertId)

	return result
}

// 단일 유저 삭제
func  (self UserService) DeleteUser(userId string) model.UserResult {
	result := model.UserResult{}
	// 셀렉트 먼저 해봄
	sql := fmt.Sprintf("SELECT id FROM user WHERE id = %s", userId)
	err := self.DB.QueryRow(sql).Scan(&result.UserData.ID)
	if err != nil {
		log.Println(err)
		
		result.Status = http.StatusNoContent
		result.Message = "등록된 회원 정보가 없습니다."
		result.Cmd = "DELETE"

		return result
	}
	sql = fmt.Sprintf("Delete FROM user WHERE id = %d", result.UserData.ID)
	_, delErr := self.DB.Query(sql)

	if delErr != nil {
		log.Println(delErr)

		result.Status = http.StatusBadRequest
		result.Message = "회원 삭제 오류"
		result.Cmd = "DELETE"

		return result
	}
	
	result.Status = http.StatusOK
	result.Message = "회원 삭제 성공"
	result.Cmd = "DELETE"

	return result
}