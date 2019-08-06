# Links for relevant APIs



## Visitors api
- http://developer.pardot.com/kb/api-version-4/visitors/
- Parameters:
    - created_after: starting date
    - created_before: end date 
    - only_identified: boolean value of true will return only prospects
    - offset: Pardot API returns only 200 results at a time - there are typically about 500-800 prospects every week that we want to be returning, so this offset will be necessary for each API call
    - will return <visitor> information tag for each visitor - use read (/api/visitor/version/4/do/read/id/<id>) with the visitor id to get other information like number of visits, minutes on site, etc
