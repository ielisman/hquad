package main

import (
	"fmt"
	"biostar/service/wiegand"
)


func testWiegand(deviceID uint32) error {
	fmt.Printf("\n===== Wiegand Config Test =====\n\n")

	origConfig, err := wiegandSvc.GetConfig(deviceID)
	if err != nil {
		return err
	}

	fmt.Printf("Original Config: %v\n", origConfig)

	test26bit(deviceID)
	test37bit(deviceID)

	err = wiegandSvc.SetConfig(deviceID, origConfig)
	if err != nil {
		return err
	}
	
	return nil
}

const (
	WIEGAND_MODE = wiegand.WiegandMode_WIEGAND_OUT_ONLY
	OUT_PULSE_WIDTH = 40
	OUT_PULSE_INTERVAL = 10000

	WIEGAND_26BIT_LENGTH = 26
  WIEGAND_26BIT_NUM_OF_FIELD = 4
  WIEGAND_26BIT_EVEN_PARITY_POS = 0
	WIEGAND_26BIT_ODD_PARITY_POS = 25
)

// 26 bit standard
// FC: 01 1111 1110 0000 0000 0000 0000 : 0x01fe0000
// ID: 00 0000 0001 1111 1111 1111 1110 : 0x0001fffe
// EP: 01 1111 1111 1110 0000 0000 0000 : 0x01ffe000, Pos 0, Type: Even
// OP: 00 0000 0000 0001 1111 1111 1110 : 0x00001ffe, Pos 25, Type: Odd

func test26bit(deviceID uint32) error {
	var bitArray_26bit = [][]byte {
		{0, 1, /**/ 1, 1, 1, 1, 1, 1, 1, 0, /**/ 0, 0, 0, 0, 0, 0, 0, 0, /**/ 0, 0, 0, 0, 0, 0, 0, 0}, // Facility Code
		{0, 0, /**/ 0, 0, 0, 0, 0, 0, 0, 1, /**/ 1, 1, 1, 1, 1, 1, 1, 1, /**/ 1, 1, 1, 1, 1, 1, 1, 0}, // ID
		{0, 1, /**/ 1, 1, 1, 1, 1, 1, 1, 1, /**/ 1, 1, 1, 0, 0, 0, 0, 0, /**/ 0, 0, 0, 0, 0, 0, 0, 0}, // Even Parity
		{0, 0, /**/ 0, 0, 0, 0, 0, 0, 0, 0, /**/ 0, 0, 0, 1, 1, 1, 1, 1, /**/ 1, 1, 1, 1, 1, 1, 1, 0}, // Odd Parity
	}
	
	std26bitFormat := &wiegand.WiegandFormat{
		Length: WIEGAND_26BIT_LENGTH,
		IDFields: [][]byte{
			bitArray_26bit[0],
			bitArray_26bit[1],
		},
		ParityFields: []*wiegand.ParityField{
			&wiegand.ParityField{
				ParityPos: WIEGAND_26BIT_EVEN_PARITY_POS,
				ParityType: wiegand.WiegandParity_WIEGAND_PARITY_EVEN,
				Data: bitArray_26bit[2],
			},
			&wiegand.ParityField{
				ParityPos: WIEGAND_26BIT_ODD_PARITY_POS,
				ParityType: wiegand.WiegandParity_WIEGAND_PARITY_ODD,
				Data: bitArray_26bit[3],
			},
		},
	}

	std26bitConfig := &wiegand.WiegandConfig{
		Mode: WIEGAND_MODE,
		OutPulseWidth: OUT_PULSE_WIDTH,
		OutPulseInterval: OUT_PULSE_INTERVAL,
		Formats: []*wiegand.WiegandFormat{
			std26bitFormat,
		},
	}

	err := wiegandSvc.SetConfig(deviceID, std26bitConfig)
	if err != nil {
		return err
	}

	config, err := wiegandSvc.GetConfig(deviceID)
	if err != nil {
		return err
	}

	fmt.Printf("\n>>> Wiegand Config with Standard 26bit Format: %v\n", config)
	return nil
}


const (
	WIEGAND_37BIT_LENGTH = 37
	WIEGAND_37BIT_NUM_OF_FIELD = 4
	WIEGAND_37BIT_EVEN_PARITY_POS = 0
	WIEGAND_37BIT_ODD_PARITY_POS = 36
)


// 37 bit HID
// FC: 0 1111 1111 1111 1111 0000 0000 0000 0000 0000 : 0x0ffff00000
// ID: 0 0000 0000 0000 0000 1111 1111 1111 1111 1110 : 0x00000ffffe
// EP: 0 1111 1111 1111 1111 1100 0000 0000 0000 0000 : 0x0ffffc0000, Pos 0, Type: Even
// OP: 0 0000 0000 0000 0000 0111 1111 1111 1111 1110 : 0x000007fffe, Pos 36, Type: Odd

func test37bit(deviceID uint32) error {
	var bitArray_37bit = [][]byte {
		{0, 1, 1, 1, 1, /**/ 1, 1, 1, 1, 1, 1, 1, 1, /**/ 1, 1, 1, 1, 0, 0, 0, 0, /**/ 0, 0, 0, 0, 0, 0, 0, 0, /**/ 0, 0, 0, 0, 0, 0, 0, 0}, // Facility Code
		{0, 0, 0, 0, 0, /**/ 0, 0 ,0 ,0, 0, 0, 0, 0, /**/ 0, 0, 0, 0, 1, 1, 1, 1, /**/ 1, 1, 1, 1, 1, 1, 1, 1, /**/ 1, 1, 1, 1, 1, 1, 1, 0}, // ID
		{0, 1, 1, 1, 1, /**/ 1, 1, 1, 1, 1, 1, 1, 1, /**/ 1, 1, 1, 1, 1, 1, 0, 0, /**/ 0, 0, 0, 0, 0, 0, 0, 0, /**/ 0, 0, 0, 0, 0, 0, 0, 0}, // Even Parity
		{0, 0, 0, 0, 0, /**/ 0, 0 ,0 ,0, 0, 0, 0, 0, /**/ 0, 0, 0, 0, 0, 1, 1, 1, /**/ 1, 1, 1, 1, 1, 1, 1, 1, /**/ 1, 1, 1, 1, 1, 1, 1, 0}, // Odd Parity
	}

	hid37bitFormat := &wiegand.WiegandFormat{
		Length: WIEGAND_37BIT_LENGTH,
		IDFields: [][]byte{
			bitArray_37bit[0],
			bitArray_37bit[1],
		},
		ParityFields: []*wiegand.ParityField{
			&wiegand.ParityField{
				ParityPos: WIEGAND_37BIT_EVEN_PARITY_POS,
				ParityType: wiegand.WiegandParity_WIEGAND_PARITY_EVEN,
				Data: bitArray_37bit[2],
			},
			&wiegand.ParityField{
				ParityPos: WIEGAND_37BIT_ODD_PARITY_POS,
				ParityType: wiegand.WiegandParity_WIEGAND_PARITY_ODD,
				Data: bitArray_37bit[3],
			},
		},
	}	

	hid37bitConfig := &wiegand.WiegandConfig{
		Mode: WIEGAND_MODE,
		OutPulseWidth: OUT_PULSE_WIDTH,
		OutPulseInterval: OUT_PULSE_INTERVAL,
		Formats: []*wiegand.WiegandFormat{
			hid37bitFormat,
		},
	}

	err := wiegandSvc.SetConfig(deviceID, hid37bitConfig)
	if err != nil {
		return err
	}

	config, err := wiegandSvc.GetConfig(deviceID)
	if err != nil {
		return err
	}

	fmt.Printf("\n>>> Wiegand Config with HID 37bit Format: %v\n", config)
	return nil
}


