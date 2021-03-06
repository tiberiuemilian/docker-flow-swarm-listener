package service

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TypesTestSuite struct {
	suite.Suite
}

func TestTypesUnitTestSuite(t *testing.T) {
	suite.Run(t, new(TypesTestSuite))
}

func (s *TypesTestSuite) Test_EqualMapStringString_SameKeys_DifferentValue() {
	a := map[string]string{"k1": "v1", "k2": "v2"}
	b := map[string]string{"k1": "v2", "k2": "v2"}

	s.False(EqualMapStringString(a, b))
	s.False(EqualMapStringString(b, a))
}

func (s *TypesTestSuite) Test_EqualMapStringString_SameKeys_SameValue() {

	a := map[string]string{"k1": "v1", "k2": "v2"}
	b := map[string]string{"k1": "v1", "k2": "v2"}

	s.True(EqualMapStringString(a, b))
	s.True(EqualMapStringString(b, a))
}

func (s *TypesTestSuite) Test_EqualMapStringString_DifferentKeys() {

	a := map[string]string{"k1": "v1", "k2": "v2"}
	b := map[string]string{"k1": "v1", "k3": "v2"}

	s.False(EqualMapStringString(a, b))
	s.False(EqualMapStringString(b, a))
}

func (s *TypesTestSuite) Test_EqualMapStringString_DifferentNumberOfValues() {

	a := map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"}
	b := map[string]string{"k1": "v1", "k2": "v2"}

	s.False(EqualMapStringString(a, b))
	s.False(EqualMapStringString(b, a))
}

func (s *TypesTestSuite) Test_Cardinality_DifferentElms() {
	a := NodeIPSet{}
	a.Add("node-1", "1.0.0.1", "id1")
	a.Add("node-2", "1.0.1.1", "id2")
	s.Equal(2, a.Cardinality())
}

func (s *TypesTestSuite) Test_Cardinality_RepeatElems() {
	a := NodeIPSet{}
	a.Add("node-1", "1.0.0.1", "id1")
	a.Add("node-1", "1.0.0.1", "id1")
	s.Equal(1, a.Cardinality())
}

func (s *TypesTestSuite) Test_NodeIPSetEqual_RepeatElems() {
	a := NodeIPSet{}
	a.Add("node-1", "1.0.0.1", "id1")
	a.Add("node-1", "1.0.0.1", "id1")
	b := NodeIPSet{}
	b.Add("node-1", "1.0.0.1", "id1")
	s.True(EqualNodeIPSet(a, b))
}

func (s *TypesTestSuite) Test_NodeIPSetEqual_LenUnequal() {
	a := NodeIPSet{}
	a.Add("node-1", "1.0.0.1", "id1")
	a.Add("node-2", "1.0.1.1", "id2")
	b := NodeIPSet{}
	b.Add("node-1", "1.0.0.1", "id1")
	b.Add("node-2", "1.0.1.1", "id2")
	b.Add("node-2", "1.0.1.2", "id2")
	s.False(EqualNodeIPSet(a, b))
}

func (s *TypesTestSuite) Test_NodeIPSetEqual_EqualSets() {
	a := NodeIPSet{}
	a.Add("node-1", "1.0.0.1", "id1")
	a.Add("node-2", "1.0.1.1", "id2")
	b := NodeIPSet{}
	b.Add("node-1", "1.0.0.1", "id1")
	b.Add("node-2", "1.0.1.1", "id2")
	s.True(EqualNodeIPSet(a, b))
}

func (s *TypesTestSuite) Test_NodeIPSetEqual_AddrNotEqual() {
	a := NodeIPSet{}
	a.Add("node-1", "1.0.0.1", "id1")
	a.Add("node-2", "1.0.1.1", "id2")
	b := NodeIPSet{}
	b.Add("node-1", "1.0.0.1", "id1")
	b.Add("node-2", "1.0.1.2", "id2")
	s.False(EqualNodeIPSet(a, b))
}

func (s *TypesTestSuite) Test_NodeIPSetEqual_NodeNameNotEqual() {
	a := NodeIPSet{}
	a.Add("node-1", "1.0.0.1", "id1")
	a.Add("node-2", "1.0.1.1", "id2")
	b := NodeIPSet{}
	b.Add("node-1", "1.0.0.1", "id1")
	b.Add("node-1", "1.0.1.1", "id1")
	s.False(EqualNodeIPSet(a, b))
}

func (s *TypesTestSuite) Test_NodeIPSetEqual_EmptySets() {
	a := NodeIPSet{}
	b := NodeIPSet{}
	s.True(EqualNodeIPSet(a, b))
}

func (s *TypesTestSuite) Test_NodeIPSetEqual_OneEmpty() {
	a := NodeIPSet{}
	b := NodeIPSet{}
	b.Add("node-1", "1.0.0.1", "id1")
	b.Add("node-1", "1.0.1.1", "id1")
	s.False(EqualNodeIPSet(a, b))
}

func (s *TypesTestSuite) Test_NodeIPMarshallJSON_EmptySet() {
	a := NodeIPSet{}
	b, err := json.Marshal(a)
	s.Require().NoError(err)

	s.Equal([]byte("[]"), b)
}

func (s *TypesTestSuite) Test_NodeIP_JSONCycle() {
	a := NodeIPSet{}
	a.Add("node-1", "1.0.0.1", "id1")
	a.Add("node-2", "1.0.1.1", "id2")
	by, err := json.Marshal(a)
	s.Require().NoError(err)

	i := NodeIPSet{}
	err = json.Unmarshal(by, &i)
	s.Require().NoError(err)

	s.True(EqualNodeIPSet(a, i))
}

func (s *TypesTestSuite) Test_NodeIPSet_Add() {
	a := NodeIPSet{}
	a.Add("node-1", "1.0.0.1", "id1")
	a.Add("node-1", "1.0.1.1", "id1")
	b := NodeIPSet{}
	b.Add("node-1", "1.0.0.1", "id1")
	b.Add("node-1", "1.0.1.1", "id1")

	s.True(EqualNodeIPSet(a, b))
}

func (s *TypesTestSuite) Test_NodeIP_WithTwoEntries() {
	nodeBytes := []byte("[[\"node-1\", \"1.0.0.1\"], [\"node-2\", \"1.0.1.1\"]]")
	ipSet := NodeIPSet{}
	err := json.Unmarshal(nodeBytes, &ipSet)
	s.Require().NoError(err)

	s.Contains(ipSet, NodeIP{Name: "node-1", Addr: "1.0.0.1"})
	s.Contains(ipSet, NodeIP{Name: "node-2", Addr: "1.0.1.1"})
}
