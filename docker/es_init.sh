#!/bin/bash

# Elasticsearch host
ES_HOST="http://localhost:9200"

# Index name
INDEX_NAME="menu"

# Create index
curl -X PUT "$ES_HOST/$INDEX_NAME" -H 'Content-Type: application/json' -d'
{
    "settings": {
        "number_of_shards": 1,
        "number_of_replicas": 1
    },
    "mappings": {
        "properties": {
            "name": {
                "type": "text"
            },
            "price": {
                "type": "float"
            }
        }
    }
}'

# Insert seeding values
curl -X POST "$ES_HOST/$INDEX_NAME/_bulk" -H 'Content-Type: application/json' -d'
{ "index" : { "_id" : "1" } } 
{ "name": "focaccia", "price": 5 } 
{ "index" : { "_id" : "2" } } 
{ "name": "biancaneve", "price": 5.5 } 
{ "index" : { "_id" : "3" } } 
{ "name": "margherita", "price": 6.5 }   
'

# curl -X POST "$ES_HOST/$INDEX_NAME/_doc/2" -H 'Content-Type: application/json' -d'
# {
#     "name": "Jane Smith",
#     "age": 25,
#     "email": "jane.smith@example.com"
# }'