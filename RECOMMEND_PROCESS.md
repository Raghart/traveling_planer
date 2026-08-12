# Process steps for new Brain

1. Gather user data for his/her desired travel.
   - **How**: TUI. 
   - **Language**: Go.
  
2. Save the data in the cache (Redis / some other cache framework).
    - **How**: Redis / PostgreSQL.
    - **Language**: Go.
   
3. If cache empty, gather the required data from the DB or from the APIs.
    - **How**: Go's routine concurrency.
    - **Language**: Go.
  
4. Execute the analyze algorithm to determine which country would be best given X user data.
    - **How**: Using the data gathered in the previous step to calculate which country matches best to the user wanted data.
    - **Language**: A language that does calculations and comparations fastly.  
  
5. Deliver the data for the user to check out.
    - **How**: Delivered the most recommended country data to show through RabbitMQ.
    - **Language**: RabbitMQ.
  
6. Inform the user about the option of inputing a traveling advice through a PDF.
    - **How**: TUI to show the data / N8N Microservice with data.
    - **Language**: Go *(TUI)* / Python *(PDF)*