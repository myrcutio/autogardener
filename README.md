### Auto Gardener

This repo deploys a serverless application on AWS IoT/DynamoDB/Lambda for a watering scheduler on a custom Arduino device.  Schedules are defined by cron spec to determine next run time, and when triggered will send a 10-second watering MQTT signal to the devices.

### Setup
run `go install` to install the command line utility to the go path.  then run `autogardener -h` to see available actions.

### CLI Actions
```shell
# use the `benny` aws profile to deploy the application
autogardener deploy --profile benny

# add a new schedule and target plant available water (PWA) level, in bars. 
# Output is a header file with configured cert and keys to be used in provisioning firmware
autogardener register --plant dramaqueen \
  --schedule weekly \
  --target-paw .3
```