package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func copyDir(src, dst string) error {
	// 确保源路径存在且是目录
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("源路径错误: %v", err)
	}
	if !srcInfo.IsDir() {
		return fmt.Errorf("源路径不是目录: %s", src)
	}

	// 创建目标目录
	err = os.MkdirAll(dst, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("创建目标目录失败: %v", err)
	}

	// 遍历源目录
	err = filepath.Walk(src, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 计算目标路径
		relPath, err := filepath.Rel(src, srcPath)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)

		// 如果是目录，创建目录
		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		// 如果是文件，复制文件内容
		return copyFile(srcPath, dstPath)
	})

	return err
}

func copyFile(src, dst string) error {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("打开源文件失败: %v", err)
	}
	defer srcFile.Close()

	// 创建目标文件
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("创建目标文件失败: %v", err)
	}
	defer dstFile.Close()

	// 复制文件内容
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("复制文件内容失败: %v", err)
	}

	// 同步文件权限
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	return os.Chmod(dst, srcInfo.Mode())
}

func main() {
	// 定义源目录和目标目录
	srcDir := "source"
	dstDir := "destination"

	// 创建测试目录和文件
	err := os.MkdirAll(srcDir+"/subdir", 0755)
	if err != nil {
		fmt.Println("创建源目录失败:", err)
		return
	}
	err = os.WriteFile(srcDir+"/file1.txt", []byte("Hello, World!"), 0644)
	if err != nil {
		fmt.Println("创建源文件失败:", err)
		return
	}

	// 执行复制
	err = copyDir(srcDir, dstDir)
	if err != nil {
		fmt.Println("复制文件夹失败:", err)
		return
	}

	fmt.Println("文件夹复制成功！")
}
