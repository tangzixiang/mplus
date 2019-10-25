package util

// IsSystem64Bit 检查当前系统是否 64 位
func IsSystem64Bit() bool {
	return SystemBit() == 64
}

// IsSystem32Bit 检查当前系统是否 32 位
func IsSystem32Bit() bool {
	return SystemBit() == 32
}

// SystemBit 计算系统位数, 32 位系统返回 32，64 位系统返回 64
func SystemBit() int {
	return 32 << (^uint(0) >> 32 & 1)
}
