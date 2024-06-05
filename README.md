# Junior Developer Technical Test

This is a technical test that has the following objectives:
1. *Enterprise auth:* the system must be engineered in a way such that only registered users can access the resources stored inside the system. Also, users should only have access to their own information. My solution is to use a permissions table for the first requirement, and a filter for the second.
2. *Electronic bill upload:* the end user must be able to upload bills in JSON. Also the system must be capable of validating the electronic bills that are being sent into the system so that they meet the minimal requirements. I'll be doing validation checks as shown in *Let's Go Further* by **Alex Edwards**.
3. *Bill Management:* the end user must be capable of doing the basic ***CRUD*** operations. So, basic *creating, reading, updating and deleting* operations must be performed by the program. Regarding how certain operations must be done, I have my own thoughts:
    - Issuing bills must be done one at a time. As a first approach, I find this to be sufficient. 
    - Reading bills must be done in batches. I think that using creation dates will be useful the long run. So I'll be using a time interval in order to access bills in batches.
    - Updating must be done one at a time. My reasons are twofold, first, this ensures that the updating of sensitive information is done slowly, and second, it's easier to implement.
    - Deleting bills should be done one at a time. Due to the sensitive information inside, I think that deleting bills one at a time reduces the risk of third actors tampering with the information.
4. *Filtering:* every user should just access the information that they own. I'll use a ```uuid``` in order to filter the information that the user has access to.

I'll elaborate my building process in detail into this document. First, I think that i'll elaborate in my implementation regarding the backend:
- Backend-wise I'll use a JWT system in order to authenticate the requests coming from the client. The other implementation details are quite simple. I use something of a style of hexagonal architecture regarding the different services used in the backend, normally I try to decouple services that depend on one another using interfaces, such that they encapsulate all the methods being used inside a service you can see them being used a lot in the handlers service and the models service. This makes it so that new features can be added into a service without necessarily changing the implementation details of other services and enables the posibility of injecting new components into an existing service with the addition of another interface.
- In addition to this, all the dependency management (services and components used in the app) is done inside de app directory. With this, the user can be sure which services are loaded as dependencies for the app. With this, we can change the implementation details of the app services or replace them entirely and the app will work in the same way due to the fact that the methods and services are declared as empty interfaces.
- Regarding the Frontend, I used React with Typescript enabled. In order to have type safety and compatibility with the UI library that was used into the project. I used ```shadcn/ui``` as a component library, that enabled me to develop things quickly and get to testing fast, reducing the struggle regarding component styling and placement. Also, I used the Tanstack library: ```Router```, ```Query``` and ```Table``` in order to build basic data dashboards and form queries. I used ```Table``` in the bill fetching view that is included inside the app. ```Query``` was used for basic data fetching and authentication, using the jwt dispatched by the server.  

## Features

![image](/screenshots/Fetch.png)

So far I have developed a system that can do basic bill management operations, which are:

- Bill Creation: It's done one at a time, mimicking the actions of a clerk inside a store. That being said, the form that does the creation can be accomodated in order to meet the details and miscelaneous information that the companies will need to store in the bill information. The authentication is done both in the client (using ```zod```) and in the backend, using validation checks regarding the desired shape that the bill fields must adhere to. 

- Bill Removal: Regarding this, it's also done one at a time. I plan to add the bulk version of this method in a future iteration of this type of preject. 

- Bill Fetching and Reading: I dedicated two separate endpoints regarding this feature. I use an endpoint that does the fetching of all the bills that were created between two dates. The reading of individual bills can be done inside the client, but I still made an endpoint in order to fetch individual bills. Regarding future upgrades, I think that the next logical step is implementing a pagination system for the api and the client.

- Bill Update: With this feature I decided to use a *PUT* endpoint, mostly due to the fact that I'll fill the other fields in the update form using placeholder values lifted directly from the bill that will be updated.  

## How to run the project

1. Regarding the backend: you can run the following commands in the root directory order to run the project:
```
go build ./cmd
```

2. Regarding the frontend: you will need to perform the following commands from the root directory of the project in order to run the client:
```
cd client_side
npm install 
npm run dev
```
