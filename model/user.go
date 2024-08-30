package model

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizfpsoares/albums/storage"
)

type User struct {
	UserId      string
	Name        string
	Surname     string
	Cpf         string
	Email       string
	DateOfBirth string
	Password    string
}

func AddUser(add *gin.Context) {
	user := User{}
	data, err := add.GetRawData()
	if err != nil {
		add.AbortWithStatusJSON(400, "Usuário não definido")
		return
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		add.AbortWithStatusJSON(400, "Entrada incorreta")
		return
	}

	_, err = storage.Db.Exec("insert into users(user_name, user_surname, user_cpf, user_email, user_date_of_birth, user_password) values ($1,$2,$3,$4,$5,$6)", user.Name, user.Surname, user.Cpf, user.Email, user.DateOfBirth, user.Password)
	if err != nil {
		fmt.Println(err)
		add.AbortWithStatusJSON(400, "Não foi possivel criar o novo usuário.")
	} else {
		add.JSON(http.StatusOK, "Usuário cadastrado com sucesso.")
	}
}

func GetUser(find *gin.Context) {
	find.Header("Content-Type", "application/json")

	rows, err := storage.Db.Query("SELECT user_id, user_name, user_surname, user_cpf, user_email, user_date_of_birth FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserId, &user.Name, &user.Surname, &user.Cpf, &user.Email, &user.DateOfBirth)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	find.IndentedJSON(http.StatusOK, users)
}
