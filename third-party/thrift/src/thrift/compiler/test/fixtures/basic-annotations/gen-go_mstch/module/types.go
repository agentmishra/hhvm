// @generated by Thrift for thrift/compiler/test/fixtures/basic-annotations/src/module.thrift
// This file is probably not the place you want to edit!

package module // thrift/compiler/test/fixtures/basic-annotations/src/module.thrift

import (
  "fmt"

  cpp "thrift/annotation/cpp"
  "github.com/facebook/fbthrift/thrift/lib/go/thrift"
)

var _ = cpp.GoUnusedProtection__

// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = thrift.ZERO


type MyId = int16

type MyEnum int32

const (
    MyEnum_MyValue1 MyEnum = 0
    MyEnum_MyValue2 MyEnum = 1
    MyEnum_DOMAIN MyEnum = 2
)

// Enum value maps for MyEnum
var (
    MyEnum_name = map[MyEnum]string {
        MyEnum_MyValue1: "MyValue1",
        MyEnum_MyValue2: "MyValue2",
        MyEnum_DOMAIN: "DOMAIN",
    }

    MyEnum_value = map[string]MyEnum {
        "MyValue1": MyEnum_MyValue1,
        "MyValue2": MyEnum_MyValue2,
        "DOMAIN": MyEnum_DOMAIN,
    }

    // Deprecated: Use MyEnum_name instead.
    MyEnumToName = MyEnum_name

    // Deprecated: Use MyEnum_value instead.
    MyEnumToValue = MyEnum_value

    // Deprecated: Use MyEnum_name instead (e.g. `for name, _ := range MyEnum_name {}`).
    MyEnumNames = []string{
        "MyValue1",
        "MyValue2",
        "DOMAIN",
    }

    // Deprecated: Use MyEnum_value instead (e.g. `for value, _ := range MyEnum_value {}`).
    MyEnumValues = []MyEnum{
        0,
        1,
        2,
    }
)

func (x MyEnum) String() string {
    if v, ok := MyEnumToName[x]; ok {
        return v
    }
    return "<UNSET>"
}

func (x MyEnum) Ptr() *MyEnum {
    return &x
}

// Deprecated: Use MyEnum_value instead (e.g. `x, ok := MyEnum_value["name"]`).
func MyEnumFromString(s string) (MyEnum, error) {
    if v, ok := MyEnum_value[s]; ok {
        return v, nil
    }
    return MyEnum(0), fmt.Errorf("not a valid MyEnum string")
}

// Deprecated: Use MyEnum.Ptr() instead.
func MyEnumPtr(v MyEnum) *MyEnum {
    return &v
}


type MyStructNestedAnnotation struct {
    Name string `thrift:"name,1" json:"name" db:"name"`
}
// Compile time interface enforcer
var _ thrift.Struct = &MyStructNestedAnnotation{}

func NewMyStructNestedAnnotation() *MyStructNestedAnnotation {
    return (&MyStructNestedAnnotation{})
}
func (x *MyStructNestedAnnotation) GetName() string {
    return x.Name
}

func (x *MyStructNestedAnnotation) SetName(name string) *MyStructNestedAnnotation {
    x.Name = name
    return x
}


func (x *MyStructNestedAnnotation) writeField1(p thrift.Protocol) error {  // Name
    if err := p.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetName()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStructNestedAnnotation) readField1(p thrift.Protocol) error {  // Name
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetName(result)
    return nil
}


// Deprecated: Use MyStructNestedAnnotation.Set* methods instead or set the fields directly.
type MyStructNestedAnnotationBuilder struct {
    obj *MyStructNestedAnnotation
}

func NewMyStructNestedAnnotationBuilder() *MyStructNestedAnnotationBuilder {
    return &MyStructNestedAnnotationBuilder{
        obj: NewMyStructNestedAnnotation(),
    }
}

func (x *MyStructNestedAnnotationBuilder) Name(name string) *MyStructNestedAnnotationBuilder {
    x.obj.Name = name
    return x
}

func (x *MyStructNestedAnnotationBuilder) Emit() *MyStructNestedAnnotation {
    var objCopy MyStructNestedAnnotation = *x.obj
    return &objCopy
}

func (x *MyStructNestedAnnotation) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("MyStructNestedAnnotation"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *MyStructNestedAnnotation) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // name
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

type MyUnion struct {
}
// Compile time interface enforcer
var _ thrift.Struct = &MyUnion{}

func NewMyUnion() *MyUnion {
    return (&MyUnion{})
}


func (x *MyUnion) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("MyUnion"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *MyUnion) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

type MyException struct {
}
// Compile time interface enforcer
var _ thrift.Struct = &MyException{}

func NewMyException() *MyException {
    return (&MyException{})
}
func (x *MyException) String() string {
    return fmt.Sprintf("%+v", x)
}

func (x *MyException) Error() string {
    return x.String()
}


// Deprecated: Use MyException.Set* methods instead or set the fields directly.
type MyExceptionBuilder struct {
    obj *MyException
}

func NewMyExceptionBuilder() *MyExceptionBuilder {
    return &MyExceptionBuilder{
        obj: NewMyException(),
    }
}

func (x *MyExceptionBuilder) Emit() *MyException {
    var objCopy MyException = *x.obj
    return &objCopy
}

func (x *MyException) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("MyException"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *MyException) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

type MyStruct struct {
    Major int64 `thrift:"major,2" json:"major" db:"major"`
    Package string `thrift:"package,1" json:"package" db:"package"`
    AnnotationWithQuote string `thrift:"annotation_with_quote,3" tag:"somevalue"`
    Class string `thrift:"class_,4" json:"class_" db:"class_"`
    AnnotationWithTrailingComma string `thrift:"annotation_with_trailing_comma,5" json:"annotation_with_trailing_comma" db:"annotation_with_trailing_comma"`
    EmptyAnnotations string `thrift:"empty_annotations,6" json:"empty_annotations" db:"empty_annotations"`
    MyEnum MyEnum `thrift:"my_enum,7" json:"my_enum" db:"my_enum"`
    CppTypeAnnotation []string `thrift:"cpp_type_annotation,8" json:"cpp_type_annotation" db:"cpp_type_annotation"`
    MyUnion *MyUnion `thrift:"my_union,9" json:"my_union" db:"my_union"`
    MyId MyId `thrift:"my_id,10" json:"my_id" db:"my_id"`
}
// Compile time interface enforcer
var _ thrift.Struct = &MyStruct{}

func NewMyStruct() *MyStruct {
    return (&MyStruct{})
}
func (x *MyStruct) GetMajor() int64 {
    return x.Major
}

func (x *MyStruct) GetPackage() string {
    return x.Package
}

func (x *MyStruct) GetAnnotationWithQuote() string {
    return x.AnnotationWithQuote
}

func (x *MyStruct) GetClass() string {
    return x.Class
}

func (x *MyStruct) GetAnnotationWithTrailingComma() string {
    return x.AnnotationWithTrailingComma
}

func (x *MyStruct) GetEmptyAnnotations() string {
    return x.EmptyAnnotations
}

func (x *MyStruct) GetMyEnum() MyEnum {
    return x.MyEnum
}

func (x *MyStruct) GetCppTypeAnnotation() []string {
    return x.CppTypeAnnotation
}

func (x *MyStruct) GetMyUnion() *MyUnion {
    return x.MyUnion
}

func (x *MyStruct) GetMyId() MyId {
    return x.MyId
}

func (x *MyStruct) SetMajor(major int64) *MyStruct {
    x.Major = major
    return x
}

func (x *MyStruct) SetPackage(package_ string) *MyStruct {
    x.Package = package_
    return x
}

func (x *MyStruct) SetAnnotationWithQuote(annotationWithQuote string) *MyStruct {
    x.AnnotationWithQuote = annotationWithQuote
    return x
}

func (x *MyStruct) SetClass(class string) *MyStruct {
    x.Class = class
    return x
}

func (x *MyStruct) SetAnnotationWithTrailingComma(annotationWithTrailingComma string) *MyStruct {
    x.AnnotationWithTrailingComma = annotationWithTrailingComma
    return x
}

func (x *MyStruct) SetEmptyAnnotations(emptyAnnotations string) *MyStruct {
    x.EmptyAnnotations = emptyAnnotations
    return x
}

func (x *MyStruct) SetMyEnum(myEnum MyEnum) *MyStruct {
    x.MyEnum = myEnum
    return x
}

func (x *MyStruct) SetCppTypeAnnotation(cppTypeAnnotation []string) *MyStruct {
    x.CppTypeAnnotation = cppTypeAnnotation
    return x
}

func (x *MyStruct) SetMyUnion(myUnion MyUnion) *MyStruct {
    x.MyUnion = &myUnion
    return x
}

func (x *MyStruct) SetMyId(myId MyId) *MyStruct {
    x.MyId = myId
    return x
}








func (x *MyStruct) IsSetCppTypeAnnotation() bool {
    return x.CppTypeAnnotation != nil
}

func (x *MyStruct) IsSetMyUnion() bool {
    return x.MyUnion != nil
}


func (x *MyStruct) writeField2(p thrift.Protocol) error {  // Major
    if err := p.WriteFieldBegin("major", thrift.I64, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetMajor()
    if err := p.WriteI64(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField1(p thrift.Protocol) error {  // Package
    if err := p.WriteFieldBegin("package", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetPackage()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField3(p thrift.Protocol) error {  // AnnotationWithQuote
    if err := p.WriteFieldBegin("annotation_with_quote", thrift.STRING, 3); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetAnnotationWithQuote()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField4(p thrift.Protocol) error {  // Class
    if err := p.WriteFieldBegin("class_", thrift.STRING, 4); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetClass()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField5(p thrift.Protocol) error {  // AnnotationWithTrailingComma
    if err := p.WriteFieldBegin("annotation_with_trailing_comma", thrift.STRING, 5); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetAnnotationWithTrailingComma()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField6(p thrift.Protocol) error {  // EmptyAnnotations
    if err := p.WriteFieldBegin("empty_annotations", thrift.STRING, 6); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetEmptyAnnotations()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField7(p thrift.Protocol) error {  // MyEnum
    if err := p.WriteFieldBegin("my_enum", thrift.I32, 7); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetMyEnum()
    if err := p.WriteI32(int32(item)); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField8(p thrift.Protocol) error {  // CppTypeAnnotation
    if !x.IsSetCppTypeAnnotation() {
        return nil
    }

    if err := p.WriteFieldBegin("cpp_type_annotation", thrift.LIST, 8); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetCppTypeAnnotation()
    if err := p.WriteListBegin(thrift.STRING, len(item)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
}
for _, v := range item {
    {
        item := v
        if err := p.WriteString(item); err != nil {
    return err
}
    }
}
if err := p.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField9(p thrift.Protocol) error {  // MyUnion
    if !x.IsSetMyUnion() {
        return nil
    }

    if err := p.WriteFieldBegin("my_union", thrift.STRUCT, 9); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := *x.GetMyUnion()
    if err := item.Write(p); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) writeField10(p thrift.Protocol) error {  // MyId
    if err := p.WriteFieldBegin("my_id", thrift.I16, 10); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetMyId()
    if err := p.WriteI16(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) readField2(p thrift.Protocol) error {  // Major
    result, err := p.ReadI64()
if err != nil {
    return err
}

    x.SetMajor(result)
    return nil
}

func (x *MyStruct) readField1(p thrift.Protocol) error {  // Package
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetPackage(result)
    return nil
}

func (x *MyStruct) readField3(p thrift.Protocol) error {  // AnnotationWithQuote
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetAnnotationWithQuote(result)
    return nil
}

func (x *MyStruct) readField4(p thrift.Protocol) error {  // Class
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetClass(result)
    return nil
}

func (x *MyStruct) readField5(p thrift.Protocol) error {  // AnnotationWithTrailingComma
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetAnnotationWithTrailingComma(result)
    return nil
}

func (x *MyStruct) readField6(p thrift.Protocol) error {  // EmptyAnnotations
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetEmptyAnnotations(result)
    return nil
}

func (x *MyStruct) readField7(p thrift.Protocol) error {  // MyEnum
    enumResult, err := p.ReadI32()
if err != nil {
    return err
}
result := MyEnum(enumResult)

    x.SetMyEnum(result)
    return nil
}

func (x *MyStruct) readField8(p thrift.Protocol) error {  // CppTypeAnnotation
    _ /* elemType */, size, err := p.ReadListBegin()
if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
}

listResult := make([]string, 0, size)
for i := 0; i < size; i++ {
    var elem string
    {
        result, err := p.ReadString()
if err != nil {
    return err
}
        elem = result
    }
    listResult = append(listResult, elem)
}

if err := p.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
}
result := listResult

    x.SetCppTypeAnnotation(result)
    return nil
}

func (x *MyStruct) readField9(p thrift.Protocol) error {  // MyUnion
    result := *NewMyUnion()
err := result.Read(p)
if err != nil {
    return err
}

    x.SetMyUnion(result)
    return nil
}

func (x *MyStruct) readField10(p thrift.Protocol) error {  // MyId
    result, err := p.ReadI16()
if err != nil {
    return err
}

    x.SetMyId(result)
    return nil
}


// Deprecated: Use MyStruct.Set* methods instead or set the fields directly.
type MyStructBuilder struct {
    obj *MyStruct
}

func NewMyStructBuilder() *MyStructBuilder {
    return &MyStructBuilder{
        obj: NewMyStruct(),
    }
}

func (x *MyStructBuilder) Major(major int64) *MyStructBuilder {
    x.obj.Major = major
    return x
}

func (x *MyStructBuilder) Package(package_ string) *MyStructBuilder {
    x.obj.Package = package_
    return x
}

func (x *MyStructBuilder) AnnotationWithQuote(annotationWithQuote string) *MyStructBuilder {
    x.obj.AnnotationWithQuote = annotationWithQuote
    return x
}

func (x *MyStructBuilder) Class(class string) *MyStructBuilder {
    x.obj.Class = class
    return x
}

func (x *MyStructBuilder) AnnotationWithTrailingComma(annotationWithTrailingComma string) *MyStructBuilder {
    x.obj.AnnotationWithTrailingComma = annotationWithTrailingComma
    return x
}

func (x *MyStructBuilder) EmptyAnnotations(emptyAnnotations string) *MyStructBuilder {
    x.obj.EmptyAnnotations = emptyAnnotations
    return x
}

func (x *MyStructBuilder) MyEnum(myEnum MyEnum) *MyStructBuilder {
    x.obj.MyEnum = myEnum
    return x
}

func (x *MyStructBuilder) CppTypeAnnotation(cppTypeAnnotation []string) *MyStructBuilder {
    x.obj.CppTypeAnnotation = cppTypeAnnotation
    return x
}

func (x *MyStructBuilder) MyUnion(myUnion *MyUnion) *MyStructBuilder {
    x.obj.MyUnion = myUnion
    return x
}

func (x *MyStructBuilder) MyId(myId MyId) *MyStructBuilder {
    x.obj.MyId = myId
    return x
}

func (x *MyStructBuilder) Emit() *MyStruct {
    var objCopy MyStruct = *x.obj
    return &objCopy
}

func (x *MyStruct) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("MyStruct"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField2(p); err != nil {
        return err
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := x.writeField3(p); err != nil {
        return err
    }

    if err := x.writeField4(p); err != nil {
        return err
    }

    if err := x.writeField5(p); err != nil {
        return err
    }

    if err := x.writeField6(p); err != nil {
        return err
    }

    if err := x.writeField7(p); err != nil {
        return err
    }

    if err := x.writeField8(p); err != nil {
        return err
    }

    if err := x.writeField9(p); err != nil {
        return err
    }

    if err := x.writeField10(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *MyStruct) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 2:  // major
            if err := x.readField2(p); err != nil {
                return err
            }
        case 1:  // package
            if err := x.readField1(p); err != nil {
                return err
            }
        case 3:  // annotation_with_quote
            if err := x.readField3(p); err != nil {
                return err
            }
        case 4:  // class_
            if err := x.readField4(p); err != nil {
                return err
            }
        case 5:  // annotation_with_trailing_comma
            if err := x.readField5(p); err != nil {
                return err
            }
        case 6:  // empty_annotations
            if err := x.readField6(p); err != nil {
                return err
            }
        case 7:  // my_enum
            if err := x.readField7(p); err != nil {
                return err
            }
        case 8:  // cpp_type_annotation
            if err := x.readField8(p); err != nil {
                return err
            }
        case 9:  // my_union
            if err := x.readField9(p); err != nil {
                return err
            }
        case 10:  // my_id
            if err := x.readField10(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

type SecretStruct struct {
    Id int64 `thrift:"id,1" json:"id" db:"id"`
    Password string `thrift:"password,2" json:"password" db:"password"`
}
// Compile time interface enforcer
var _ thrift.Struct = &SecretStruct{}

func NewSecretStruct() *SecretStruct {
    return (&SecretStruct{})
}
func (x *SecretStruct) GetId() int64 {
    return x.Id
}

func (x *SecretStruct) GetPassword() string {
    return x.Password
}

func (x *SecretStruct) SetId(id int64) *SecretStruct {
    x.Id = id
    return x
}

func (x *SecretStruct) SetPassword(password string) *SecretStruct {
    x.Password = password
    return x
}



func (x *SecretStruct) writeField1(p thrift.Protocol) error {  // Id
    if err := p.WriteFieldBegin("id", thrift.I64, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetId()
    if err := p.WriteI64(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *SecretStruct) writeField2(p thrift.Protocol) error {  // Password
    if err := p.WriteFieldBegin("password", thrift.STRING, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetPassword()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *SecretStruct) readField1(p thrift.Protocol) error {  // Id
    result, err := p.ReadI64()
if err != nil {
    return err
}

    x.SetId(result)
    return nil
}

func (x *SecretStruct) readField2(p thrift.Protocol) error {  // Password
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetPassword(result)
    return nil
}


// Deprecated: Use SecretStruct.Set* methods instead or set the fields directly.
type SecretStructBuilder struct {
    obj *SecretStruct
}

func NewSecretStructBuilder() *SecretStructBuilder {
    return &SecretStructBuilder{
        obj: NewSecretStruct(),
    }
}

func (x *SecretStructBuilder) Id(id int64) *SecretStructBuilder {
    x.obj.Id = id
    return x
}

func (x *SecretStructBuilder) Password(password string) *SecretStructBuilder {
    x.obj.Password = password
    return x
}

func (x *SecretStructBuilder) Emit() *SecretStruct {
    var objCopy SecretStruct = *x.obj
    return &objCopy
}

func (x *SecretStruct) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("SecretStruct"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := x.writeField2(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *SecretStruct) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // id
            if err := x.readField1(p); err != nil {
                return err
            }
        case 2:  // password
            if err := x.readField2(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}
