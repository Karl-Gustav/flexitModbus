package main

func GetAllHoldingRegisters() map[string]FlexitRegister {
	return map[string]FlexitRegister{
		"SupplyAirSpeed1": FlexitRegister{
			Address:     0,
			Name:        "SupplyAirSpeed1",
			Format:      "INT 16",
			Description: "Speed 1 (sa)",
			Write:       true,
			Unit:        "%",
			Default:     50,
			Min:         20,
			Max:         100,
		},

		"SupplyAirSpeed2": FlexitRegister{
			Address:     1,
			Name:        "SupplyAirSpeed2",
			Format:      "INT 16",
			Description: "Speed 2 (sa)",
			Write:       true,
			Unit:        "%",
			Default:     75,
			Min:         20,
			Max:         100,
		},

		"SupplyAirSpeed3": FlexitRegister{
			Address:     2,
			Name:        "SupplyAirSpeed3",
			Format:      "INT 16",
			Description: "Speed 3 (sa)",
			Write:       true,
			Unit:        "%",
			Default:     100,
			Min:         20,
			Max:         100,
		},

		"SupplyAirSpeed4": FlexitRegister{
			Address:     3,
			Name:        "SupplyAirSpeed4",
			Format:      "INT 16",
			Description: "Speed 4 (sa)",
			Write:       true,
			Unit:        "",
			Default:     3,
			Min:         1,
			Max:         3,
		},

		"ExtractAirSpeed1": FlexitRegister{
			Address:     4,
			Name:        "ExtractAirSpeed1",
			Format:      "INT 16",
			Description: "Speed 1 (ea)",
			Write:       true,
			Unit:        "%",
			Default:     50,
			Min:         20,
			Max:         100,
		},

		"ExtractAirSpeed2": FlexitRegister{
			Address:     5,
			Name:        "ExtractAirSpeed2",
			Format:      "INT 16",
			Description: "Speed 2 (ea)",
			Write:       true,
			Unit:        "%",
			Default:     75,
			Min:         20,
			Max:         100,
		},

		"ExtractAirSpeed3": FlexitRegister{
			Address:     6,
			Name:        "ExtractAirSpeed3",
			Format:      "INT 16",
			Description: "Speed 3 (ea)",
			Write:       true,
			Unit:        "%",
			Default:     100,
			Min:         20,
			Max:         100,
		},

		"ExtractAirSpeed4": FlexitRegister{
			Address:     7,
			Name:        "ExtractAirSpeed4",
			Format:      "INT 16",
			Description: "Speed 4 (ea)",
			Write:       true,
			Unit:        "",
			Default:     1,
			Min:         1,
			Max:         3,
		},

		"SetAirTemperature": FlexitRegister{
			Address:     8,
			Name:        "SetAirTemperature",
			Format:      "INT 16",
			Description: "Air temperature",
			Write:       true,
			Unit:        "0.1°C",
			Default:     200,
			Min:         100,
			Max:         300,
		},

		"SupplyAirMinTemp": FlexitRegister{
			Address:     9,
			Name:        "SupplyAirMinTemp",
			Format:      "INT 16",
			Description: "Min supply air temp.",
			Write:       true,
			Unit:        "0.1°C",
			Default:     160,
			Min:         50,
			Max:         250,
		},

		"SupplyAirMaxTemp": FlexitRegister{
			Address:     10,
			Name:        "SupplyAirMaxTemp",
			Format:      "INT 16",
			Description: "Max supply air temp.",
			Write:       true,
			Unit:        "0.1°C",
			Default:     350,
			Min:         150,
			Max:         450,
		},

		"CoolingOutdoorAirMinTemp": FlexitRegister{
			Address:     11,
			Name:        "CoolingOutdoorAirMinTemp",
			Format:      "INT 16",
			Description: "CO, Min outdoor temp.",
			Write:       true,
			Unit:        "0.1°C",
			Default:     170,
			Min:         50,
			Max:         250,
		},

		"ForcedVentSpeed": FlexitRegister{
			Address:     12,
			Name:        "ForcedVentSpeed",
			Format:      "INT 16",
			Description: "Speed to set during forced ventilation (max timer)",
			Write:       true,
			Unit:        "",
			Default:     3,
			Min:         1,
			Max:         3,
		},

		"ForcedVentTime": FlexitRegister{
			Address:     13,
			Name:        "ForcedVentTime",
			Format:      "INT 16",
			Description: "Forced ventilation (max timer) activation time",
			Write:       true,
			Unit:        "min",
			Default:     30,
			Min:         0,
			Max:         360,
		},

		"AirRegulationType": FlexitRegister{
			Address:     14,
			Name:        "AirRegulationType",
			Format:      "INT 16",
			Description: "Air regulation type",
			Write:       true,
			Unit:        "bool",
			Default:     1,
			Min:         0,
			Max:         1,
		},

		"CoolingActive": FlexitRegister{
			Address:     15,
			Name:        "CoolingActive",
			Format:      "INT 16",
			Description: "Cooling (CO)",
			Write:       true,
			Unit:        "bool",
			Default:     0,
			Min:         0,
			Max:         1,
		},

		"ForcedVentilation": FlexitRegister{
			Address:     16,
			Name:        "ForcedVentilation",
			Format:      "INT 16",
			Description: "Forced ventilation activate/deactivate",
			Write:       true,
			Unit:        "bool",
			Default:     0,
			Min:         0,
			Max:         1,
		},

		"SetAirSpeed": FlexitRegister{
			Address:     17,
			Name:        "SetAirSpeed",
			Format:      "INT 16",
			Description: "A set value that panels use to write wanted air speed to CU.",
			Write:       true,
			Unit:        "",
			Default:     2,
			Min:         0,
			Max:         3,
		},

		"Time": FlexitRegister{
			Address:     18,
			Name:        "Time",
			Format:      "UINT 32",
			Description: "Real time clock value",
			Write:       false,
			Unit:        "s",
			Default:     0,
			Min:         0,
			Max:         0xffffffff,
		},

		"FireSmokeMode": FlexitRegister{
			Address:     21,
			Name:        "FireSmokeMode",
			Format:      "INT 16",
			Description: "Fire/Smoke mode",
			Write:       true,
			Unit:        "",
			Default:     1,
			Min:         1,
			Max:         4,
		},
	}
}
