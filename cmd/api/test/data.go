package test

var CreateUser1InputBody = map[string]string{"id": "user_id_1_testing"}

const CreateUser1Answer = "{\"result\":{\"id\":\"user_id_1_testing\",\"status\":\"ok\"}}"
const CreateExitingUserAnswer = "{\"result\":{\"error\":\"creating user. user with this ID is already exists\"}}"

const Data1Path = "../../../test_data/data_1.xlsx"
const PostData1Answer = "{\"result\":{\"count_deleted_products\":0,\"count_new_products\":3,\"count_updated_products\":0,\"not_processed\":[]}}"
const GetData1Answer = "{\"result\":{\"count\":3,\"data\":[{\"name\":\"телевизор\",\"offer_id\":\"offer_id_1\",\"price\":10000,\"quality\":5},{\"name\":\"телефон\",\"offer_id\":\"offer_id_2\",\"price\":20000,\"quality\":20},{\"name\":\"монитор\",\"offer_id\":\"offer_id_3\",\"price\":5000,\"quality\":10}]}}"

const Data12Path = "../../../test_data/data_1_2.xlsx"
const PostData12Answer = "{\"result\":{\"count_deleted_products\":1,\"count_new_products\":1,\"count_updated_products\":1,\"not_processed\":[]}}"

const PostDataOfNotExistedUser = "{\"result\":{\"error\":\"user with input ID does not exist\"}}"

const BadDataPath = "../../../test_data/bad_data.xlsx"
const PostBadDataAnswer = "{\"result\":{\"count_deleted_products\":0,\"count_new_products\":0,\"count_updated_products\":0,\"not_processed\":[{\"error\":\"available is not converted to bool\",\"indexes_in_xlsx\":[4]},{\"error\":\"name is required\",\"indexes_in_xlsx\":[1]},{\"error\":\"offer ID is required\",\"indexes_in_xlsx\":[0]},{\"error\":\"price is less than 0\",\"indexes_in_xlsx\":[2]},{\"error\":\"quality is less than 0\",\"indexes_in_xlsx\":[3]}]}}"
