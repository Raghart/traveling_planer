# Process steps for new Brain

1. Gather user data for his/her desired travel.
   - **How**: TUI. 
   - **Language**: Go.
   - **Questions**:
     - Where are you from?
     - When are you planning to travel?
     - Do you prefer *hotter* or *colder* enviroments?
     - Which factor calls your attention when choosing for a place to travel? (*Nature-wilderness*, *City-life style*, *History-culture*)

2. Check **if** there is a similar request in the cache.

    1. If the cache doesn't have a match, gather the required data from the DB or from the APIs and save it.
       - **How**: Go's routine concurrency / Redis or other cache framework to save the data.
       - **Language**: Go / Redis (*Cache*).
  
    2. If cache has a match, deliver the required data to the algorithm.
       - **How**: Redis / PostgreSQL.
       - **Language**: Go.
  
3. Execute the analyze algorithm to determine which country would be best given X user data.
    - **How**: Using the data gathered in the previous step to calculate which country matches best to the user wanted data.
    - **Language**: A language that does calculations and comparations fastly.  
  
4. Deliver the data for the user to check out.
    - **How**: Delivered the most recommended country data to show through RabbitMQ.
    - **Language**: RabbitMQ.
  
5. Inform the user about the option of inputing a traveling advice through a PDF.
    - **How**: TUI to show the data / N8N Microservice with data.
    - **Language**: Go *(TUI)* / Python *(PDF)*