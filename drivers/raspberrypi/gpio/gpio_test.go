package gpio

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/teran/microgpio/models"
)

type GPIOTestSuite struct {
	suite.Suite

	dir string
}

func (s *GPIOTestSuite) TestAll() {
	p := NewWithPrefix(1, s.dir)
	err := p.Export()
	s.Require().NoError(err)

	err = p.Unexport()
	s.Require().NoError(err)

	err = p.Input()
	s.Require().NoError(err)

	m, err := p.Mode()
	s.Require().NoError(err)
	s.Require().Equal(models.ModeIn, m)

	err = p.Output()
	s.Require().NoError(err)

	m, err = p.Mode()
	s.Require().NoError(err)
	s.Require().Equal(models.ModeOut, m)

	err = p.High()
	s.Require().NoError(err)

	v, err := p.Value()
	s.Require().NoError(err)
	s.Require().Equal(1, v)

	err = p.Low()
	s.Require().NoError(err)

	v, err = p.Value()
	s.Require().NoError(err)
	s.Require().Equal(0, v)

	err = p.Input()
	s.Require().NoError(err)

	m, err = p.Mode()
	s.Require().NoError(err)
	s.Require().Equal(models.ModeIn, m)
}

func (s *GPIOTestSuite) SetupSuite() {
	dir, err := ioutil.TempDir("", "gpiotestsuite")
	s.Require().NoError(err)

	err = os.MkdirAll(path.Join(dir, path.Dir(ExportFilepath)), 0777)
	s.Require().NoError(err)

	err = os.MkdirAll(path.Join(dir, path.Dir(fmt.Sprintf(DirectionFilepath, 1))), 0777)
	s.Require().NoError(err)

	s.dir = dir
}

func (s *GPIOTestSuite) TearDownSuite() {
	err := os.RemoveAll(s.dir)
	s.Require().NoError(err)
}

func TestGPIOTestSuite(t *testing.T) {
	suite.Run(t, &GPIOTestSuite{})
}
