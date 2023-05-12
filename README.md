# fakeS3

FakeS3 is a RESTful API project written in Golang using the Fiber framework.
It replicates some endpoints of AWS S3 API, such as create bucket, put object, list objects, and get object.
It uses MongoDB as database management system and can be run with Docker using docker-compose (MongoDb will be executed togheter with application in a separate container).

## Installation

To install the project, follow these instructions:
1. Clone the repository git clone https://github.com/antoniocapizzi95/fakeS3.git
2. Go to the project root directory `cd projectname`
3. Create a new `.env` file by copying the `.env.sample` file with command `cp .env.sample .env`
4. Run the project with Docker using command `docker-compose up`
5. With running project on Docker, it's possible to execute unit tests using this command `docker exec fakeS3 go test ./...`


## Exposed endpoints

### Create a Bucket
Endpoint:

#### PUT /bucket

Input:

    bucket: a string that represents the name of the bucket to be created.

Output:

If successful, the HTTP response code should be 200 OK. No output body is returned.


### Add an Object inside a Bucket
Endpoint:

#### PUT /bucket/key+

Input:

    bucket: a string that represents the name of the bucket where the object should be added to.
    key+: the path of the object being added.
    
    Body: the file to be uploaded in binary format

Output:

If successful, the HTTP response code should be 200 OK, with the indication of object Etag in the response headers. No output body is returned.


### List Objects inside a Bucket
Endpoint:

#### GET /bucket

Input:

    bucket: a string that represents the name of the bucket whose objects should be listed.
    
    Params: marker, max-keys and prefix. But the only parameter that will be considered is "prefix", it allows to return all objects in a bucket that have, in their key, the required prefix.

Output:

If successful, the HTTP response code should be 200 OK. The body of the response should contain a list of objects within the bucket, it's formatted in XML, like the one below:

    <ListObjectsOutput>
        <Name>testBucket</Name>
        <Prefix></Prefix>
        <Marker></Marker>
        <MaxKeys>0</MaxKeys>
        <IsTruncated>false</IsTruncated>
        <Contents>
            <Key>folder/image.png</Key>
            <ETag>d41d8cd98f00b204e9800998ecf8427e</ETag>
            <CreationDate>2023-05-12T09:57:51.397Z</CreationDate>
            <LastModified>2023-05-12T09:57:51.397Z</LastModified>
            <Size>204</Size>
        </Contents>
        <Contents>
            <Key>folder/image2.png</Key>
            <ETag>d41d8cd98f00b204e9800998ecf8427e</ETag>
            <CreationDate>2023-05-12T09:58:00.927Z</CreationDate>
            <LastModified>2023-05-12T09:58:10.383Z</LastModified>
            <Size>315</Size>
        </Contents>
    </ListObjectsOutput>

### Get an Object
Endpoint:

#### GET /bucket/key+
Input:

    bucket: a string that represents the name of the bucket where the object can be found.
    key+: the path of the object to be retrieved.

    Headers: Range - it is used to download only a fraction of the requested file, a string of this type "bytes:start-end" must be entered, where "start" and "end" are the range of bytes to be retrieved. For example "bytes:1-150" this means that only the first 150 bytes of the requested file will be downloaded.
Output:

If successful, the HTTP response code should be 200 OK. The body will contain the binary format file that was requested.

## Compatibility with AWS CLI
When the application is running, it can also be used with the AWS CLI, following are some example commands:

### Create a Bucket

    aws s3api \
    --no-sign-request \
    --endpoint-url http://localhost:8080 \
    create-bucket \
    --bucket <bucket-name>

### Add an Object inside a Bucket

    aws s3api \
    --no-sign-request \
    --endpoint-url http://localhost:8080 \
    put-object \
    --bucket <bucket-name> \
    --key <key> \
    --body <path of the file to upload>

### List Objects inside a Bucket
    aws s3api \
    --no-sign-request \
    --endpoint-url http://localhost:8080 \
    list-objects \
    --bucket <bucket-name>

Or with prefix:

    aws s3api \
    --no-sign-request \
    --endpoint-url http://localhost:8080 \
    list-objects \
    --bucket <bucket-name> \
    --prefix <prefix>

### Get an Object

    aws s3api \
    --no-sign-request \
    --endpoint-url http://localhost:8080 \
    get-object \
    --bucket <bucket-name> \
    --key <key> <path-to-save-file-locally>

Or with Range:

    aws s3api \
    --no-sign-request \
    --endpoint-url http://localhost:8080 \
    get-object \
    --bucket <bucket-name> \
    --key <key>  \
    --range <range "bytes=start-end"> <path-to-save-file-locally>

## Some implementation details
Considering that this project makes use of a no-sql DBMS (specifically MongoDB which is document-oriented), in order to avoid having to make joins in queries, the database structure has been "denormalized," in fact for each Bucket only one document is saved that has an array of Objects inside it (with a relational database it would have been necessary to have two separate tables).

In the main function of the project the Bucket repository implemented for Mongo is injected inside the requests controller, so if in the future the DBMS should change, just reimplement the repository interface and inject the new repository into the controller. It is important to note that inside the controller the use of Mongo rather than another DBMS is completely transparent.

A rigid folder structure was not followed to organize the project files, in Go there are different schools of thought, in this case an organized but simplified structure was used.

Finally, some unit tests were added only on the most logical functions that do not require mocking of dependencies.
