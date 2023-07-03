start cmd /k go run node.go -h localhost:8000 -n localhost:8001
start cmd /k go run node.go -h localhost:8001 -p localhost:8000 -n localhost:8002
start cmd /k go run node.go -h localhost:8002 -p localhost:8001 -n localhost:8003
start cmd /k go run node.go -h localhost:8003 -p localhost:8002 -n localhost:8004
start cmd /k go run node.go -h localhost:8004 -p localhost:8003 -n localhost:8005
start cmd /k go run node.go -h localhost:8005 -p localhost:8004 -n localhost:8006
start cmd /k go run node.go -h localhost:8006 -p localhost:8005
REM -n localhost:8007
REM start cmd /k go run node.go -h localhost:8007 -p localhost:8006 -n localhost:8008
REM start cmd /k go run node.go -h localhost:8008 -p localhost:8007 -n localhost:8009
REM start cmd /k go run node.go -h localhost:8009 -p localhost:8008 -n localhost:8010
REM start cmd /k go run node.go -h localhost:8010 -p localhost:8009 -n localhost:8011
REM start cmd /k go run node.go -h localhost:8011 -p localhost:8010 -n localhost:8012
REM start cmd /k go run node.go -h localhost:8012 -p localhost:8011 -n localhost:8013
REM start cmd /k go run node.go -h localhost:8013 -p localhost:8012 -n localhost:8014
REM start cmd /k go run node.go -h localhost:8014 -p localhost:8013 -n localhost:8015
REM start cmd /k go run node.go -h localhost:8015 -p localhost:8014 -n localhost:8016
REM start cmd /k go run node.go -h localhost:8016 -p localhost:8015