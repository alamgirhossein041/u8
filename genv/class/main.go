package main

import (
	"fmt"
	"log"

	"github.com/dop251/goja"
)

func main() {
	vm := goja.New()
	vm.Set(`println`, func(args ...interface{}) {
		fmt.Println(args...)
	})
	v, e := vm.RunString(`
function __extends(d, b) {
    Object.setPrototypeOf(d.prototype, b.prototype);
}
__extends;
`)
	if e != nil {
		log.Fatalln(e)
	}
	var jsExtends func(d, b *goja.Object)
	e = vm.ExportTo(v, &jsExtends)
	if e != nil {
		log.Fatalln(e)
	}
	createChildCtor := func(superObj *goja.Object, factory func(func(goja.FunctionCall) goja.Value) func(call goja.ConstructorCall) *goja.Object) *goja.Object {
		var super func(goja.FunctionCall) goja.Value
		vm.ExportTo(superObj, &super)
		ctor := vm.ToValue(factory(super)).(*goja.Object)
		jsExtends(ctor, superObj)
		return ctor
	}

	animal := vm.ToValue(func(call goja.ConstructorCall) *goja.Object {
		obj := call.This
		obj.Set(`name`, call.Argument(0).String())
		obj.Set(`eat`, func(call goja.FunctionCall) goja.Value {
			self := call.This.(*goja.Object)
			fmt.Println(self.Get(`name`), `eat`)
			return nil
		})
		return nil
	}).(*goja.Object)
	vm.Set(`Animal`, animal)
	vm.Set("Cat", createChildCtor(animal, func(super func(goja.FunctionCall) goja.Value) func(call goja.ConstructorCall) *goja.Object {
		return func(call goja.ConstructorCall) *goja.Object {
			self := call.This

			// call super()
			v := super(goja.FunctionCall{
				This:      self,
				Arguments: call.Arguments,
			})
			if o, ok := v.(*goja.Object); ok {
				self = o
			}

			// add subclass method
			self.Set(`speak`, func(call goja.FunctionCall) goja.Value {
				self := call.This.(*goja.Object)
				return vm.ToValue(self.Get(`name`).String() + ` speak`)
			})
			return self
		}
	}))

	_, e = vm.RunScript(`main.js`, `

var cat = new Cat('cat')
println(cat instanceof Animal)  // true
println(cat instanceof Cat)  // true
cat.eat()
cat.speak()
    `)
	if e != nil {
		log.Fatalln(e)
	}
}
