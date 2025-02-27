// @generated by Thrift for thrift/compiler/test/fixtures/includes/src/IncludesAlso.thrift
// This file is probably not the place you want to edit!

package IncludesAlso // thrift/compiler/test/fixtures/includes/src/IncludesAlso.thrift

import (
  "fmt"

  "github.com/facebook/fbthrift/thrift/lib/go/thrift"
)


// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = thrift.ZERO


type Also struct {
}
// Compile time interface enforcer
var _ thrift.Struct = &Also{}

func NewAlso() *Also {
    return (&Also{})
}

// Deprecated: Use Also.Set* methods instead or set the fields directly.
type AlsoBuilder struct {
    obj *Also
}

func NewAlsoBuilder() *AlsoBuilder {
    return &AlsoBuilder{
        obj: NewAlso(),
    }
}

func (x *AlsoBuilder) Emit() *Also {
    var objCopy Also = *x.obj
    return &objCopy
}

func (x *Also) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("Also"); err != nil {
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

func (x *Also) Read(p thrift.Protocol) error {
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
