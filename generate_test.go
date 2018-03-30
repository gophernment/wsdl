package wsdl

import (
	"fmt"
	"testing"
)

type Elastic struct{}

func (Elastic) Location() string {
	return "http://localhost:1323/elastic"
}
func (Elastic) OperationName() string {
	return "Elastic"
}

func (Elastic) InputType() Type {
	return ElasticInput{}
}
func (Elastic) OutputType() Type {
	return ElasticOutput{}
}
func (Elastic) ErrorType() Type {
	return ElasticError{}
}

type ElasticInput struct{}

func (ElasticInput) MessageName() string {
	return "ElasticInput"
}
func (ElasticInput) TypeName() string {
	return "Elastic"
}
func (ElasticInput) SingleFields() []string {
	return []string{"ID", "RowID", "CustNo", "SubrNo", "ListName"}
}

type ElasticOutput struct{}

func (ElasticOutput) MessageName() string {
	return "ElasticOutput"
}
func (ElasticOutput) TypeName() string {
	return "ElasticResponse"
}
func (ElasticOutput) SingleFields() []string {
	return []string{"Index", "Type", "ID", "Version", "Created"}
}

type ElasticError struct{}

func (ElasticError) MessageName() string {
	return "ElasticError"
}
func (ElasticError) TypeName() string {
	return "ElasticFault"
}
func (ElasticError) SingleFields() []string {
	return []string{"En", "Th", "Code"}
}

func TestGenEasyWSDLFromXMLString(t *testing.T) {
	oper := NewOperation(Elastic{})
	wsdlString, err := WSDL(oper)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(wsdlString)

}
