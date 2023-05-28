package cocuhbase

import (
	"github.com/gofiber/utils"
	"strings"
	"testing"
)

func TestGet_ShouldReturnNoError(t *testing.T) {
	testStorage := New(
		Config{
			Username: "admin",
			Password: "123456",
			Host:     "127.0.0.1:8091",
			Bucket:   "fiber_storage",
		})

	err := testStorage.Set("test", []byte("test"), 0)

	utils.AssertEqual(t, nil, err)

}

func TestSetAndGet_GetShouldReturn_SettedValueWithoutError(t *testing.T) {
	testStorage := New(
		Config{
			Username: "admin",
			Password: "123456",
			Host:     "127.0.0.1:8091",
			Bucket:   "fiber_storage",
		})

	err := testStorage.Set("test", []byte("fiber_test_value"), 0)

	utils.AssertEqual(t, nil, err)

	val, err := testStorage.Get("test")

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, val, []byte("fiber_test_value"))

}

func TestSetAndDelete_ShouldReturn_NoError(t *testing.T) {
	testStorage := New(
		Config{
			Username: "admin",
			Password: "123456",
			Host:     "127.0.0.1:8091",
			Bucket:   "fiber_storage",
		})

	err := testStorage.Set("test", []byte("fiber_test_value"), 0)
	utils.AssertEqual(t, nil, err)

	err = testStorage.Delete("test")
	utils.AssertEqual(t, nil, err)

	_, err = testStorage.Get("test")

	errStr := err.Error()
	utils.AssertEqual(t, true, strings.Contains(errStr, "document not found"))

}

func TestSetAndReset_ResetShouldReturn_NoError(t *testing.T) {
	testStorage := New(
		Config{
			Username: "admin",
			Password: "123456",
			Host:     "127.0.0.1:8091",
			Bucket:   "fiber_storage",
		})

	err := testStorage.Set("test", []byte("fiber_test_value"), 0)
	utils.AssertEqual(t, nil, err)

	err = testStorage.Reset()
	utils.AssertEqual(t, nil, err)

	_, err = testStorage.Get("test")

	errStr := err.Error()
	utils.AssertEqual(t, true, strings.Contains(errStr, "document not found"))

}

func TestClose_CloseShouldReturn_NoError(t *testing.T) {
	testStorage := New(
		Config{
			Username: "admin",
			Password: "123456",
			Host:     "127.0.0.1:8091",
			Bucket:   "fiber_storage",
		})

	err := testStorage.Close()
	utils.AssertEqual(t, nil, err)

}
