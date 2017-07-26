package main

import (
	"errors"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/gaia-adm/pumba/action"
	"github.com/gaia-adm/pumba/container"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli"
)

//---- MOCK: Chaos Iterface

type ChaosMock struct {
	mock.Mock
}

func (m *ChaosMock) StopContainers(c container.Client, n []string, p string, cmd interface{}) error {
	args := m.Called(c, n, p, cmd)
	return args.Error(0)
}

func (m *ChaosMock) KillContainers(c container.Client, n []string, p string, cmd interface{}) error {
	args := m.Called(c, n, p, cmd)
	return args.Error(0)
}

func (m *ChaosMock) RemoveContainers(c container.Client, n []string, p string, cmd interface{}) error {
	args := m.Called(c, n, p, cmd)
	return args.Error(0)
}

func (m *ChaosMock) PauseContainers(c container.Client, n []string, p string, cmd interface{}) error {
	args := m.Called(c, n, p, cmd)
	return args.Error(0)
}

func (m *ChaosMock) NetemDelayContainers(c container.Client, n []string, p string, cmd interface{}) error {
	args := m.Called(c, n, p, cmd)
	return args.Error(0)
}

func (m *ChaosMock) NetemLossRandomContainers(c container.Client, n []string, p string, cmd interface{}) error {
	args := m.Called(c, n, p, cmd)
	return args.Error(0)
}

func (m *ChaosMock) NetemLossStateContainers(c container.Client, n []string, p string, cmd interface{}) error {
	args := m.Called(c, n, p, cmd)
	return args.Error(0)
}

func (m *ChaosMock) NetemLossGEmodelContainers(c container.Client, n []string, p string, cmd interface{}) error {
	args := m.Called(c, n, p, cmd)
	return args.Error(0)
}

func (m *ChaosMock) NetemRateContainers(c container.Client, n []string, p string, cmd interface{}) error {
	args := m.Called(c, n, p, cmd)
	return args.Error(0)
}

//---- TESTS

type mainTestSuite struct {
	suite.Suite
}

func (s *mainTestSuite) SetupSuite() {
	gTestRun = true
}

func (s *mainTestSuite) TearDownSuite() {
}

func (s *mainTestSuite) SetupTest() {
}

func (s *mainTestSuite) TearDownTest() {
}

func (s *mainTestSuite) Test_main() {
	os.Args = []string{"pumba", "-v"}
	main()
}

func (s *mainTestSuite) Test_getNames() {
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.Parse([]string{"c1", "c2", "c3"})
	c := cli.NewContext(nil, globalSet, nil)
	names, pattern := getNamesOrPattern(c)
	assert.True(s.T(), len(names) == 3)
	assert.True(s.T(), pattern == "")
}

func (s *mainTestSuite) Test_getSingleName() {
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.Parse([]string{"single"})
	c := cli.NewContext(nil, globalSet, nil)
	names, pattern := getNamesOrPattern(c)
	assert.True(s.T(), len(names) == 1)
	assert.True(s.T(), names[0] == "single")
	assert.True(s.T(), pattern == "")
}

func (s *mainTestSuite) Test_getPattern() {
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.Parse([]string{"re2:^test"})
	c := cli.NewContext(nil, globalSet, nil)
	names, pattern := getNamesOrPattern(c)
	assert.True(s.T(), len(names) == 0)
	assert.True(s.T(), pattern == "^test")
}

func (s *mainTestSuite) Test_getPattern2() {
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.Parse([]string{"re2:(.+)test"})
	c := cli.NewContext(nil, globalSet, nil)
	names, pattern := getNamesOrPattern(c)
	assert.True(s.T(), len(names) == 0)
	assert.True(s.T(), pattern == "(.+)test")
}

func (s *mainTestSuite) Test_beforeCommand_NoInterval() {
	// prepare
	set := flag.NewFlagSet("test", 0)
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("test", "me", "doc")
	parseErr := set.Parse([]string{})
	globalCtx := cli.NewContext(nil, globalSet, nil)
	c := cli.NewContext(nil, set, globalCtx)
	// invoke command
	err := beforeCommand(c)
	// asserts
	assert.NoError(s.T(), parseErr)
	assert.NoError(s.T(), err)
}

func (s *mainTestSuite) Test_beforeCommand_BadInterval() {
	// prepare
	set := flag.NewFlagSet("test", 0)
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("interval", "BAD", "doc")
	parseErr := set.Parse([]string{})
	globalCtx := cli.NewContext(nil, globalSet, nil)
	c := cli.NewContext(nil, set, globalCtx)
	// invoke command
	err := beforeCommand(c)
	// asserts
	assert.NoError(s.T(), parseErr)
	assert.Error(s.T(), err)
	assert.EqualError(s.T(), err, "time: invalid duration BAD")
}

func (s *mainTestSuite) Test_beforeCommand_EmptyArgs() {
	// prepare
	set := flag.NewFlagSet("test", 0)
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("interval", "10s", "doc")
	parseErr := set.Parse([]string{})
	globalCtx := cli.NewContext(nil, globalSet, nil)
	c := cli.NewContext(nil, set, globalCtx)
	// invoke command
	err := beforeCommand(c)
	names, pattern := getNamesOrPattern(c)
	// asserts
	assert.NoError(s.T(), parseErr)
	assert.NoError(s.T(), err)
	assert.True(s.T(), len(names) == 0)
	assert.True(s.T(), pattern == "")
}

func (s *mainTestSuite) Test_beforeCommand_Re2Args() {
	// prepare
	set := flag.NewFlagSet("test", 0)
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("interval", "10s", "doc")
	parseErr := set.Parse([]string{"re2:^c"})
	globalCtx := cli.NewContext(nil, globalSet, nil)
	c := cli.NewContext(nil, set, globalCtx)
	// invoke command
	err := beforeCommand(c)
	names, pattern := getNamesOrPattern(c)
	// asserts
	assert.NoError(s.T(), parseErr)
	assert.NoError(s.T(), err)
	assert.True(s.T(), len(names) == 0)
	assert.True(s.T(), pattern == "^c")
}

func (s *mainTestSuite) Test_beforeCommand_2Args() {
	// prepare
	set := flag.NewFlagSet("test", 0)
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("interval", "10s", "doc")
	parseErr := set.Parse([]string{"c1", "c2"})
	globalCtx := cli.NewContext(nil, globalSet, nil)
	c := cli.NewContext(nil, set, globalCtx)
	// invoke command
	err := beforeCommand(c)
	names, pattern := getNamesOrPattern(c)
	// asserts
	assert.NoError(s.T(), parseErr)
	assert.NoError(s.T(), err)
	assert.True(s.T(), len(names) == 2)
	assert.True(s.T(), pattern == "")
}

func (s *mainTestSuite) Test_handleSignals() {
	gWG.Add(1)
	handleSignals()
	gWG.Done()
}

func (s *mainTestSuite) Test_killSucess() {
	// prepare
	set := flag.NewFlagSet("kill", 0)
	set.String("signal", "SIGTERM", "doc")
	c := cli.NewContext(nil, set, nil)
	// set interval to 1ms
	gInterval = 1 * time.Millisecond
	// setup mock
	chaosMock := &ChaosMock{}
	chaos = chaosMock
	command := action.CommandKill{
		Signal: "SIGTERM",
	}
	chaosMock.On("KillContainers", nil, []string{}, "", command).Return(nil)
	// invoke command
	err := kill(c)
	// asserts
	// (!)WAIT till called action is completed (Sleep > Timer), it's executed in separate go routine
	time.Sleep(2 * time.Millisecond)
	assert.NoError(s.T(), err)
	chaosMock.AssertExpectations(s.T())
}

func (s *mainTestSuite) Test_killBadSignal() {
	// prepare
	set := flag.NewFlagSet("kill", 0)
	set.String("signal", "UNKNOWN", "doc")
	c := cli.NewContext(nil, set, nil)
	// invoke command
	err := kill(c)
	// asserts
	assert.EqualError(s.T(), err, "Unexpected signal: UNKNOWN")
}

func (s *mainTestSuite) Test_killError() {
	// prepare
	set := flag.NewFlagSet("kill", 0)
	set.String("signal", "SIGTERM", "doc")
	c := cli.NewContext(nil, set, nil)
	// set interval to 1ms
	gInterval = 1 * time.Millisecond
	// setup mock
	chaosMock := &ChaosMock{}
	chaos = chaosMock
	command := action.CommandKill{
		Signal: "SIGTERM",
	}
	chaosMock.On("KillContainers", nil, []string{}, "", command).Return(errors.New("ERROR"))
	// invoke command
	err := kill(c)
	// asserts
	// (!)WAIT till called action is completed (Sleep > Timer), it's executed in separate go routine
	time.Sleep(2 * time.Millisecond)
	assert.NoError(s.T(), err)
	chaosMock.AssertExpectations(s.T())
}

func (s *mainTestSuite) Test_pauseSucess() {
	// prepare
	set := flag.NewFlagSet("pause", 0)
	set.String("duration", "10s", "doc")
	c := cli.NewContext(nil, set, nil)
	// set interval to 1ms
	gInterval = 1 * time.Millisecond
	// setup mock
	chaosMock := &ChaosMock{}
	chaos = chaosMock
	cmd := action.CommandPause{
		Duration: time.Duration(10 * time.Second),
		StopChan: gStopChan,
	}
	chaosMock.On("PauseContainers", nil, []string{}, "", cmd).Return(nil)
	// invoke command
	err := pause(c)
	// asserts
	// (!)WAIT till called action is completed (Sleep > Timer), it's executed in separate go routine
	time.Sleep(2 * time.Millisecond)
	assert.NoError(s.T(), err)
	chaosMock.AssertExpectations(s.T())
}

func (s *mainTestSuite) Test_pauseMissingDuraation() {
	// prepare
	set := flag.NewFlagSet("pause", 0)
	c := cli.NewContext(nil, set, nil)
	// set interval to 1ms
	gInterval = 1 * time.Millisecond
	// invoke command
	err := pause(c)
	// asserts
	assert.EqualError(s.T(), err, "Undefined duration interval")
}

func (s *mainTestSuite) Test_pauseBadDuraation() {
	// prepare
	set := flag.NewFlagSet("pause", 0)
	set.String("duration", "BAD", "doc")
	c := cli.NewContext(nil, set, nil)
	// set interval to 1ms
	gInterval = 1 * time.Millisecond
	// invoke command
	err := pause(c)
	// asserts
	assert.EqualError(s.T(), err, "time: invalid duration BAD")
}

func (s *mainTestSuite) Test_stopSucess() {
	// prepare
	set := flag.NewFlagSet("stop", 0)
	set.Int("time", 10, "doc")
	c := cli.NewContext(nil, set, nil)
	// set interval to 1ms
	gInterval = 1 * time.Millisecond
	// setup mock
	cmd := action.CommandStop{WaitTime: 10}
	chaosMock := &ChaosMock{}
	chaos = chaosMock
	chaosMock.On("StopContainers", nil, []string{}, "", cmd).Return(nil)
	// invoke command
	err := stop(c)
	// asserts
	// (!)WAIT till called action is completed (Sleep > Timer), it's executed in separate go routine
	time.Sleep(2 * time.Millisecond)
	assert.NoError(s.T(), err)
	chaosMock.AssertExpectations(s.T())
}

func (s *mainTestSuite) Test_stopError() {
	// prepare
	set := flag.NewFlagSet("stop", 0)
	set.Int("time", 10, "doc")
	c := cli.NewContext(nil, set, nil)
	// set interval to 1ms
	gInterval = 1 * time.Millisecond
	// setup mock
	cmd := action.CommandStop{WaitTime: 10}
	chaosMock := &ChaosMock{}
	chaos = chaosMock
	chaosMock.On("StopContainers", nil, []string{}, "", cmd).Return(errors.New("ERROR"))
	// invoke command
	err := stop(c)
	// asserts
	// (!)WAIT till called action is completed (Sleep > Timer), it's executed in separate go routine
	time.Sleep(2 * time.Millisecond)
	assert.NoError(s.T(), err)
	chaosMock.AssertExpectations(s.T())
}

func (s *mainTestSuite) Test_removeSucess() {
	// prepare
	set := flag.NewFlagSet("stop", 0)
	set.Bool("force", true, "doc")
	set.Bool("links", true, "doc")
	set.Bool("volumes", true, "doc")
	c := cli.NewContext(nil, set, nil)
	// set interval to 1ms
	gInterval = 1 * time.Millisecond
	// setup mock
	cmd := action.CommandRemove{Force: true, Links: true, Volumes: true}
	chaosMock := &ChaosMock{}
	chaos = chaosMock
	chaosMock.On("RemoveContainers", nil, []string{}, "", cmd).Return(nil)
	// invoke command
	err := remove(c)
	// asserts
	// (!)WAIT till called action is completed (Sleep > Timer), it's executed in separate go routine
	time.Sleep(2 * time.Millisecond)
	assert.NoError(s.T(), err)
	chaosMock.AssertExpectations(s.T())
}

func (s *mainTestSuite) Test_netemDelaySucess() {
	// prepare test data
	// netem flags
	netemSet := flag.NewFlagSet("netem", 0)
	netemSet.String("duration", "10ms", "doc")
	netemSet.String("interface", "test0", "doc")
	netemCtx := cli.NewContext(nil, netemSet, nil)
	// delay flags
	delaySet := flag.NewFlagSet("delay", 0)
	delaySet.Int("time", 200, "doc")
	delaySet.Int("jitter", 20, "doc")
	delaySet.Float64("correlation", 1.5, "doc")
	delaySet.String("distribution", "normal", "doc")
	delaySet.Parse([]string{"c1", "c2", "c3"})
	delayCtx := cli.NewContext(nil, delaySet, netemCtx)
	// set interval to 1ms
	gInterval = 20 * time.Millisecond
	// setup mock
	cmd := action.CommandNetemDelay{
		NetInterface: "test0",
		Duration:     10 * time.Millisecond,
		Time:         200,
		Jitter:       20,
		Correlation:  1.5,
		Distribution: "normal",
		StopChan:     gStopChan,
	}
	chaosMock := &ChaosMock{}
	chaos = chaosMock
	chaosMock.On("NetemDelayContainers", nil, []string{"c1", "c2", "c3"}, "", cmd).Return(nil)
	// invoke command
	err := netemDelay(delayCtx)
	// asserts
	// (!)WAIT till called action is completed (Sleep > Timer), it's executed in separate go routine
	time.Sleep(2 * time.Millisecond)
	assert.NoError(s.T(), err)
	chaosMock.AssertExpectations(s.T())
}

func (s *mainTestSuite) Test_netemDelayNoDuration() {
	// prepare test data
	// netem flags
	netemSet := flag.NewFlagSet("netem", 0)
	netemSet.String("interface", "test0", "doc")
	netemCtx := cli.NewContext(nil, netemSet, nil)
	// delay flags
	delaySet := flag.NewFlagSet("delay", 0)
	delaySet.Int("Time", 200, "doc")
	delaySet.Int("jitter", 20, "doc")
	delaySet.Float64("correlation", 10, "doc")
	delaySet.Parse([]string{"c1", "c2", "c3"})
	delayCtx := cli.NewContext(nil, delaySet, netemCtx)
	// invoke command
	err := netemDelay(delayCtx)
	// asserts
	assert.EqualError(s.T(), err, "Undefined duration interval")
}

func (s *mainTestSuite) Test_netemDelayBadDuration() {
	// prepare test data
	// netem flags
	netemSet := flag.NewFlagSet("netem", 0)
	netemSet.String("interface", "test0", "doc")
	netemSet.String("duration", "BAD", "doc")
	netemCtx := cli.NewContext(nil, netemSet, nil)
	// delay flags
	delaySet := flag.NewFlagSet("delay", 0)
	delaySet.Int("time", 200, "doc")
	delaySet.Int("jitter", 20, "doc")
	delaySet.Float64("correlation", 10, "doc")
	delaySet.Parse([]string{"c1", "c2", "c3"})
	delayCtx := cli.NewContext(nil, delaySet, netemCtx)
	// invoke command
	err := netemDelay(delayCtx)
	// asserts
	assert.EqualError(s.T(), err, "time: invalid duration BAD")
}

func (s *mainTestSuite) Test_netemDelayBigDuration() {
	// prepare test data
	gInterval = 1 * time.Second
	// netem flags
	netemSet := flag.NewFlagSet("netem", 0)
	netemSet.String("interface", "test0", "doc")
	netemSet.String("duration", "10s", "doc")
	netemCtx := cli.NewContext(nil, netemSet, nil)
	// delay flags
	delaySet := flag.NewFlagSet("delay", 0)
	delaySet.Int("time", 200, "doc")
	delaySet.Int("jitter", 20, "doc")
	delaySet.Float64("correlation", 10, "doc")
	delaySet.Parse([]string{"c1", "c2", "c3"})
	delayCtx := cli.NewContext(nil, delaySet, netemCtx)
	// invoke command
	err := netemDelay(delayCtx)
	// asserts
	assert.EqualError(s.T(), err, "Duration cannot be bigger than interval")
}

func (s *mainTestSuite) Test_netemDelayBadNetInterface() {
	// prepare test data
	gInterval = 1 * time.Second
	// netem flags
	netemSet := flag.NewFlagSet("netem", 0)
	netemSet.String("interface", "hello test", "doc")
	netemSet.String("duration", "10ms", "doc")
	netemCtx := cli.NewContext(nil, netemSet, nil)
	// delay flags
	delaySet := flag.NewFlagSet("delay", 0)
	delaySet.Int("time", 200, "doc")
	delaySet.Int("jitter", 20, "doc")
	delaySet.Float64("correlation", 10, "doc")
	delaySet.Parse([]string{"c1", "c2", "c3"})
	delayCtx := cli.NewContext(nil, delaySet, netemCtx)
	// invoke command
	err := netemDelay(delayCtx)
	// asserts
	assert.EqualError(s.T(), err, "Bad network interface name. Must match '[a-zA-Z]+[0-9]{0,2}'")
}

func (s *mainTestSuite) Test_netemDelayInvalidJitter() {
	// prepare test data
	// netem flags
	netemSet := flag.NewFlagSet("netem", 0)
	netemSet.String("interface", "test0", "doc")
	netemSet.String("duration", "10ms", "doc")
	netemCtx := cli.NewContext(nil, netemSet, nil)
	// delay flags
	delaySet := flag.NewFlagSet("delay", 0)
	delaySet.Int("time", 200, "doc")
	delaySet.Int("jitter", -10, "doc")
	delaySet.Float64("correlation", 10, "doc")
	delaySet.Parse([]string{"c1", "c2", "c3"})
	delayCtx := cli.NewContext(nil, delaySet, netemCtx)
	// invoke command
	err := netemDelay(delayCtx)
	// asserts
	assert.EqualError(s.T(), err, "Invalid delay jitter")
}

func (s *mainTestSuite) Test_netemDelayInvalidTime() {
	// prepare test data
	// netem flags
	netemSet := flag.NewFlagSet("netem", 0)
	netemSet.String("interface", "test0", "doc")
	netemSet.String("duration", "10ms", "doc")
	netemCtx := cli.NewContext(nil, netemSet, nil)
	// delay flags
	delaySet := flag.NewFlagSet("delay", 0)
	delaySet.Int("time", -20, "doc")
	delaySet.Int("jitter", 20, "doc")
	delaySet.Float64("correlation", 101.5, "doc")
	delaySet.Parse([]string{"c1", "c2", "c3"})
	delayCtx := cli.NewContext(nil, delaySet, netemCtx)
	// invoke command
	err := netemDelay(delayCtx)
	// asserts
	assert.EqualError(s.T(), err, "Invalid delay time")
}

func (s *mainTestSuite) Test_netemDelayInvalidCorrelation() {
	// prepare test data
	// netem flags
	netemSet := flag.NewFlagSet("netem", 0)
	netemSet.String("interface", "test0", "doc")
	netemSet.String("duration", "10ms", "doc")
	netemCtx := cli.NewContext(nil, netemSet, nil)
	// delay flags
	delaySet := flag.NewFlagSet("delay", 0)
	delaySet.Int("time", 200, "doc")
	delaySet.Int("jitter", 20, "doc")
	delaySet.Float64("correlation", 101.5, "doc")
	delaySet.Parse([]string{"c1", "c2", "c3"})
	delayCtx := cli.NewContext(nil, delaySet, netemCtx)
	// invoke command
	err := netemDelay(delayCtx)
	// asserts
	assert.EqualError(s.T(), err, "Invalid delay correlation: must be between 0.0 and 100.0")
}

func (s *mainTestSuite) Test_netemDelayInvalidDistribution() {
	// prepare test data
	// netem flags
	netemSet := flag.NewFlagSet("netem", 0)
	netemSet.String("interface", "test0", "doc")
	netemSet.String("duration", "10ms", "doc")
	netemCtx := cli.NewContext(nil, netemSet, nil)
	// delay flags
	delaySet := flag.NewFlagSet("delay", 0)
	delaySet.Int("time", 200, "doc")
	delaySet.Int("jitter", 10, "doc")
	delaySet.Float64("correlation", 10, "doc")
	delaySet.String("distribution", "INVALID", "doc")
	delaySet.Parse([]string{"c1", "c2", "c3"})
	delayCtx := cli.NewContext(nil, delaySet, netemCtx)
	// invoke command
	err := netemDelay(delayCtx)
	// asserts
	assert.EqualError(s.T(), err, "Invalid delay distribution: must be one of {uniform | normal | pareto |  paretonormal}")
}

func (s *mainTestSuite) Test_netemLossRandomSucess() {
	// prepare test data
	// netem flags
	netemSet := flag.NewFlagSet("netem", 0)
	netemSet.String("duration", "10ms", "doc")
	netemSet.String("interface", "test0", "doc")
	netemCtx := cli.NewContext(nil, netemSet, nil)
	// delay flags
	delaySet := flag.NewFlagSet("loss", 0)
	delaySet.Float64("percent", 20, "doc")
	delaySet.Float64("correlation", 1.5, "doc")
	delaySet.Parse([]string{"c1", "c2", "c3"})
	delayCtx := cli.NewContext(nil, delaySet, netemCtx)
	// set interval to 20ms
	gInterval = 20 * time.Millisecond
	// setup mock
	cmd := action.CommandNetemLossRandom{
		NetInterface: "test0",
		Duration:     10 * time.Millisecond,
		Percent:      20.0,
		Correlation:  1.5,
		StopChan:     gStopChan,
	}
	chaosMock := &ChaosMock{}
	chaos = chaosMock
	chaosMock.On("NetemLossRandomContainers", nil, []string{"c1", "c2", "c3"}, "", cmd).Return(nil)
	// invoke command
	err := netemLossRandom(delayCtx)
	// asserts
	// (!)WAIT till called action is completed (Sleep > Timer), it's executed in separate go routine
	time.Sleep(2 * time.Millisecond)
	assert.NoError(s.T(), err)
	chaosMock.AssertExpectations(s.T())
}

func (s *mainTestSuite) Test_netemLossStateSucess() {
	// prepare test data
	// netem flags
	netemSet := flag.NewFlagSet("netem", 0)
	netemSet.String("duration", "10ms", "doc")
	netemSet.String("interface", "test0", "doc")
	netemCtx := cli.NewContext(nil, netemSet, nil)
	// delay flags
	delaySet := flag.NewFlagSet("loss-state", 0)
	delaySet.Float64("p13", 17.5, "doc")
	delaySet.Float64("p31", 79.26, "doc")
	delaySet.Float64("p32", 1.5, "doc")
	delaySet.Float64("p23", 7.5, "doc")
	delaySet.Float64("p14", 9.31, "doc")
	delaySet.Parse([]string{"c1", "c2", "c3"})
	delayCtx := cli.NewContext(nil, delaySet, netemCtx)
	// set interval to 20ms
	gInterval = 20 * time.Millisecond
	// setup mock
	cmd := action.CommandNetemLossState{
		NetInterface: "test0",
		Duration:     10 * time.Millisecond,
		P13:          17.5,
		P31:          79.26,
		P32:          1.5,
		P23:          7.5,
		P14:          9.31,
		StopChan:     gStopChan,
	}
	chaosMock := &ChaosMock{}
	chaos = chaosMock
	chaosMock.On("NetemLossStateContainers", nil, []string{"c1", "c2", "c3"}, "", cmd).Return(nil)
	// invoke command
	err := netemLossState(delayCtx)
	// asserts
	// (!)WAIT till called action is completed (Sleep > Timer), it's executed in separate go routine
	time.Sleep(2 * time.Millisecond)
	assert.NoError(s.T(), err)
	chaosMock.AssertExpectations(s.T())
}

func (s *mainTestSuite) Test_netemLossGEmodelSucess() {
	// prepare test data
	// netem flags
	netemSet := flag.NewFlagSet("netem", 0)
	netemSet.String("duration", "10ms", "doc")
	netemSet.String("interface", "test0", "doc")
	netemCtx := cli.NewContext(nil, netemSet, nil)
	// delay flags
	delaySet := flag.NewFlagSet("loss-state", 0)
	delaySet.Float64("pg", 7.5, "doc")
	delaySet.Float64("pb", 92.1, "doc")
	delaySet.Float64("one-h", 82.34, "doc")
	delaySet.Float64("one-k", 8.32, "doc")
	delaySet.Parse([]string{"c1", "c2", "c3"})
	delayCtx := cli.NewContext(nil, delaySet, netemCtx)
	// set interval to 20ms
	gInterval = 20 * time.Millisecond
	// setup mock
	cmd := action.CommandNetemLossGEmodel{
		NetInterface: "test0",
		Duration:     10 * time.Millisecond,
		PG:           7.5,
		PB:           92.1,
		OneH:         82.34,
		OneK:         8.32,
		StopChan:     gStopChan,
	}
	chaosMock := &ChaosMock{}
	chaos = chaosMock
	chaosMock.On("NetemLossGEmodelContainers", nil, []string{"c1", "c2", "c3"}, "", cmd).Return(nil)
	// invoke command
	err := netemLossGEmodel(delayCtx)
	// asserts
	// (!)WAIT till called action is completed (Sleep > Timer), it's executed in separate go routine
	time.Sleep(2 * time.Millisecond)
	assert.NoError(s.T(), err)
	chaosMock.AssertExpectations(s.T())
}

func (s *mainTestSuite) Test_netemRateSucess() {
	// prepare test data
	// netem flags
	netemSet := flag.NewFlagSet("netem", 0)
	netemSet.String("duration", "10ms", "doc")
	netemSet.String("interface", "test0", "doc")
	netemCtx := cli.NewContext(nil, netemSet, nil)
	// rate flags
	rateSet := flag.NewFlagSet("rate", 0)
	rateSet.String("rate", "300kbit", "doc")
	rateSet.Int("packetoverhead", 10, "doc")
	rateSet.Int("cellsize", 20, "doc")
	rateSet.Int("celloverhead", 30, "doc")
	rateSet.Parse([]string{"c1", "c2", "c3"})
	rateCtx := cli.NewContext(nil, rateSet, netemCtx)
	// set interval to 1ms
	gInterval = 20 * time.Millisecond
	// setup mock
	cmd := action.CommandNetemRate{
		NetInterface:   "test0",
		Duration:       10 * time.Millisecond,
		Rate:           "300kbit",
		PacketOverhead: 10,
		CellSize:       20,
		CellOverhead:   30,
		StopChan:       gStopChan,
	}
	chaosMock := &ChaosMock{}
	chaos = chaosMock
	chaosMock.On("NetemRateContainers", nil, []string{"c1", "c2", "c3"}, "", cmd).Return(nil)
	// invoke command
	err := netemRate(rateCtx)
	// asserts
	// (!)WAIT till called action is completed (Sleep > Timer), it's executed in separate go routine
	time.Sleep(2 * time.Millisecond)
	assert.NoError(s.T(), err)
	chaosMock.AssertExpectations(s.T())
}

func (s *mainTestSuite) Test_netemRateInvalidRate() {
	// prepare test data
	// netem flags
	netemSet := flag.NewFlagSet("netem", 0)
	netemSet.String("duration", "10ms", "doc")
	netemSet.String("interface", "test0", "doc")
	netemCtx := cli.NewContext(nil, netemSet, nil)
	// rate flags
	rateSet := flag.NewFlagSet("rate", 0)
	rateSet.String("rate", "300", "doc")
	rateSet.Int("packetoverhead", 10, "doc")
	rateSet.Int("cellsize", 20, "doc")
	rateSet.Int("celloverhead", 30, "doc")
	rateSet.Parse([]string{"c1", "c2", "c3"})
	rateCtx := cli.NewContext(nil, rateSet, netemCtx)
	// invoke command
	err := netemRate(rateCtx)
	// asserts
	assert.EqualError(s.T(), err, "Invalid rate. Must match '[0-9]+[gmk]?bit'")
}

func (s *mainTestSuite) Test_netemRateEmptyRate() {
	// prepare test data
	// netem flags
	netemSet := flag.NewFlagSet("netem", 0)
	netemSet.String("duration", "10ms", "doc")
	netemSet.String("interface", "test0", "doc")
	netemCtx := cli.NewContext(nil, netemSet, nil)
	// rate flags
	rateSet := flag.NewFlagSet("rate", 0)
	rateSet.String("rate", "", "doc")
	rateSet.Int("packetoverhead", 10, "doc")
	rateSet.Int("cellsize", -20, "doc")
	rateSet.Int("celloverhead", 30, "doc")
	rateSet.Parse([]string{"c1", "c2", "c3"})
	rateCtx := cli.NewContext(nil, rateSet, netemCtx)
	// invoke command
	err := netemRate(rateCtx)
	// asserts
	assert.EqualError(s.T(), err, "Undefined rate limit")
}

func (s *mainTestSuite) Test_netemRateInvalidCellSize() {
	// prepare test data
	// netem flags
	netemSet := flag.NewFlagSet("netem", 0)
	netemSet.String("duration", "10ms", "doc")
	netemSet.String("interface", "test0", "doc")
	netemCtx := cli.NewContext(nil, netemSet, nil)
	// rate flags
	rateSet := flag.NewFlagSet("rate", 0)
	rateSet.String("rate", "300kbit", "doc")
	rateSet.Int("packetoverhead", 10, "doc")
	rateSet.Int("cellsize", -20, "doc")
	rateSet.Int("celloverhead", 30, "doc")
	rateSet.Parse([]string{"c1", "c2", "c3"})
	rateCtx := cli.NewContext(nil, rateSet, netemCtx)
	// invoke command
	err := netemRate(rateCtx)
	// asserts
	assert.EqualError(s.T(), err, "Invalid cell size: must be a non-negative integer")
}

func TestMainTestSuite(t *testing.T) {
	//suite.Run(t, new(mainTestSuite))
}
