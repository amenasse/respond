#!/bin/bash

aws lambda create-function --function-name respond --zip-file fileb://function.zip --handler lambda --role arn:aws:iam::172428475609:role/LambdaRespond --runtime go1.x
