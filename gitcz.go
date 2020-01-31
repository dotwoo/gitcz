package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type CzType struct {
	Type    string
	Message string
}

type CzCommit struct {
	Type       *CzType
	Scope      *string
	Subject    *string
	Body       *string
	References *string
}

var StdinInput = bufio.NewReader(os.Stdin)

var (
	InputTypePrompt    = "选择或输入一个提交类型(必填): "
	InputScopePrompt   = "说明本次提交的影响范围(必填): "
	InputSubjectPrompt = "对本次提交进行简短描述(必填): "
	InputBodyPrompt    = `详细说明:
	- 用于解释提交任务的内容和原因，而不是方法
	- 可以列出要点
	- 要点使用空格加上连字符，中间用空行分隔.
	- 连续两个换行结束输入
	- 选填:`
	InputReferencesPrompt = "跟踪的问题或需求(选填):"
)

var CzTypeList = []CzType{
	{
		Type:    "feat",
		Message: "新的功能",
	},
	{
		Type:    "fix",
		Message: "修补错误",
	},
	{
		Type:    "docs",
		Message: "文档修改",
	},
	{
		Type:    "style",
		Message: "格式、分号缺失等，代码无变动",
	},
	{
		Type:    "refactor",
		Message: "重构代码",
	},
	{
		Type:    "perf",
		Message: "性能提高",
	},
	{
		Type:    "test",
		Message: "测试添加、测试重构等，生产代码无变动",
	},
	{
		Type:    "chore",
		Message: "构建任务更新、程序包管理器配置等，生产代码无变动。",
	},
}

func main() {
	author := flag.Bool(
		"author",
		false,
		"关于本软件开发者",
	)
	isFull := flag.Bool("f", false, "完整信息,包含正文和引用")
	isRefer := flag.Bool("r", false, "补充关联问题或需求")

	flag.Parse()
	if *author {
		Author()
		return
	}
	czCommit := &CzCommit{}
	czCommit.Type = InputType()
	czCommit.Scope = InputScope()
	czCommit.Subject = InputSubject()

	if *isFull {
		czCommit.Body = InputBody()
	}

	if *isFull || *isRefer {
		czCommit.References = InputReferences()
	}

	commit := GenerateCommit(czCommit)
	if err := GitCommit(commit); err != nil {
		log.Println(err)
	}
}

func Author() {
	fmt.Println("welcome to our website https://aite.xyz/")
	fmt.Println("----------------------------------------")
	fmt.Println("腾讯扣扣：88966001")
	fmt.Println("电子邮箱：xiaoqidun@gmail.com")
	fmt.Println("----------------------------------------")
	fmt.Println("Copyright (c) 2020 xiaoqidun@gmail.com")
}

func NewLine() {
	fmt.Println()
}

func GitCommit(commit string) (err error) {
	tempFile, err := ioutil.TempFile("", "git_commit_")
	if err != nil {
		return
	}
	defer func() {
		_ = tempFile.Close()
		_ = os.Remove(tempFile.Name())
	}()
	if _, err = tempFile.WriteString(commit); err != nil {
		return
	}
	cmd := exec.Command("git")
	cmd.Args = []string{"git", "commit", "-F" + tempFile.Name()}
	result, err := cmd.CombinedOutput()
	if err != nil && !strings.ContainsAny(err.Error(), "exit status") {
		return
	}
	fmt.Println(string(bytes.TrimSpace(result)))

	return nil
}

func InputType() *CzType {
	typeNum := len(CzTypeList)
	for i := 0; i < typeNum; i++ {
		fmt.Printf("[%d] %s:\t%s\n", i+1, CzTypeList[i].Type, CzTypeList[i].Message)
	}
	fmt.Print(InputTypePrompt)
	text, _ := StdinInput.ReadString('\n')
	text = strings.TrimSpace(text)
	selectID, err := strconv.Atoi(text)
	if err == nil && (selectID > 0 && selectID <= typeNum) {
		NewLine()
		return &CzTypeList[selectID-1]
	}
	for i := 0; i < typeNum; i++ {
		if text == CzTypeList[i].Type {
			NewLine()
			return &CzTypeList[i]
		}
	}
	return InputType()
}

func InputScope() *string {
	fmt.Print(InputScopePrompt)
	text, _ := StdinInput.ReadString('\n')
	text = strings.TrimSpace(text)
	if text != "" {
		NewLine()
		return &text
	}
	return InputScope()
}

func InputSubject() *string {
	fmt.Print(InputSubjectPrompt)
	text, _ := StdinInput.ReadString('\n')
	text = strings.TrimSpace(text)
	if text != "" {
		NewLine()
		return &text
	}
	return InputSubject()
}

func InputBody() *string {
	fmt.Print(InputBodyPrompt)
	conLine := 0
	buf := bytes.NewBufferString("")
	for {
		text, _ := StdinInput.ReadString('\n')
		if text == "\n" {
			conLine++
		} else {
			buf.WriteString(text)
		}

		if conLine >= 2 {
			break
		}
	}

	out := buf.String()
	if out != "" {
		NewLine()
		return &out
	}
	return InputBody()
}

func InputReferences() *string {
	fmt.Print(InputReferencesPrompt)
	text, _ := StdinInput.ReadString('\n')
	text = strings.TrimSpace(text)
	if text != "" {
		NewLine()
		return &text
	}
	return InputReferences()
}

func GenerateCommit(czCommit *CzCommit) string {
	commit := fmt.Sprintf(
		"%s(%s): %s\n",
		czCommit.Type.Type,
		*czCommit.Scope,
		*czCommit.Subject,
	)
	newLine := "\n  "
	commit += newLine
	if czCommit.Body != nil {
		commit += *czCommit.Body
	}
	commit += newLine
	if czCommit.References != nil {
		commit += *czCommit.References
	}
	return commit
}
