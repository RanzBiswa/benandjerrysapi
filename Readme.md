#Read me

Few things need to be set up before testing the APIs
1. DB Server needs to be set up.
   -  In my case, I have taken Mongo DB
2. Redis Server needs to be set up
   -  This will help us to save the access token.
2. Go API server needs to be set up..
   -  Run the goapi.exe as a service


Installation of Mongo DB in Windows
1. Go To the folder where your mongo is installed.
   For eg: cd D:\mongodb-win32-x86_64-2008plus-ssl-4.0.4
2. Mongo DB requires a data directory to store all data.
   Create the folder by writing "md data"
3. Create the path by using   "mongod.exe --dbpath D:\\DATA". This starts the mongo db server
4. Once Step 3 is completed successfully. Open a new command prompt shell and write "mongo.exe" to connect the mongo db server
5. Eureka :). We are all done with the database set up.



