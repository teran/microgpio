package server

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/teran/microgpio/client"
	"github.com/teran/microgpio/drivers/fake"
	"github.com/teran/microgpio/server"
)

type ClientTestSuite struct {
	suite.Suite
}

func (s *ClientTestSuite) TestAll() {
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

	client := client.New(srv.URL)

	err := client.Ping()
	s.Require().NoError(err)

	err = client.Output(1)
	s.Require().NoError(err)

	err = client.Low(1)
	s.Require().NoError(err)

	err = client.High(1)
	s.Require().NoError(err)
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, &ClientTestSuite{})
}
