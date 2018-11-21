# Ben & Jerry's API


Few things need to be set up before testing the APIs
1. DB Server needs to be set up.
   -  In my case, I have taken Mongo DB
2. Redis Server needs to be set up
   -  This will help us to save the access token.
2. Go API server needs to be set up..
   -  Run the goapi.exe as a service

### Installation Steps
Open swagger/swagger.yaml file in editor.swagger.io .

Click on the Authenticate button , a pop up opens asking you enter a value.

OOPS !. What is the value? ..... Thinking !!!

Go to bin folder, Check for the config file, and set the port number(For now, let it point to localhost)

Double Click the icecream.exe file. This opens the cmd prompt. Hurry :). Server is started.

In editor.swagger.io, Call the v1/authorize/token API.

Enter user id and password as "biswa" and "iamcoming" respectively. This would return a access token.This is maintained in clients.json file

Hurray :). We are almost there.

Go to Step 2. Authorize yourself with the below string . For eg: Access token is "xxxxxxxxx", then kindly enter "Bearer xxxxxxxxx" . Yes, you are in now

Now, you are authorized to access the APIs .

While trying any API, Press Execute Button. OOPS !. 500 Internal Server Error.

No more issues please. What did we missed? Thinking ...........

OOPS !. Where is my database set up?

Refer ReadMeDB.md for database set up.

Once the set up is done, you are all set to use the APIs.

