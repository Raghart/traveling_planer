# Process steps for new Brain

- Gather user data for his/her desired travel.
  
- Save the data in the cache (Redis / some other cache framework).
   
- If cache empty, gather the required data from the DB or from the APIs.
  
- Execute the analyze algorithm to determine which country would be best given X user data.
  
- Deliver the data for the user to check out.
  
- Inform the user about the option of inputing a traveling advice through a PDF.