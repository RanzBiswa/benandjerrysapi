# Ben & Jerry's API

This repository contains a collection of Go programs and libraries that demonstrates the implementation of CRUD operations for Ben & Jerry's API. :)

##### Things Used
Go Programming Language, MongoDB and mgo(MongoDB driver for go Lang)


#### Pre Requisites
Few things need to be set up before testing the APIs
1.  DB Server needs to be set up. I have coded for Mongo DB. Mongo DB needs to be set up.
2.  GO APIs needs to be hosted. Either of two options mentioned below needs to be implemented.
    -  Either Go API server needs to be set up using the file "benandjerrysapi/bin/benandjerrysapi.exe" 
       OR 
    -  Run the project benandjerrysapi/bin/benandjerrysapi.exe . Check the benandjerrysapi/bin/benandjerrysapi.cfg file for checking the ports.

### Installation Steps
1. Open benandjerrysapi/swagger/swagger.yaml file in editor.swagger.io .
2. Once it's open. Click on the Authorize button, a pop up opens asking you enter a value.
3. OOPS !! What is the value ? . Thinking !!!
4. Go to bin folder, Check for the config file, and set the port number(For now, let it point to localhost)
5. Double Click the icecream.exe file. This opens the cmd prompt. Hurry :). Server is started.
6. In editor.swagger.io, Call the v1/authorize/token API.
7. Enter user id and password as "biswa" and "biswa1234" respectively. 
    -  This would return a access token. 
8. Hurray :). We are almost there.
9. Go to Step 2. 
    -  Authorize yourself with the below string . 
    -  For eg: Access token is "xxxxxxxxx", then kindly enter "Bearer xxxxxxxxx". 
    -  Yes, you are almost there :)
10. Now, you are authorized to access the APIs .
11. While trying any API, Click on Execute Button. 
12. OOPS !. 500 Internal Server Error.
    -  No more issues please. What did we missed? Thinking ...........
13. What did I miss?  OOPS !. Where is my database set up?  
    - Set up the mongo database. 
    - Copy "benandjerrysapi/readmefiles/mongoInsertScript.txt" and run in mongo instance.
    - Database set up with data is completed.
    - If you would like to create everything fresh, feel free to call create api with json present in "benandjerrysapi/readmefiles/mongoCreateScript.json"
14. Once the set up is done, you are all set to use the APIs.


## Note
User id and password's are maintained in clients.json file.