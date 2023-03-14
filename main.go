package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"log"
	"reflect"
	"time"
)

type User struct {
	UserName string `validate:"minReg" reg_error_info:"用户名至少6个字符"` //通过reg_error_info标签记录
	//reg_error_info也可以是标记错误的唯一标识，通过传入的local_language 从库中或者缓存中找到对应国家的错误提示信息
	Password string `validate:"minReg" reg_error_info:"密码至少6个字符"`
}

//自定义的校验规则，可以使用正则表达式进行匹配，这里仅仅使用了长度判断
func minRegFun(f validator.FieldLevel) bool {
	value := f.Field().String()
	log.Println(f)
	if len(value) < 6 {
		return false
	} else {
		return true
	}

}

type CreateQuotationRequest struct {
	FastQuote FastQuote `json:"fast_quote"`
}

type FastQuote struct {
	Insured Insured `json:"insured"`
}

type Insured struct {
	TravelInfos []TravelInfo `json:"travel_infos"`
}

type TravelInfo struct {
	Row                     string `json:"row"`
	Address                 string `json:"address"`
	NomineeName             string `json:"nominee_name"`
	NomineeNRIC             string `json:"nominee_nric"`
	NomineeAddress          string `json:"nominee_address"`
	MobileNumber            string `json:"mobile_number"`
	RelationshipWithInsured string `json:"relationship_with_insured"`

	TravelType      string          `json:"travel_type,omitempty"`
	TravellerType   string          `json:"traveller_type,omitempty"`
	TravelStartDate string          `json:"travel_start_date,omitempty"`
	TravelEndDate   string          `json:"travel_end_date,omitempty"`
	Area            string          `json:"area,omitempty"`
	TravelFrom      string          `json:"travel_from,omitempty"`
	TravelTo        string          `json:"travel_to,omitempty"`
	PlanType        string          `json:"plan_type,omitempty"`
	ExtremActivity  string          `json:"extrem_activity,omitempty"`
	Relationship    string          `json:"relationship,omitempty"`
	PrimaryNric     string          `json:"primary_nric,omitempty"`
	FirstName       string          `json:"first_name,omitempty"`
	LastName        string          `json:"last_name,omitempty"`
	DateOfBirth     string          `json:"date_of_birth,omitempty"`
	Nric            string          `json:"nric,omitempty"`
	Gender          string          `json:"gender,omitempty"`
	Ethnicity       string          `json:"ethnicity,omitempty"`
	Email           string          `json:"email,omitempty"`
	IsStudent       string          `json:"is_student,omitempty"`
	Premium         decimal.Decimal `json:"premium"`
	Error           string          `json:"error"`
}

func main() {
	a := make(chan int, 1)
	fmt.Println("before")
	a <- 1
	fmt.Println("after")
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(<-a)
	}()
	a <- 1
	fmt.Println(2)
}

func processErr(u interface{}, err error) string {
	if err == nil { //如果为nil 说明校验通过
		return ""
	}

	invalid, ok := err.(*validator.InvalidValidationError) //如果是输入参数无效，则直接返回输入参数错误
	if ok {
		return "输入参数错误：" + invalid.Error()
	}
	validationErrs := err.(validator.ValidationErrors) //断言是ValidationErrors
	for _, validationErr := range validationErrs {
		fieldName := validationErr.Field()                    //获取是哪个字段不符合格式
		field, ok := reflect.TypeOf(u).FieldByName(fieldName) //通过反射获取filed
		if ok {
			errorInfo := field.Tag.Get("reg_error_info") //获取field对应的reg_error_info tag值
			return fieldName + ":" + errorInfo           //返回错误
		} else {
			return "缺失reg_error_info"
		}
	}
	return ""
}
