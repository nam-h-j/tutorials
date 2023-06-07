package service

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"goBaseProj/model"
)

type UserService struct {
	DB *sql.DB
}

// 목록가져오기
func (self UserService) GetUserList() model.UserListResult {
	result := model.UserListResult{}

	// prepared statements
	getUserListQuery := "SELECT id, f_name, l_name, email, created_at FROM user"
	getUserListStmt, err := self.DB.Prepare(getUserListQuery)
	if err != nil {
		fmt.Println(err)
		result.Status = http.StatusInternalServerError
		result.Message = "내부 서버 오류"
		result.Cmd = "ERROR"
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
func (self UserService) GetUser(userId string) model.UserResult {
	result := model.UserResult{}

	// 파라미터를 확인한다.
	fmt.Println("userID : ", userId)

	// prepared statement
	getUserQuery := "SELECT id, f_name, l_name, email, created_at FROM user WHERE id = ?"
	getUserStmt, err := self.DB.Prepare(getUserQuery)
	if err != nil {
		fmt.Println(err)
		result.Status = http.StatusInternalServerError
		result.Message = "내부 서버 오류"
		result.Cmd = "ERROR"
		return result
	}
	defer getUserStmt.Close()

	// run query, Scan Data
	// single row를 가져올땐 QueryRow를 사용한다.
	tmp_item := model.User{}
	err = getUserStmt.QueryRow(userId).Scan(
		&tmp_item.ID,
		&tmp_item.FirstName,
		&tmp_item.LastName,
		&tmp_item.Email,
		&tmp_item.CreatedAt)
	if err != nil {
		fmt.Println(err)
		result.Status = http.StatusBadRequest
		result.Message = "등록된 회원 정보가 없습니다."
		result.Cmd = "ERROR"
		return result
	}

	// success resp
	result.UserData = tmp_item
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
		result.Status = http.StatusInternalServerError
		result.Message = "내부 서버 오류"
		result.Cmd = "ERROR"
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
	result.UserData = body

	log.Println("정보 등록 성공", lastInsertId)

	return result
}

// 업데이트
func (self UserService) PutUser(body model.User) model.UserResult {
	result := model.UserResult{}

	// prepared statement
	putUserQuery := "UPDATE user SET f_name = ?, l_name = ? WHERE id = ?"
	putUserStmt, err := self.DB.Prepare(putUserQuery)
	if err != nil {
		fmt.Println(err)
		result.Status = http.StatusInternalServerError
		result.Message = "내부 서버 오류"
		result.Cmd = "ERROR"
		return result
	}
	defer putUserStmt.Close()

	// post 요청 body 값 확인
	fmt.Println(
		"[*** PUT Req Body ***] \n",
		"id : ", body.ID, "\n",
		"f_name : ", body.FirstName, "\n",
		"l_name : ", body.LastName,
	)

	// UPDATE 쿼리요청
	res, err := putUserStmt.Exec(body.FirstName, body.LastName, body.ID)
	if err != nil {
		result.Status = http.StatusBadRequest
		result.Message = "정보 수정 실패"
		result.Cmd = "UPDATE"
		return result
	}

	// 수정된 row 개수 확인
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		result.Status = http.StatusInternalServerError
		result.Message = "정보 수정 실패"
		result.Cmd = "UPDATE"
		return result
	}

	// 성공 응답 리턴
	result.Status = http.StatusOK
	result.Message = fmt.Sprintf("정보 수정에 성공하였습니다. (%d 개 데이터 수정 됨)", rowsAffected)
	result.Cmd = "UPDATE"

	return result
}

// 단일 유저 삭제
func (self UserService) DeleteUser(userId string) model.UserResult {
	result := model.UserResult{}

	// Prepared Statement
	getUserQuery := "SELECT id FROM user WHERE id = ?"
	getUserStmt, err := self.DB.Prepare(getUserQuery)
	if err != nil {
		fmt.Println(err)
		result.Status = http.StatusInternalServerError
		result.Message = "내부 서버 오류"
		result.Cmd = "ERROR"
		return result
	}
	defer getUserStmt.Close()

	err = getUserStmt.QueryRow(userId).Scan(&result.UserData.ID)
	if err != nil {
		result.Status = http.StatusBadRequest
		result.Message = "등록된 회원 정보가 없습니다."
		result.Cmd = "DELETE"
		return result
	}

	// 삭제하기
	delUserQuery := "Delete FROM user WHERE id = ?"
	delUserStmt, err := self.DB.Prepare(delUserQuery)
	if err != nil {
		fmt.Println(err)
		result.Status = http.StatusInternalServerError
		result.Message = "내부 서버 오류"
		result.Cmd = "ERROR"
		return result
	}
	defer delUserStmt.Close()

	_, delErr := delUserStmt.Query(result.UserData.ID)

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
