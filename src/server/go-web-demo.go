package main

import (
	"utils"
	"net/http"
	"fmt"
	"html/template"
	"io"
	"time"
	"os"
	"bytes"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	utils.LogInfoln(utils.GetRequestInfo(r))
	userInfoMap := utils.GetStagedUserInfo(w, r)
	loginName := userInfoMap[utils.LOGIN_NAME_KEY]
	attrMap := utils.GenerateBasicAttrMap(r, (len(loginName) > 0))
	attrMap[utils.LOGIN_NAME_KEY] = loginName
	currentPage := r.FormValue("page")
	if len(currentPage) == 0 {
		currentPage = utils.HOME_PAGE
	}
	t := template.New("welcome page")
	t.Funcs(template.FuncMap{
		"equal": utils.SimpleEqual,
		"match": utils.MatchString,
	})
    t, err := t.ParseFiles(utils.GeneratePagePath(currentPage),
	    utils.GeneratePagePath("header"),
		utils.GeneratePagePath("footer"),
	    utils.GeneratePagePath("navbar"))
	if err != nil {
		utils.LogErrorln("ParseFilesErr:", err)
	}
	attrMap["currentPage"] = currentPage
	err = t.ExecuteTemplate(w, "page", attrMap)
    if err != nil {
		utils.LogErrorln("ExecuteTemplateErr:", err)
    }
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	utils.LogInfoln(utils.GetRequestInfo(r))
	userInfoMap := utils.GetStagedUserInfo(w, r)
	loginName := userInfoMap[utils.LOGIN_NAME_KEY]
	if r.Method == "GET" {
		tokenKey := utils.GenerateTokenKey(loginName, r)
		utils.LogInfoln("TokenKey:", tokenKey)
		token := utils.GenerateToken()
		utils.LogInfo("Token:", token)
		utils.SetToken(tokenKey, token)
		attrMap := utils.GenerateBasicAttrMap(r, false)
		attrMap["token"] = token
		t, err := template.ParseFiles(utils.GeneratePagePath("login"))
		if err != nil {
			utils.LogErrorln("TemplateParseErr:", err)
		}
		err = t.Execute(w, attrMap)
        if err != nil {
			utils.LogErrorln("PageWriteErr:", err)
        }
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		utils.LogInfoln("Token:", token)
		validToken := false
		if token != "" {
			tokenKey := utils.GenerateTokenKey(loginName, r)
			utils.LogInfoln("TokenKey:", tokenKey)
			storedToken := utils.GetToken(tokenKey)
			utils.LogInfoln("StoredToken:", storedToken)
			if len(token) > 0 && len(storedToken)> 0 && token == storedToken {
				validToken = true
			}
		}
		loginName = template.HTMLEscapeString(r.Form.Get(utils.LOGIN_NAME_KEY))
		utils.LogInfoln("login - loginName:", loginName)
		password := template.HTMLEscapeString(r.Form.Get(utils.PASSWORD_KEY))
		utils.LogInfoln("login - password:", password)
		rememberMe := r.Form.Get("remember-me")
		utils.LogInfoln("login - remember-me:", rememberMe)
		validLogin, err := utils.VerifyUser(loginName, password)
		utils.LogInfoln("Verify user:", validLogin)
		if err != nil {
			utils.LogErrorf("VerifyUserError (loginName=%s): %s\n", loginName, err)
		} else {
			rememberMeTag := r.Form.Get("remember-me")
			if validLogin {
				if validToken {
					userInfoMap[utils.LOGIN_NAME_KEY] = loginName
					onlySession := len(rememberMeTag) == 0 || rememberMeTag != "y"
					utils.SetUserInfoToStage(userInfoMap, w, r, onlySession)
				}
			}
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	utils.LogInfoln(utils.GetRequestInfo(r))
	userInfoMap := utils.GetStagedUserInfo(w, r)
	loginName := userInfoMap[utils.LOGIN_NAME_KEY]
	if len(loginName) > 0 {
		utils.RemoveUserInfoFromStage(userInfoMap, w, r)
		utils.LogInfoln("Logout: The user '%s' has  logout.\n", loginName)
	} else {
		utils.LogInfoln("Logout: The user '%s' has yet login.\n", loginName)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	utils.LogInfoln(utils.GetRequestInfo(r))
	if r.Method == "GET" {
		attrMap := utils.GenerateBasicAttrMap(r, false)
		encodedHint := r.FormValue("hint")
		if len(encodedHint) > 0 {
			hint := utils.UrlDecoding(encodedHint)
			attrMap["hint"] = hint
		}
		t, _ := template.ParseFiles(utils.GeneratePagePath("register"))
		err := t.Execute(w, attrMap)
		if err != nil {
			utils.LogErrorln("PageWriteErr:", err)
		}
	} else {
		fieldMap, invalidFields := utils.VerifyRegisterForm(r)
		utils.LogInfoln("The field map:", fieldMap)
		if len(invalidFields) > 0 {
			hint := fmt.Sprintln("There are some invalid fields of '':", invalidFields, ".")
			utils.LogInfoln(hint)
			encodedHint := utils.UrlEncoding(hint)
			redirectUrl := "/register?hint=" + encodedHint
			http.Redirect(w, r, redirectUrl, http.StatusFound)
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	utils.LogInfoln(utils.GetRequestInfo(r))
	if r.Method == "GET" {
		token := r.Form.Get("token")
		t, _ := template.ParseFiles(utils.GeneratePagePath("upload"))
		err := t.Execute(w, token)
        if err != nil {
			utils.LogErrorln("PageWriteErr:", err)
        }
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			utils.LogErrorln("UploadFileParsError:", err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		var buffer bytes.Buffer
		buffer.WriteString(os.TempDir())
		buffer.WriteString("/")
		buffer.WriteString(handler.Filename)
		tempFilePath := buffer.String()
		f, err := os.OpenFile(tempFilePath, os.O_WRONLY | os.O_CREATE, 0666)
		if err != nil {
			utils.LogErrorln(err)
			return
		}
		defer f.Close()
		utils.LogInfoln("Receive a file & save to %s ...\n", tempFilePath)
		io.Copy(f, file)
		go utils.DeleteTempFile(time.Duration(time.Minute * 5), tempFilePath)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("web"))
    http.Handle("/css/", fileServer)
    http.Handle("/js/", fileServer)
    http.Handle("/img/", fileServer)
    http.HandleFunc("/", welcome)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/upload", upload)
	conf, err := utils.ReadConfig(true)
	if err != nil {
		utils.LogFatalln("ConfigLoadError: ", err)
	} else {
		addr := ":" + fmt.Sprintf("%v", conf.ServerPort)
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			utils.LogFatalln("ListenAndServeError: ", err)
		}
	}
}
