## gomail


### 描述
配置邮件，可用于消息发送、错误报警...

目前仅支持SMTP服务器发送邮件。


使用方法：
```go
defailtMailer := email.NewEmail(&email.SMTPInfo{
Host:     "Host",
Port:     "Port",
IsSSL:    "IsSSL",
UserName: "UserName",
Password: "Password",
From:     "From",
})

err := defailtMailer.SendMail(
global.EmailSetting.To,
fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
fmt.Sprintf("错误信息: %v", err),
)
if err != nil {
global.Logger.Panicf(c, "mail.SendMail err: %v", err)
}

```