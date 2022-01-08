/*
   Package errors for database
   containing custom error for our app
*/
package errors

const (
    // ErrDataEmpty is error code for empty data result
    ErrDataIsEmpty = iota + 800 

    // ErrDataNotFound is error code for not found data 
    ErrDataNotFound

    // ErrGettingData is error code for fail to get/ retreive data
    ErrGettingData

    // ErrSaveDataFail is error code for 'failling' on saving data
    ErrSaveDataFail

    // ErrDataCouldNotUpdate is error code for 'failling' on update data
    ErrUpdateDataFail

    // ErrDeleteData is error code for failing to delete data
    ErrDeleteDataFail

    // ErrDataAlreadyExist is error code when triying to save data on an already exist data
    // for example 'Primary Key' or 'Unique Constraint' already exist 
    ErrDataAlreadyExist

)

const (
    // ErrDataEmptyMsg is error code for empty data result
    ErrDataIsEmptyMsg = "data is empty"

    // ErrDataNotFoundMsg is error code for not found data 
    ErrDataNotFoundMsg = "data not found"

    // ErrGettingDataMsg is error message for fail to get/ retreive data
    ErrGettingDataMsg = "could not retreive data"
    // ErrSaveDataFailMsg is error code for 'failling' on saving data
    ErrSaveDataFailMsg = "could not save data"

    // ErrDataCouldNotUpdateMsg is error code for 'failling' on update data
    ErrUpdateDataFailMsg = "could not update data"

    // ErrDeleteDataMsg is error code for failing to delete data
    ErrDeleteDataFailMsg = "could not delete data"

    // ErrDataExist is error code when triying to save data on an already exist data
    ErrDataAlreadyExistMsg = "data already exist"


)
