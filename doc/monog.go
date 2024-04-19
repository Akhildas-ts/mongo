package main 

// PROGRAMM 

// 1.CREATE COLLECTION 
// db.createCollection("product")


// 2.CHECK DOCUMENT IN THE COLLECITON 
// db.product.find()


// 3.ADD A NEW FIELD DOCUMENT / UPDATE  
// db.product.updateOne(
// { _id : ObjectId( 'unique id ') },
// { $set : {price : 500 } )   //ADD/UPDATE what i need to add into the document 


// 4.DELETE DOCUMENT 
// db.product.remove({product_name: "lenova 3"})  

// 5. INSERTDOCUMENT 
// db.product.insert({ proudct_name : "hpTail", descrption: "hello world", price: 500},)

// 6. REMOVE ALL IN THE COLLECTION
// db .product.remove ({})  

// 7.LIMIT AND SKIP 
// db.product.find().limit(2).skip(2).pretty()

// 8.FIND EQUALS 
// db.product.find({pice: {$eq: 500}})

// 9.FIND GREATER THAN 
// db.product.find({price: {$gt: 500}}) //  find greater than document in these price .. 

//   find greater or equal to 
// db.product.find({price: {$gte: 500}}) //

// 10.GET DOCUMENT WHAT ARE IN 
// db.product.find({product_name : {$in: ["HP laptop", "lenova gamming",]}) //useing array , we got the document where the vlaues is availbale .

// 11. LES THAN // LESTHAN OR EQUAL 
// db.product.find({pice: {$lt: 5500} })
//        equal                             // get less than or equal to the document 
// db.product.find({price: {$lte: 5500}} )


// 12. FIND WITH MULTIPLE CONDITION
// db.product.find({$and : [{product_name : {$eq : "HP laptop"}}, {price :{$eq: 500}} ] })

// 	i
// 13. NOT EQUAL . 
// db.find.product.find({ product_name: {$not : {$eq: "HP laptop"}} })	

// 14. NOT THE VALUE (opposite of and )
// db.find.product.find({$nor: [ {product_name: {$eq: "lenkscart"} }, {price: {$eq: 2000}} ]})


// 15.OR 
// db.find.product.find({$or: [ {product_name: {$eq: "apple"} }, {price: {$eq: 500}}, ]})


// 16 NOt IN 
// db.find.product.find({product_name: {$nin: ["dumbajalle", "apple kanndi"] } } ) 

// 17 UPDATE WITH ALL DOCUMENT 
// db.games.updateMany({}, {$set: {age: 0} })


// 18 DELTE ALL filed At same time in  every document

// db.games.updateMany({},{$unset: {age: ""} })

// 19. delete document with specific contion : 
// db.games.deleteOne({name: "BGMI'} )

// 20. 


