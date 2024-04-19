package main

import "fmt"

func main() {

	fmt.Println("MONGO DB ")
}



// AD MONOGO


// 1.capped collection in mongoDb 
// .specail types of collection with a specific size, its working like a when the size is limit ,when the new document overwrite the oldest  doucments  	
// syntax : db.createCollection("games",{capped: true,size: __)



// 2. embedded document 
// . embedded document is a document that is nested within another document  
  

// 3. namespace 
// .  (databasename.collectionname)
// . combination of the database name and collectionname  of a database ,which uniquely indentify a collection with a mongodb instance
// .   identify the unique collection from the database 



// 4.mongoshell 
// . is an intreative js interface for mongodb , its allow to  user to intreact with mongo db and its allow to use the opreation in the mognodb, such as insrtion, deletion, updation querrying like that, and its offer addminstrative commands for manage  mongodb ,

// 5 . upsert 
// . when we update a field in the document if the field is not there, so that time when we using the upsert , it will create a new document and insert the value ,

// 6. BSON
// . Binary encoded javascript object noation (json)
// . its widely used to store data 
// .json is easly understand to human, but bson its not easy 
// . bson : build in binary/ j - text formate
//        : slow to read but faster to build and scan/j - faseter to read but slower to build 
//        : encoded before storeing and decoded before displaying /j - without encoding we can send to the api's
//        : database is use bson to store / used to send data through the network 

// .data type : minkey,maxkey,binary data,objectID 

// 7. Distinct 

// . it is method from mondo db , and its return  array of unique values 
//                     or 
// .it is method is used  to find distinct value from a specific field of a collection 


// 8. gridFs 
// . gridFs is native featur of mongo db , it used to storing a largefiles in mongo,like image ,vidos bson document like 16 megabyte data like these data. its divided into chunks each of the chunks inserted into the doucment of a collection 

// . when we working with gridFs you intract with fs bucket ,which is vircual container for storing and retrieving  files 

// 9. Advantage of the mongoDb
// . sharding 
// . replication , are there so we can easly scale up and down , 
// . we can find data with patteren 
// . we can parition data (sharding)
// . support primarynode and secondary index on any field


// 10.Document 

// .document is like rows in sql, in document we store data  json like formate , but its store in bson 


// 11. collection 
// . its is group of document  we can call collection , its like table from the sql , 

// 12. 12. What are the databases in mongoDb
// . admin,local, config, these database are already craete by the mongodb 


//13 what is mongo Shell 
// . 



// $unwind : make single data 
// $lookup : join 
// $regex  : to find with pattern {}




