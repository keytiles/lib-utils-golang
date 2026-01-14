package kt_utils_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/keytiles/lib-utils-golang/pkg/kt_utils"
	"github.com/stretchr/testify/assert"
)

type privateStruct struct {
	strField string
}

type RecipientsRole string

type ContainerUserRecipientFilters struct {
	HasAnyOfRoles *[]RecipientsRole `json:"hasAnyOfRoles" yaml:"hasAnyOfRoles"`
}

type ContainerUserRecipients struct {
	ContainerIds *[]string                      `json:"containerIds" yaml:"containerIds"`
	Filters      *ContainerUserRecipientFilters `json:"filters" yaml:"filters"`
}

type Recipients struct {
	ContainerUsers *ContainerUserRecipients `json:"containerUsers" yaml:"containerUsers"`
	UserIdTerms    *[]string                `json:"userIdTerms" yaml:"userIdTerms"`
}

type Notification struct {
	CreatedAt   *int    `json:"createdAt" yaml:"createdAt"`
	HtmlContent *string `json:"htmlContent" yaml:"htmlContent"`
	Id          string  `json:"id" yaml:"id"`
	Severity    string  `json:"severity" yaml:"severity"`
	Title       string  `json:"title" yaml:"title"`
}

type SendNotificationRequest struct {
	Notification Notification `json:"notification" yaml:"notification"`
	Recipients   *Recipients  `json:"recipients" yaml:"recipients"`
}

type UserWithToStringFunc struct {
	Name string
	Age  int
}

func (s UserWithToStringFunc) String() string {
	return fmt.Sprintf("User{name: '%s', Age: %d}", s.Name, s.Age)
}

func getSendNotifRequestStruct() SendNotificationRequest {
	st := SendNotificationRequest{
		Notification: Notification{
			Id:       "notif-id",
			Title:    "Notif Title",
			Severity: "info",
		},
		Recipients: kt_utils.Ptr(Recipients{
			UserIdTerms: kt_utils.Ptr([]string{"user1", "user2"}),
			ContainerUsers: kt_utils.Ptr(ContainerUserRecipients{
				Filters: kt_utils.Ptr(ContainerUserRecipientFilters{
					HasAnyOfRoles: kt_utils.Ptr([]RecipientsRole{RecipientsRole("admin"), RecipientsRole("view")}),
				}),
			}),
		}),
	}
	return st
}

func TestVarPrinterStruct_withToString(t *testing.T) {

	// ---- GIVEN
	st := UserWithToStringFunc{Name: "John Smith", Age: 36}
	vp := kt_utils.VarPrinter{TheVar: st}
	// ---- WHEN
	stStr := vp.String()
	// ---- THEN
	expected := "User{name: 'John Smith', Age: 36}"
	assert.Equal(t, expected, stStr)
	//fmt.Printf("PrintVarS prettyPrint result for struct is: %s\n", stStr)
}

func TestVarPrinterStruct_privateStruct(t *testing.T) {

	// ---- GIVEN
	st := privateStruct{strField: "priv string value"}
	vp := kt_utils.VarPrinter{TheVar: st}
	// ---- WHEN
	stStr := vp.String()
	// ---- THEN
	expected := "kt_utils_test.privateStruct{strField:\"priv string value\"}"
	assert.Equal(t, expected, stStr)
	//fmt.Printf("PrintVarS prettyPrint result for struct is: %s\n", stStr)
}

func TestVarPrinterStruct_defaults(t *testing.T) {

	// ---- GIVEN
	st := getSendNotifRequestStruct()
	vp := kt_utils.VarPrinter{TheVar: st}
	// ---- WHEN
	stStr := vp.String()
	// ---- THEN
	expected := "kt_utils_test.SendNotificationRequest{Notification:kt_utils_test.Notification{CreatedAt:nil,HtmlContent:nil,Id:\"notif-id\",Severity:\"info\",Title:\"Notif Title\"},Recipients:&kt_utils_test.Recipients{ContainerUsers:&kt_utils_test.ContainerUserRecipients{ContainerIds:nil,Filters:&kt_utils_test.ContainerUserRecipientFilters{HasAnyOfRoles:&[]kt_utils_test.RecipientsRole{\"admin\",\"view\"}}},UserIdTerms:&[]string{\"user1\",\"user2\"}}}"
	assert.Equal(t, expected, stStr)
	//fmt.Printf("PrintVarS prettyPrint result for struct is: %s\n", stStr)

	// ---- GIVEN
	// nil test
	var stPtr *SendNotificationRequest
	vp = kt_utils.VarPrinter{TheVar: stPtr}
	// ---- WHEN
	stStr = vp.String()
	// ---- THEN
	expected = "nil"
	assert.Equal(t, expected, stStr)
	//fmt.Printf("PrintVarS prettyPrint result for nil struct ptr is: %s\n", stStr)

}

func TestVarPrinterStruct_prettyPrint(t *testing.T) {

	// ---- GIVEN
	// printing a complex struct
	st := getSendNotifRequestStruct()
	vp := kt_utils.VarPrinter{TheVar: st, PrettyPrint: true}
	// ---- WHEN
	stStr := vp.String()
	// ---- THEN
	expected, err := getFileContentUnixStyle("testdata/sendNotifRequest_prettyprint_result.txt")
	assert.NoError(t, err)
	assert.Equal(t, expected, stStr)
	//fmt.Printf("PrintVarS prettyPrint result for struct is: %s\n", stStr)

	// ---- GIVEN
	// same, but now we have a pointer
	st = getSendNotifRequestStruct()
	vp = kt_utils.VarPrinter{TheVar: &st, PrettyPrint: true}
	// ---- WHEN
	stStr = vp.String()
	// ---- THEN
	expected, err = getFileContentUnixStyle("testdata/sendNotifRequest_ptr_prettyprint_result.txt")
	assert.NoError(t, err)
	assert.Equal(t, expected, stStr)
	//fmt.Printf("PrintVarS prettyPrint result for struct is: %s\n", stStr)

	// ---- GIVEN
	// nil test
	var stPtr *SendNotificationRequest
	vp = kt_utils.VarPrinter{TheVar: stPtr, PrettyPrint: true}
	// ---- WHEN
	stStr = vp.String()
	// ---- THEN
	expected = "nil"
	assert.Equal(t, expected, stStr)
	//fmt.Printf("PrintVarS prettyPrint result for nil struct ptr is: %s\n", stStr)

}

func getFileContentUnixStyle(filePath string) (string, error) {
	contentBytes, err := os.ReadFile(filePath)
	if err == nil {
		content := string(contentBytes)
		content = strings.ReplaceAll(content, "\r", "")
		return content, nil
	}
	return "", err
}
