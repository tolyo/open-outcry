package registration

import (
	"testing"

	"open-outcry/pkg/conf"
	"open-outcry/pkg/db"
	sqlpkg "open-outcry/sql"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateAppEntity(t *testing.T) {
	conf.LoadTestConfig()
	require.NoError(t, db.SetupInstance())
	defer func() {
		_ = sqlpkg.MigrateDown()
		if db.Instance() != nil {
			_ = db.Instance().Close()
		}
	}()

	require.NoError(t, sqlpkg.MigrateUp())

	count := db.GetCount("app_entity")
	res := CreateAppEntity("test")

	require.NotEmpty(t, res)
	_, err := uuid.Parse(string(res))
	require.NoError(t, err)
	require.Equal(t, count+1, db.GetCount("app_entity"))

	res2 := CreateAppEntity("test")
	require.NotEqual(t, res, res2)
	require.Equal(t, count+2, db.GetCount("app_entity"))
}
