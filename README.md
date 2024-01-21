# Basic CRM Using Fiber Poject:

This is a simple Customer Relationship Management (CRM) application built using the Fiber framework and SQLite database. 
It allows you to manage leads with basic information like name, company, email, and phone number.

## PICTURES:
##### running code:
<img width="400" alt="running_Code" src="https://github.com/Masanbat12/Basic_CRM_Using_Fiber/assets/93978448/468f9799-3cef-482a-8891-0269cd77c324">

<img width="400" alt="running_code_port_7000" src="https://github.com/Masanbat12/Basic_CRM_Using_Fiber/assets/93978448/9f8f2916-b772-4536-9004-f2e5aac1c0a2">


##### Insert new lead:
<img width="600" alt="insertion_of_data" src="https://github.com/Masanbat12/Basic_CRM_Using_Fiber/assets/93978448/3f94ef86-ef63-4c8d-aa79-6910292d2c0f">

#### Adding Id:
###### BEFORE:
<img width="400" alt="Adding_22_for_id_to_change" src="https://github.com/Masanbat12/Basic_CRM_Using_Fiber/assets/93978448/45b79f3d-6921-4990-89d6-7282da0b2ffc">

###### AFTER:
<img width="400" alt="Adding_2_worked" src="https://github.com/Masanbat12/Basic_CRM_Using_Fiber/assets/93978448/b1caf42a-fef1-473f-8e8b-726050aafa62">

##### delete a lead:
<img width="400" alt="Deleting_Object_Success" src="https://github.com/Masanbat12/Basic_CRM_Using_Fiber/assets/93978448/71ae0398-826f-4558-8ab6-80daf94398e1">

## Application Structure:
#### main.go: 
This is the entry point of the application. 
It sets up the Fiber web server, initializes the database, defines API routes, and listens for incoming requests.

#### lead/lead.go: 
This file contains the definition of the Lead struct, which represents the data structure for leads. 
It also includes functions to handle CRUD (Create, Read, Update, Delete) operations for leads.

#### database/database.go: 
This file establishes the database connection using GORM (Object-Relational Mapping),
and defines a global database variable to be used throughout the application.

## Database Initialization:
When the application is started, the initDatabase function in main.go is called.
This function establishes a connection to an SQLite database named leads.db located in the project directory. 
It also performs an auto migration to ensure that the Lead model is properly mapped to the database.

## API Endpoints
#### The application provides the following API endpoints:
##### 1. Get All Leads
Endpoint:

/api/v1/lead
Method: GET
Description: Retrieves a list of all leads from the database using the GetLeads function defined in lead/lead.go. 
The result is returned as a JSON response.

##### 2. Get Lead by ID
Endpoint:

/api/v1/lead/:id
Method: GET
Description: Retrieves a lead by its unique ID from the database using the GetLead function defined in lead/lead.go. 
The result is returned as a JSON response.

##### 3. Create a New Lead
Endpoint: 

/api/v1/lead
Method: POST
Description: Creates a new lead by parsing the JSON request body and using the NewLead function defined in lead/lead.go. 
The newly created lead is returned as a JSON response.

##### 4. Delete Lead by ID
Endpoint: 

/api/v1/lead/:id
Method: DELETE
Description: Deletes a lead by its unique ID from the database using the DeleteLead function defined in lead/lead.go. 
A success message is sent as a response.

## TOOLS (And examples how to use in postman):
You can use tools like Postman or curl to interact with the API endpoints.
Here some examples to Get All Leads:
localhost:7000/api/v1/lead,
For creating new lead (need to add in JSON):
'{
    "name": "Jane Sm",
    "company": "XYZ Corp",
    "email": "jane.sm@example.com",
    "phone": 1234567890
}'
And then POST : localhost:7000/api/v1/lead
Also to change id : localhost:7000/api/v1/lead/2 <- adding after 'lead/' any number.

if you chose to use postman then you can test it like this: 

To interact with the CRM application using Postman:
Open Postman (and download postman agent).
click on "new" button.
Import or manually create requests using the base URL http: //localhost:7000. Or any port you get.
Send requests to the various endpoints (e.g., GET, POST, DELETE) to interact with the CRM and observe the responses.

## Setup:
Make sure you have Go and Fiber installed on your system.
go get -u github.com/gofiber/fiber/v2

## Installation:
Clone this repository.
make sure you work in c enviroment.

## Initialize the database:
go run main.go.

This will create an SQLite database named leads.db in your project directory.
Also Start the CRM application.
The application will start, and you can access it at http: //localhost:7000.

## Contributing:
Contributions to enhance the functionality or design of the application are welcome. Please follow standard pull request procedures.
