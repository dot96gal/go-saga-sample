package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dot96gal/go-saga-sample/saga"
)

type HogeState interface {
	GetHoge() string
	SetHoge(value string)
}

var _ saga.Step = (*HogeStep)(nil)

type HogeStep struct {
	state HogeState
}

func (s *HogeStep) Invoke(ctx context.Context) error {
	fmt.Println("invoke hoge")
	s.state.SetHoge("hoge")

	return nil
}

func (s *HogeStep) Compensate(ctx context.Context) error {
	fmt.Println("compensate hoge")
	return nil
}

type FugaState interface {
	GetFuga() string
	SetFuga(value string)
}

var _ saga.Step = (*FugaStep)(nil)

type FugaStep struct {
	state FugaState
}

func (s *FugaStep) Invoke(ctx context.Context) error {
	fmt.Println("invoke fuga")
	s.state.SetFuga("fuga")

	return nil
}

func (s *FugaStep) Compensate(ctx context.Context) error {
	fmt.Println("compensate fuga")
	return nil
}

var _ HogeState = (*MyState)(nil)
var _ FugaState = (*MyState)(nil)

type MyState struct {
	saga.State
}

func NewMyState() MyState {
	return MyState{saga.NewState()}
}

func (s MyState) GetHoge() string {
	v, err := s.Get("hoge")
	if err != nil {
		panic(err)
	}

	return v.(string)
}

func (s MyState) SetHoge(value string) {
	s.Set("hoge", value)
}

func (s MyState) GetFuga() string {
	v, err := s.Get("fuga")
	if err != nil {
		panic(err)
	}

	return v.(string)
}

func (s MyState) SetFuga(value string) {
	s.Set("fuga", value)
}

func main() {
	orchestrator := saga.NewOrchestrator()

	ctx := context.Background()
	state := NewMyState()

	hogeStep := HogeStep{state: state}
	orchestrator.AddStep(
		&hogeStep,
	)

	fugaStep := FugaStep{state: state}
	orchestrator.AddStep(
		&fugaStep,
	)

	err := orchestrator.Run(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(state.GetHoge())
	fmt.Println(state.GetFuga())
}
