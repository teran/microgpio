package server

import (
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"

	"github.com/teran/microgpio/client"
	"github.com/teran/microgpio/drivers/fake"
	"github.com/teran/microgpio/server"
)

type ClientTestSuite struct {
	suite.Suite
}

func (s *ClientTestSuite) TestSuccessFlow() {
	driver := &fake.FakeDriver{
		LowFunc: func(id int) error {
			return nil
		},
		HighFunc: func(id int) error {
			return nil
		},
		OutputFunc: func(id int) error {
			return nil
		},
	}

	srv := httptest.NewServer(server.New(driver))
	defer srv.Close()

	c := client.New(srv.URL)

	err := c.Ping()
	s.Require().NoError(err)

	err = c.Output(1)
	s.Require().NoError(err)

	err = c.Low(1)
	s.Require().NoError(err)

	err = c.High(1)
	s.Require().NoError(err)
}

func (s *ClientTestSuite) Test5xxOnDriverError() {
	sampleError := errors.New("test error")

	driver := &fake.FakeDriver{
		LowFunc: func(id int) error {
			return sampleError
		},
		HighFunc: func(id int) error {
			return sampleError
		},
		OutputFunc: func(id int) error {
			return sampleError
		},
	}

	srv := httptest.NewServer(server.New(driver))
	defer srv.Close()

	c := client.New(srv.URL)

	err := c.Ping()
	s.Require().NoError(err)

	err = c.Output(1)
	s.Require().Error(err)
	s.Require().Equal(client.ErrUnexpectedStatusCode, errors.Cause(err))

	err = c.Low(1)
	s.Require().Error(err)
	s.Require().Equal(client.ErrUnexpectedStatusCode, errors.Cause(err))

	err = c.High(1)
	s.Require().Error(err)
	s.Require().Equal(client.ErrUnexpectedStatusCode, errors.Cause(err))
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, &ClientTestSuite{})
}
