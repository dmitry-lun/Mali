package main

import (
	"crypto/rand"
	"encoding/binary"
	"os"
)

func createSafePE(filename string) {
	file, _ := os.Create(filename)
	defer file.Close()

	dosHeader := make([]byte, 64)
	dosHeader[0] = 'M'
	dosHeader[1] = 'Z'
	binary.LittleEndian.PutUint32(dosHeader[0x3C:0x40], 0x40)
	file.Write(dosHeader)

	peSig := []byte{'P', 'E', 0, 0}
	file.Write(peSig)

	coffHeader := make([]byte, 20)
	binary.LittleEndian.PutUint16(coffHeader[0:2], 0x014C)
	binary.LittleEndian.PutUint16(coffHeader[2:4], 2)
	binary.LittleEndian.PutUint16(coffHeader[16:18], 0xE0)
	file.Write(coffHeader)

	optHeader := make([]byte, 224)
	binary.LittleEndian.PutUint16(optHeader[0:2], 0x10B)
	binary.LittleEndian.PutUint32(optHeader[16:20], 0x1000)
	binary.LittleEndian.PutUint32(optHeader[28:32], 0x400000)
	binary.LittleEndian.PutUint32(optHeader[32:36], 0x1000)
	binary.LittleEndian.PutUint32(optHeader[36:40], 0x200)
	binary.LittleEndian.PutUint32(optHeader[56:60], 0x3000)
	binary.LittleEndian.PutUint32(optHeader[60:64], 0x400)
	file.Write(optHeader)

	section1 := make([]byte, 40)
	copy(section1[0:8], []byte(".text\x00\x00\x00"))
	binary.LittleEndian.PutUint32(section1[8:12], 0x1000)
	binary.LittleEndian.PutUint32(section1[12:16], 0x1000)
	binary.LittleEndian.PutUint32(section1[16:20], 0x200)
	binary.LittleEndian.PutUint32(section1[20:24], 0x400)
	binary.LittleEndian.PutUint32(section1[36:40], 0x60000020)
	file.Write(section1)

	section2 := make([]byte, 40)
	copy(section2[0:8], []byte(".data\x00\x00\x00"))
	binary.LittleEndian.PutUint32(section2[8:12], 0x500)
	binary.LittleEndian.PutUint32(section2[12:16], 0x2000)
	binary.LittleEndian.PutUint32(section2[16:20], 0x200)
	binary.LittleEndian.PutUint32(section2[20:24], 0x600)
	binary.LittleEndian.PutUint32(section2[36:40], 0xC0000040)
	file.Write(section2)

	file.Write(make([]byte, 0x200))
	file.Write(make([]byte, 0x200))
	file.Sync()
}

func createLowRiskPE(filename string) {
	file, _ := os.Create(filename)
	defer file.Close()

	dosHeader := make([]byte, 64)
	dosHeader[0] = 'M'
	dosHeader[1] = 'Z'
	binary.LittleEndian.PutUint32(dosHeader[0x3C:0x40], 0x40)
	file.Write(dosHeader)

	peSig := []byte{'P', 'E', 0, 0}
	file.Write(peSig)

	coffHeader := make([]byte, 20)
	binary.LittleEndian.PutUint16(coffHeader[0:2], 0x014C)
	binary.LittleEndian.PutUint16(coffHeader[2:4], 2)
	binary.LittleEndian.PutUint16(coffHeader[16:18], 0xE0)
	file.Write(coffHeader)

	optHeader := make([]byte, 224)
	binary.LittleEndian.PutUint16(optHeader[0:2], 0x10B)
	binary.LittleEndian.PutUint32(optHeader[16:20], 0x1000)
	binary.LittleEndian.PutUint32(optHeader[28:32], 0x400000)
	binary.LittleEndian.PutUint32(optHeader[32:36], 0x1000)
	binary.LittleEndian.PutUint32(optHeader[36:40], 0x200)
	binary.LittleEndian.PutUint32(optHeader[56:60], 0x3000)
	binary.LittleEndian.PutUint32(optHeader[60:64], 0x400)
	file.Write(optHeader)

	section1 := make([]byte, 40)
	copy(section1[0:8], []byte(".text\x00\x00\x00"))
	binary.LittleEndian.PutUint32(section1[8:12], 0x1000)
	binary.LittleEndian.PutUint32(section1[12:16], 0x1000)
	binary.LittleEndian.PutUint32(section1[16:20], 0x200)
	binary.LittleEndian.PutUint32(section1[20:24], 0x400)
	binary.LittleEndian.PutUint32(section1[36:40], 0x60000020)
	file.Write(section1)

	section2 := make([]byte, 40)
	copy(section2[0:8], []byte(".code\x00\x00\x00"))
	binary.LittleEndian.PutUint32(section2[8:12], 0x500)
	binary.LittleEndian.PutUint32(section2[12:16], 0x2000)
	binary.LittleEndian.PutUint32(section2[16:20], 0x200)
	binary.LittleEndian.PutUint32(section2[20:24], 0x600)
	binary.LittleEndian.PutUint32(section2[36:40], 0x60000020)
	file.Write(section2)

	file.Write(make([]byte, 0x200))
	mediumEntropy := make([]byte, 0x200)
	rand.Read(mediumEntropy)
	file.Write(mediumEntropy)
	file.Sync()
}

func createMediumRiskPE(filename string) {
	file, _ := os.Create(filename)
	defer file.Close()

	dosHeader := make([]byte, 64)
	dosHeader[0] = 'M'
	dosHeader[1] = 'Z'
	binary.LittleEndian.PutUint32(dosHeader[0x3C:0x40], 0x40)
	file.Write(dosHeader)

	peSig := []byte{'P', 'E', 0, 0}
	file.Write(peSig)

	coffHeader := make([]byte, 20)
	binary.LittleEndian.PutUint16(coffHeader[0:2], 0x014C)
	binary.LittleEndian.PutUint16(coffHeader[2:4], 3)
	binary.LittleEndian.PutUint16(coffHeader[16:18], 0xE0)
	file.Write(coffHeader)

	optHeader := make([]byte, 224)
	binary.LittleEndian.PutUint16(optHeader[0:2], 0x10B)
	binary.LittleEndian.PutUint32(optHeader[16:20], 0x1000)
	binary.LittleEndian.PutUint32(optHeader[28:32], 0x400000)
	binary.LittleEndian.PutUint32(optHeader[32:36], 0x1000)
	binary.LittleEndian.PutUint32(optHeader[36:40], 0x200)
	binary.LittleEndian.PutUint32(optHeader[56:60], 0x4000)
	binary.LittleEndian.PutUint32(optHeader[60:64], 0x400)
	file.Write(optHeader)

	section1 := make([]byte, 40)
	copy(section1[0:8], []byte(".text1\x00\x00"))
	binary.LittleEndian.PutUint32(section1[8:12], 0x1000)
	binary.LittleEndian.PutUint32(section1[12:16], 0x1000)
	binary.LittleEndian.PutUint32(section1[16:20], 0x200)
	binary.LittleEndian.PutUint32(section1[20:24], 0x400)
	binary.LittleEndian.PutUint32(section1[36:40], 0x60000020)
	file.Write(section1)

	section2 := make([]byte, 40)
	copy(section2[0:8], []byte(".data2\x00\x00"))
	binary.LittleEndian.PutUint32(section2[8:12], 0x500)
	binary.LittleEndian.PutUint32(section2[12:16], 0x2000)
	binary.LittleEndian.PutUint32(section2[16:20], 0x200)
	binary.LittleEndian.PutUint32(section2[20:24], 0x600)
	binary.LittleEndian.PutUint32(section2[36:40], 0xC0000040)
	file.Write(section2)

	section3 := make([]byte, 40)
	copy(section3[0:8], []byte(".code\x00\x00\x00"))
	binary.LittleEndian.PutUint32(section3[8:12], 0x500)
	binary.LittleEndian.PutUint32(section3[12:16], 0x3000)
	binary.LittleEndian.PutUint32(section3[16:20], 0x200)
	binary.LittleEndian.PutUint32(section3[20:24], 0x800)
	binary.LittleEndian.PutUint32(section3[36:40], 0x60000020)
	file.Write(section3)

	file.Write(make([]byte, 0x200))
	file.Write(make([]byte, 0x200))
	mediumEntropy := make([]byte, 0x400)
	for i := 0; i < len(mediumEntropy); i++ {
		if i%5 == 0 {
			mediumEntropy[i] = byte(i % 256)
		} else {
			mediumEntropy[i] = byte((i*13 + i%7) % 256)
		}
	}
	file.Write(mediumEntropy)
	file.Sync()
}

func createHighRiskPE(filename string) {
	file, _ := os.Create(filename)
	defer file.Close()

	dosHeader := make([]byte, 64)
	dosHeader[0] = 'M'
	dosHeader[1] = 'Z'
	binary.LittleEndian.PutUint32(dosHeader[0x3C:0x40], 0x40)
	file.Write(dosHeader)

	peSig := []byte{'P', 'E', 0, 0}
	file.Write(peSig)

	coffHeader := make([]byte, 20)
	binary.LittleEndian.PutUint16(coffHeader[0:2], 0x014C)
	binary.LittleEndian.PutUint16(coffHeader[2:4], 3)
	binary.LittleEndian.PutUint16(coffHeader[16:18], 0xE0)
	file.Write(coffHeader)

	optHeader := make([]byte, 224)
	binary.LittleEndian.PutUint16(optHeader[0:2], 0x10B)
	binary.LittleEndian.PutUint32(optHeader[16:20], 0x1000)
	binary.LittleEndian.PutUint32(optHeader[28:32], 0x400000)
	binary.LittleEndian.PutUint32(optHeader[32:36], 0x1000)
	binary.LittleEndian.PutUint32(optHeader[36:40], 0x200)
	binary.LittleEndian.PutUint32(optHeader[56:60], 0x5000)
	binary.LittleEndian.PutUint32(optHeader[60:64], 0x400)
	file.Write(optHeader)

	section1 := make([]byte, 40)
	copy(section1[0:8], []byte("UPX0\x00\x00\x00\x00"))
	binary.LittleEndian.PutUint32(section1[8:12], 0x2000)
	binary.LittleEndian.PutUint32(section1[12:16], 0x1000)
	binary.LittleEndian.PutUint32(section1[16:20], 0x0)
	binary.LittleEndian.PutUint32(section1[20:24], 0x400)
	binary.LittleEndian.PutUint32(section1[36:40], 0xE0000020)
	file.Write(section1)

	section2 := make([]byte, 40)
	copy(section2[0:8], []byte("UPX1\x00\x00\x00\x00"))
	binary.LittleEndian.PutUint32(section2[8:12], 0x1000)
	binary.LittleEndian.PutUint32(section2[12:16], 0x3000)
	binary.LittleEndian.PutUint32(section2[16:20], 0x1000)
	binary.LittleEndian.PutUint32(section2[20:24], 0x600)
	binary.LittleEndian.PutUint32(section2[36:40], 0xE0000020)
	file.Write(section2)

	section3 := make([]byte, 40)
	copy(section3[0:8], []byte(".text1\x00\x00"))
	binary.LittleEndian.PutUint32(section3[8:12], 0x1000)
	binary.LittleEndian.PutUint32(section3[12:16], 0x4000)
	binary.LittleEndian.PutUint32(section3[16:20], 0x1000)
	binary.LittleEndian.PutUint32(section3[20:24], 0x1600)
	binary.LittleEndian.PutUint32(section3[36:40], 0x60000020)
	file.Write(section3)

	file.Write(make([]byte, 0x200))
	highEntropy1 := make([]byte, 0x1000)
	rand.Read(highEntropy1)
	file.Write(highEntropy1)

	highEntropy2 := make([]byte, 0x1000)
	rand.Read(highEntropy2)
	file.Write(highEntropy2)
	file.Sync()
}

func main() {
	os.MkdirAll("samples", 0755)

	createSafePE("samples/test_safe.exe")
	createLowRiskPE("samples/test_low.exe")
	createMediumRiskPE("samples/test_medium.exe")
	createHighRiskPE("samples/test_high.exe")

	println("✓ Created test PE files:")
	println("  - samples/test_safe.exe   (expected: SAFE)")
	println("  - samples/test_low.exe    (expected: LOW)")
	println("  - samples/test_medium.exe (expected: MEDIUM)")
	println("  - samples/test_high.exe   (expected: HIGH)")
}
