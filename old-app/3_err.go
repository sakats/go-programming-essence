package main

import "fmt"


// 例: Goでのエラー伝達
// Goではエラーが発生し得る関数を呼び出した場合には、まずerrをチェックする
// 即座に呼び出し元に返却することが推奨されている
func FindUser(name string) (*User, error){
	user, err : = findUserFromList(users, name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// この例ではFindUser関数から帰ってきたerrの内容に応じてログなどの処理を呼び出し元が行う
func main(){
	user, err := FindUser("Satoshi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.Name)
}