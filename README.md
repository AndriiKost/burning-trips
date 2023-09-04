## Burning Trips

### Summary

Do you like to travel? One of the most time and energy-consuming tasks when traveling is what to visit. assuming you aren't at the resort or similar "one-stop place". 

Rather than spending your time on figuring out exactly what is out there worth visiting, provide a preference on what you like, and let an AI-powered travel planning system do that for you. Lastly, you will select the places from the recommendation list.

Let's say you have only one day dedicated to exploring and enjoying the entertainment in the desired location. Save your time and make the most out of your time by taking a closer look at "Burning Routes". This handpicked list of stops will save lots of your time while making sure you are visiting the best of the best landmarks that are out there.

You do not have to rearrange each stop until you find the most efficient way to visit all of them, this has been taken care of by our algorithm.


### Technical description

Multiple Docker services are being initialized by docker-compose.

- Front-end application - `burning-client`. Modern javascript framework Vue.js with typescript superset.
- Server application - `burning-gateway`. API gateway is written in Go.
- Reverse proxy - `burning-nginx`. Routing within the instance using nginx.
- Automated SSL certification renewal image - `certbot` that utilizes Let's Encrypt. Configured to work on aws ec2 instance completely autonomously.
- Web scraper - looks for information and details for over 100 000 landmarks around the globe.
- Database - PostgreSQL - AWS RDS.

The application was hosted on aws (currently offline) using services: EC2, Route 53, S3, Load Balancer, Serverless Lambda and RDS.
