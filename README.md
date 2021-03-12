HACKATHON
Hackathon is a service which expose REST APIs to create and collabrate challenges for the employees.

Steps to setup Database
1. Install postrges
2. create new database i.e. hackathon
3. execute database scripts from "dbScripts/hackathon.sql"
 

Steps to Run Application 
1. Update configuration for database in "config/config.yaml" file
2. Build docker image using below command
     docker build -t hackathon .
3. Run container using below command
    docker run -p 8080:8080 "hackathon"

