# fakeS3

FakeS3 is a RESTful API project written in Golang using the Fiber framework.
It replicates some endpoints of AWS S3 API, such as create bucket, put object, list objects, and get object.
It uses MongoDB as database management system and can be run with Docker using docker-compose up command.

## Installation

To install the project, follow these instructions:
1. Clone the repository git clone https://github.com/antoniocapizzi95/fakeS3.git
2. Go to the project root directory `cd projectname`
3. Create a new `.env` file by copying the `.env.sample` file with command `cp .env.sample .env`
4. Run the project with Docker using command `docker-compose up`