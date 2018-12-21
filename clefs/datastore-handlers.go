package clefs

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ikasamt/zapp/zapp"
	"net/http"
	"strings"

	"cloud.google.com/go/datastore"
)

type Anything struct {//generic.Type
	Name string //generic.Type
	Value string //generic.Type
	Errors []error  //generic.Type
} //generic.Type

func (any *Anything) NameKey(k string) *datastore.Key {
	return datastore.NameKey(any.Kind(), k, nil)
}

func (any *Anything) DefaultNameKey() *datastore.Key {
	return any.NameKey(any.NameKeyFormat())
}


func (any *Anything) Kind() string { //generic.Type
	return `Anything`  //generic.Type
}  //generic.Type


func (any Anything) NameKeyFormat() string { //generic.Type
	return fmt.Sprintf("%s", any.Name) //generic.Type
} //generic.Type


func (any *Anything) Setter(c *gin.Context) {//generic.Type
}//generic.Type

func (any *Anything) Validations() error {//generic.Type
}//generic.Type


var ProjectID = "YOUR_PROJECT_ID" //generic.Type
var TemplateDir = "admin" //generic.Type

const URLSep = `/`

func SaveAnything(any *Anything) error {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, ProjectID)
	if err != nil {
		return err
	}

	// Save
	nameKey := any.DefaultNameKey()
	if _, err := client.Put(ctx, nameKey, any); err != nil {
		return err
	}

	return nil
}


func FindOneAnything(nameKey string) (Anything, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, ProjectID)
	if err != nil {
		return Anything{}, err
	}

	any := Anything{}
	if err := client.Get(ctx, any.NameKey(nameKey), &any); err != nil {
		return Anything{}, err
	}

	return any, nil
}


func AnythingShowHandler(c *gin.Context) {
	anyIDStr := c.Param(`ID`)
	any, err:= FindOneAnything(anyIDStr)
	if err!= nil{
		c.String(http.StatusOK, fmt.Sprintf("%v", err))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("%v", any))
}

func AnythingNewHandler(c *gin.Context) {
	instance := Anything{}
	variables := gin.H{"instance": instance}
	zapp.Render(c, TemplateDir, variables)
}

func AnythingCreateHandler(c *gin.Context) {
	instance := &Anything{}
	instance.Setter(c)
	instance.Validations()
	if instance.Errors != nil {
		variables := gin.H{"instance": instance}
		zapp.Render(c, TemplateDir, variables, `new`)
		return
	}

	// 値の更新
	SaveAnything(instance)

	// 完了ページへリダイレクト
	message := fmt.Sprintf("%v 追加しました", instance)
	zapp.SetFlashMessage(c, message)

	paths := strings.Split(c.Request.URL.Path, URLSep)
	backURL := strings.Join(paths[0:len(paths)-1], URLSep)
	c.Redirect(http.StatusFound, backURL)
}
