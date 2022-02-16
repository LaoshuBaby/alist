package controllers

import (
	"encoding/base64"
	"fmt"
	"github.com/Xhofe/alist/conf"
	"github.com/Xhofe/alist/server/common"
	"github.com/Xhofe/alist/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func Favicon(c *gin.Context) {
	c.Redirect(302, conf.GetStr("favicon"))
}

func Plist(c *gin.Context) {
	data := c.Param("data")
	bytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		common.ErrorResp(c, err, 500)
		return
	}
	u := string(bytes)
	name := utils.Base(u)
	name = strings.TrimRight(name, ".ipa")
	plist := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?><!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <key>items</key>
        <array>
            <dict>
                <key>assets</key>
                <array>
                    <dict>
                        <key>kind</key>
                        <string>software-package</string>
                        <key>url</key>
                        <string>%s</string>
                    </dict>
                </array>
                <key>metadata</key>
                <dict>
                    <key>bundle-identifier</key>
                     <string>ci.nn.%s</string>
                     <key>bundle-version</key>
                    <string>1.0.0</string>
                    <key>kind</key>
                    <string>software</string>
                    <key>title</key>
                    <string>%s</string>
                </dict>
            </dict>
        </array>
    </dict>
</plist>`, u, name, name)
	c.Status(200)
	c.Header("Content-Type", "application/xml;charset=utf-8")
	_, _ = c.Writer.WriteString(plist)
}
