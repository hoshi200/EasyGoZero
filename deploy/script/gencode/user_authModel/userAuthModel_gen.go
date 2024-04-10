// Code generated by goctl. DO NOT EDIT.

package user_authModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userAuthFieldNames          = builder.RawFieldNames(&UserAuth{}, true)
	userAuthRows                = strings.Join(userAuthFieldNames, ",")
	userAuthRowsExpectAutoSet   = strings.Join(stringx.Remove(userAuthFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	userAuthRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(userAuthFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cachePublicUserAuthIdPrefix                     = "cache:public:userAuth:id:"
	cachePublicUserAuthIdentityKindIdentifierPrefix = "cache:public:userAuth:identityKind:identifier:"
	cachePublicUserAuthUserIdIdentityKindPrefix     = "cache:public:userAuth:userId:identityKind:"
)

type (
	userAuthModel interface {
		Insert(ctx context.Context, data *UserAuth) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserAuth, error)
		FindOneByIdentityKindIdentifier(ctx context.Context, identityKind int64, identifier string) (*UserAuth, error)
		FindOneByUserIdIdentityKind(ctx context.Context, userId int64, identityKind int64) (*UserAuth, error)
		Update(ctx context.Context, data *UserAuth) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserAuthModel struct {
		sqlc.CachedConn
		table string
	}

	UserAuth struct {
		Id           int64          `db:"id"`
		UserId       int64          `db:"user_id"`
		IdentityKind int64          `db:"identity_kind"`
		Identifier   string         `db:"identifier"`
		Credential   string         `db:"credential"`
		CreatedAt    time.Time      `db:"created_at"`
		UpdatedAt    time.Time      `db:"updated_at"`
		DeletedAt    sql.NullTime   `db:"deleted_at"`
		Verified     int64          `db:"verified"`
		Ip           sql.NullString `db:"ip"`
		IpLocation   sql.NullString `db:"ip_location"`
	}
)

func newUserAuthModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserAuthModel {
	return &defaultUserAuthModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."user_auth"`,
	}
}

func (m *defaultUserAuthModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	publicUserAuthIdKey := fmt.Sprintf("%s%v", cachePublicUserAuthIdPrefix, id)
	publicUserAuthIdentityKindIdentifierKey := fmt.Sprintf("%s%v:%v", cachePublicUserAuthIdentityKindIdentifierPrefix, data.IdentityKind, data.Identifier)
	publicUserAuthUserIdIdentityKindKey := fmt.Sprintf("%s%v:%v", cachePublicUserAuthUserIdIdentityKindPrefix, data.UserId, data.IdentityKind)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicUserAuthIdKey, publicUserAuthIdentityKindIdentifierKey, publicUserAuthUserIdIdentityKindKey)
	return err
}

func (m *defaultUserAuthModel) FindOne(ctx context.Context, id int64) (*UserAuth, error) {
	publicUserAuthIdKey := fmt.Sprintf("%s%v", cachePublicUserAuthIdPrefix, id)
	var resp UserAuth
	err := m.QueryRowCtx(ctx, &resp, publicUserAuthIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", userAuthRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) FindOneByIdentityKindIdentifier(ctx context.Context, identityKind int64, identifier string) (*UserAuth, error) {
	publicUserAuthIdentityKindIdentifierKey := fmt.Sprintf("%s%v:%v", cachePublicUserAuthIdentityKindIdentifierPrefix, identityKind, identifier)
	var resp UserAuth
	err := m.QueryRowIndexCtx(ctx, &resp, publicUserAuthIdentityKindIdentifierKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where identity_kind = $1 and identifier = $2 limit 1", userAuthRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, identityKind, identifier); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) FindOneByUserIdIdentityKind(ctx context.Context, userId int64, identityKind int64) (*UserAuth, error) {
	publicUserAuthUserIdIdentityKindKey := fmt.Sprintf("%s%v:%v", cachePublicUserAuthUserIdIdentityKindPrefix, userId, identityKind)
	var resp UserAuth
	err := m.QueryRowIndexCtx(ctx, &resp, publicUserAuthUserIdIdentityKindKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where user_id = $1 and identity_kind = $2 limit 1", userAuthRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, identityKind); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) Insert(ctx context.Context, data *UserAuth) (sql.Result, error) {
	publicUserAuthIdKey := fmt.Sprintf("%s%v", cachePublicUserAuthIdPrefix, data.Id)
	publicUserAuthIdentityKindIdentifierKey := fmt.Sprintf("%s%v:%v", cachePublicUserAuthIdentityKindIdentifierPrefix, data.IdentityKind, data.Identifier)
	publicUserAuthUserIdIdentityKindKey := fmt.Sprintf("%s%v:%v", cachePublicUserAuthUserIdIdentityKindPrefix, data.UserId, data.IdentityKind)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8)", m.table, userAuthRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.IdentityKind, data.Identifier, data.Credential, data.DeletedAt, data.Verified, data.Ip, data.IpLocation)
	}, publicUserAuthIdKey, publicUserAuthIdentityKindIdentifierKey, publicUserAuthUserIdIdentityKindKey)
	return ret, err
}

func (m *defaultUserAuthModel) Update(ctx context.Context, newData *UserAuth) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	publicUserAuthIdKey := fmt.Sprintf("%s%v", cachePublicUserAuthIdPrefix, data.Id)
	publicUserAuthIdentityKindIdentifierKey := fmt.Sprintf("%s%v:%v", cachePublicUserAuthIdentityKindIdentifierPrefix, data.IdentityKind, data.Identifier)
	publicUserAuthUserIdIdentityKindKey := fmt.Sprintf("%s%v:%v", cachePublicUserAuthUserIdIdentityKindPrefix, data.UserId, data.IdentityKind)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, userAuthRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Id, newData.UserId, newData.IdentityKind, newData.Identifier, newData.Credential, newData.DeletedAt, newData.Verified, newData.Ip, newData.IpLocation)
	}, publicUserAuthIdKey, publicUserAuthIdentityKindIdentifierKey, publicUserAuthUserIdIdentityKindKey)
	return err
}

func (m *defaultUserAuthModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicUserAuthIdPrefix, primary)
}

func (m *defaultUserAuthModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", userAuthRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserAuthModel) tableName() string {
	return m.table
}
