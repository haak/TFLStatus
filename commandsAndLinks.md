https://tfl.gov.uk/tube-dlr-overground/status/#

http://cloud.tfl.gov.uk/TrackerNet/PredictionDetailed/H/BST


https://tfl.gov.uk/info-for/open-data-users/unified-api?intcmp=29422

https://tfl.gov.uk/info-for/open-data-users/open-data-policy

https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Line/Line_GetByMode

https://tfl.gov.uk/info-for/open-data-users/api-documentation#on-this-page-3



///Mode/{mode}/Arrivals

// take a station name
// show timetable for that station

//This section is going to represent the TracketNet API that is available at
// 4 urls
// prediction service summary
// prediction services detailed
// station status (same as line status usage)
// line status (this could be used to change text colour of station maybe red if its behind or something)

// Prediction Summeary takes a line


// func (m Mode) String() string {
// 	output :=
// }



mode object

Client object

NewClient()

error object



ModeService object 

NewModeService()


getModes() takes client returns mode



## OLD API
http://cloud.tfl.gov.uk/TrackerNet/PredictionDetailed/H/BST
http://cloud.tfl.gov.uk/TrackerNet/PredictionSummary/H



## NEW API
naptan id == stopcode i think



## STORAGE 
each line has a thing
each station could also store info
might be easier  to store it by line 

there is an api call that gives geo information 

## JOURNEY PLANNER - (JOURNEY)


/Journey/Meta/Modes
/Journey/JourneyResults/{from}/to/{to} - takes naptan id or somthing


## LINE 

Line/Meta/Modes

/Line/Meta/ServiceTypes

/Line/{ids} - dont know what ID

/Line/Mode/{modes} - 

/Line/Route
https://api.tfl.gov.uk/Line/Route


/Line/{IDS}/Route - https://api.tfl.gov.uk/Line/bakerloo/Route

/Line/{id}/Route/Sequence/{direction} - 



/Line/{id}/StopPoints - gives all stop points related to a line which is LOADS



https://api.tfl.gov.uk/Line/Bakerloo/Arrivals/940GZZLUKEN

/Line/{ids}/Arrivals/{Stoppoint} -- this could be partially used to find trains

use /link/linename/Arrivals



## MODE 
/Mode/{Mode}/Arrivals -- all tube stations status



NOw i need the lat and long of a stop id for the  station

i can use /StopPoint/{point} to get lat and long

this can then be plotted on a map somehow

## TODO
* Store list of tube line names and codes in struct and save in json file
* make way of generating the above file

* take 


## Command Line Arguments
-modes

-mode

-lines
-line

-stations
need to supply both a line and -stations to get the stations for a line