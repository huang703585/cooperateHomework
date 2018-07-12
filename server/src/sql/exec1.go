package sql
import (
       "os/exec"
       "fmt"
       "strings"
       "bytes"
       "strconv"
       "time"
)
var db string
var cmd *exec.Cmd
var workplace string

func init(){
     //workplace = "C:\\Users/asus-pc/gowork/Homework"
     //workplace = "../../.."
     workplace = ".."
     fmt.Println("sql init at ./src/sql")
     var err error
     db, err = exec.LookPath("SQLCMD.exe")	
     if err != nil {
     	fmt.Println("no sql server"+err.Error())
     }
     
     cmd = exec.Command(db)
     cmd.Stdin = strings.NewReader("DROP DATABASE HOMEWORK")
     err = cmd.Run()
     if err != nil {
     	fmt.Println("no drop database"+err.Error())
     }
     
     cmd = exec.Command(db)
     cmd.Stdin = strings.NewReader("CREATE DATABASE HOMEWORK")
     err = cmd.Run()
     if err != nil {
     	fmt.Println("no create database"+err.Error())
     }

     sqlpath := fmt.Sprintf("%s/database/server_sql.txt", workplace)
     cmd = exec.Command("SQLCMD", "-d", "HOMEWORK", "-i", sqlpath)
     err = cmd.Run()
     if err != nil {
     	fmt.Println("create table error"+err.Error())
     }
     
     sqlinit := fmt.Sprintf("%s/database/init_data.txt", workplace)
     cmd = exec.Command("SQLCMD", "-d", "HOMEWORK", "-i", sqlinit)
     err = cmd.Run()
     if err != nil {
     	fmt.Println("init data error"+err.Error())
     }
}

func getResult(input string) string{
     cmd = exec.Command(db, "-d", "HOMEWORK")
     cmd.Stdin = strings.NewReader(input)
     var out bytes.Buffer
     cmd.Stdout = &out

     err := cmd.Run()
     if err != nil {
     	fmt.Printf(err.Error())
     }
    
     output := out.String()
     temp := strings.Split(output, "\n")
     respart := temp[2]
     if respart == "" {
     	return ""
     }
     var preindex, posindex int
     var leftbound = strings.Count(respart, "")-2
     for preindex = 0; respart[preindex] == ' ' && preindex < leftbound; preindex++{
     	 
     }
     for posindex = preindex; respart[posindex] != ' ' && posindex < leftbound; posindex++ {
     	 
     }
     
     res := respart[preindex:posindex]
     return res
}
/*
func main(){
     fmt.Println(GetMaintain("huang"))
     ret := GetAnnouncement("2018-01-01")
     fmt.Println(ret)
}*/

func GetSignIn(username, password string) bool{
     input := fmt.Sprintf("SELECT u_name FROM [USR] WHERE u_name = '%s'", username)
     output := getResult(input)
     if output == "" {
     	return false
     }

     input = fmt.Sprintf("SELECT u_pwd AS Password FROM [USR] WHERE u_name = '%s'", username)
     output = getResult(input)
   
     return output==password
}

func GetSignUp(username, password string) bool {
     input := fmt.Sprintf("SELECT u_name FROM [USR] WHERE u_name = '%s'", username)
     output := getResult(input)
     if output != "" {
     	return false
     }
     
     input = fmt.Sprintf("SELECT u_id FROM [USR] WHERE u_id IN(SELECT MAX(u_id) FROM [USR])")
     output = getResult(input)	  
     
     newuid, _ := strconv.Atoi(output)
     newuid = newuid + 1

     cmd = exec.Command(db, "-d", "homework")
     input = fmt.Sprintf("INSERT INTO [USR] VALUES(%d,'%s','%s','','','0')", newuid, username, password)
     cmd.Stdin = strings.NewReader(input)
     err := cmd.Run()
     if err != nil {
     	fmt.Printf(err.Error())
     }	
     return true
}

func GetMaintain(username string) string{
     input := fmt.Sprintf("SELECT r_id FROM REPAIR WHERE r_id IN(SELECT MAX(r_id) FROM REPAIR)")
     output := getResult(input)	  
     newrid, _ := strconv.Atoi(output)
     newrid = newrid + 1

     cmd = exec.Command(db, "-d", "homework")
     t := time.Now()
     yy, mm ,dd := t.Date()
     date := fmt.Sprintf("%d-%d-%d", yy, mm, dd)
     input = fmt.Sprintf("INSERT INTO REPAIR(r_id,r_what,r_who,r_date) VALUES(%d,'door','%s','%s')", newrid, username, date)
     cmd.Stdin = strings.NewReader(input)
     err := cmd.Run()
     if err != nil {
     	fmt.Printf(err.Error())
     }	

     input = fmt.Sprintf("SELECT r_what FROM REPAIR ORDER BY r_id DESC")
     what := getResult(input)
     output = fmt.Sprintf("%d %s %s", newrid, what, date)

     PayUp(username, 50.10)
     addAnnouncement(username, output)
     
     return output
}

func PayUp(username string, bill float64){
     newbill := GetAccount(username) + bill
     input := fmt.Sprintf("UPDATE [USR] SET u_bill = %f WHERE u_name = '%s'", newbill, username)
     
     cmd = exec.Command(db, "-d", "homework")
     cmd.Stdin = strings.NewReader(input)
     err := cmd.Run()
     if err != nil {
     	fmt.Printf(err.Error())
     }
}

func GetAnnouncement(lasttime string) []string{
     input := fmt.Sprintf("SELECT n_what FROM NOTICE WHERE n_date > '%s'", lasttime)
     cmd = exec.Command(db, "-d", "homework")
     cmd.Stdin = strings.NewReader(input)
     var out bytes.Buffer
     cmd.Stdout = &out

     err := cmd.Run()
     if err != nil {
     	fmt.Printf(err.Error())
     }
    
     output := out.String()
     temp := strings.Split(output,"\n")
     res := temp[2:]

     input = fmt.Sprintf("SELECT n_id FROM NOTICE WHERE n_id IN(SELECT MAX(n_id) FROM NOTICE)")
     output = getResult(input)
     num, _ := strconv.Atoi(output)
     /*fmt.Println(num)
     for j := 0; j < num ; j++ {
     	 fmt.Println(res[j])
     }*/

    /* var lastindex int
     for i, s := range res {
     	 if strings.Index(s, "#") == 0 {
	    lastindex = i
	    break
	 }
     }*/
     ret := res[:num]
     ret = append(ret, "end")
     return ret
}

func GetAccount(username string) float64{
     input := fmt.Sprintf("SELECT u_bill AS payment FROM [USR] WHERE u_name = '%s'", username)
     output := getResult(input)	  
     
     res, _ := strconv.ParseFloat(output, 64)
     return res
}

func addAnnouncement(username, s string){
     input := fmt.Sprintf("SELECT n_id FROM NOTICE WHERE n_id IN(SELECT MAX(n_id) FROM NOTICE)")
     output := getResult(input)	  
     newnid, _ := strconv.Atoi(output)
     newnid = newnid + 1

     cmd = exec.Command(db, "-d", "homework")
     t := time.Now()
     yy, mm ,dd := t.Date()
     date := fmt.Sprintf("%d-%d-%d", yy, mm, dd)
     what := fmt.Sprintf("#"+s)
     input = fmt.Sprintf("INSERT INTO NOTICE(n_id,n_title,n_what,n_date) VALUES(%d,'%s','%s','%s')", newnid, username, what, date)
     cmd.Stdin = strings.NewReader(input)
     err := cmd.Run()
     if err != nil {
     	fmt.Printf(err.Error())
     }	
}
