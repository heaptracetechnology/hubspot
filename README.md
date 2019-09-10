# _Hubspot_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/hubspot.svg?branch=master)](https://travis-ci.com/heaptracetechnology/hubspot)
[![codecov](https://codecov.io/gh/heaptracetechnology/hubspot/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/hubspot)


An OMG service for hubspot, to create, update, delete contacts and to create deals and tickets in hubspot.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Create Or Update Contact
```coffee
hubspot createOrUpdate email="someone@example.com" properties={"properties":[{"property": "phone", "value": "789456xxxx"},{"property": "firstname", "value": "john"},{"property": "lastname", "value": "smith"}]}
```
##### Get Contact
```coffee
hubspot getContact vid=651
```
##### Delete Contact
```coffee
hubspot deleteContact vid=651
```
##### Create deal
```coffee
hubspot createDeal properties={"associations": {"associatedVids":[601]},"properties": [{"value": "Tims Newer Deal","name": "dealname"},{"value": "appointmentscheduled","name": "dealstage"},{"value": "default","name": "pipeline"},{"value": 1409443200000,"name": "closedate"},{"value": "60000","name": "amount"},{"value": "newbusiness","name": "dealtype"}]}
```
##### Create Ticket
```coffee
hubspot createTicket ticketProperties=[{"name": "subject","value": "This is an example ticket"},{"name": "content","value": "Here are the details of the ticket."},{"name": "hs_pipeline","value": "0"},{"name": "hs_pipeline_stage","value": "1"}]
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Create Or Update Contact:
```shell
$ omg run createOrUpdate -a email=<CONTACT_EMAIL_ADDRESS> -a properties=<MAPPED_CONTACT_PROPERTIES> -e API_KEY=<API_KEY>
```
##### Get Contact:
```shell
$ omg run getContact -a vid=<CONTACT_VID> -e API_KEY=<API_KEY>
```
##### Delete Contact:
```shell
$ omg run deleteContact -a vid=<CONTACT_VID> -e API_KEY=<API_KEY>
```
##### Create deal:
```shell
$ omg run createDeal -a properties=<MAPPED_DEAL_PROPERTIES> -e API_KEY=<API_KEY>
```
##### Create Ticket:
```shell
$ omg run createTicket -a ticketProperties=<MAPPED_TICKET_PROPERTIES> -e API_KEY=<API_KEY>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/godaddy/blob/master/LICENSE).
