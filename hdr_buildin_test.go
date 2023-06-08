package hdrhistogram

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "os"
    "testing"
)

func TestToBytes(t *testing.T) {
    h := New(1, 10000000, 3)
    for i := 0; i < 1000000; i++ {
        if err := h.RecordValue(int64(i)); err != nil {
            t.Fatal(err)
        }
    }

    countsBytes, err := h.fillBufferFromCountsArray()
    if err != nil {
        panic(err)
    }

    buffer := new(bytes.Buffer)
    err = binary.Write(buffer, binary.BigEndian, int32(len(countsBytes))) // 3-7
    if err != nil {
        panic(err)
    }
    err = binary.Write(buffer, binary.BigEndian, h.getNormalizingIndexOffset()) // 8-11
    if err != nil {
        panic(err)
    }
    // numberOfSignificantValueDigits
    err = binary.Write(buffer, binary.BigEndian, int32(h.significantFigures)) // 12-15
    if err != nil {
        panic(err)
    }
    err = binary.Write(buffer, binary.BigEndian, h.lowestDiscernibleValue) // 16-23
    if err != nil {
        panic(err)
    }
    err = binary.Write(buffer, binary.BigEndian, h.highestTrackableValue) // 24-31
    if err != nil {
        panic(err)
    }
    err = binary.Write(buffer, binary.BigEndian, h.getIntegerToDoubleValueConversionRatio()) // 32-39
    if err != nil {
        panic(err)
    }
    err = binary.Write(buffer, binary.BigEndian, countsBytes)
    if err != nil {
        panic(err)
    }

    b := buffer.Bytes()
    fmt.Println(len(countsBytes))
    fmt.Println(len(b))
    os.WriteFile("D:\\files\\hdr-go.bin", b, os.ModePerm)
}

func TestToBytesBuildBin(t *testing.T) {
    h := New(1, 10000000, 3)
    for i := 0; i < 1000000; i++ {
        if err := h.RecordValue(int64(i)); err != nil {
            t.Fatal(err)
        }
    }

    buffer, _ := h.encodeIntoByteBuffer()
    os.WriteFile("D:\\files\\hdr-go-buildin.bin", buffer.Bytes(), os.ModePerm)
    h.PercentilesPrint(os.Stdout, 5, 1)
}

func TestToCompressedBytesBuildBin(t *testing.T) {
    h := New(1, 10000000, 3)
    for i := 0; i < 1000000; i++ {
        if err := h.RecordValue(int64(i)); err != nil {
            t.Fatal(err)
        }
    }

    data, _ := h.dumpV2CompressedEncodingWithLevel(-1)
    os.WriteFile("D:\\files\\hdr-go-compress-buildin.bin", data, os.ModePerm)
    h.PercentilesPrint(os.Stdout, 5, 1)
}
