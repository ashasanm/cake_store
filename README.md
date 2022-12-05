# Cake Store

REST API Cake Store

This is an API to Create, Update, Delete and Get a Cake from Cake Store API

The entire application is build with golang programming language

## Requirements
To run this api server, docker need to be installed first:
### Docker:
  #### Linux Ubuntu
      read installation docs below
      https://docs.docker.com/engine/install/ubuntu/

  #### Mac Os
      to install docker on mac, visit docker link below
      https://docs.docker.com/desktop/install/mac-install/

  #### Windows
      to install docker on Windows visit docker link below
      https://docs.docker.com/desktop/install/windows-install/

## REST API
### Get List of Cake

#### Request
`GET /cakes/`
    curl -i -H 'Accept: application/json' http://localhost:8080/cakes/

#### Response

    {
      "data": [
        {
          "id": 5,
          "title": "Green Tea Cake",
          "description": "A Cake with green tea flavour",
          "rating": 10,
          "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
          "created_at": "2022-12-04T15:31:36Z",
          "updated_at": "2022-12-04T15:31:36Z"
        }, ...
      ],
      "total_data": 5
    }
### Get Cake by Id

#### Request
`GET /cakes/:id`
      curl -i -H 'Accept: application/json' http://localhost:8080/cakes/5
    
 #### Response
      {
        "id": 5,
        "title": "Green Tea Cake",
        "description": "A Cake with green tea flavour",
        "rating": 10,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
        "created_at": "2022-12-04T15:31:36Z",
        "updated_at": "2022-12-04T15:31:36Z"
      }
 
### Update Cake
`PATCH /cakes/:id`
     curl -i -H 'Accept: application/json' -X PATCH http://localhost:8080/cakes/1
#### Request Body
      {
        "title": "Cookies",
        "description": "A Cookies made out of chocolate and cream",
        "rating": 5.5,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
      }
#### Response
      {
        "message": "Cake successfully updated!",
        "status_code": 200
      }
      
### Create Cake
`POST /cakes/:id`
     curl -i -H 'Accept: application/json' -X POST http://localhost:8080/cakes
#### Request Body
      {
        "title": "Cookies",
        "description": "A Cookies made out of chocolate and cream",
        "rating": 5.5,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
      }
#### Response
      {
        "message": "Cake successfully created!",
        "status_code": 200
      }

### Delete Cake
`DELETE /cakes/:id`
     curl -i -H 'Accept: application/json' -X DELETE http://localhost:8080/cakes/5
     
#### Response
      {
        "message": "Cake successfully deleted!",
        "status_code": 200
      }
