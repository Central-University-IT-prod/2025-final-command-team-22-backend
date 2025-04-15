package business_test

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"loyalit/internal/adapters/repository/postgres/ent"
	entbusiness "loyalit/internal/adapters/repository/postgres/ent/business"
	"loyalit/test/testhelper"
	"testing"
)

type BusinessServiceTestSuite struct {
	suite.Suite
	pgContainer *testhelper.PostgresContainer
	db          *ent.Client
	ctx         context.Context
	businessID  uuid.UUID
}

func (suite *BusinessServiceTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := testhelper.CreatePostgresContainer(suite.ctx)
	if err != nil {
		panic(err)
	}
	suite.pgContainer = pgContainer

	db, err := ent.Open(dialect.Postgres, suite.pgContainer.ConnStr)
	if err != nil {
		panic(err)
	}

	suite.db = db

	if errMigrate := suite.db.Schema.Create(
		context.Background(),
	); errMigrate != nil {
		panic(err)
	}

}

func (suite *BusinessServiceTestSuite) TearDownSuite() {
	if err := suite.db.Close(); err != nil {
		panic(err)
	}
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		panic(err)
	}
}

func (suite *BusinessServiceTestSuite) SetupTest() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("TestPass"), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	business, err := suite.db.Business.Create().
		SetName("Тест").
		SetEmail("test@example.com").
		SetPassword(string(hashedPassword)).
		SetDescription("Тест").
		Save(suite.ctx)
	if err != nil {
		panic(err)
	}
	suite.businessID = business.ID
}

func (suite *BusinessServiceTestSuite) TearDownTest() {
	_, err := suite.db.Business.Delete().Where(entbusiness.ID(suite.businessID)).Exec(suite.ctx)
	if err != nil {
		panic(err)
	}
}

func TestBusinessServiceTestSuite(t *testing.T) {
	suite.Run(t, new(BusinessServiceTestSuite))
}
