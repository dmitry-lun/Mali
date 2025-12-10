package main

import (
	"encoding/binary"
	"os"
)

func main() {
	file, _ := os.Create("samples/minimal_pe.exe")
	defer file.Close()

	// DOS Header
	dosHeader := make([]byte, 64)
	dosHeader[0] = 'M'
	dosHeader[1] = 'Z'
	binary.LittleEndian.PutUint32(dosHeader[0x3C:0x40], 0x40)
	file.Write(dosHeader)

	// PE Signature
	peSig := []byte{'P', 'E', 0, 0}
	file.Write(peSig)

	// COFF Header (20 bytes)
	coffHeader := make([]byte, 20)
	binary.LittleEndian.PutUint16(coffHeader[0:2], 0x8664) // x64
	binary.LittleEndian.PutUint16(coffHeader[2:4], 1)      // 1 section
	binary.LittleEndian.PutUint16(coffHeader[16:18], 0xF0) // SizeOfOptionalHeader
	file.Write(coffHeader)

	// Optional Header (240 bytes for x64)
	optHeader := make([]byte, 240)
	binary.LittleEndian.PutUint16(optHeader[0:2], 0x20B)         // PE32+ magic
	binary.LittleEndian.PutUint32(optHeader[16:20], 0x1000)      // AddressOfEntryPoint
	binary.LittleEndian.PutUint64(optHeader[24:32], 0x140000000) // ImageBase
	binary.LittleEndian.PutUint32(optHeader[32:36], 0x1000)      // SectionAlignment
	binary.LittleEndian.PutUint32(optHeader[36:40], 0x200)       // FileAlignment
	binary.LittleEndian.PutUint32(optHeader[56:60], 0x2000)      // SizeOfImage
	binary.LittleEndian.PutUint32(optHeader[60:64], 0x400)       // SizeOfHeaders
	file.Write(optHeader)

	// Section Header (40 bytes)
	sectionHeader := make([]byte, 40)
	copy(sectionHeader[0:8], []byte(".text\x00\x00\x00"))
	binary.LittleEndian.PutUint32(sectionHeader[8:12], 0x1000)      // VirtualSize
	binary.LittleEndian.PutUint32(sectionHeader[12:16], 0x1000)     // VirtualAddress
	binary.LittleEndian.PutUint32(sectionHeader[16:20], 0x200)      // SizeOfRawData
	binary.LittleEndian.PutUint32(sectionHeader[20:24], 0x400)      // PointerToRawData
	binary.LittleEndian.PutUint32(sectionHeader[36:40], 0x60000020) // Characteristics (executable, readable)
	file.Write(sectionHeader)

	// Section data (padding to align)
	padding := make([]byte, 0x200)
	file.Write(padding)

	file.Sync()
}
