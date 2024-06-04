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

I'll elaborate my building process in detail into this document. First, I think that the 
