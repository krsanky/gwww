package model

//Limit
//Specify the max number of records to retrieve
//
//db.Limit(3).Find(&users)
////// SELECT * FROM users LIMIT 3;
//
//// Cancel limit condition with -1
//db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
////// SELECT * FROM users LIMIT 10; (users1)
////// SELECT * FROM users; (users2)

//Offset
//Specify the number of records to skip before starting to return the records
//
//db.Offset(3).Find(&users)
////// SELECT * FROM users OFFSET 3;
//
//// Cancel offset condition with -1
//db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
////// SELECT * FROM users OFFSET 10; (users1)
////// SELECT * FROM users; (users2)
