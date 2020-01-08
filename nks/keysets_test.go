package nks

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

var keysetTestID int
var keysetTest = Keyset{
	Name:       "Test Go SDK " + GetTicks(),
	Category:   "user_ssh",
	Workspaces: []int{},
	IsDefault:  false,
	Keys:       []Key{},
}

func TestKeysets(t *testing.T) {
	defer gock.Off()
	t.Run("get keysets", func(t *testing.T) {
		t.Run("create keyset", testCreateKeyset)
		t.Run("keysets", testGetKeysets)
		t.Run("keyset", testGetKeyset)
		t.Run("delete keyset", testDeleteKeyset)
	})
}

func testCreateKeyset(t *testing.T) {
	c, err := NewTestClientFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	idRsaPubPath, err := GetValueFromEnv("NKS_ID_RSA_PUB_PATH")
	if err != nil {
		t.Error(err)
	}

	idRsaPubPath, err = GetAbsPath(idRsaPubPath)
	if err != nil {
		t.Error(err)
	}

	content, err := ioutil.ReadFile(idRsaPubPath)
	if err != nil {
		t.Error(err)
	}

	testKeyset.Keys = append(testKeyset.Keys, Key{
		Type:  "pub",
		Value: string(content),
	})

	keyset, err := c.CreateKeyset(orgID, testKeyset)
	if err != nil {
		t.Error(err)
	}

	testKeysetLiveID = keyset.ID

	assert.Equal(t, testKeyset.Name, keyset.Name, "Name should be equal")
	assert.NotNil(t, len(testKeyset.Keys), 1, "One key should be present")
	assert.Equal(t, testKeyset.Keys[0].Type, "pub", "A key should be pub")
}

func testGetKeysets(t *testing.T) {
	fmt.Println("GetKeysets testing")
	c, err := NewTestClientFromEnv()
	if err != nil {
		t.Error(err)
	}

	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	keys, err := c.GetKeysets(orgID)
	if err != nil {
		t.Error(err)
	}
	if len(keys) == 0 {
		fmt.Println("No keysets found, but no error")
	}
}

func testGetKeyset(t *testing.T) {
	fmt.Println("GetKeyset testing")
	c, err := NewTestClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	keyset, err := c.GetKeyset(orgID, testKeysetLiveID)
	if err != nil {
		t.Error(err)
	}

	if keyset == nil {
		fmt.Println("No keyset found, but no error")
	}

	assert.Equal(t, keysetTest.Name, keyset.Name, "Name should be equal")
}

func testDeleteKeyset(t *testing.T) {
	c, err := NewTestClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	err = c.DeleteKeyset(orgID, testKeysetLiveID)
	if err != nil {
		t.Error(err)
	}
}
