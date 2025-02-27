// @generated by Thrift for thrift/compiler/test/fixtures/serialization_field_order/src/module.thrift
// This file is probably not the place you want to edit!

package module // thrift/compiler/test/fixtures/serialization_field_order/src/module.thrift

import (
  "fmt"

  thrift0 "thrift/annotation/thrift"
  "github.com/facebook/fbthrift/thrift/lib/go/thrift"
)

var _ = thrift0.GoUnusedProtection__

// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = thrift.ZERO


type Foo struct {
    Field1 int32 `thrift:"field1,3" json:"field1" db:"field1"`
    Field2 int32 `thrift:"field2,1" json:"field2" db:"field2"`
    Field3 int32 `thrift:"field3,2" json:"field3" db:"field3"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Foo{}

func NewFoo() *Foo {
    return (&Foo{})
}
func (x *Foo) GetField1() int32 {
    return x.Field1
}

func (x *Foo) GetField2() int32 {
    return x.Field2
}

func (x *Foo) GetField3() int32 {
    return x.Field3
}

func (x *Foo) SetField1(field1 int32) *Foo {
    x.Field1 = field1
    return x
}

func (x *Foo) SetField2(field2 int32) *Foo {
    x.Field2 = field2
    return x
}

func (x *Foo) SetField3(field3 int32) *Foo {
    x.Field3 = field3
    return x
}




func (x *Foo) writeField3(p thrift.Protocol) error {  // Field1
    if err := p.WriteFieldBegin("field1", thrift.I32, 3); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetField1()
    if err := p.WriteI32(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Foo) writeField1(p thrift.Protocol) error {  // Field2
    if err := p.WriteFieldBegin("field2", thrift.I32, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetField2()
    if err := p.WriteI32(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Foo) writeField2(p thrift.Protocol) error {  // Field3
    if err := p.WriteFieldBegin("field3", thrift.I32, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetField3()
    if err := p.WriteI32(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Foo) readField3(p thrift.Protocol) error {  // Field1
    result, err := p.ReadI32()
if err != nil {
    return err
}

    x.SetField1(result)
    return nil
}

func (x *Foo) readField1(p thrift.Protocol) error {  // Field2
    result, err := p.ReadI32()
if err != nil {
    return err
}

    x.SetField2(result)
    return nil
}

func (x *Foo) readField2(p thrift.Protocol) error {  // Field3
    result, err := p.ReadI32()
if err != nil {
    return err
}

    x.SetField3(result)
    return nil
}


// Deprecated: Use Foo.Set* methods instead or set the fields directly.
type FooBuilder struct {
    obj *Foo
}

func NewFooBuilder() *FooBuilder {
    return &FooBuilder{
        obj: NewFoo(),
    }
}

func (x *FooBuilder) Field1(field1 int32) *FooBuilder {
    x.obj.Field1 = field1
    return x
}

func (x *FooBuilder) Field2(field2 int32) *FooBuilder {
    x.obj.Field2 = field2
    return x
}

func (x *FooBuilder) Field3(field3 int32) *FooBuilder {
    x.obj.Field3 = field3
    return x
}

func (x *FooBuilder) Emit() *Foo {
    var objCopy Foo = *x.obj
    return &objCopy
}

func (x *Foo) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("Foo"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField3(p); err != nil {
        return err
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

func (x *Foo) Read(p thrift.Protocol) error {
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
        case 3:  // field1
            if err := x.readField3(p); err != nil {
                return err
            }
        case 1:  // field2
            if err := x.readField1(p); err != nil {
                return err
            }
        case 2:  // field3
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

type Foo2 struct {
    Field1 int32 `thrift:"field1,3" json:"field1" db:"field1"`
    Field2 int32 `thrift:"field2,1" json:"field2" db:"field2"`
    Field3 int32 `thrift:"field3,2" json:"field3" db:"field3"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Foo2{}

func NewFoo2() *Foo2 {
    return (&Foo2{})
}
func (x *Foo2) GetField1() int32 {
    return x.Field1
}

func (x *Foo2) GetField2() int32 {
    return x.Field2
}

func (x *Foo2) GetField3() int32 {
    return x.Field3
}

func (x *Foo2) SetField1(field1 int32) *Foo2 {
    x.Field1 = field1
    return x
}

func (x *Foo2) SetField2(field2 int32) *Foo2 {
    x.Field2 = field2
    return x
}

func (x *Foo2) SetField3(field3 int32) *Foo2 {
    x.Field3 = field3
    return x
}




func (x *Foo2) writeField3(p thrift.Protocol) error {  // Field1
    if err := p.WriteFieldBegin("field1", thrift.I32, 3); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetField1()
    if err := p.WriteI32(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Foo2) writeField1(p thrift.Protocol) error {  // Field2
    if err := p.WriteFieldBegin("field2", thrift.I32, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetField2()
    if err := p.WriteI32(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Foo2) writeField2(p thrift.Protocol) error {  // Field3
    if err := p.WriteFieldBegin("field3", thrift.I32, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetField3()
    if err := p.WriteI32(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Foo2) readField3(p thrift.Protocol) error {  // Field1
    result, err := p.ReadI32()
if err != nil {
    return err
}

    x.SetField1(result)
    return nil
}

func (x *Foo2) readField1(p thrift.Protocol) error {  // Field2
    result, err := p.ReadI32()
if err != nil {
    return err
}

    x.SetField2(result)
    return nil
}

func (x *Foo2) readField2(p thrift.Protocol) error {  // Field3
    result, err := p.ReadI32()
if err != nil {
    return err
}

    x.SetField3(result)
    return nil
}


// Deprecated: Use Foo2.Set* methods instead or set the fields directly.
type Foo2Builder struct {
    obj *Foo2
}

func NewFoo2Builder() *Foo2Builder {
    return &Foo2Builder{
        obj: NewFoo2(),
    }
}

func (x *Foo2Builder) Field1(field1 int32) *Foo2Builder {
    x.obj.Field1 = field1
    return x
}

func (x *Foo2Builder) Field2(field2 int32) *Foo2Builder {
    x.obj.Field2 = field2
    return x
}

func (x *Foo2Builder) Field3(field3 int32) *Foo2Builder {
    x.obj.Field3 = field3
    return x
}

func (x *Foo2Builder) Emit() *Foo2 {
    var objCopy Foo2 = *x.obj
    return &objCopy
}

func (x *Foo2) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("Foo2"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField3(p); err != nil {
        return err
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

func (x *Foo2) Read(p thrift.Protocol) error {
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
        case 3:  // field1
            if err := x.readField3(p); err != nil {
                return err
            }
        case 1:  // field2
            if err := x.readField1(p); err != nil {
                return err
            }
        case 2:  // field3
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
