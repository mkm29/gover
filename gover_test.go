package main

import (
	"fmt"
	"gover/internal/utils"
	"gover/pkg/config"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
)

var evs = []map[string]string{
	{
		"CI_DEFAULT_BRANCH": "master",
	},
	{
		"CI_MERGE_REQUEST_TARGET_BRANCH_NAME": "test",
	},
	{
		"CI_PIPELINE_IID": "1",
	},
}

type UnitTestSuite struct {
	suite.Suite
	Cmd    *cobra.Command
	Config *config.Config
}

func (s *UnitTestSuite) SetupSuite() {
	s.T().Log("Setting up test suite...")
	// clear environment variables
	os.Clearenv()
	// create temporary testdata directory
	os.Mkdir("testdata", 0755)
	// asserts := s.Assert()
	requires := s.Require()
	// s.setupEnvironmentVariables()

	vc := `MAJOR=1
MINOR=2
PATCH=3
#ADDOPTS=ignore_me`

	err := createFile("version", vc)
	requires.NoError(err)
	// s.T().Logf("Version file: %s\n", s.VersionFile)

	cc := `CI_DEFAULT_BRANCH=master
CI_MERGE_REQUEST_TARGET_BRANCH_NAME=test
CI_PIPELINE_IID=1`
	err = createFile("config.env", cc)
	requires.NoError(err)
	// s.T().Logf("Config file: %s\n", s.ConfigFile)

}

func createFile(fn string, s string) error {
	// file, err := os.CreateTemp("testdata", fn)
	fp := fmt.Sprintf("testdata/%s", fn)
	if err := os.WriteFile(fp, []byte(s), 0644); err != nil {
		return err
	}
	return nil
}

// function to set environment variables
func (s *UnitTestSuite) setupEnvironmentVariables() {
	for _, e := range evs {
		for k, v := range e {
			s.T().Setenv(k, v)
		}
	}
}

func TestUnitTestSuite(t *testing.T) {
	// run
	suite.Run(t, &UnitTestSuite{})
}

func (s *UnitTestSuite) TestLoadConfig() {
	requires := s.Require()
	config, err := config.LoadConfig("testdata", "config.env")
	requires.NoError(err)
	s.Config = config
	asserts := s.Assert()
	asserts.Equal("master", config.Variables.DefaultBranch)
	asserts.Equal("test", config.Variables.MergeRequestTargetBranchName)
	asserts.Equal(1, config.Variables.PipelineIid)
}

func (s *UnitTestSuite) TestGetVersion() {
	// requires := s.Require()
	config, err := config.LoadConfig("testdata", "config.env")
	s.Require().NoError(err)
	config.VersionFile = "testdata/version"
	version := utils.GetVersion(config)
	// make sure version contains 1.2.3
	s.Assert().Contains(version, "1.2.3")
}

func (s *UnitTestSuite) TearDownSuite() {
	// run after all tests
	s.T().Log("Tearing down test suite...")
	os.Remove("testdata/version")
	os.Remove("testdata/config.env")
}

// func (s *UnitTestSuite) TestRootCmd() {
// 	cmd := s.Cmd
// 	b := bytes.NewBufferString("")
// 	cmd.SetOut(b)
// 	cmd.SetArgs([]string{"help"})
// 	cmd.Execute()
// 	out, err := io.ReadAll(b)
// 	require := s.Require()
// 	// assert := s.Assert()
// 	require.NoError(err)
// 	s.T().Logf("out: %s", out)
// 	// assert that Usage is contained in output
// 	// assert.Contains(string(out), "Usage:")
// }
