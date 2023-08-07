/*
 * Author : Pradana Novan Rianto (https://github.com/pradananovanr)
 * Created on : Thu Apr 13 2023
 * Copyright : Pradana Novan Rianto © 2023. All rights reserved
 */

package entity

type Customer struct {
	ID_Customer int    `json:"id_customer"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}

type CustomerLogin struct {
	ID_Customer int    `json:"id_customer"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
