// func consecutively_add(data *[25]uint64)
TEXT ·consecutively_add(SB),$0-24
	MOVD	data+0(FP), R0

	VLD1.P	32(R0), [V0.D2, V1.D2]
	VLD1.P	32(R0), [V2.D2, V3.D2]
	VLD1.P	32(R0), [V4.D2, V5.D2]
	VLD1.P	32(R0), [V6.D2, V7.D2]
	VLD1.P	32(R0), [V8.D2, V9.D2]
	VLD1.P	32(R0), [V10.D2, V11.D2]
	VLD1.P	32(R0), [V12.D2, V13.D2]
	VLD1.P	32(R0), [V14.D2, V15.D2]
	VLD1.P	32(R0), [V16.D2, V17.D2]
	VLD1.P	32(R0), [V18.D2, V19.D2]
	VLD1.P	32(R0), [V20.D2, V21.D2]
	VLD1.P	32(R0), [V22.D2, V23.D2]
	VLD1 (R0), [V24.D2]

	SUB	$384, R0, R0 // set R0 to the beginning of array "data"

	VADD	V0.D2, V1.D2, V0.D2
	VADD	V2.D2, V3.D2, V2.D2
	VADD	V4.D2, V5.D2, V4.D2
	VADD	V6.D2, V7.D2, V6.D2
	VADD	V8.D2, V9.D2, V8.D2
	VADD	V10.D2, V11.D2, V10.D2
	VADD	V12.D2, V13.D2, V12.D2
	VADD	V14.D2, V15.D2, V14.D2
	VADD	V16.D2, V17.D2, V16.D2
	VADD	V18.D2, V19.D2, V18.D2
	VADD	V20.D2, V21.D2, V20.D2
	VADD	V22.D2, V23.D2, V22.D2
	
	VST1.P	[V0.D2, V1.D2], 32(R0)
	VST1.P	[V2.D2, V3.D2], 32(R0)
	VST1.P	[V4.D2, V5.D2], 32(R0)
	VST1.P	[V6.D2, V7.D2], 32(R0)
	VST1.P	[V8.D2, V9.D2], 32(R0)
	VST1.P	[V10.D2, V11.D2], 32(R0)
	VST1.P	[V12.D2, V13.D2], 32(R0)
	VST1.P	[V14.D2, V15.D2], 32(R0)
	VST1.P	[V16.D2, V17.D2], 32(R0)
	VST1.P	[V18.D2, V19.D2], 32(R0)
	VST1.P	[V20.D2, V21.D2], 32(R0)
	VST1.P	[V22.D2, V23.D2], 32(R0)
	VST1	[V24.D2], (R0)

	RET
