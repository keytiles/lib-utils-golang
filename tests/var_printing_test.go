package kt_utils_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/keytiles/lib-utils-golang/v2/pkg/kt_utils"
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

// Lazy []byte-as-string formatting via VarPrinter.BytesAsString: text output, nil/empty slices,
// PrintS dispatch, flag ignored for non-[]byte TheVar, and regression that raw []byte still uses litter dump.
func TestVarPrinter_BytesAsString(t *testing.T) {

	t.Run("formats byte slice as string", func(t *testing.T) {
		// ---- GIVEN
		// text payload with BytesAsString flag for lazy conversion
		vp := kt_utils.VarPrinter{TheVar: []byte("hello"), BytesAsString: true}
		// ---- WHEN
		stStr := vp.String()
		// ---- THEN
		expected := "hello"
		assert.Equal(t, expected, stStr)
	})

	t.Run("nil slice yields empty string", func(t *testing.T) {
		// ---- GIVEN
		// nil byte slice
		vp := kt_utils.VarPrinter{TheVar: []byte(nil), BytesAsString: true}
		// ---- WHEN
		stStr := vp.String()
		// ---- THEN
		expected := ""
		assert.Equal(t, expected, stStr)
	})

	t.Run("empty slice yields empty string", func(t *testing.T) {
		// ---- GIVEN
		// empty but non-nil byte slice
		vp := kt_utils.VarPrinter{TheVar: []byte{}, BytesAsString: true}
		// ---- WHEN
		stStr := vp.String()
		// ---- THEN
		expected := ""
		assert.Equal(t, expected, stStr)
	})

	t.Run("uses PrintS path not PrintVarS on raw TheVar", func(t *testing.T) {
		// ---- GIVEN
		// BytesAsString is on VarPrinter — PrintVarS on TheVar alone does not apply the flag
		b := []byte("payload")
		vp := kt_utils.VarPrinter{TheVar: b, BytesAsString: true}
		// ---- WHEN
		stStr := vp.PrintS(false)
		// ---- THEN
		expected := "payload"
		assert.Equal(t, expected, stStr)
	})

	t.Run("flag ignored when TheVar is not byte slice", func(t *testing.T) {
		// ---- GIVEN
		// BytesAsString set but TheVar is a string — falls through to normal PrintVarS
		vp := kt_utils.VarPrinter{TheVar: "hello", BytesAsString: true}
		// ---- WHEN
		stStr := vp.String()
		// ---- THEN
		expected := "\"hello\""
		assert.Equal(t, expected, stStr)
	})

	t.Run("raw byte slice still uses litter dump", func(t *testing.T) {
		// ---- GIVEN
		// raw []byte without BytesAsString — unchanged v1.2 litter behavior
		vp := kt_utils.VarPrinter{TheVar: []byte("hello")}
		// ---- WHEN
		stStr := vp.String()
		// ---- THEN
		expected := "[]uint8{104,101,108,108,111}"
		assert.Equal(t, expected, stStr)
	})
}

type valueWithLogAndString struct {
	Secret string
}

func (v valueWithLogAndString) String() string {
	return "stringer:" + v.Secret
}

func (v valueWithLogAndString) LogString() string {
	return "log:masked"
}

type stringerAndError struct {
	msg string
}

func (e stringerAndError) String() string {
	return "stringer:" + e.msg
}

func (e stringerAndError) Error() string {
	return "error:" + e.msg
}

// Lazy Transform on VarPrinter, PrintVarSWithTransform, BytesAsString-after-transform,
// VarPrinter-only LogStringer, and Stringer-before-error regression.
func TestVarPrinter_TransformAndLogStringer(t *testing.T) {

	t.Run("Transform runs only when String is invoked", func(t *testing.T) {
		// ---- GIVEN
		calls := 0
		vp := kt_utils.VarPrinter{
			TheVar: "raw",
			Transform: func(v any) any {
				calls++
				return "transformed"
			},
		}
		assert.Equal(t, 0, calls)
		// ---- WHEN
		stStr := vp.String()
		// ---- THEN
		assert.Equal(t, 1, calls)
		assert.Equal(t, "\"transformed\"", stStr)
	})

	t.Run("Transform result is what litter sees", func(t *testing.T) {
		// ---- GIVEN
		vp := kt_utils.VarPrinter{
			TheVar: privateStruct{strField: "secret"},
			Transform: func(v any) any {
				return privateStruct{strField: "masked"}
			},
		}
		// ---- WHEN
		stStr := vp.String()
		// ---- THEN
		expected := "kt_utils_test.privateStruct{strField:\"masked\"}"
		assert.Equal(t, expected, stStr)
	})

	t.Run("BytesAsString applies after Transform", func(t *testing.T) {
		// ---- GIVEN
		vp := kt_utils.VarPrinter{
			TheVar:        "ignored",
			BytesAsString: true,
			Transform: func(v any) any {
				return []byte("from-transform")
			},
		}
		// ---- WHEN
		stStr := vp.String()
		// ---- THEN
		assert.Equal(t, "from-transform", stStr)
	})

	t.Run("VarPrinter prefers LogStringer over Stringer", func(t *testing.T) {
		// ---- GIVEN
		v := valueWithLogAndString{Secret: "top-secret"}
		vp := kt_utils.VarPrinter{TheVar: v}
		// ---- WHEN
		stStr := vp.String()
		// ---- THEN
		assert.Equal(t, "log:masked", stStr)
	})

	t.Run("PrintVarS ignores LogStringer and uses Stringer", func(t *testing.T) {
		// ---- GIVEN
		v := valueWithLogAndString{Secret: "top-secret"}
		// ---- WHEN
		stStr := kt_utils.PrintVarS(v, false)
		// ---- THEN
		assert.Equal(t, "stringer:top-secret", stStr)
	})

	t.Run("PrintVarSWithTransform applies transform then Stringer", func(t *testing.T) {
		// ---- GIVEN
		// ---- WHEN
		stStr := kt_utils.PrintVarSWithTransform("x", false, func(v any) any {
			return UserWithToStringFunc{Name: "Ann", Age: 20}
		})
		// ---- THEN
		assert.Equal(t, "User{name: 'Ann', Age: 20}", stStr)
	})

	t.Run("Stringer wins over error", func(t *testing.T) {
		// ---- GIVEN
		v := stringerAndError{msg: "boom"}
		// ---- WHEN
		viaPrintVarS := kt_utils.PrintVarS(v, false)
		viaVarPrinter := kt_utils.VarPrinter{TheVar: v}.String()
		// ---- THEN
		assert.Equal(t, "stringer:boom", viaPrintVarS)
		assert.Equal(t, "stringer:boom", viaVarPrinter)
	})
}
