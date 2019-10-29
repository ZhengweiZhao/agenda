/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"io/ioutil"
	"encoding/json"
	"log"
)

type userinfo struct{
	Name string
	Password string
	Email string
	Phone string
}

type user struct{
	Id []userinfo
}

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		str := userinfo{
			Name: "",
			Password: "",
			Email: "",
			Phone: "",
		}
		
		username, _:=cmd.Flags().GetString("user")
		password, _:=cmd.Flags().GetString("password")
		email, _:=cmd.Flags().GetString("email")
		phone, _:=cmd.Flags().GetString("phone")
		fmt.Println("register name : "+username)
		fmt.Println("password : "+password)
		fmt.Println("email : "+email)
		fmt.Println("phone : "+phone)
		fmt.Println("Register success!")
		str.Name=username
		str.Password=password
		str.Email=email
		str.Phone=phone
		filename:="./entity/log.log"
		logfile, errr:=os.OpenFile(filename,os.O_RDWR|os.O_APPEND,7)
		if(errr!=nil){
			fmt.Println("openfile fail")
		}
		defer logfile.Close()
		debuglog:=log.New(logfile,"",log.LstdFlags)

		if checkuser(username)==true {
			fmt.Println("username exist, create account fail")
			debuglog.Println("register: username "+username+" exist, create account fail")
		} else {
			debuglog.Println("register: username: "+username+" password: "+password+" email: "+email+" phone:"+phone);
			savecuruser(str)
			input := readinfo()
			input.Id=append(input.Id,str)
			data, _:= json.Marshal(input)
			saveinfo(data)
		}
	},
}


func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("user", "u", "Anonymous", "help message for username")
	registerCmd.Flags().StringP("password", "p", "123", "help message for password")
	registerCmd.Flags().StringP("email", "e", "123@xxx.com", "help message for email")
	registerCmd.Flags().StringP("phone", "f", "13611263068", "help message for phone")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func savecuruser(s userinfo){
	fp, _:=os.OpenFile("./entity/curUser.txt", os.O_RDWR, 7)
	defer fp.Close()
	fp.WriteString("Usrname: "+s.Name+" Password: "+s.Password+" Email: "+s.Email+" Phone: "+s.Phone)
}

func checkuser(username string) bool{
	input := readinfo()
	l :=len(input.Id)
	for i:=0;i<l;i++ {
		if(input.Id[i].Name==username){
			return true
		}
	}
	return false
}

func readinfo() user{
	data, err := ioutil.ReadFile("./entity/data.json")
	if err != nil{
	}
	var user1 user
	err = json.Unmarshal(data, &user1)
	if err!=nil {
		fmt.Println(err)
	}
	return user1
}

func saveinfo(data []byte){
	fp, err := os.OpenFile("./entity/data.json",os.O_WRONLY,0755)
	if err!=nil {
		fmt.Println(err)
	}
	
	_, err = fp.Write(data)
	if err!=nil {
		fmt.Println(err)
	}
	
	defer fp.Close()
}


