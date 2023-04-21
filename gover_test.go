package main

import (
	"fmt"
	"gover/internal/config"
	"gover/internal/utils"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"
)

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
	if err := os.MkdirAll("testdata", os.ModeSticky|os.ModePerm); err != nil {
		s.T().Fatalf("Unable to create testdata directory: %v", err)
	}
	requires := s.Require()

	vc := `MAJOR=1
MINOR=2
PATCH=3
#ADDOPTS=ignore_me`

	err := createFile("version", vc)
	requires.NoError(err)

	cc := `CI_DEFAULT_BRANCH=master
CI_MERGE_REQUEST_TARGET_BRANCH_NAME=test
CI_PIPELINE_IID=1
CI_PIPELINE_SOURCE=merge_event
CI_PIPELINE_TRIGGERED=true
CI_MERGE_REQUEST_IID=1
CI_COMMIT_BRANCH=test-branch`
	err = createFile("config.env", cc)
	requires.NoError(err)
	// initialize config
	config.Init()
}

func createFile(fn, s string) error {
	fp := fmt.Sprintf("testdata/%s", fn)
	if err := os.WriteFile(fp, []byte(s), os.FileMode(0600)); err != nil {
		return err
	}
	return nil
}

func TestUnitTestSuite(t *testing.T) {
	// run
	suite.Run(t, &UnitTestSuite{})
}

func (s *UnitTestSuite) TestLoadConfig() {
	requires := s.Require()
	cfg, err := config.LoadConfig("testdata", "config.env")
	requires.NoError(err)
	s.Config = cfg
	asserts := s.Assert()
	asserts.Equal("master", cfg.Variables.DefaultBranch)
	asserts.Equal("test", cfg.Variables.MergeRequestTargetBranchName)
	asserts.Equal(1, cfg.Variables.PipelineIid)
}

func (s *UnitTestSuite) TestParseError() {
	// intentionally add error in config file
	requires := s.Require()
	cc := `CI_DEFAULT_BRANCH?master`
	err := createFile("error.env", cc)
	requires.NoError(err)
	_, err = config.LoadConfig("testdata", "error.env")
	// expects error
	requires.Error(err)
}

func (s *UnitTestSuite) TestGetVersion() {
	cfg, err := config.LoadConfig("testdata", "config.env")
	s.Require().NoError(err)
	cfg.VersionFile = "testdata/version"
	version := utils.GetVersion(cfg)
	// make sure version equals 1.2.3-test-branch+1
	s.Assert().Equal("1.2.3-test-branch+1", version)
}

func (s *UnitTestSuite) TearDownSuite() {
	// run after all tests
	s.T().Log("Tearing down test suite...")
	// remove the testdata directory
	if err := os.RemoveAll("testdata"); err != nil {
		s.T().Fatalf("Unable to remove testdata directory: %v", err)
	}
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
