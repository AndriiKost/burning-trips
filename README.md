## Burning Trips

### Technical description

Multiple Docker services are being initialized by docker-compose.

- Front-end application - `burning-client`. With help  of modern javascript framework - Vue.js, and a superset for js language - Typescript, the client application follows most of the modern best practices.
- Server application - `burning-gateway`. Written in modern language Golang with a good test coverege
- Reverse proxy - `burning-nginx`. 
- Automated SSL certification renewal image - `certbot` that utilizes Let's Encrypt. Configured to work on aws ec2 instance completely autonomously.
- Scraper runs only when needed - scrapes information and details for over 100 000 landmarks around the globe.
- PostgreSQL database is hosted on AWS RDS.

The application is hosted on aws (currently offline) using services: EC2, Route 53, S3, Load Balancer, Serverless Lambda and RDS.

### Application Summary

Traveling somewhere? Check out our application to get the best places to see, rated by travelers - for travelers.

We introduce the ranked list of most loved landmarks in the area you wish to visit curated by people who visited them.

Let's say you have only one day dedicated to exploring and enjoying the entertainment in the desired location. Save your time and make the most out of your time by taking a closer look at "Burning Routes".
This handpicked list of stops will save lots of your time while making sure you are visiting best of the best landmarks that are out there.

You do not have to rearrange each stop until you find the most efficient way to visit all of them, this has been taken care of by our algorithm.

