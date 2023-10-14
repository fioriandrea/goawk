// Code generated by "stringer -type=Opcode,AugOp,BuiltinOp"; DO NOT EDIT.

package compiler

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Nop-0]
	_ = x[Num-1]
	_ = x[Str-2]
	_ = x[Dupe-3]
	_ = x[Drop-4]
	_ = x[Swap-5]
	_ = x[Roll-6]
	_ = x[Field-7]
	_ = x[FieldInt-8]
	_ = x[FieldByName-9]
	_ = x[FieldByNameStr-10]
	_ = x[Global-11]
	_ = x[Local-12]
	_ = x[Special-13]
	_ = x[ArrayGlobal-14]
	_ = x[ArrayLocal-15]
	_ = x[InGlobal-16]
	_ = x[InLocal-17]
	_ = x[AssignField-18]
	_ = x[AssignGlobal-19]
	_ = x[AssignLocal-20]
	_ = x[AssignSpecial-21]
	_ = x[AssignArrayGlobal-22]
	_ = x[AssignArrayLocal-23]
	_ = x[Delete-24]
	_ = x[DeleteAll-25]
	_ = x[IncrField-26]
	_ = x[IncrGlobal-27]
	_ = x[IncrLocal-28]
	_ = x[IncrSpecial-29]
	_ = x[IncrArrayGlobal-30]
	_ = x[IncrArrayLocal-31]
	_ = x[AugAssignField-32]
	_ = x[AugAssignGlobal-33]
	_ = x[AugAssignLocal-34]
	_ = x[AugAssignSpecial-35]
	_ = x[AugAssignArrayGlobal-36]
	_ = x[AugAssignArrayLocal-37]
	_ = x[Regex-38]
	_ = x[IndexMulti-39]
	_ = x[ConcatMulti-40]
	_ = x[Add-41]
	_ = x[Subtract-42]
	_ = x[Multiply-43]
	_ = x[Divide-44]
	_ = x[Power-45]
	_ = x[Modulo-46]
	_ = x[Equals-47]
	_ = x[NotEquals-48]
	_ = x[Less-49]
	_ = x[Greater-50]
	_ = x[LessOrEqual-51]
	_ = x[GreaterOrEqual-52]
	_ = x[Concat-53]
	_ = x[Match-54]
	_ = x[NotMatch-55]
	_ = x[Not-56]
	_ = x[UnaryMinus-57]
	_ = x[UnaryPlus-58]
	_ = x[Boolean-59]
	_ = x[Jump-60]
	_ = x[JumpFalse-61]
	_ = x[JumpTrue-62]
	_ = x[JumpEquals-63]
	_ = x[JumpNotEquals-64]
	_ = x[JumpLess-65]
	_ = x[JumpGreater-66]
	_ = x[JumpLessOrEqual-67]
	_ = x[JumpGreaterOrEqual-68]
	_ = x[Next-69]
	_ = x[Nextfile-70]
	_ = x[Exit-71]
	_ = x[ExitStatus-72]
	_ = x[ForIn-73]
	_ = x[BreakForIn-74]
	_ = x[CallBuiltin-75]
	_ = x[CallLengthArray-76]
	_ = x[CallSplit-77]
	_ = x[CallSplitSep-78]
	_ = x[CallSprintf-79]
	_ = x[CallUser-80]
	_ = x[CallNative-81]
	_ = x[Return-82]
	_ = x[ReturnNull-83]
	_ = x[Nulls-84]
	_ = x[Print-85]
	_ = x[Printf-86]
	_ = x[Getline-87]
	_ = x[GetlineField-88]
	_ = x[GetlineGlobal-89]
	_ = x[GetlineLocal-90]
	_ = x[GetlineSpecial-91]
	_ = x[GetlineArray-92]
	_ = x[EndOpcode-93]
}

const _Opcode_name = "NopNumStrDupeDropSwapRollFieldFieldIntFieldByNameFieldByNameStrGlobalLocalSpecialArrayGlobalArrayLocalInGlobalInLocalAssignFieldAssignGlobalAssignLocalAssignSpecialAssignArrayGlobalAssignArrayLocalDeleteDeleteAllIncrFieldIncrGlobalIncrLocalIncrSpecialIncrArrayGlobalIncrArrayLocalAugAssignFieldAugAssignGlobalAugAssignLocalAugAssignSpecialAugAssignArrayGlobalAugAssignArrayLocalRegexIndexMultiConcatMultiAddSubtractMultiplyDividePowerModuloEqualsNotEqualsLessGreaterLessOrEqualGreaterOrEqualConcatMatchNotMatchNotUnaryMinusUnaryPlusBooleanJumpJumpFalseJumpTrueJumpEqualsJumpNotEqualsJumpLessJumpGreaterJumpLessOrEqualJumpGreaterOrEqualNextNextfileExitExitStatusForInBreakForInCallBuiltinCallLengthArrayCallSplitCallSplitSepCallSprintfCallUserCallNativeReturnReturnNullNullsPrintPrintfGetlineGetlineFieldGetlineGlobalGetlineLocalGetlineSpecialGetlineArrayEndOpcode"

var _Opcode_index = [...]uint16{0, 3, 6, 9, 13, 17, 21, 25, 30, 38, 49, 63, 69, 74, 81, 92, 102, 110, 117, 128, 140, 151, 164, 181, 197, 203, 212, 221, 231, 240, 251, 266, 280, 294, 309, 323, 339, 359, 378, 383, 393, 404, 407, 415, 423, 429, 434, 440, 446, 455, 459, 466, 477, 491, 497, 502, 510, 513, 523, 532, 539, 543, 552, 560, 570, 583, 591, 602, 617, 635, 639, 647, 651, 661, 666, 676, 687, 702, 711, 723, 734, 742, 752, 758, 768, 773, 778, 784, 791, 803, 816, 828, 842, 854, 863}

func (i Opcode) String() string {
	if i < 0 || i >= Opcode(len(_Opcode_index)-1) {
		return "Opcode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Opcode_name[_Opcode_index[i]:_Opcode_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AugOpAdd-0]
	_ = x[AugOpSub-1]
	_ = x[AugOpMul-2]
	_ = x[AugOpDiv-3]
	_ = x[AugOpPow-4]
	_ = x[AugOpMod-5]
}

const _AugOp_name = "AugOpAddAugOpSubAugOpMulAugOpDivAugOpPowAugOpMod"

var _AugOp_index = [...]uint8{0, 8, 16, 24, 32, 40, 48}

func (i AugOp) String() string {
	if i < 0 || i >= AugOp(len(_AugOp_index)-1) {
		return "AugOp(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AugOp_name[_AugOp_index[i]:_AugOp_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BuiltinAtan2-0]
	_ = x[BuiltinClose-1]
	_ = x[BuiltinCos-2]
	_ = x[BuiltinExp-3]
	_ = x[BuiltinFflush-4]
	_ = x[BuiltinFflushAll-5]
	_ = x[BuiltinGsub-6]
	_ = x[BuiltinIndex-7]
	_ = x[BuiltinInt-8]
	_ = x[BuiltinLength-9]
	_ = x[BuiltinLengthArg-10]
	_ = x[BuiltinLog-11]
	_ = x[BuiltinMatch-12]
	_ = x[BuiltinRand-13]
	_ = x[BuiltinSin-14]
	_ = x[BuiltinSqrt-15]
	_ = x[BuiltinSrand-16]
	_ = x[BuiltinSrandSeed-17]
	_ = x[BuiltinSub-18]
	_ = x[BuiltinSubstr-19]
	_ = x[BuiltinSubstrLength-20]
	_ = x[BuiltinSystem-21]
	_ = x[BuiltinTolower-22]
	_ = x[BuiltinToupper-23]
}

const _BuiltinOp_name = "BuiltinAtan2BuiltinCloseBuiltinCosBuiltinExpBuiltinFflushBuiltinFflushAllBuiltinGsubBuiltinIndexBuiltinIntBuiltinLengthBuiltinLengthArgBuiltinLogBuiltinMatchBuiltinRandBuiltinSinBuiltinSqrtBuiltinSrandBuiltinSrandSeedBuiltinSubBuiltinSubstrBuiltinSubstrLengthBuiltinSystemBuiltinTolowerBuiltinToupper"

var _BuiltinOp_index = [...]uint16{0, 12, 24, 34, 44, 57, 73, 84, 96, 106, 119, 135, 145, 157, 168, 178, 189, 201, 217, 227, 240, 259, 272, 286, 300}

func (i BuiltinOp) String() string {
	if i < 0 || i >= BuiltinOp(len(_BuiltinOp_index)-1) {
		return "BuiltinOp(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BuiltinOp_name[_BuiltinOp_index[i]:_BuiltinOp_index[i+1]]
}
