/*
   Package database 
   General database for CRUD operation
*/
package database

import "gorm.io/gorm"


// Create will create new record into model and save it to database
func Create(db *gorm.DB, model interface{}) error {
    return db.Create(model).Error
}

// Save will save new record into model and save it to database
func Save(db *gorm.DB, model interface{}) error {
    return db.Save(model).Error
}

// Update will update model record into the database
func Update(db *gorm.DB, condition, model interface{}) error {
    return db.Model(condition).Updates(model).Error
}

// Delete will delete record(s) from model with given condition
// it will returning 'count' of affected record and error status
func Delete(db *gorm.DB, condition, model interface{}) (count int, err error) {
    record := db.Where(condition).Delete(model)
    err = record.Error
    if err != nil {
        return
    }

    count = int(record.RowsAffected)

    return
}

// First will fetch model and its associations based on given condition
func First(db *gorm.DB, condition, model interface{}, associations []string) (
    isNotFound bool, err error) {
    // loop for associations
    for _, a := range associations {
        db = db.Preload(a)
    }

    err = db.Where(condition).First(model).Error
    if err != nil {
        isNotFound = gorm.ErrRecordNotFound == err
    }

    return
}

// Find will get model and its associations records based on given condition
func Find(db *gorm.DB, condition, model interface{}, associations []string) (
    isEmpty bool, err error) {
    // loop for association
    for _, a := range associations {
        db = db.Preload(a)
    }

    err = db.Where(condition).Find(model).Error
    if err != nil {
        isEmpty = gorm.ErrRecordNotFound == err
    }

    return
}
