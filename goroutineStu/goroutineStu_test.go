package goroutineStu

import (
	"fmt"
	"testing"
)

func TestStudyChannelOne(t *testing.T) {
	StudyChannelOne()
}

func TestStart(t *testing.T) {
	Start()
}

func TestStudyChannelTwo(t *testing.T) {
	StudyChannelTwo()
}

func TestForRangeChannel(t *testing.T) {
	ForRangeChannel()
}

func TestWriteReadData(t *testing.T) {
	WriteReadData()
}

func TestSelectTest(t *testing.T) {
	SelectTest()
}

func TestSyntaxStu(t *testing.T) {
	SyntaxStu()
}

func TestSyntaxOne(t *testing.T) {
	SyntaxOne()
}

func TestWaitGroupStu(t *testing.T) {
	WaitGroupStu()
}

func TestSelectStudyTwo(t *testing.T) {
	SelectStudyTwo()
}

func TestRandomGenerator(t *testing.T) {
	generator := RandomGenerator()
	fmt.Println(<-generator)
	fmt.Println(<-generator, <-generator)
}

func TestRandomGenerator2(t *testing.T) {
	RandomGenerator2()
}

func TestMultiRoutineGenerator(t *testing.T) {
	generator := MultiRoutineGenerator()
	fmt.Println(<-generator, <-generator)
}