package debt

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
	"path/filepath"
	"strconv"
)

// FileStore struct for storing debt data
type store struct {
	Name string
	Type string
}

var memoryDebts []GoBill

// NewDebtStore creates a new DebtStore
// where name is the name of the file
// and formatType is the type of file
func NewDebtStore(name, formatType string) *store {
	return &store{
		Name: name,
		Type: formatType,
	}
}

// Load loads the data from the file
func Load(ds *store) []GoBill {
	switch ds.Type {
	case "json":
		return LoadJsonFile(ds.Name)
	case "sqlite":
		return LoadSqliteDB(ds.Name)
	default:
		log.Warnln("Type not implemented. Using memory instead.")
		memoryDebts = make([]GoBill, 0)
		return memoryDebts
	}
}

// Save saves the data to the file
func Save(ds *store, debt []GoBill) {
	switch ds.Type {
	case "json":
		SaveJsonFileDebts(ds.Name, debt)
	case "sqlite":
		SaveSQLiteDebts(ds.Name, debt)
	default:
		memoryDebts = debt
	}
}

func Delete(ds *store, debtID int) {
	switch ds.Type {
	case "json":
		log.Error("Use Save with []GoBill without GoBill element to be deleted instead of Delete for json files.")
	case "sqlite":
		DeleteSQLiteDebts(ds.Name, debtID)
	default:
		for i, debt := range memoryDebts {
			if debt.ID == debtID {
				memoryDebts = append(memoryDebts[:i], memoryDebts[i+1:]...)
			}
		}
	}
}

// ConvertOldJsonToSqlite converts the old json file with OldDebtItem and saves it as a sqlite db with GoBill
func ConvertOldJsonToSqlite(jsonFile, sqliteFile string) {
	oldDebts := OldLoadJsonFile(jsonFile)
	var debts []GoBill
	for _, oldDebt := range oldDebts {
		log.Infof("Adding %s to sqlite db...", oldDebt.Name)
		due, err := strconv.Atoi(oldDebt.Due)
		if err != nil {
			log.Error(err)
			due = 0
		}
		debts = append(debts, *NewDebtItem(
			0,
			due,
			oldDebt.Total,
			oldDebt.Monthly,
			oldDebt.Monthly,
			oldDebt.Interest,
			oldDebt.Name,
			oldDebt.Type,
		))
	}
	log.Infof("Adding total of %d debts to sqlite db...", len(debts))
	SaveSQLiteDebts(sqliteFile, debts)
}

// LoadSqliteDB creates the sqlite db with gorm and returns the data
func LoadSqliteDB(name string) []GoBill {
	var db *gorm.DB
	dbFile := name + ".db"
	home, _ := os.UserHomeDir()
	dbPath := filepath.Join(home, dbFile)
	if _, err := os.Stat(dbPath); err == nil {
		log.Debugln("File exists")
		db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		if err != nil {
			log.Error(err)
		}
	} else {
		log.Warnln("File does not exist")
		db = CreateSQLiteStore(dbPath)
		log.Infof("%s created\n", dbFile)
		dataFile := "debt"
		home, _ := os.UserHomeDir()
		if _, err := os.Stat(filepath.Join(home, dataFile+".json")); err == nil {
			log.Infof("Converting %s to sqlite\n", dataFile+".json")
			ConvertOldJsonToSqlite(dataFile, name)
		} else {
			log.Warnf("No %s found. Creating empty db.\n", filepath.Join(home, dataFile))
		}
	}

	return GetSQLiteDebts(db)
}

// CreateSQLiteStore creates a new sqlite store
func CreateSQLiteStore(dbName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}

	if !db.Migrator().HasTable(&GoBill{}) {
		err := db.AutoMigrate(&GoBill{})
		if err != nil {
			return nil
		}
	}

	return db
}

// GetSQLiteDebts from the store
func GetSQLiteDebts(db *gorm.DB) []GoBill {
	var debts []GoBill
	db.Find(&debts)
	return debts
}

// SaveSQLiteDebts to the store
func SaveSQLiteDebts(name string, debt []GoBill) {
	log.Infof("Saving to sqlite db at %s...", name)
	dbFile := name + ".db"
	home, _ := os.UserHomeDir()
	dbPath := filepath.Join(home, dbFile)
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}
	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&debt)
}

// DeleteSQLiteDebts from the store
func DeleteSQLiteDebts(name string, debtID int) {
	dbFile := name + ".db"
	home, _ := os.UserHomeDir()
	dbPath := filepath.Join(home, dbFile)
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}
	res := db.Where("id = ?", debtID).Delete(&GoBill{})
	if res.Error != nil {
		log.Error(res.Error)
	}
}

// LoadJsonFile loads the data from the json file
func LoadJsonFile(name string) []GoBill {
	dataFile := name + ".json"
	home, _ := os.UserHomeDir()
	if _, err := os.Stat(filepath.Join(home, dataFile)); err == nil {
		log.Debugln("File exists")
		return GetJsonFileDebts(home, dataFile)

	} else {
		log.Warnln("File does not exist")
		CreateJsonFileStore(home, dataFile)
		log.Infof("%s created\n", dataFile)
		return make([]GoBill, 0)
	}
}

// OldLoadJsonFile loads the data from the json file
func OldLoadJsonFile(name string) []OldDebtItem {
	dataFile := name + ".json"
	home, _ := os.UserHomeDir()
	return OldGetJsonFileDebts(home, dataFile)

}

// CreateJsonFileStore creates a new json file store
func CreateJsonFileStore(home, dataFile string) {
	err := os.WriteFile(filepath.Join(home, dataFile), nil, 0600)
	if err != nil {
		log.Error(err)
	}
}

// GetJsonFileDebts from the store
func GetJsonFileDebts(home, dataFile string) []GoBill {
	file, err := os.ReadFile(filepath.Join(home, dataFile))
	if err != nil {
		return nil
	}
	debts := make([]GoBill, 0)
	err = json.Unmarshal(file, &debts)
	if err != nil {
		log.Error(err)
	}
	return debts
}

func OldGetJsonFileDebts(home, dataFile string) []OldDebtItem {
	file, err := os.ReadFile(filepath.Join(home, dataFile))
	if err != nil {
		return nil
	}
	debts := make([]OldDebtItem, 0)
	err = json.Unmarshal(file, &debts)
	if err != nil {
		log.Error(err)
	}
	log.Infof("Loaded %d debts from %s\n", len(debts), dataFile)
	return debts
}

// SaveJsonFileDebts to the store
func SaveJsonFileDebts(name string, debt []GoBill) {
	DataFile := name + ".json"
	home, err := os.UserHomeDir()
	if err != nil {
		log.Error(err)
	}
	file, err := json.MarshalIndent(debt, "", "    ")
	if err != nil {
		log.Error(err)
	}
	err = os.WriteFile(filepath.Join(home, DataFile), file, 0600)
	if err != nil {
		log.Error(err)
	}
}
