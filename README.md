# go-url-shortener
## URL Shortener API Built Using GO

### Requirements
1. Application uses Docker Compose. User must have Docker Compose installed. 
    - Instruction can be found here: 
       https://docs.docker.com/compose/install/

### Installation and Use
1. Download Zip file and extract or clone repo
2. Navigate to the directory that was extracted or cloned in a terminal
3. Run ` # docker-compose up`. This will start the API running at `localhost:9888`
    - Open a web browser window and navigate to `localhost:9888/`. A response should be displayed.
4. The API can be interacted with using POST and GET commands from the terminal 
    
    ##### Examples: 
    
      ###### POST request to shorten a URL     
            
    ` curl --request POST --data '{"long_url": "https://www.cyclingnews.com"}'   http://localhost:9888/api/v1/shorten/ `
    
      ###### GET request to expand shortened URL
      
     ` GET http://localhost:9888/aKC66gdK `

