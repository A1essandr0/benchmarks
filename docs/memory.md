## go-vs-faust
Performance under load for basic Go and Faust applications.

Comparing senders:   
faust-sender (FastAPI, Faust, 100 locust users):  58% CPU usage, 0.4% memory usage   
goka-sarama-sender (FastAPI, goka,  100 locust users):  6.5% CPU usage, 0.1% memory usage   

Comparing collectors:   
goka-sarama-collector (basic net/http) uses 3-4 times less CPU resources than fastapi-collector under the same load (goka-sarama-sender, 100 locust users)


## uvicorn vs uvicorn[standard]
Comparing uvicorn vs uvicorn[standard]    
fastapi-collector (pure python uvicorn impl):   40-45% CPU usage, 0.3% memory usage    
fastapi-collector-uvloop (cython uvicorn impl): 30-35% CPU usage, 0.3% memory usage    


## grpc vs rest
Comparing REST/http Go service and gRPC/tcp Go service

REST (100 locust users)    
goka-sarama-collector 10-12% CPU usage, 0.2% memory usage   
goka-sarama-sender    3% CPU usage, 0.3% memory usage    
400 RPS    

gRPC (100 locust users via grpc_interceptor)    
grpc-librd-collector 66% CPU usage, 2% memory usage    
goka-sarama-sender         12-13% CPU usage, 0.3% memory usage    
3000 RPS    


## librd vs franz
librd-collector 18% CPU usage, 0.4% memory usage    
librd-sender 6% CPU usage, 0.3% memory usage    

franz-collector 26% CPU usage, 0.1% memory usage    
franz-sender 22% CPU usage, 0.1% memory usage