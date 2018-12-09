package v1
/*
import (
	"strings"
	"time"
	"os"
	//"fmt"
	"log"
	"encoding/csv"
	"io"
	//"net/http"
	//"github.com/gin-gonic/gin"
)

//csv文件读写操作
func CsvHandle(user_id int, csv_filename string, task_column_number int, task_project_name string, is_append bool, number_labels int) {
	log.Println("begin to work！")
	start_time := time.Now().Unix()

	file_in, err := os.Open(csv_filename) //读取的csv文件
	if err != nil {
		panic(err)
	}
	defer file_in.Close()

	file_out_name := task_project_name + "_" + csv_filename
	file_out, err := os.Create(file_out_name)  //创建一个新写入文件

	if err != nil {
		panic(err)
	}
	defer file_out.Close()

	reader := csv.NewReader(file_in)
	writer := csv.NewWriter(file_out)
	for {				//读取第一行表头
		row, err := reader.Read()
		if err == io.EOF{
			break
		} else if err != nil {
			panic(err)	
		}
		text := row[task_column_number]
		tae_text := make(map[string]interface{})
		tae_text["body"] = text

		cc_result := tae_client.getResult(tae_text, []string {task_project_name}, 'all')   //从tea服务返回标签结果,处理的函数名有待修改

		csv_line_list := []string {}
		if is_append == false {
			if err != nil {
				csv_line_list = csv_line_list
			} else {
				for i:=0; i<len(row); i++ {
					csv_line_list = append([]string {row[i]}, csv_line_list)
				}
			}
				
		}	

		if len(cc_result) == 0 {
			continue
		} else {
			for cache, branch := range cc_result {
				top := branch[0]
				sentence := branch[1]
				if number_labels == 1 {
					var rows [][]string
					rows[0] = row[0]
					rows[1] = task_project_name
					rows[2] = sentence
					e_label := strings.Split(top, "/")[1:]
					rows = append(e_label, rows)
					
				} else {
					rows := []string {row[0], task_project_name, sentence}
				}
			}
		}
		if file_add := 0 {
			var result [][]string
			rows = rows[1:]
			result = append(csv_line_list, result)
			result = append(rows, result)
			writer.WriteAll(result)
		} else {
			writer.WriteAll(rows)
		}

		if err := writer.Error(); err != nil {
			log.Fatalln("写csv文件出错：", err)
		}
	end_time := time.Now().Unix()
	log.Println("the tasks have done")
	log.Println("time use %f" % (end_time-start_time))
	}
	task_result := make(map[string]interface{})
	task_result["taskStatus"] = "SUCCESS"
	task_result["user_id"] = user_id
	task_result["file_name"] = file_name
	task_result["tree"] = task_project_name
	task_result["task_project_name"] = task_project_name
	task_result["task_column_number"] = task_column_number
	task_result["file_add"] = file_add
	task_result["file_out_name"] = file_out_name
	return task_result
	*/