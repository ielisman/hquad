package flow

import (
	"biostar/service/connect"
	"fmt"
	"os"
)

type EnrollErrorType uint8

const (
	RETRIEVE_IMAGE  EnrollErrorType = 1
	NORMALIZE_IMAGE EnrollErrorType = 2
	ENROLL_USER     EnrollErrorType = 3
)

type EnrollError struct {
	ErrorType    EnrollErrorType
	ErrorMessage string
}

func ErrorFlowStandard(err error) bool {
	if err != nil {
		fmt.Printf(" %v\n", err)
		return false
	} else {
		fmt.Println(" Ok")
	}
	return true
}

func ErrorFlowStandardNone(err error) bool {
	if err != nil {
		fmt.Printf(" %v\n", err)
		return false
	}
	return true
}

func ErrorFlowStandardExit(err error) {
	if err != nil {
		fmt.Printf(" %v\n", err)
		os.Exit(0)
	} else {
		fmt.Println(" Ok")
	}
}

func ErrorFlowStandardMessage(err error, message string) bool {
	if err != nil {
		fmt.Printf(" %v\n", err)
		return false
	} else {
		fmt.Println(message)
	}
	return true
}

func IfErrorExitElseOkId(err error, closeConn func(uint32) error, device uint32) error { //func IfErrorExitElseOk(err error, closeConn func(uint32) error, deviceId uint32) error {
	if err != nil {
		fmt.Printf(" %v\n", err)
		if closeConn != nil {
			closeErr := closeConn(device)
			if closeErr == nil {
				fmt.Println(" Ok")
			}
		}
		os.Exit(0)
	} else {
		fmt.Println(" Ok")
	}

	return nil
}

func IfErrorExitElseOk(err error, closeConn func(*connect.SearchDeviceInfo) error, device *connect.SearchDeviceInfo) error { //func IfErrorExitElseOk(err error, closeConn func(uint32) error, deviceId uint32) error {
	if err != nil {
		fmt.Printf(" %v\n", err)
		if closeConn != nil {
			closeErr := closeConn(device)
			if closeErr == nil {
				fmt.Println(" Ok")
			}
		}
		os.Exit(0)
	} else {
		fmt.Println(" Ok")
	}

	return nil
}

func IfErrorExitElseMessage(err error, message string, closeConn func(*connect.SearchDeviceInfo) error, device *connect.SearchDeviceInfo) error {
	if err != nil {
		fmt.Printf(" %v\n", err)
		if closeConn != nil {
			closeErr := closeConn(device)
			if closeErr == nil {
				fmt.Println(" Ok")
			}
		}
		os.Exit(0)
	} else {
		fmt.Println(message)
	}

	return nil
}

func EnrollErrorString(value EnrollErrorType) string {
	switch value {
	case RETRIEVE_IMAGE:
		return "RETRIEVE_IMAGE"
	case NORMALIZE_IMAGE:
		return "NORMALIZE_IMAGE"
	case ENROLL_USER:
		return "ENROLL_USER"
	default:
		return ""
	}
}
