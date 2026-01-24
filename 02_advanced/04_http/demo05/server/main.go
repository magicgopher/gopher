package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// 服务端示例
// 处理POST请求（multipart/form-data）格式的数据

type Response struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 信息
	Data    interface{} `json:"data"`    // 数据
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// 设置公共响应头 (JSON格式)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 校验请求方法
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Response{
			Code:    http.StatusMethodNotAllowed,
			Message: "只支持 POST 请求",
		})
		return
	}
	// 限制文件大小 (10MB)
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	// 解析 Multipart 表单
	if err := r.ParseMultipartForm(32 << 10); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Code:    http.StatusBadRequest,
			Message: "文件太大或解析失败: " + err.Error(),
		})
		return
	}
	// 获取上传的文件
	// 注意：这里的 "file" 是前端表单的 key 名字 (name="file")
	// 如果你用 Postman 测试，Key 必须填 "file"；如果你用之前的 Client 代码，这里要改成 "my_file"
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Code:    http.StatusBadRequest,
			Message: "无法获取文件，请检查表单字段名是否为 'file'",
		})
		return
	}
	// 记得关闭上传文件的流
	defer file.Close()

	// 创建本地目标文件
	// 这里指定文件名为 upload_file.txt
	dst, err := os.Create("upload_file.txt")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Code:    http.StatusInternalServerError,
			Message: "服务器创建文件失败",
		})
		log.Printf("创建文件失败: %v", err)
		return
	}
	// 记得关闭本地文件的流
	defer dst.Close()
	// 拷贝数据 (将上传的文件流 copy 到本地文件流)
	// io.Copy 是流式拷贝，不会把大文件一次性加载到内存，很安全
	written, err := io.Copy(dst, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Code:    http.StatusInternalServerError,
			Message: "文件保存过程中出错",
		})
		log.Printf("文件拷贝失败: %v", err)
		return
	}
	fmt.Printf("成功保存文件: upload_file.txt, 大小: %d 字节\n", written)
	// 返回 JSON 成功信息
	w.WriteHeader(http.StatusOK)
	resp := Response{
		Code:    http.StatusOK,
		Message: "上传并保存成功",
		Data: map[string]interface{}{
			"filename":      "upload_file.txt",
			"original_name": header.Filename,
			"size":          written,
		},
	}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.Handle("/api/v1/upload", http.HandlerFunc(uploadHandler))
	log.Println("服务端启动成功！")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("服务端启动失败: %v\n", err)
	}
}
