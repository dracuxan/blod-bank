package handler_test

import (
	"context"
	"testing"

	blodBank "github.com/dracuxan/blod-bank/proto"
	"github.com/dracuxan/blod-bank/server/handler"
	"github.com/dracuxan/blod-bank/server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to in-memory db: %v", err)
	}
	// auto-migrate schema
	if err := db.AutoMigrate(&models.Configs{}); err != nil {
		t.Fatalf("failed to migrate schema: %v", err)
	}
	return db
}

func TestRegisterConfig_Success(t *testing.T) {
	db := SetupTestDB(t)
	s := handler.NewServer(db)

	req := &blodBank.ConfigItem{
		Name:    "test.conf",
		Content: "foo=bar",
	}

	resp, err := s.RegisterConfig(context.Background(), req)
	if err != nil {
		t.Fatalf("RegisterConfig failed: %v", err)
	}

	if resp.Status != "Registered new config" {
		t.Errorf("expected status 'Registered new config', got %q", resp.Status)
	}

	var count int64
	if err := db.Model(&models.Configs{}).Where("name = ?", "test.conf").Count(&count).Error; err != nil {
		t.Fatalf("failed to query db: %v", err)
	}
	if count != 1 {
		t.Errorf("expected 1 record in DB, got %d", count)
	}
}
