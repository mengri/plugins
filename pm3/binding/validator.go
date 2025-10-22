package binding

import "github.com/gin-gonic/gin/binding"

type replaceValidator struct {
	binding.StructValidator
}

var (
	sourceValidator binding.StructValidator
)

func init() {
	sourceValidator = binding.Validator
	binding.Validator = &replaceValidator{sourceValidator}
}
func (v *replaceValidator) ValidateStruct(obj any) error {
	return nil
}
func (v *replaceValidator) Engine() any {
	return v.StructValidator.Engine()
}
