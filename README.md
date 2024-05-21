# RedAbierta Technical Test

This is a technical test that has the following objectives:
1. *Enterprise auth:* the end user must be authenticated with an username and password that correctly gives permissions to said user in order to access a subset of the bills stored into the system. This also means that only the users that are registered by the administrator can make use of the resources inside the system. My implementation achieves this using a permissions table, so that the normal users just have permissions regarding to the bills, and the administrators have the permissions over the bills and the users.
2. *Electronic bill upload:* the end user must be able to upload the bills into the system using a JSON format. Also the system must be capable of validating the electronic bills that are being sent into the system. My idea behind the validation is doing validation checks as *Let's Go Further* by **Alex Edwards** displays in his book. The advantage of this system is that the validation is done at the domain level of the application and it is decoupled from the rest of the application, ensuring modularity at a fine level.
3. *Bill Management:* the end user must be capable of doing the basic ***CRUD*** operations. So the basic functionality to send a bill, update the bill if the need comes, reading the bill, and deleting bills, must be operations that need to be done with the product. In this sense I think I have my own ideas regarding the functionality:
- Creating bills must be done one at a time, most stores send bills in this manner, so as a first approach, I think that this is the minimal functionality needed as of now.
